
// SPDX-License-Identifier: MIT

pragma solidity >=0.8.8 <0.9.0;


import "./libraries/UniCompHelper.sol";

interface itest {
	
	
	function rebalance() external;
	function getAmounts() external;
	
}


contract test1  {
	UniCompHelper.StratInfo info;
	constructor( UniCompHelper.StratInfo memory _info) {
		info = _info;
	}
	
	function rebalance() public {}
	function getAmounts() public view {}
	
}


contract test2 {
	UniCompHelper.StratInfo2 info;
	constructor( UniCompHelper.StratInfo2 memory _info) {
		info = _info;
	}
	
	function rebalance() public  {}
	function getAmounts() public view {}
	
}

contract testMe {
	itest mytest;
	
	constructor(address _test) {
		mytest = itest(_test);
	}
	
	function rebalance() public {
		mytest.rebalance();
	}
	
}
