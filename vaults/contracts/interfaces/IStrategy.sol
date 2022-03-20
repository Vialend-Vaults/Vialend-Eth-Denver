// SPDX-License-Identifier: MIT

pragma solidity >=0.8.8;

interface IStrategy
{
    function getPrice() external view returns(uint256);
	function getTotalAmounts() external returns(uint256 , uint256) ;
	function vaultWithdraw(address to, uint256 shares, uint256 totalSupply) external returns(bool);
	function rebalance(int24 newLow, int24 newHigh ) external returns(uint256, uint256, uint256,uint256);
	function rebalanceByVault(	int24 newLow, int24 newHigh	) external;
	function callFunds() external;
	
}