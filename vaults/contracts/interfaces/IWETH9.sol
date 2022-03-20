// SPDX-License-Identifier: MIT

pragma solidity >=0.5.0;

interface IWETH9 
{
	function  balanceOf(address account) external returns(uint);
	function allowance( address owner, address spender) external returns(uint);

	fallback() external payable ;
	 receive() external payable ;

	function deposit() external payable ;
	function withdraw(uint wad) external ;
	function totalSupply() external view returns (uint) ;
	 
	function approve(address guy, uint wad) external returns (bool) ;
	
	function transfer(address dst, uint wad) external returns (bool) ;
	
	function transferFrom(address src, address dst, uint wad) external returns (bool);
}
