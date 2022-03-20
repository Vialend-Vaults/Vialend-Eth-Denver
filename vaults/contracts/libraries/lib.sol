// SPDX-License-Identifier: MIT
pragma solidity >=0.5.0;



library lib  {
 		
	 // Returns approximate square root of 256-bit value
	  function sqrt(uint256 x) internal pure returns (uint256 y) {
		
	    uint256 z = (x + 1) >> 1;
	    y = x;
	
	    while (z < y) {
	      y = z;
	      z = (x / z + z) >> 1;
	    }
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
       
 /// @dev Rounds tick down towards negative infinity so that it's a multiple
    /// of `tickSpacing`.
    function _floor(int24 tick,int24 tickSpacing) internal pure returns (int24) {
        int24 compressed = tick / tickSpacing;
        if (tick < 0 && tick % tickSpacing != 0) compressed--;
        return compressed * tickSpacing;
    }

    /// @dev Casts uint256 to uint128 with overflow check.
    function _toUint128(uint256 x) internal pure returns (uint128) {
        assert(x <= type(uint128).max);
        return uint128(x);
    }
 
}