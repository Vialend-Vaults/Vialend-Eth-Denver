// SPDX-License-Identifier: Unlicense

pragma solidity ^0.7.0;


interface Ai {
    
    function addToken(address newToken, uint8 decimal) external;

    function setParam0(uint256 value) external;

    function checkTotalSupply() external returns( address, uint256) ;

}