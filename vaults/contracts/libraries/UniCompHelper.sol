// SPDX-License-Identifier: MIT
pragma solidity >=0.8.8;

library UniCompHelper{ 

	struct StratInfo1 {
        uint256 _vaultCap;
        uint256 _individualCap;
        uint128  _quoteAmount;
        uint8 _decimal0;
        uint8 _decimal1;
        uint8  _uniPortion;         
        uint8  _compPortion;        
        uint8  _creatorFee;     
        uint8  _adminFee;       
        uint24  _feetier;
        int24  _maxTwapDeviation;
        uint32  _twapDuration;
        uint32  _threshold;
       }
}


