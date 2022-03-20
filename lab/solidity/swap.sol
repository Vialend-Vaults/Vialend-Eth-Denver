import "./../interfaces/IUniswapV2Router.sol"; 

contract MyContract {

   address UNISWAP_V2_ROUTER = 0x....;

    function swap(...) external {
        uint[] memory out_amounts = IUniswapV2Router02(UNISWAP_V2_ROUTER).swapExactTokensForTokens(
            amount, // amountIn
            minOut, //The minimum amount we want to receive, taking into account slippage
            path, // Path
            beneficiary, // To address
            block.timestamp + 60 // Deadline
        );

    }

}