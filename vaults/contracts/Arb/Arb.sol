// SPDX-License-Identifier: GPL-3.0
/*
solc --optimize --overwrite --abi arb.sol -o ../../build/
solc --optimize --overwrite --bin arb.sol -o ../../build/

abigen --abi=../../build/arb_sol_arb.abi --bin=../../build/arb_sol_arb.bin --pkg=api --out=../../deploy/arb/arb.go


*/

pragma solidity <0.9.0;


contract arb {
        
    uint256 public number;
    uint256 public number1;
    uint256 public number2;
    uint256 public number3;
    
    uint256[] public array;

    constructor() {
    }
    
    function ethBalance(address _addr ) public view returns (uint256) {
        return ( address(_addr).balance );
    }
    
}