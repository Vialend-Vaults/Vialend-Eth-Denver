// SPDX-License-Identifier: MIT

pragma solidity >=0.8.8 <0.9.0;

import "@openzeppelin/contracts/token/ERC20/SafeERC20.sol";
import "@openzeppelin/contracts/math/Math.sol";
import "@openzeppelin/contracts/utils/ReentrancyGuard.sol";
import "@uniswap/v3-core/contracts/libraries/FullMath.sol";
import "@uniswap/v3-core/contracts/interfaces/IUniswapV3Factory.sol";
import "@uniswap/v3-core/contracts/interfaces/IUniswapV3Pool.sol";
import "@uniswap/v3-periphery/contracts/libraries/LiquidityAmounts.sol";
import "@uniswap/v3-periphery/contracts/libraries/PositionKey.sol";
import "@uniswap/v3-core/contracts/libraries/TickMath.sol";
import "@uniswap/v3-core/contracts/interfaces/callback/IUniswapV3MintCallback.sol";
import "@uniswap/v3-core/contracts/interfaces/callback/IUniswapV3SwapCallback.sol";
import "@uniswap/v3-periphery/contracts/interfaces/ISwapRouter.sol";
import "@uniswap/swap-router-contracts/contracts/interfaces/ISwapRouter02.sol";
import "@uniswap/v3-periphery/contracts/libraries/TransferHelper.sol";
import  { IWETH9 , IController, IPriceOracleGetter}  from "./interfaces/IViaProtocols.sol";

import "./UniCompFees.sol";
import "./TwapGetter.sol";
import "./AaveHelper.sol";
//import "./libraries/UniCompHelper.sol"; 
import "./libraries/Debugger.sol"; 

import "./interfaces/IViaVault.sol";
import "./interfaces/IVaultFactory.sol";
import "./interfaces/AggregatorV3Interface.sol";

/*
Error code:
'OLD' getTwap error. rebalance cannot proceed. solution: setTwapDuratio = 0
'E5' not creator or admin
'E6' amount <= 0 
'E7' invalid tokenIn
'A0' Aave usdc amount is 0
'L0' require(leverage > 0 && leverage < 6, );
*/

/// @author  ViaLend
/// @title   strategy Uni + oSqth  + HedgeVault ( AAve Flashloan)
/// @notice  A Smart Contract that helps liquidity providers managing their funds on Uniswap V3 .

struct UniSqueethParam {
	address Uni3Factory;
	address VaultFactory;
	address HedgeVault;
	address Protocol; 			// protocol fee collector address 
	address Creator;			// who created this strategy when used as template
	address Token0;
	address Token1;
	address WETH;
	address OSqth;
	address ChainLinkProxy;
	uint8 Token0Decimals;
	uint8 Token1Decimals;
	uint8 ProtocolFeeRate;  
	uint8 MotivatorFeeRate; 
	uint24 FeeTier;
	int24 MaxTwapDeviation;
	uint32 TwapDuration;	// initial twap durantion

}

struct RebalanceParam {
	int24 newLow;
    int24 newHigh;
    uint32 sqthPercent; 
	uint32 uniPortionRate;
	uint32 sqthPortionRate;
	uint32 shortPortionRate; 
	uint32 shortLever;   
}

