// SPDX-License-Identifier: MIX

pragma solidity >=0.8.10;

// future feature after 0.8.10 usage: if (balance < amount)  revert NotEnoughFunds(amount, balance);
error NotEnoughFunds(uint requested, uint available);  
error OraclePrice(uint available);  

	