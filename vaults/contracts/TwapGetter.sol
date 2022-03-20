// SPDX-License-Identifier: MIT

pragma solidity >=0.8.8;

import "@uniswap/v3-core/contracts/interfaces/IUniswapV3Pool.sol";
import "@uniswap/v3-core/contracts/libraries/TickMath.sol";
import "@uniswap/v3-core/contracts/libraries/FixedPoint96.sol";
import "@uniswap/v3-core/contracts/libraries/FullMath.sol";

contract TwapGetter {

   function getTwap(address _pool , uint32 twapInterval) public view returns (int24 tick) {
        
		if (twapInterval == 0) {
            // return the current tick if twapDuration == 0
            (, tick , , , , , ) = IUniswapV3Pool(_pool).slot0();
        } else {
	        uint32[] memory secondsAgo = new uint32[](2);
	        secondsAgo[0] = twapInterval;
	        secondsAgo[1] = 0;
	
	        (int56[] memory tickCumulatives, ) = IUniswapV3Pool(_pool).observe(secondsAgo);
	        int56 tickDelta = tickCumulatives[1] - tickCumulatives[0];
	        tick = int24(tickDelta / int56(int32( twapInterval )));
	        
	        if (tickDelta < 0 && (tickDelta % int56(int32(twapInterval)) != 0)) tick--;
        }
    }

   function getQuoteAtTick(
        int24 tick,
        uint128 baseAmount,
        address baseToken,
        address quoteToken
    ) public pure returns (uint256 quote) {
    	
        uint160 sqrtRatioX96 = TickMath.getSqrtRatioAtTick(tick);

        // Calculate quoteamount with better precision if it doesn't overflow when multiplied by itself
        if (sqrtRatioX96 <= type(uint128).max) {
            uint256 ratioX192 = uint256(sqrtRatioX96) * sqrtRatioX96;
            quote = baseToken < quoteToken
                ? FullMath.mulDiv(ratioX192, baseAmount, 1 << 192)
                : FullMath.mulDiv(1 << 192, baseAmount, ratioX192);
        } else {
            uint256 ratioX128 = FullMath.mulDiv(sqrtRatioX96, sqrtRatioX96, 1 << 64);
            quote = baseToken < quoteToken
                ? FullMath.mulDiv(ratioX128, baseAmount, 1 << 128)
                : FullMath.mulDiv(1 << 128, baseAmount, ratioX128);
        }
    }



	
/** /// reference: cal price by sqrtx96
    function getSqrtTwapX96(address _pool, uint32 twapInterval) public view returns (uint160 sqrtPriceX96) {
        if (twapInterval == 0) {
            // return the current price if twapInterval == 0
            (sqrtPriceX96, , , , , , ) = IUniswapV3Pool(_pool).slot0();
        } else {
            uint32[] memory secondsAgos = new uint32[](2);
            secondsAgos[0] = twapInterval; // from (before)
            secondsAgos[1] = 0; // to (now)

            (int56[] memory tickCumulatives, ) = IUniswapV3Pool(_pool).observe(secondsAgos);

            // tick(imprecise as it's an integer) to price
            sqrtPriceX96 = TickMath.getSqrtRatioAtTick(
                int24((tickCumulatives[1] - tickCumulatives[0]) / int32(twapInterval))
            );
        }
    }

    function getPriceX96FromSqrtPriceX96(uint160 sqrtPriceX96) public pure returns(uint256 priceX96) {
        return FullMath.mulDiv(sqrtPriceX96, sqrtPriceX96, FixedPoint96.Q96);
    }
*/
}