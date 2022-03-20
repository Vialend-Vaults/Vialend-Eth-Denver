// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.4.0 <0.9.0;
/*
solc --optimize --overwrite --abi vaultBridge.sol -o ../../build/
solc --optimize --overwrite --bin vaultBridge.sol -o ../../build/

abigen --abi=../../build/vaultBridge.abi --bin=../../build/vaultBridge.bin --pkg=api --out=../../deploy/vaultBridge/vaultBridge.go

abigen --abi=../../build/vaultAdmin.abi --bin=../../build/vaultAdmin.bin --pkg=api --out=../../deploy/vaultAdmin/vaultAdmin.go

v2.x  currently use
deployed on remix  
0x428EeA0B87f8E0f5653155057f58aaaBb667A3ec

2022/jan/14 can change owner
0x601D3992aB70ffEe3858730b12d94AAB35A4a60E
*/


/*
	VaultBridge 
	store vault address indexed by a mapping var
		0: weth/usdc vault
		1: weth/dai  vault
*/

contract VaultBridge  {

	address public owner  ;
    
    mapping(address => uint ) public permit;

    mapping(uint  => address) public vaults;


  	constructor()  {
		owner = msg.sender;
	}
        
   	modifier onlyOwner {
         require(msg.sender == owner, "owner");
        _;
    }


    function getAddress(uint ind) external view returns( address) {
    	return vaults[ind];
    }

    function setAddress(address newAddress,  uint ind ) external onlyOwner {
		vaults[ind] = newAddress;
    	
    }

    function getPermit(address addr ) external view returns(uint) {
		return permit[addr];
    	
    }

    function setPermit(address addr,  uint level ) external onlyOwner {
		permit[addr] = level;
    	
    }
    
    function changeOwner(address newOwner ) external onlyOwner {
		owner = newOwner;
    	
    }
   

}