contract VaultStrategy2 
	is 
	ReentrancyGuard , 
	UniCompFees,   
	TwapGetter,
	AaveHelper
{
	
	using SafeERC20 for IERC20;

	AggregatorV3Interface internal priceFeed;


	address constant public UNIV3_FACTORY = 0x1F98431c8aD98523631AE4a59f267346ea31F984;
	address constant public SWAP_ROUTER = 0x68b3465833fb72A70ecDF485E0e4C7bD8665Fc45;   //ISwapRouter02

 
	address public aavePoolProvider = 0xA55125A90d75a95EC00130E8E8C197dB5641Eb19; // rinkey 
	address public shortCallback = 0xa55eC76678548b4Eff64A09fA693cC18F7c6fc8c;	
    address public unwindCallback = 0x5fC0Ed674B2dA337Cf9BCfd5f17DA1F52Eb2F661;    
    address public aaveUSDC = 0x5B8B635c2665791cf62fe429cB149EaB42A3cEd8; // rinkey 
    address public aaveETH = 0x98a5F1520f7F7fb1e83Fe3398f9aBd151f8C65ed ; // rinkey 
	address public aStableDebt = 0x1E2192C406c6a53056E923A1aCe9e05b0090a531;  // rinkeby aStableDebt 
	address public aTokenUSDC = 0xD624c05a873B9906e5F1afD9c5d6B2dC625d36c3; // aToken USDC rinkeby
	
    	
	address public creator;
	address public immutable factory;
	address payable public immutable 	_WETH;
    address public immutable token0;         // underlying token0
    address public immutable token1;         // underlying token1
	address public protocol;			// where fee cuts to protocol 
	address public immutable oSqth;
	address public oSqthEthPool;
	address public _USDC;


    IUniswapV3Pool public immutable pool;        // get by uni factory, token0, token1, feetier

	IUniswapV3Pool internal swapPool;

   
// 	uint8 public decimal0;
//    uint8 public decimal1;
    uint8 public protocolFeeRate;		// 0 - 20% of profit
	uint8 public motivatorFeeRate;		// 0- 10% from profit for keeping system running by press buttons
	uint32 public twapDuration;        // oracle twap durantion
//    uint32 public threshold;	    // initial range

    int24 public tickSpacing;
    int24 public cLow;
    int24 public cHigh;
    int24 public maxTwapDeviation;     // for twap     


    bool private isEmergency = false;  // only canbe changed within emergency

    mapping (uint => address[] ) public motivator;		// who helped to triggering buttons e.g. motivator[1].push( address )  1=rebalance
	mapping (address => uint8) public Decimals;
    
	uint256 public lastCount;
	address[] public motivators;
	mapping(address => uint) motivatorCounter;
	
	uint256 public baseTokenId; 

    constructor( UniSqueethParam memory params) {
		factory = params.VaultFactory; 
		
        priceFeed = AggregatorV3Interface(params.ChainLinkProxy); 

		protocol =  params.Protocol;
		creator = params.Creator;
		pool = IUniswapV3Pool(IUniswapV3Factory(UNIV3_FACTORY).getPool( params.Token0, params.Token1, params.FeeTier)); 
		
		// compare swap methods (direct pool , router , 3rd protocols 1inch/para swap/ )
        oSqth =  params.OSqth;

        // token0 & token1 sort could be changed by the uni v3 pool 
        token0 = pool.token0();  
        token1 = pool.token1();
        

		_WETH = payable (params.WETH);	

        oSqthEthPool = IUniswapV3Factory(UNIV3_FACTORY).getPool(oSqth, _WETH, 3000); 
		
		//todo
		_USDC = token0; 
		
		baseTokenId = ( _WETH == token0 ) ? 0 : 1 ;
		
		tickSpacing = pool.tickSpacing();
        
        protocolFeeRate = params.ProtocolFeeRate;

        twapDuration = params.TwapDuration;   // mainnet should be 2
        maxTwapDeviation = params.MaxTwapDeviation;
        motivatorFeeRate = params.MotivatorFeeRate;
		Decimals[token0] = params.Token0Decimals;        
		Decimals[token1] = params.Token1Decimals;
		Decimals[oSqth] = 18;

    }
    
    event MintUniV3Liquidity(int24 indexed newLow, int24 indexed newHigh, uint128 indexed liquidity, uint256 minied0, uint256 minted1 );
    event BurnUniV3Liquidity(int24 indexed cLow, int24 indexed cHigh, uint128 indexed liquidity);
    event Rebalance(address indexed,uint256 u0,  uint256 u1, uint256 c0,  uint256 c1);
    event AllocFees(address indexed, uint256 Fees0, uint256 Fees1);
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
		
        //(address _token0, address _token1) = abi.decode(data, (address,address));
        address _caller = abi.decode(data, (address));

		// make sure the callback is from the predefined pool 
        require(_caller == address(this) && msg.sender == address(swapPool), 'PS2');    
        
        
		if (amount0Delta > 0) {
	        IERC20(swapPool.token0()).safeTransfer(msg.sender, uint256(amount0Delta));
        } else {
            assert(amount1Delta > 0);
	        IERC20(swapPool.token1()).safeTransfer(msg.sender, uint256(amount1Delta));
        }        

    }

	function swapDirectPool(
		address _tokenIn, 
		address _tokenOut, 
		uint24 fee, 
		uint160 sqrtPriceLimitX96Shift, // >=1
		uint256 amount
	) public nonReentrant {  
		

		if (amount ==0) return;

		//require (_tokenIn == token0 || _tokenIn == token1 || _tokenIn == oSqth || _tokenIn == aStableDebt, 'E7');
		
		swapPool = IUniswapV3Pool(IUniswapV3Factory(UNIV3_FACTORY).getPool( _tokenIn, _tokenOut,fee)); 
		address _token0 = swapPool.token0();  
//        address _token1 = swapPool.token1();  

		bool zeroForOne;
		uint160 sqrtPriceLimitX96;

		if (_token0 == _tokenIn) {
			//swapExact0For1  
			zeroForOne = true;		
			sqrtPriceLimitX96 = TickMath.MIN_SQRT_RATIO + sqrtPriceLimitX96Shift;
		} else {
			// swapExact1For0  
			zeroForOne = false;	
			sqrtPriceLimitX96 = TickMath.MAX_SQRT_RATIO - sqrtPriceLimitX96Shift;
		}

	    if (amount != 0) {
            swapPool.swap(
                address(this),
                zeroForOne,
                int256(amount),
                sqrtPriceLimitX96,
                abi.encode(address(this))   // to authenticate it only strategy from the callback.
            );
        }
	}
        
	//## debugging function, will be remvoed on product
	function swapTest (
		address	_tokenIn,
        address _tokenOut,
        uint24 _fee,
        uint256 _amountIn
    ) external onlyAdmin nonReentrant returns (uint256)  {
	     return swap(_tokenIn, _tokenOut, _fee, _amountIn );
    }
    
    

    function swap(
        address	_tokenIn,
        address _tokenOut,
        uint24 _fee,
        uint256 _amountIn
    ) internal returns (uint256) {
    	
        ISwapRouter02 router = ISwapRouter02(SWAP_ROUTER);
        
        // approve
        TransferHelper.safeApprove(_tokenIn, SWAP_ROUTER, _amountIn);
        
		ISwapRouter02.ExactInputSingleParams memory params =  IV3SwapRouter.ExactInputSingleParams(
                _tokenIn,
                _tokenOut,
                _fee,
                address(this),		//recipient
	            //block.timestamp + 20,	//deadline  **  this attribute has been removed from V3 swaprouter
                _amountIn,
                0,   				//amountOutMinimum:  TODO: Does this value need to be set.
                0 					// sqrtPriceLimitX96:
            );
            
        // do swap and return
        return router.exactInputSingle(params);
    } 
	

	function getSqueethNorm(address _controller  ) external view returns (uint256){
		
		uint256 norm =  IController(_controller).getExpectedNormalizationFactor(); // 813118587869342996; 
		return ( norm );
	}
	
	///@notice called by viaVault. send funds back to visaVault
	function callFunds() external onlyVault {
		
		uint256 a0 =IERC20(token0).balanceOf(address(this));
		uint256 a1 =IERC20(token1).balanceOf(address(this));
		uint256 a2 =IERC20(oSqth).balanceOf(address(this)); 
		uint256 a3 =IERC20(aTokenUSDC).balanceOf(address(this)); 
		
        if(a0>0) IERC20(token0).safeTransfer(myVault(), a0);
        if(a1>0) IERC20(token1).safeTransfer(myVault(), a1);
        if(a2>0) IERC20(oSqth).safeTransfer(myVault(), a2);

   		// note: have to repay loan or this would fail.
		//if(a3>0) IERC20(aTokenUSDC).safeTransfer(myVault(), a3);
	}

	function myVault() internal view returns(address){
       return IVaultFactory(factory).getPair0(address(this));
    }

