// SPDX-License-Identifier: MIT

pragma solidity >=0.8.8 <0.9.0;

import "@openzeppelin/contracts/token/ERC20/SafeERC20.sol";
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/utils/ReentrancyGuard.sol";
import "./interfaces/IVaultFactory.sol";
import "./Ownable.sol";
import "./interfaces/IStrategy.sol";

/// @author  ViaLend Team
/// @title   Vault for Strategy Uni + Compound
/// @notice  A Smart Contract that helps liquidity providers managing their funds on Uniswap V3 and Compound.

contract ViaVault is 
	ERC20,
	ReentrancyGuard,
	Ownable
{
    using SafeERC20 for IERC20;
    
    address public immutable factory;  
    address public immutable token0;   // token0
	address public immutable token1;   // token1
	address public immutable baseToken;   // token1 or token0 for different network

	address public oSqth;

	uint256 public vaultCap;
	uint256 public individualCap;

	struct Withdrawal {
		address recipient;
		uint8 percent;
	}
	
	Withdrawal[] public wds; 	
   	
	constructor(
    	address _factory,
        address _token0,
		address _token1,
		address _baseToken,
		address _oSqth,
		uint256 _vaultCap,
		uint256 _individualCap
    ) ERC20("ViaLend Uni Compound Token","VUC0") {
    	
		require(_baseToken == _token0 || _baseToken == _token1, "b0");

		factory = _factory;
        token0 = _token0;
        token1 = _token1;
		baseToken = _baseToken;
        oSqth =  _oSqth;
        vaultCap = _vaultCap;
		individualCap = _individualCap;


    }

	event Deposit(
        address indexed sender,
        uint256 shares,
        uint256 amount0,
        uint256 amount1
    );

    event Withdraw(
        address indexed sender,
        uint256 shares,
        uint256 amount0,
        uint256 amount1
     
    );
    
    event MintFees(address indexed, uint256 share, uint256 f0, uint256 f1);

	event PendingWithdraw(address to, bool pending);

 	modifier onlyAdmin {
        //require (msg.sender == admin, 'admin');  
        require(IVaultFactory(factory).getAdmin() == msg.sender, 'only admin');
        _;
    }
    
    modifier onlyStrategy {
		require( IVaultFactory(factory).getPair0(address(this)) == msg.sender, "not strat");
        //require (msg.sender == strategy, 'strat');
        _;
    }
    
    modifier onlyActive {
        require ( IVaultFactory(factory).getStat( address(this)) == 1, 'not active');
        _;
    }
    
    
    modifier onlyFactory {
    	require(msg.sender == factory, 'factory');
		_;
    }
    
    modifier onlyEmergency {
        require ( IVaultFactory(factory).getStat(address(this)) == 3 , 'emergency');
		_;
    }
    
	/// fallback function has been split into receive() and fallback(). It is a new change of the compiler.
	receive() external payable {}
	fallback() external payable {}


	function setVaultCap(uint256 newMax) external onlyAdmin {
			vaultCap = newMax;
	}

	function setIndividualCap(uint256 newMax) external onlyAdmin {
			individualCap = newMax;
	}

	function setSqueethAddress(address _oSqth) external onlyAdmin {
			oSqth =  _oSqth;
	}
    
    function EmergencyWithdraw() external onlyEmergency {
    	withdrawProc(msg.sender, 100);
    }

    
    ///@notice call funds back to vault
    function callFunds() external  {
    	IStrategy(myStrategy()).callFunds();
    }
    
		
 	function sweep( address _token, uint256 amount) external  onlyAdmin {
		require (msg.sender != address(0), "s0x");
		require(_token != token0 && _token != token1, 'tok');
        IERC20(_token).safeTransfer(msg.sender, amount);
    }	

    function deposit( uint256 amountIn0, uint256 amountIn1) external onlyActive {
    	
		require(amountIn0>0 || amountIn1>0, 'amt0');
		require(msg.sender != address(0), 'd0x');
       	(uint256 shares, uint256 amount0, uint256 amount1) = calcPositionShares(amountIn0, amountIn1);

		require (shares > 0, 'share<=0'); //affirm the shares
		
        _mint( msg.sender, shares);
        if (amount0 > 0) IERC20(token0).safeTransferFrom(msg.sender, address(this), amount0);
        if (amount1 > 0) IERC20(token1).safeTransferFrom(msg.sender, address(this), amount1);
        
        Assetholder[msg.sender].deposit0 += amount0;		// add or increase msg.sender's asset0
		Assetholder[msg.sender].deposit1 += amount1;		// add or increase msg.sender's asset1
		
        emit Deposit(msg.sender, shares, amount0, amount1);
        
        
        
    }     
    
    ///@notice get the strategy address from factory
    function myStrategy() internal view returns(address){
       return IVaultFactory(factory).getPair0(address(this));
    }
    
    function wdsLen() public view returns(uint256) {
    	return wds.length;
    }
    
	function withdrawLoop() public onlyStrategy {
		
		//consider max gas fee one block allowed
		
		while (wds.length > 0) {
			uint i = wds.length-1;
			withdrawProc(wds[i].recipient, wds[i].percent);
			wds.pop();
		}
		
	}
	
	///@notice this is pending withdraw to be put in a queue and its fund will be sent to user in next rebalance.
    function withdraw( uint8 percent ) external onlyActive
	{
		//10% reserve for withdraw
		
		bool pending = true;
		require ( percent <= 100 , 'p100');
		
		uint256 shares = balanceOf(msg.sender) * percent / 100;
		uint256 _totalSupply = totalSupply();
		
		uint256 reserve0 = getbalance0(); 
		uint256 reserve1 = getbalance1(); 

		( uint256 total0, uint256 total1 ) = IStrategy(myStrategy()).getTotalAmounts();
		total0 += reserve0;
		total1 += reserve1;
		
		uint256 t0 = total0 * shares / _totalSupply;
		uint256 t1 = total1 * shares / _totalSupply;
		
		if (reserve0 >= t0 && reserve1 >= t1) {
			withdrawFinal(msg.sender, shares,t0, t1);
			pending = false;
		}
		
		if ( pending ) {
			int id = checkExist(msg.sender);
			// if pending withdraw exists, updates the withdraw data.
			if (id < 0 ) {
				wds.push(Withdrawal({recipient:msg.sender, percent:percent}));
			} else {
				wds[uint(id)].percent = percent;
			}
		}
		
		emit PendingWithdraw(msg.sender, pending);
		
    }
    
    function checkExist(address addr) internal view returns(int){
    	for (uint i=0; i< wds.length; i++ ) {
			if (wds[i].recipient == addr) {
				return int(i);
			}
		}
		return -1;
    }

	
    function withdrawProc(address to, uint8 percent) internal    	
	{

        uint256 shares = balanceOf(to) * percent / 100;
		uint256 _totalSupply = totalSupply();
		
		if ( to == address(0) ||  shares > _totalSupply ) 
			return;
        
        (uint256 total0, uint256 total1) = ( getbalance0(), getbalance1());

        withdrawFinal(
        	to, 
        	shares,
        	total0 * shares / _totalSupply,
			total1 * shares / _totalSupply
		);
        
    }
    
    function withdrawFinal(address to, uint256 shares, uint256 amount0, uint256 amount1) internal  {
    	
		//log withdraw
        emit Withdraw(to, shares, amount0, amount1);
		
		if (shares == 0 ) {
			return;	
		}
		
        _burn(to, shares);	// burn user's share

		if(amount0>0) { IERC20(token0).safeTransfer(to, amount0); }
		if(amount1>0) { IERC20(token1).safeTransfer(to, amount1); }
		
		// update user's total deposits 
		 Assetholder[to].deposit0  = ( Assetholder[to].deposit0 > amount0  ) ? 
				Assetholder[to].deposit0 -= amount0 : 0;
		 Assetholder[to].deposit1  = ( Assetholder[to].deposit1 > amount1  ) ? 
				Assetholder[to].deposit1 -= amount1 : 0;

		
		
    }

	function calcPositionShares( uint256 amountIn0, uint256 amountIn1) 
		public 
		returns(uint256 shares, uint256 amount0, uint256 amount1)
	{
		(amount0, amount1) = (amountIn0, amountIn1);


		uint256 totalSupply = totalSupply();
		if (totalSupply == 0 ) {
			shares = calcShare( amount0, amount1);  
			
		} else {
			( uint256 total0, uint256 total1 ) = IStrategy(myStrategy()).getTotalAmounts();

			total0 += getbalance0();
			total1 += getbalance1();
			
			uint256 estTotalSupply = calcShare( total0, total1);
			uint256 estShare =   calcShare( amountIn0, amountIn1); 
			shares = estShare * totalSupply / estTotalSupply ;
		}
	}
	
	function mintFees(address to, uint256 fee0, uint256 fee1) external onlyStrategy {
		uint256 share = calcShare(fee0, fee1); 
		emit MintFees(to, share, fee0, fee1);
		_mint(to, share);
	}
	
	function calcShare(uint256 a0, uint256 a1) public view returns(uint256){
		
		uint256 p = IStrategy(myStrategy()).getPrice(); 
		
		require(p>0,'p0');
		
		// share =(weth + usdc/price * 1e18)  //--> price = xxxx per weth
		return   (baseToken  == token0 ) ? (a0 + a1/p*1e18) : (a1 + a0/p*1e18); 
	}
	
    
    //move funds to strategy
    function moveFunds() external onlyStrategy {
    	//idea: factory.getStrategy() instead of onlyStrategy
		uint256 a0 = getbalance0();
		uint256 a1 = getbalance1();
		uint256 a2 = IERC20(oSqth).balanceOf(address(this));
		
		if (a0 > 0) IERC20(token0).safeTransfer(myStrategy() , a0);
        if (a1 > 0) IERC20(token1).safeTransfer(myStrategy() , a1);
        if (a2 > 0) IERC20(oSqth).safeTransfer(myStrategy() , a2);
    }
    
    function getbalance0() public view returns(uint256) {
    	return(IERC20(token0).balanceOf(address(this)));
    }
  
    function getbalance1() public view returns(uint256) {
    	return(IERC20(token1).balanceOf(address(this)));
    }
  

}