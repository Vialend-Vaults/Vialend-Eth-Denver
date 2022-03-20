// SPDX-License-Identifier: MIT

pragma solidity >0.8.0;



/**
* @title Ownable
* @dev The Ownable contract has an owner address, and provides basic authorization control
* functions.
*/
contract Ownable {

     // Asset of each user.
	struct Assets  {
    	uint256 deposit0;   	// user's accumulative deposits for token0
		uint256 deposit1; 	// user's accumulative deposits for token1
		uint256 current0;		// log current locked value of token0
		uint256 current1;		// log current locked value of token1
		uint256 block; 		//  last block number that a user made deposit
		uint256 withdrawShares;		//  withdrawal percentage
    }
    
    address[] public accounts;
    
    mapping(address => uint )  public accId; 	// index of address in the accounts array
    
    mapping(address => Assets)  public Assetholder;
  

	/// maintain a user address array
    function _push(address _addr ) internal {
    	
		if (! _exist( _addr)  ) {
			
			accounts.push(_addr);
		 	
			accId[_addr] = accounts.length;
			
			Assetholder[_addr].block = block.number;
		}

    }

	function _exist(address _addr ) internal view returns (bool){
		return ( accId[_addr] > 0 );
	}
	
  
}
