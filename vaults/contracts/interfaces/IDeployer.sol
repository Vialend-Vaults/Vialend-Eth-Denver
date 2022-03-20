// SPDX-License-Identifier: MIT

pragma solidity >=0.8.8;

interface IDeployer
{
   function deployVault(
		address _factory,
    	address	_admin,
		address	token0,
		address	token1,
		uint	_vaultCap,
		uint	_individualCap
    ) external returns(address _vault);
    

 	function deployStrategy(
    	address[11] memory _contracts,
		uint8  _uniPortion,
		uint8  _compPortion, 
		uint8  _protocolFee,
		uint24 _feetier,
		uint128  _quoteAmount
	) external returns(address _strat);
	    	
}