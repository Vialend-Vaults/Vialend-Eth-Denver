// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.4.0 <0.9.0;

import "https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v2.5.0/contracts/math/SafeMath.sol";


contract test {
        
    using SafeMath for uint256;

    uint256 public number;
    uint256 public number1;
    uint256 public number2;

    
    
    function getShare ( uint256 num1, uint256 num2 ) public returns(uint256 , uint256, uint256) {
        
        
        number =  _sqrt( num1.mul(1e18).mul(num2.mul(1e18)) );
        
        number1=num1;
        number2=num2;
        
        
        return (number, num1, num2);
        
    }
    
      function testDiv ( uint256 share, uint256 totalshare, uint256 totalvalue  ) public returns(uint256 , uint256, uint256) {
        
        number = totalvalue.mul(share).div(totalshare);
    }

    
    function _sqrt(uint256 x) internal pure returns(uint256) {
	  
		if (x == 0)  return 0;   // Avoid zero divide crash    
		
		uint256 z = (x + 1 ) / 2;
		uint256 y = x;
		while(z < y){
		  y = z;
		  z = ( x / z + z ) / 2;
		}
		
		// require(y*y == x , "sqrt");  // cause crash
		
		return y;
	}



	 function calcSahres(uint256 x, uint256 y) public  returns(uint256){
    	
    	(uint256 shares) = cantor_pair_calculate(x,y);
    	
    	number = shares;

		return (shares);
    }

	 function fetchAmounts(uint256 myshare) public  returns(uint256,uint256){
    	
    	(uint256 d0,uint256 d1) = cantor_pair_reverse(myshare);
    	
    	number1 = d0;
    	number2 = d1;
		
		return ( d0, d1);
    }
 
 
	function cantor_pair_calculate(uint256 x , uint256 y ) internal pure returns (uint256 ) {
	
	
		uint256 result = ((x+y)*(x+y+1))/2 + y;

		return result;
	}

  
	 function cantor_pair_reverse (uint256 z ) internal pure returns (uint256 ,uint256 ) {
	
	 	uint256 t =  ( _sqrt(1+ 8 * z ) -1 )  / 2;
		
		uint256 x = t*(t+3)/2 - z;
		
		uint256 y = z - t*(t+1)/2;
	
	 	return (x, y);
	
	 }
	           
    
}