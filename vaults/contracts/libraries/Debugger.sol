 // SPDX-License-Identifier: MIT
pragma solidity >=0.5.0;

library Debugger  {

	function uint2str(
	  uint256 _i
	)
	  internal
	  pure
	  returns (string memory str)
	{
	  if (_i == 0)
	  {
	    return "0";
	  }
	  uint256 j = _i;
	  uint256 length;
	  while (j != 0)
	  {
	    length++;
	    j /= 10;
	  }
	  bytes memory bstr = new bytes(length);
	  uint256 k = length;
	  j = _i;
	  while (j != 0)
	  {
	    bstr[--k] = bytes1(uint8(48 + j % 10));
	    j /= 10;
	  }
	  str = string(bstr);
	}


	function append2(string memory a1, string  memory a2) public pure returns (string memory) {
	
	    return string(abi.encodePacked(a1, a2));
	
	}

}