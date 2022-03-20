// SPDX-License-Identifier: MIT
pragma solidity >=0.5.0;


interface ICEth {
    function mint() external payable;

	function exchangeRateStored() external  view returns (uint256);	

    function exchangeRateCurrent() external  returns (uint256);

    function supplyRatePerBlock() external  returns (uint256);

    function redeem(uint) external returns (uint);

    function transfer(address to, uint value) external returns (bool);

    function redeemUnderlying(uint) external returns (uint);
    
    function balanceOf(address) external view returns (uint256); 
}