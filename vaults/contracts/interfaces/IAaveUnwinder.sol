/ SPDX-License-Identifier: MIT

pragma solidity >=0.8.8;

interface IAaveShorter {
    
    function unwind(
        address _aavePool,
		address _colateral,
        address _unwinding,
		uint256 _amount,
        unit256 _unwindSize

    ) external returns(bool);
}