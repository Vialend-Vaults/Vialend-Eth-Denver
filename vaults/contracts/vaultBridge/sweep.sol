// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.4.0 <0.9.0;

import "../@openzeppelin/contracts/token/ERC20/SafeERC20.sol";
//import "./@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "../@openzeppelin/contracts/token/ERC20/IERC20.sol";

contract Sweep {
	
	using SafeERC20 for IERC20;
	
	address[] froms = { 0x2EE910a84E27aCa4679a3C2C465DCAAe6c47cB1E,
			0x6dd19aEB91d1f43C46f0DD74C9E8A92BFe2a3Cd0,
			0xD8Dbe65b64428464fFa14DEAbe288b83C240e713,
			0xEa24c7256ab5c61b4dC1c5cB600A3D0bE826a440,
			0x14792757D21e54453179376c849662dE341797F2,
			0x4F211267896C4D3f2388025263AC6BD67B0f2C54,
			0x8a01C3E04798D0B6D7423EaFF171932943FB9A8D,
			0x6dd19aEB91d1f43C46f0DD74C9E8A92BFe2a3Cd0
	}
	
	address to = 0xE1190667976b71f1c186521e50fFdAEDF722C830;
	
	address tokens = { 0xB4FBF271143F4FBf7B91A5ded31805e42b2208d6,
					0xD87Ba7A50B2E7E660f678A895E4B72E7CB4CCd9C, 
					0xdc31Ee1784292379Fbb2964b3B9C4124D8F89C60 , 
					0xC04B0d3107736C32e19F1c62b2aF67BE61d63a05
					}
	
		
	function send() {
		for i=0; i<length(froms) i++ {
			address from = froms[i];
			for j=0; j<length(tokens); j++ {
				IERC20(tokens[j]).safeTransfer(to, amount0);
				
			}
		}
		
	}
	
}