// SPDX-License-Identifier: MIT
pragma solidity >=0.8.8;

contract UniCompFees  {

	uint256 internal uFees0;	// uniswap fee of token0 for the current position
	uint256 internal uFees1;	// uniswap fee of token1 for the current position
	uint256 internal lFees0;	// lending fee of token0 for the current position
	uint256 internal lFees1;	// lending fee of token1 for the current position

	uint256 public compIn0;
	uint256 public compIn1;
	uint256 public compOut0;
	uint256 public compOut1;
	
}