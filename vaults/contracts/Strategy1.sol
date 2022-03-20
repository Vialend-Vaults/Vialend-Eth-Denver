// SPDX-License-Identifier: MIT

pragma solidity >=0.8.8 <0.9.0;

import "./@openzeppelin/contracts/token/ERC20/SafeERC20.sol";
import "./@openzeppelin/contracts/math/Math.sol";
import "./@openzeppelin/contracts/utils/ReentrancyGuard.sol";
import "./@uniswap/v3-core/contracts/libraries/FullMath.sol";
import "./@uniswap/v3-core/contracts/interfaces/IUniswapV3Factory.sol";
import "./@uniswap/v3-core/contracts/interfaces/IUniswapV3Pool.sol";
import "./@uniswap/v3-periphery/contracts/libraries/LiquidityAmounts.sol";
import "./@uniswap/v3-periphery/contracts/libraries/PositionKey.sol";
import "./@uniswap/v3-core/contracts/libraries/TickMath.sol";
import "./@uniswap/v3-core/contracts/interfaces/callback/IUniswapV3MintCallback.sol";
import "./@uniswap/v3-core/contracts/interfaces/callback/IUniswapV3SwapCallback.sol";
import  { ICErc20, ICEth,IWETH9 }  from "./interfaces/IViaProtocols.sol";

import "./UniCompFees.sol";
import "./TwapGetter.sol";
//import "./libraries/UniCompHelper.sol"; 
import "./libraries/Debugger.sol"; 

import "./interfaces/IViaVault.sol";
import "./interfaces/IVaultFactory.sol";

/*
Error code:
'OLD' getTwap error. rebalance cannot proceed. solution: setTwapDuratio = 0
'E5' not creator or admin

*/

/// @author  ViaLend
/// @title   strategy Uni + Compound
/// @notice  A Smart Contract that helps liquidity providers managing their funds on Uniswap V3 .

struct UniCompParam {
	address Uni3Factory;
	address VaultFactory;
	address Protocol; 
	address Creator;
	address Token0;
	address Token1;
	address CToken0;
	address CToken1;
	address WETH;
	address CETH;
	uint8 Token0Decimals;
	uint8 Token1Decimals;
	uint8 UniPortionRate;  
	uint8 CompPortionRate;  
	uint8 ProtocolFeeRate;  
	uint8 MotivatorFeeRate; 
	uint24 FeeTier;
	uint32 TwapDuration;
	int24 MaxTwapDeviation;
	uint256 VaultCap;
	uint256 IndividualCap;
}

