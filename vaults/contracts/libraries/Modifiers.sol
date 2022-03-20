// SPDX-License-Identifier: MIT
pragma solidity >=0.8.8;

library Modifiers {
	
   modifier onlyActive {
        require (IVaultFactory(factory).checkActive( address(this),1 ), 'not active');
        _;
    }
    
 	modifier onlyAdmin {
        require (msg.sender == admin, 'admin');  
        _;
    }
  } 