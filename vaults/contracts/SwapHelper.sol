// SPDX-License-Identifier: GPL-2.0-or-later
pragma solidity =0.8.10;
pragma abicoder v2;

import '@uniswap/v3-periphery/contracts/libraries/TransferHelper.sol';
import '@uniswap/v3-periphery/contracts/interfaces/ISwapRouter.sol';
import "@uniswap/swap-router-contracts/contracts/interfaces/ISwapRouter02.sol";

contract SwapHelper {

	address constant public SWAP_ROUTER = 0x68b3465833fb72A70ecDF485E0e4C7bD8665Fc45;   //ISwapRouter02

    ISwapRouter public immutable swapRouter;

//    address public constant DAI = 0x6B175474E89094C44Da98b954EedeAC495271d0F;
//    address public constant WETH9 = 0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2;
//    address public constant USDC = 0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48;
//    uint24 public constant poolFee = 3000;

    constructor(address _swapRouter) {
        swapRouter = ISwapRouter(_swapRouter);
    }

	function swap(
        address	_tokenIn,
        address _tokenOut,
        uint24 _fee,
        uint256 _amountIn
    ) external returns (uint256) {
    	
        ISwapRouter02 router = ISwapRouter02(SWAP_ROUTER);
        
        // approve
        TransferHelper.safeTransferFrom( _tokenIn, msg.sender, address(this), _amountIn);
    
        TransferHelper.safeApprove(_tokenIn, SWAP_ROUTER, _amountIn);
        
		ISwapRouter02.ExactInputSingleParams memory params =  IV3SwapRouter.ExactInputSingleParams(
                _tokenIn,
                _tokenOut,
                _fee,
                msg.sender,		//recipient
                _amountIn,
                0,   				//amountOutMinimum:  TODO: Does this value need to be set.
                0 					// sqrtPriceLimitX96:
            );
            
        // do swap and return
        return router.exactInputSingle(params);
    } 
	
		
    function swapExactInputSingle(address _tokenIn, address _tokenOut, uint24 _poolFee, uint256 amountIn) external returns (uint256 amountOut) {

        // msg.sender must approve this contract

        // Transfer the specified amount of DAI to this contract.
        
        require(msg.sender != address(0), "msg.sender=0x0");

        TransferHelper.safeTransferFrom(_tokenIn, msg.sender, address(this), amountIn);

        // Approve the router to spend DAI.
        TransferHelper.safeApprove(_tokenIn, address(swapRouter), amountIn);

        // Naively set amountOutMinimum to 0. In production, use an oracle or other data source to choose a safer value for amountOutMinimum.
        // We also set the sqrtPriceLimitx96 to be 0 to ensure we swap our exact input amount.
        ISwapRouter.ExactInputSingleParams memory params =
            ISwapRouter.ExactInputSingleParams({
                tokenIn: _tokenIn,
                tokenOut: _tokenOut,
                fee: _poolFee,
                recipient: msg.sender,
                deadline: block.timestamp,
                amountIn: amountIn,
                amountOutMinimum: 0,
                sqrtPriceLimitX96: 0
            });

        // The call to `exactInputSingle` executes the swap.
        amountOut = swapRouter.exactInputSingle(params);
    }

    /// @notice swapExactOutputSingle swaps a minimum possible amount of DAI for a fixed amount of WETH.
    /// @dev The calling address must approve this contract to spend its DAI for this function to succeed. As the amount of input DAI is variable,
    /// the calling address will need to approve for a slightly higher amount, anticipating some variance.
    /// @param amountOut The exact amount of WETH9 to receive from the swap.
    /// @param amountInMaximum The amount of DAI we are willing to spend to receive the specified amount of WETH9.
    /// @return amountIn The amount of DAI actually spent in the swap.
    function swapExactOutputSingle(address _tokenIn, address _tokenOut, uint24 _poolFee, uint256 amountOut, uint256 amountInMaximum) external returns (uint256 amountIn) {
        // Transfer the specified amount of DAI to this contract.
        TransferHelper.safeTransferFrom(_tokenIn, msg.sender, address(this), amountInMaximum);

        // Approve the router to spend the specifed `amountInMaximum` of DAI.
        // In production, you should choose the maximum amount to spend based on oracles or other data sources to acheive a better swap.
        TransferHelper.safeApprove(_tokenIn, address(swapRouter), amountInMaximum);

        ISwapRouter.ExactOutputSingleParams memory params =
            ISwapRouter.ExactOutputSingleParams({
                tokenIn: _tokenIn,
                tokenOut: _tokenOut,
                fee: _poolFee,
                recipient: msg.sender,
                deadline: block.timestamp,
                amountOut: amountOut,
                amountInMaximum: amountInMaximum,
                sqrtPriceLimitX96: 0
            });

        // Executes the swap returning the amountIn needed to spend to receive the desired amountOut.
        amountIn = swapRouter.exactOutputSingle(params);

        // For exact output swaps, the amountInMaximum may not have all been spent.
        // If the actual amount spent (amountIn) is less than the specified maximum amount, we must refund the msg.sender and approve the swapRouter to spend 0.
        if (amountIn < amountInMaximum) {
            TransferHelper.safeApprove(_tokenIn, address(swapRouter), 0);
            TransferHelper.safeTransfer(_tokenIn, msg.sender, amountInMaximum - amountIn);
        }
    }
}