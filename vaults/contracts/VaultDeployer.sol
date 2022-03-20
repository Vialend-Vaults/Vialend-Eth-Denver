// SPDX-License-Identifier: MIT

pragma solidity >=0.8.8 <0.9.0;

import "./ViaVault.sol";

contract VaultDeployer {  
	
	event DeployedVault(address sender, address _vault);

    function deployVault(
    	address _factory,
		address	_admin,
		address	token0,
		address	token1,
		uint	_vaultCap,
		uint	_individualCap
    ) external returns(address _vault){
    	_vault = payable( address(new ViaVault{salt:keccak256(abi.encode(block.difficulty,block.timestamp))}
				(	_factory,
					_admin,
					token0,
					token1,
					_vaultCap,
					_individualCap
				)));
				
			
		emit DeployedVault(msg.sender, _vault);	
		
		return(_vault);
    }
}