/*	

	** off chain calculation:
	1. check uni portion : amount0, amount1
	2. check vault's squeeth portion: oSquth Amount
	3. calculate short hedging amount with eth price. 
		( e.g. 
				@eth price = $3000
				uni portion: $700k usdc :$700k eth (333 eth) = $1.4M usdc

				squeeth portion (5.5x uni eth portion -- based on carlos data):  566 eth (333 *1.7) = 566*3000 = $1.7M usdc 
				 
				.85x * uni portion = .85 * 333eth   = 283 eth swap to buy osth  -> funding fee paid -> 111% apy
				
				  => does it fully cover uni portion 700k  impermant loss   ???
				
				.85 * uni portion = .85 * 333eth   = 283 eth -> sqth/eth pool -> earned fees 0.3% based on volume unfixed fees -> earned ~60% apy 
				
				net loss apy = 111-60 = -51% 
				
				
				for hedging, suppose the borrowing power is 75%, the amount needs to short is 333+566=899 eth. the collateral =  899 * 1.25 = 1123 eth * 3000 = $3.37M usdc 
				flash loan 3.37M and borrow 899 eth. sold 899 eth for 2.7M usdc. the capital for repay flashloan = 3.37M-2.7M = $670k usdc
				
				the total capital assets required in vault: 
					$700k usdc + (899eth)*3000 + $670k usdc = 700k + 2.7M + $670k = $3.4M usdc 
				
				@rebalance:
				@eth price fell to $2000, uni portion is out of range , all usdc is converted to eth. 
				uni portion before rebalance: $0 usdc: 341 eth +   333 eth = 675 eth  (roughly calc) = 1.35M   , il = 1.4m-1.35m= 50k usdc
				squeeth portion value: (roughly pnl = -40%)  = 566*3000*60% = $1M usdc   (pnl = 1M - 1.7M = -$700k in usdc)
				hedge portion value: eth price down 33% , profit = 899 $2.7M* 33% = $891k usdc  (sold 2.7M at 3000, unrealized pl = 2.7M*33%)
				
				net pnl = 891k - 700k (squeeth loss) - 50k(uni loss) = $141k  
				
				rebalance position status:
					uni: 0 usdc, 675 eth
					squeeth:  ~= $1M usdc  (before $1.7M)
					hedge: (payoff debt  899eth by flash loan -> take out $3.37M usdc collateral -> buy 899 eth with 1.8m usdc -> 1.57M usdc left in hedge vault.
					
					total value: = 657 * 2000 + 1.7M + 1.57M = 4.5M

				Monitor: 
					1. uni position  ( or just uni position for now)
					2. squeeth position 
					3. hedge position
					
				rebalance Steps:
					1. remove hedging, 
						. flash loan X eth
						. pay debt X eth and redeem Y usdc collateral
						. swap Yx usdc for X eth 
						. repay flash loan X eth
						
						+usdc or -usdc
						
					2 calc uni portion
					3 calc squeth portion
					4 calc hedge portion
					
					5 allocate assets 
						squeeth - 
							sell or buy more squeeth (swap)
							get funds from hedge vault
						hedge - 
							send / receive funds
							flash loan, 
							borrow
							swap, 
							repay
						uni - 
							swap, 
							transfer, 
							mint position
							
	
				
				initial fees:
				1. flashloan 0.09%   * $3.37M = $3033
				2. borrowing apy 6% / 365 * 7 (7days) = 899 eth * 6% /365 * 7 = 1.03 eth ~= $3000 usdc
				3. swap fee (eth): 0.05% * 899 = 0.45 eth ~= 1348  
				4. swap fee (squeeth): 0.3% * 566 = 1.7 eth ~= 5094 ( one time )
				5. tx gas ~= 600 
				
				total init fees:  $7381 + $5094 ~= $12k
				
				rebalance fees:
				
				1. flashloan 0.09%   * $3.37M = $3033
				2. 
				2. borrowing apy 6% / 365 * 7 (7days) = 899 eth * 6% /365 * 7 = 1.03 eth ~= $3000 usdc
				3. swap fee (eth): 0.05% * 899 = 0.45 eth ~= 1348  
				4. swap fee (squeeth): 0.3% * 566 = 1.7 eth ~= 5094 ( one time )
				5. tx gas ~= 600 
				
				total init fees:  $7381 + $5094 ~= $12k
					
				
				Please note: this is just an example on how to calc the values in each portion. the sqth portion and pnl will need to be adjusted in this example to reflect a closer net pnl. 
				
		)
				
	
	
*/	

	function rebalance(RebalanceParam memory rebalParam) 
		external onlyActive
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
		
		// move all funds back to strat from vault
		if ( IERC20(token0).balanceOf(myVault()) > 0 ||  IERC20(token1).balanceOf(myVault()) > 0 ) {
			IViaVault(myVault()).moveFunds();
		}
		
        _rebalance(  rebalParam); 
        
	}
	
	function moveFunds() external onlyAdmin {
		IViaVault(myVault()).moveFunds();
	}
	
	// make sure all position has been removed before doing rebalance. 
	// when newLow and newHigh is 0, calc new range with current cLow and cHigh
   /*
	rebalance creteria :
    	new deposit
		value in uniswap + squeeth + aave  (when hedging and price in squeeth is not perfect)
		fees out
		fees in
		
		-> calculate the best situation to do the rebalance
		
		
    */

	function _rebalance( RebalanceParam memory rebalParam) 
		internal 
	{

		(	,int24 tick, , , , , ) 	= pool.slot0();

  		_validRange(rebalParam.newLow, rebalParam.newHigh, tick);  

		uint256 currentBalance0 = getBalance(token0);
		uint256 currentBalance1 = getBalance(token1);

		// since all assets should be now available, if balance is 0, no need to continue
		if (currentBalance0 == 0 && currentBalance1 == 0 ) return;
		
        //#calculate the amounts to supply into uniswap
        (uint256 u0, uint256 u1, uint256 sqthPortionEth, uint256 shortPortionEth, uint256 ethSwap, uint256 usdcSwap) 
        	= makePortionSize(
					rebalParam, 
					currentBalance0, 
					currentBalance1 );

		// mint uni portion 
        (cLow, cHigh) = (rebalParam.newLow, rebalParam.newHigh);    
		(uint256 minted0, uint256 minted1) = mintUniV3( cLow, cHigh, u0, u1);

		//require(false, Debugger.uint2str(u0));

		// squeeth portion; 
		currentBalance1 = getBalance(_WETH);  
		require(currentBalance1 > sqthPortionEth , 'C0');  // make sure there is enough balance of weth
		swapDirectPool(_WETH, oSqth, 3000, 1, sqthPortionEth);

		// short portion; call aave short 
		shortHelper( shortPortionEth, rebalParam.shortLever );

		
		emit Rebalance(address(this), minted0,minted1,sqthPortionEth, shortPortionEth);

	}
	
	/// do the aave v3 Shorting  
	function shortHelper(uint256 ausdcAmountInEth, uint32 leverage) internal {
		
		require( ausdcAmountInEth > 0, "AS0");
		
		require(leverage >= 1 && leverage <= 4, 'L0');

		// my weth -> eth unwrap
		_unwrap(_WETH, ausdcAmountInEth);
		
		// eth -> aaveETH wrap
		_wrap(aaveETH, ausdcAmountInEth);
		
		// swap aaveETH for aaUsdc (workaround for rinkeby only) 
		// todo: swap in previous step to save this swap
		swapDirectPool(aaveETH, aaveUSDC, 3000, 1, ausdcAmountInEth);

		uint256 amountInUsdc  = IERC20(aaveUSDC).balanceOf(address(this));
		
		require(amountInUsdc > 0, "A0");  // be sure there is enough funds

        TransferHelper.safeApprove(aaveUSDC, shortCallback, amountInUsdc + IERC20(aaveUSDC).allowance(address(this), shortCallback));
        AaveHelper.short(aavePoolProvider, shortCallback, aaveUSDC, aaveETH, amountInUsdc, amountInUsdc * 250563768  * leverage );
    }
	
	
	///@notice calculate best portion to put into uniswap
	/// todo: formula of parts in uniswap, squeeth, and hedging
	function makePortionSize(
		
			RebalanceParam memory rebalParam,
			uint256 usdcBalance, 
			uint256 ethBalance
		) 
			internal 
			returns(
				uint256 uniAmount0, 
				uint256 uniAmount1, 
				uint256 sqthAmount,
				uint256 shortAmount,
				uint256 ethSwap,
				uint256 usdcSwap
			)
	{

	/* count 
			default portion rate:
			26.65% uni   2665 / 10000
			46.75% sqth	4675 / 10000
			26.59% short  2659 / 10000

			uint256 uniPortionRate =  2665 ; 
			uint256 sqthPortionRate =  4675;
			uint256 shortPortionRate= 2659; 
	*/ 


		// todo: more efficiently, by assigning _USDC 
		uint256 ethPrice = getEthPriceFromPool(_USDC, 500);

		uint256 totalEth = ethBalance + usdcBalance * 1e18 / ethPrice ;
		uint256 uniPortionEth = totalEth * rebalParam.uniPortionRate / 10000;
	
		uniAmount1 = uniPortionEth / 2; 			// eth for uni
		uniAmount0 = uniAmount1 * ethPrice / 1e18; 	// usdc for uni 
		sqthAmount = totalEth  * rebalParam.sqthPortionRate / 10000  *  rebalParam.sqthPercent / 100 ;
		shortAmount = totalEth * rebalParam.shortPortionRate / 10000;

		if (usdcBalance < uniAmount0 ) {
			// not enough usdc, need to swap some eth for usdc
			ethSwap = uniAmount1 - (usdcBalance * 1e18 / ethPrice );
			require(ethSwap > 0, "C1");
			swapDirectPool(_WETH, _USDC, 500, 1, ethSwap);

		}  else {
			//  if eth is enough, just swap all remaining usdc to eth. 
			usdcSwap = usdcBalance - uniAmount0;
			require(usdcSwap > 0, "C2");
			swapDirectPool(_USDC, _WETH, 500, 1, usdcSwap);
		}

	}
	
	function getEthPriceFromPool(address _token, uint24 _fee) public view returns(uint256 price) {
		
		// uint256 ethPrice = getPrice(); 
		
		address _ethPool = IUniswapV3Factory(UNIV3_FACTORY).getPool( _token, _WETH, _fee);
		
    	// old method: uniswap Twap oracle 
		// new method: use chainlink price feed 

        //int24 twaptick =  getTwap(_ethPool, _twapDuration); // twap on testnet throw error 'OLD'  
		(	,int24 twaptick, , , , , ) 	= IUniswapV3Pool(_ethPool).slot0();
		
        address _token0 = IUniswapV3Pool(_ethPool).token0();
        address _token1 = IUniswapV3Pool(_ethPool).token1();
        
       	price =  (_token0 == _WETH)? getQuoteAtTick(twaptick, uint128(1e18), _token0, _token1) : getQuoteAtTick(twaptick, uint128(1e18), _token1, _token0 );

	}
	
	function getPriceFromAavePool() public view returns(uint256 price) {

		address _aavePool = IUniswapV3Factory(UNIV3_FACTORY).getPool( aaveETH, aaveUSDC, 3000); //todo: hardcoded solver
		
        //int24 twaptick =  getTwap(_ethPool, _twapDuration); // twap on testnet throw error 'OLD'  
		(	,int24 twaptick, , , , , ) 	= IUniswapV3Pool(_aavePool).slot0();
		
        address _token0 = IUniswapV3Pool(_aavePool).token0();
        address _token1 = IUniswapV3Pool(_aavePool).token1();
        
       	price =  (_token0 == aaveETH)? getQuoteAtTick(twaptick, uint128(1e18), _token0, _token1) : getQuoteAtTick(twaptick, uint128(1e18), _token1, _token0 );

	}
	
	
	function getBalance(address _token) internal view returns(uint256){
		return (IERC20(_token).balanceOf(address(this)));

	}

	/// get oracle price from ChainLink 	
    function getPrice() public view returns(uint256) {
        //( uint80 roundID, int price, uint startedAt, uint timeStamp, uint80 answeredInRound) = priceFeed.latestRoundData();
        //( , int price,,,) = priceFeed.latestRoundData();
		//return uint256( price );
		
		return getEthPriceFromPool(_USDC,500);
    }