contract VaultStrategy 
	is 
	ReentrancyGuard , 
	UniCompFees,   
	TwapGetter
{
	
	using SafeERC20 for IERC20;


	address constant public UNIV3_FACTORY = 0x1F98431c8aD98523631AE4a59f267346ea31F984;
	address public creator;
	address public immutable factory;
	address payable public immutable 	_WETH;
    address public immutable token0;         // underlying token0
    address public immutable token1;         // underlying token1
	address public protocol;			// where fee cuts to protocol 

	ICEth public immutable _CETH;
    IUniswapV3Pool public immutable pool;        // get by uni factory, token0, token1, feetier
    uint128 public immutable quoteamount;  		// for calc price, based on token0 decimal, ie: 1e18 for eth, 1e8 for wbtc
   
// 	uint8 public decimal0;
//    uint8 public decimal1;
    uint8 public uniPortion;       // uniswap portion ratio
    uint8 public compPortion;       // compound portion ratio
    uint8 public protocolFeeRate;		// 0 - 20% of profit
	uint8 public motivatorFeeRate;		// 0- 10% from profit for keeping system running by press buttons
	uint32 public twapDuration;        // oracle twap durantion
//    uint32 public threshold;	    // initial range

    int24 public tickSpacing;
    int24 public cLow;
    int24 public cHigh;
    int24 public maxTwapDeviation;     // for twap     

    mapping (address => address) public _CTOKEN;
    
    uint256 public vaultCap;	   	// 0: no cap
    uint256 public individualCap;	//  0 : no cap      

    bool private isEmergency = false;  // only canbe changed within emergency

    mapping (uint => address[] ) public motivator;		// who helped to triggering buttons e.g. motivator[1].push( address )  1=rebalance
	mapping (address => uint8) public Decimals;
    
	uint256 public lastCount;
	address[] public motivators;
	mapping(address => uint) motivatorCounter;

    constructor( UniCompParam memory params) {
		factory = params.VaultFactory; 

		protocol =  params.Protocol;
		creator = params.Creator;
		pool = IUniswapV3Pool(IUniswapV3Factory(UNIV3_FACTORY).getPool( params.Token0, params.Token1, params.FeeTier)); 

        // token0 & token1 sort could be changed by the uni v3 pool 
        token0 = pool.token0();  
        token1 = pool.token1();
        
        if ( token0 == params.Token1 ) { 	
        	// tokens order changed , the ctokens order must change accordingly.
			_CTOKEN[token0] = params.CToken1;	
			_CTOKEN[token1] = params.CToken0; 
        } else {
			_CTOKEN[token0] = params.CToken0;	
			_CTOKEN[token1] = params.CToken1; 
        }

		_WETH = payable (params.WETH);	
        _CETH = ICEth(params.CETH);	
		
		tickSpacing = pool.tickSpacing();
		uniPortion =  params.UniPortionRate;
        compPortion =  params.CompPortionRate;
        
        protocolFeeRate = params.ProtocolFeeRate;

        twapDuration = params.TwapDuration;   // mainnet should be 2
        maxTwapDeviation = params.MaxTwapDeviation;
        motivatorFeeRate = params.MotivatorFeeRate;
		Decimals[token0] = params.Token0Decimals;        
		Decimals[token1] = params.Token1Decimals;
		quoteamount = uint128(10 ** Decimals[token0]);

    }
    
    event MintUniV3Liquidity(int24 indexed newLow, int24 indexed newHigh, uint128 indexed liquidity, uint256 minied0, uint256 minted1 );
    event BurnUniV3Liquidity(int24 indexed cLow, int24 indexed cHigh, uint128 indexed liquidity);
    event Rebalance(address indexed,uint256 u0,  uint256 u1, uint256 c0,  uint256 c1);
    event AllocFees(address indexed, uint256 uFees0, uint256 uFees1,uint256 lFees0,uint256 lFees1);
    event MyLog(string, uint256);
    
    
    /// check status == 1
    modifier onlyActive {
        require (IVaultFactory(factory).getStat( address(this)) == 1 , 'not active');
        _;
    }
    
    modifier onlyCreator {
        require (msg.sender == creator ||  msg.sender == IVaultFactory(factory).getAdmin()  ,'E5');  //only creator
        _;
    }


    modifier onlyAdmin {
        require(IVaultFactory(factory).getAdmin() == msg.sender, 'only admin');
        _;
    }

    modifier onlyVault {

		require( IVaultFactory(factory).getPair0(address(this)) == msg.sender, "not vault");

    	//require(IVaultFactory(factory).onlyPair(address(this), msg.sender), "not vault");
        //require (msg.sender == vault, 'not vault');
        _;
    }    
    
    
	receive() external payable {}
	fallback() external payable {}


    	
 /// @dev Callback for Uniswap V3 pool
    function uniswapV3MintCallback(
        uint256 amount0,
        uint256 amount1,
        bytes calldata data
    ) external  {
        require(msg.sender == address(pool),"PS");

        IERC20(token0).safeTransfer(msg.sender, amount0);
        IERC20(token1).safeTransfer(msg.sender, amount1);

 		assert(data.length == 0);
    }

    /// @dev Callback for Uniswap V3 pool
    function uniswapV3SwapCallback(
        int256 amount0Delta,
        int256 amount1Delta,
        bytes calldata data
    ) external  {
        require(msg.sender == address(pool),"PS2");
        
        IERC20(token0).safeTransfer(msg.sender, uint256(amount0Delta));
        IERC20(token1).safeTransfer(msg.sender, uint256(amount1Delta));

        assert(data.length == 0);
    }
	
	
	
	///@notice send funds back to vault
	function callFunds() external onlyVault {
		
		alloc();
		
		uint256 a0 =IERC20(token0).balanceOf(address(this));
		uint256 a1 =IERC20(token1).balanceOf(address(this));
		
        if(a0>0) IERC20(token0).safeTransfer(myVault(), a0);
        if(a1>0) IERC20(token1).safeTransfer(myVault(), a1);
	}

	function myVault() internal view returns(address){
       return IVaultFactory(factory).getPair0(address(this));
    }

/*	struct RebalanceParam {
		uint8 uniPortion,
		uint8 lendPortion,
		uint8 lendingPool,
		int24 baseLow,
		int24 baseHigh,
		int24 limitLow,
		int24 limitHigh,
		int24 hedging, 
		int24 options
	}
	
	
	function rebalanceReady() public view {
		
		
	}
*/	
	function rebalance(
		int24 newLow,
        int24 newHigh
		) external onlyActive
	{
		
		alloc();
		allocFees();

		//process pending withdrawals
		uint _wdslength = IViaVault(myVault()).wdsLen();
		if (_wdslength>0) {
			//send all funds back to vault
			IERC20(token0).safeTransfer(myVault(), getBalance(token0));
		    IERC20(token1).safeTransfer(myVault(), getBalance(token1));		
			IViaVault(myVault()).withdrawLoop();
		}
		
		// move funds back from vault
		if ( IERC20(token0).balanceOf(myVault()) > 0 ||  IERC20(token1).balanceOf(myVault()) > 0 ) {
			IViaVault(myVault()).moveFunds();
		}

        _rebalance(newLow, newHigh); 
        
	}
	

	// make sure all position has been removed before doing rebalance. 
	// when newLow and newHigh is 0, calc new range with current cLow and cHigh
	function _rebalance(int24 newLow, int24 newHigh ) 
		internal 
	{

		(	,int24 tick, , , , , ) 	= pool.slot0();
		
    	if (newLow==0 && newHigh==0) {

			if (cHigh == 0 && cLow ==0) {
                return;  // cannot do rebalance if cLow and cHigh is 0
            }
			int24 hRange = ( cHigh - cLow ) / 2;
			newHigh = (tick + hRange) / tickSpacing * tickSpacing;
			newLow  = (tick - hRange) / tickSpacing * tickSpacing;
		}

  		_validRange(newLow, newHigh, tick);  // passed 1200 , 2100, 18382

		uint256 currentBalance0 = getBalance(token0);
		uint256 currentBalance1 = getBalance(token1);

		// since all assets should be now available, if balance is 0, no need to continue
		if (currentBalance0 == 0 && currentBalance1 == 0 ) return;

		uint256 p = getPrice();
        //#calculate the amounts to supply into uniswap
        (uint256 u0, uint256 u1) = calcUniPortionAmounts(currentBalance0, currentBalance1, p);
		
        if ( uniPortion > 0 ) {
			(uint256 minted0, uint256 minted1) = mintUniV3( newLow, newHigh, u0, u1 );
            (cLow, cHigh) = (newLow, newHigh);    

			// update the current balances should there be minted amounts on uniswap 
			currentBalance0 -= minted0;
			currentBalance1 -= minted1;
		}
	
		if (compPortion > 0 ) {
			compIn0 = currentBalance0 * compPortion / 100 ;
			compIn1 = currentBalance1 * compPortion / 100 ;

	        mintCompound(token0,compIn0);
	        mintCompound(token1,compIn1);
			
		}	
		
		// todo if ( hedging > 0) {}
		// todo if (aavePortion> 0 ) {}

		
		emit Rebalance(address(this), u0,u1,compIn0,compIn1);

	}
	
	function getBalance(address _token) internal view returns(uint256){
		return (IERC20(_token).balanceOf(address(this)));
	}
	
    function getPrice() public view returns(uint256) {
    	
    	// Check price has not moved a lot recently. This mitigates price
        // manipulation during rebalance and also prevents placing orders
        // when it's too volatile.

        int24 tick = getTwap(address(pool), twapDuration);

    	return( getQuoteAtTick(tick, uint128(quoteamount), token0, token1) );

    }
    
/*    
    function getTickPrice() public view returns(uint256) {
    	
        (uint160 sqrtPriceX96,,,,,,) =  pool.slot0();
        uint256 p =  ( uint256(sqrtPriceX96) * uint256(sqrtPriceX96) >> (96 * 2) ) * quoteamount;
        return p;
    }
*/
    
    function setMaxTwapDeviation(int24 _maxTwapDeviation) external onlyCreator {
         require(_maxTwapDeviation > 0, "maxTwapDeviation");
        maxTwapDeviation = _maxTwapDeviation;
    }

    function setTwapDuration(uint32 _twapDuration) external onlyCreator {
        twapDuration = _twapDuration;
    }
    
    function setPortionRatio(uint8 uni, uint8 comp) external onlyCreator {
    	require(uni <=100 && comp <=100,'100');
		(uniPortion, compPortion ) =  ( uni, comp );
    }

	///@notice set protocol address to collect fees
	function setProtocol(address _protocol) external onlyAdmin {
        protocol = _protocol;
    }
   

 	function mintUniV3(
        int24 newLow,
        int24 newHigh,
        uint256 amount0,
        uint256 amount1
    ) internal returns(uint256 minied0, uint256 minted1) { 
	  
        (uint160 sqrtRatioX96, , , , , , ) = pool.slot0();
        
        uint128 liquidity =
            LiquidityAmounts.getLiquidityForAmounts(
                sqrtRatioX96,
                TickMath.getSqrtRatioAtTick(newLow),
                TickMath.getSqrtRatioAtTick(newHigh),
                amount0,
                amount1
            );
            
        (minied0, minted1 ) = pool.mint(address(this), newLow, newHigh, liquidity, "");
        emit MintUniV3Liquidity(newLow, newHigh, liquidity,minied0, minted1);
	}

	

	
	function getUniAmounts(int24 tickLower, int24 tickUpper)
        public
        view
        returns (uint256 amount0, uint256 amount1)
    {
		if (cLow == 0 && cHigh == 0 ) { return (0,0); }

        (uint128 liquidity, , , uint128 tokensOwed0, uint128 tokensOwed1) =
            _position(tickLower, tickUpper);
        (amount0, amount1) = _amountsForLiquidity(tickLower, tickUpper, liquidity);
        amount0 = amount0 + uint256(tokensOwed0);
        amount1 = amount1 + uint256(tokensOwed1);
    }
    
    
    function redeemCEth(uint256 amount, bool redeemType) internal {
		uint r;
        if (redeemType == true) {
            // amount=cETH的数量
            r = _CETH.redeem(amount);
			if (r==0) _wrap(address(this).balance);
        } else {
            // amount=要赎回的ETH的数量
            r = _CETH.redeemUnderlying(amount);
			if (r==0) _wrap(amount);
        }

		
        if (r != 0) {	//something wrong. Ctoken may be stuck in Comp
			if (!isEmergency) {
				revert("Ceth not 0");  
			} else {
				// emergency 
			}
        }	
	}

	function redeemCErc20(address underlying, uint256 amount, bool redeemType) internal {
	
		
		if (underlying == _WETH) {
			redeemCEth(amount, redeemType);
			return;
		}
		
        uint256 r;
        if (redeemType == true) {
            // amount=归还cERC20的数量
            r = ICErc20(_CTOKEN[underlying]).redeem(amount);
        } else {
            // amount=要赎回的ERC20的数量
            r = ICErc20(_CTOKEN[underlying]).redeemUnderlying(amount);
        }
        
        if (r != 0) {	//something wrong. Ctoken may be stuck in Comp
			if (!isEmergency) {
				revert('Ct');  
			} else {
				// emergency 
			}
        }	
        
	}	

    /// remove all positions from protocols, if any failed, an emergency maybe required.
    function removePositions() internal  {
	    removeUniswap();
		removeLending();
	}
	
	function removeUniswap() internal {
       
    	if (cLow == 0 && cHigh ==0)  return;
		
       (uint128 liquidity, , , , ) = _position(cLow, cHigh);   // get liquidity by current ticklow, tickhigh
       
 		(uint256 burned0, uint256 burned1, uint256 collect0, uint256 collect1) = _burnAndCollectUnis(cLow, cHigh, liquidity);
    
    	if (liquidity > 0) {
	        uFees0 = collect0 - burned0;
	       	uFees1 = collect1 - burned1;
       }
	   	(liquidity, , , , ) = _position(cLow, cHigh);	// should be 0 otherwise, there is problem
	}

        	
	/// remove lending position; gether fees
	function removeLending() internal {
	
		uint256 b0 = getBalance(token0);
		uint256 b1 = getBalance(token1);

        (uint256 c0, uint256 c1 ) = getCAmounts();
        if(c0>0) redeemCErc20(token0, c0, true);
        if(c1>0) redeemCErc20(token1, c1, true);
        
        compOut0 = getBalance(token0);
        compOut1 = getBalance(token1);
        compOut0 = (compOut0 > b0 )? compOut0 - b0 : compIn0;
        compOut1 = (compOut1 > b1 )? compOut1 - b1 : compIn1;
        
	}
	
	function _position(int24 tickLower, int24 tickUpper)
        internal
        view
        returns ( uint128, uint256, uint256,  uint128, uint128)
    {
        bytes32 positionKey = PositionKey.compute(address(this), tickLower, tickUpper);
        return pool.positions(positionKey);
    }
    

 function _burnAndCollectUnis(
        int24 tickLow,
        int24 tickHigh,
        uint128 liquidity 
        )
        internal
        returns ( uint256 burned0, uint256 burned1, uint256 collect0,uint256 collect1)
    {
		
        if ( liquidity > 0) {
        	( burned0, burned1) =  pool.burn(tickLow, tickHigh, liquidity) ;
        } 
        
         // Collect all owed tokens including earned fees
        (collect0,  collect1) =
            pool.collect(
                address(this),
                tickLow,
                tickHigh,
                type(uint128).max,
                type(uint128).max
            );
    }   

	///@notice calculate best portion to put into uniswap
	function calcUniPortionAmounts(uint256 total0, uint256 total1, uint256 price) internal view returns(uint256 amount0, uint256 amount1){

		//uint256 p = getPrice();
		//uint256 p = getTickPrice();
		 uint256  p = price;
		
		uint256 total0In1 = total0 * p;   // to have same quantity of token0 as token1 upon current price
		
		if (total0In1 > total1)  {    
			// amount of token0 > amount of token 1, then use token0 to calculate the larger part of supply  
			amount0 = total0 * uniPortion / 100;
			amount1 = Math.min( total1 * uniPortion /100, amount0 * p );
		} else {
			// amount of token1 > amount of token 0, then use token1 to calculate the larger part of supply  
			amount1 = total1 * uniPortion / 100;	
			amount0 = Math.min( total0 * uniPortion/100, amount1/p );
		}
		
	}


	function mintCompound(address _underlying, uint256 amount) internal {
        if (_underlying == _WETH ) {
   	        _unwrap(amount);  
			_CETH.mint{gas:250000,value:amount}();
        } else {

			IERC20(_underlying).approve(_CTOKEN[_underlying], amount);
			uint mintResult = ICErc20(_CTOKEN[_underlying]).mint(amount);
			if (mintResult != 0) {
				emit MyLog("mintResult is not 0, there was an error", mintResult);
			}
        }
	}

	function _wrap(uint256 amount) internal {
	    if (amount > 0) {
	        IWETH9(_WETH).deposit{value:amount}();
	    }   
	}

	function _unwrap(uint256 amount) internal {
	    if (amount > 0) {
	        IWETH9(_WETH).withdraw(amount);
	    }
	}   
  

    function _amountsForLiquidity(
        int24 tickLow,
        int24 tickHigh,
        uint128 liquidity
    ) internal view returns (uint256, uint256) {
        (uint160 sqrtRatioX96, , , , , , ) = pool.slot0();
        return
            LiquidityAmounts.getAmountsForLiquidity(
                sqrtRatioX96,
                TickMath.getSqrtRatioAtTick(tickLow),
                TickMath.getSqrtRatioAtTick(tickHigh),
                liquidity
            );
    }
    
    function allocFees() internal {
    	
		lFees0 = compOut0-compIn0;
		lFees1 = compOut1-compIn1;
		
		uint256 allfees0 = uFees0 +lFees0;
		uint256 allfees1 = uFees1 +lFees1;

		if ( protocolFeeRate > 0 ) {
			uint256 feecut0 = allfees0 * protocolFeeRate /100;
			uint256 feecut1 = allfees1 * protocolFeeRate /100;
			address team = IVaultFactory(factory).getTeam();
			IViaVault(myVault()).mintFees( team, feecut0 , feecut1);
			
			if ( motivatorFeeRate > 0 ) {
				uint256 dfeecut0 = feecut0 * motivatorFeeRate /100;
				uint256 dfeecut1 = feecut1 * motivatorFeeRate /100;
				IViaVault(myVault()).mintFees( msg.sender, dfeecut0 , dfeecut1);
				//allocMotivatorFees( feecut0, feecut1);
			}
		}
		
		emit AllocFees(address(this), uFees0,uFees1,lFees0,lFees1);
		
		compOut0 = 0;
		compIn0 = 0;
		compOut1=0;
		compIn1=0;
		uFees0 =0;
		lFees0 =0;
		uFees1 =0;
		lFees1 =0;
		
    }

    function alloc() internal returns ( bool ){
    	
		removePositions();		// get all assets back to vault
		return(true);  // fees are allocated
    }


	function allocMotivatorFees(uint256 feecut0, uint256 feecut1) internal {
		while (motivators.length > 0) {
			uint i = motivators.length-1;
			uint icount = motivatorCounter[motivators[i]];
			if (icount > 0) {
				IViaVault(myVault()).mintFees( motivators[i], feecut0 * icount / lastCount, feecut1 * icount / lastCount);
				motivatorCounter[motivators[i]] = 0;
				motivators.pop();
			}
		}
		lastCount = 0;
	}
	function countMotivator(address addr) internal {
		if (motivatorCounter[addr] > 0) {
			motivatorCounter[addr]++;
		} else {
			motivatorCounter[addr] = 1;
			motivators.push(addr);
		}
		lastCount++;
	}

	function getCAmounts() public view returns (uint256 amount0, uint256 amount1) {
		amount0 = ICErc20(_CTOKEN[token0]).balanceOf( address(this) );
		amount1 = ICErc20(_CTOKEN[token1]).balanceOf( address(this) );
	}

	function getCtokenUnderlyingAmounts() internal returns (uint256 amount0, uint256 amount1) {
		amount0 = ICErc20(_CTOKEN[token0]).balanceOfUnderlying( address(this) );
		amount1 = ICErc20(_CTOKEN[token1]).balanceOfUnderlying( address(this) );
	}
	
	function getTotalAmounts() public view returns(uint256 , uint256) {
		
		uint256 b0 = IERC20(token0).balanceOf(address(this));
		uint256 b1 = IERC20(token1).balanceOf(address(this));

		(uint256 a0, uint256 a1) = getUniAmounts(cLow, cHigh);
		(uint256 c0, uint256 c1 ) = getAmountsInComp();
		
		return (a0+b0+c0, a1+b1+c1);
	}
	
 	function getAmountsInComp() public view returns(uint256 , uint256 ){

    	(uint256 cAmount0, uint256 cAmount1) = getCAmounts();
		
		if (cAmount0 == 0 && cAmount1 == 1 ) {return (0,0);	}

    	ICErc20 ctoken0 = ICErc20(_CTOKEN[token0]);
		ICErc20 ctoken1 = ICErc20(_CTOKEN[token1]);
		
		uint256 exchangeRate0 = ctoken0.exchangeRateStored();  // (1* 10 ** 18 + Decimals[token0] -8 )));   // the sampled calculation get a wrong result
		uint256 exchangeRate1 = ctoken1.exchangeRateStored();  // (1* 10 ** 18 + Decimals[token1] -8 ))); 
				
        uint256 amount0 = cAmount0 * exchangeRate0 / 1e18;
        uint256 amount1 = cAmount1 * exchangeRate1 / 1e18;

		return(amount0,amount1);
    }	
	
  	function _validRange(int24 _lower, int24 _upper, int24 _tick) internal view {
        require(_lower < _upper, "V1");
        require(_lower < _tick , "V2");
        require(_upper > _tick , "V3");
        require(_lower >= TickMath.MIN_TICK, "V4");
        require(_upper <= TickMath.MAX_TICK, "V5");
        require(_lower % tickSpacing == 0, "V6");
        require(_upper % tickSpacing == 0, "V7");
    }
    
    
}