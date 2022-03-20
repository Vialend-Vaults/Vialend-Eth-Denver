// SPDX-License-Identifier: MIT

pragma solidity >=0.8.8 <0.9.0;

import "./VaultStrategy.sol";

contract StratDeployer {
	
	event DeployedStrat(address sender, address _strat);
	   
    function deployStrategy(
    	address[11] memory _contracts,
		uint8  _uniPortion,
		uint8  _compPortion, 
		uint8  _protocolFee,
		uint24 _feetier,
		uint128  _quoteAmount
	) external returns(address _strat){
			_strat = payable( address(new VaultStrategy{salt:keccak256(abi.encode(block.difficulty,block.timestamp))}
				(	_contracts,
					_uniPortion,
					_compPortion, 
					_protocolFee, 
					_feetier,    
					_quoteAmount
					
				)));
				
			emit DeployedStrat(msg.sender, _strat);
			
			return(_strat);
			
    }
}