/*    
    function setChainlinkProxy(address newProxy) external onlyAdmin {
        priceFeed = AggregatorV3Interface(newProxy); 
    }
    
    ///@notice in case liquidity range get screwed. 
    function setHighLow(int24 _high, int24 _low ) external onlyAdmin {
        cHigh = _high;
        cLow = _low;
    }
    
    
    function setMaxTwapDeviation(int24 _maxTwapDeviation) external onlyCreator {
         require(_maxTwapDeviation > 0, "maxTwapDeviation");
        maxTwapDeviation = _maxTwapDeviation;
    }

    function setTwapDuration(uint32 _twapDuration) external onlyCreator {
        twapDuration = _twapDuration;
    }

	///@notice set protocol address to collect fees
	function setProtocol(address _protocol) external onlyAdmin {
        protocol = _protocol;
    }
   
*/    

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
        returns (uint256 amount0, uint256 amount1, uint256 myliquidity)
    {
		if (cLow == 0 && cHigh == 0 ) { return (0,0,0); }

        (uint128 liquidity, , , uint128 tokensOwed0, uint128 tokensOwed1) =
            _position(tickLower, tickUpper);
        
        if ( liquidity > 0 ) {
	        (amount0, amount1) = _amountsForLiquidity(tickLower, tickUpper, liquidity);
	        amount0 = amount0 + uint256(tokensOwed0);
	        amount1 = amount1 + uint256(tokensOwed1);
	        myliquidity = uint256(liquidity);
        }
    }

	function removeUniswap() public {
       
    	if (cLow == 0 && cHigh ==0)  return;
		
       (uint128 liquidity, , , , ) = _position(cLow, cHigh);   // get liquidity by current ticklow, tickhigh
       
 		(uint256 burned0, uint256 burned1, uint256 collect0, uint256 collect1) = _burnAndCollectUnis(cLow, cHigh, liquidity);
    
    	if (liquidity > 0) {
	        uFees0 = collect0 - burned0;
	       	uFees1 = collect1 - burned1;
       }
	   	(liquidity, , , , ) = _position(cLow, cHigh);	// should be 0 otherwise, there is problem
	}
    
	
	function removeSqueeth() public {
		// remove osqth portion
		uint256 sqthAmount = IERC20(oSqth).balanceOf(address(this));
		if ( sqthAmount == 0 ) return;
		swapDirectPool(oSqth, _WETH, 3000, 1, sqthAmount);
	}
	
	function removeShort() public {
        
		uint256 ethPerUsdc =  250563768; // 1e18 / getPriceFromAavePool();

		uint256 aAmount = IERC20(aTokenUSDC).balanceOf(address(this));
		
		if (aAmount == 0) return;

		AaveHelper.unwind(
        	aavePoolProvider,
			unwindCallback,
        	aaveUSDC,
        	aaveETH,
        	ethPerUsdc
    	);
		
        uint256 aUsdcAmount = IERC20(aaveUSDC).balanceOf(address(this));
		
		// swap aaUsdc for aaveETH  (workaround for rinkeby only) 
		swapDirectPool(aaveUSDC, aaveETH, 3000, 1, aUsdcAmount);

		uint256 aETHamount  = IERC20(aaveETH).balanceOf(address(this));

		//  aaveWeth -> eth unwrap
		_unwrap(aaveETH, aETHamount);
		
		// eth -> WETH wrap
		_wrap(_WETH, aETHamount);
    
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


	function _wrap(address _weth, uint256 amount) internal {
	    if (amount > 0) {
	        IWETH9(_weth).deposit{value:amount}();
	    }   
	}

	function _unwrap(address _weth, uint256 amount) internal {
	    if (amount > 0) {
	        IWETH9(_weth).withdraw(amount);
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
    
    /// make fee cut before compounding
    function allocFees() internal {
    	
		return;  //todo
		
		uint256 allfees0 = uFees0;
		uint256 allfees1 = uFees1;

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

	    emit AllocFees( msg.sender,  allfees0,  allfees1);

		uFees0 =0;
		uFees1 =0;

    }

    function alloc() public {
		removeUniswap();		// get all assets back to vault
		removeSqueeth() ;
		removeShort() ;
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

	function getTotalAmounts() public view returns(uint256 , uint256) {
		
		(uint256 u0, uint256 u1, ) = getUniAmounts(cLow, cHigh);

		uint256 usdc1 = IERC20(token0).balanceOf(address(this));
		uint256 usdc2 = IERC20(aTokenUSDC).balanceOf(address(this));

		uint256 eth1 = IERC20(token1).balanceOf(address(this));
		uint256 oe = IERC20(oSqth).balanceOf(address(this)) ;
		// convert squeeth to eth
		uint256 eth2 = oe * 1e18 / getEthPriceFromPool(oSqth,3000);
		
		uint256 v0 = IERC20(token0).balanceOf(myVault());
		uint256 v1 = IERC20(token1).balanceOf(myVault());
		
		return (u0+usdc1+usdc2, u1+eth1+eth2);
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