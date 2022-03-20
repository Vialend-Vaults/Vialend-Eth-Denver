// SPDX-License-Identifier: MIT

pragma solidity ^0.8.10;

interface IViaVault
{
	function moveFunds() external;
	function wdsLen() external view returns(uint256);
	function withdrawLoop() external ;
	function mintFees(address , uint256, uint256) external ;
}
