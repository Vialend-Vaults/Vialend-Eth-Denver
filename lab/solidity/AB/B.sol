// SPDX-License-Identifier: MIT
pragma solidity ^0.7.0;

import "./@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "./Ai.sol";
import "./A.sol";




contract Deployer {

    function deploy() internal returns (address _A) {
        uint256 mintable = 2000;
        address token0 = address(0xB4FBF271143F4FBf7B91A5ded31805e42b2208d6);
        address token1 = address(0xD87Ba7A50B2E7E660f678A895E4B72E7CB4CCd9C);
        uint8 decimal0 = 18;
        uint8 decimal1 = 8;
        
        _A = address(new A{salt:keccak256(abi.encode(block.difficulty,block.timestamp, token0,token1))}(mintable,token0, decimal0, token1, decimal1) );
        
        delete mintable;
    }
}

contract B is Ai, Bi,Deployer{
    
    Ai public immutable a;

    string public sparam;
    uint256 public uintparam;
    address public addrparam;


    constructor ()  {

        a = Ai (deploy());
        
        
    }

    function checkTotalSupply() public override returns( address, uint256) {
        
        uint256 Atotal = IERC20(address(a)).totalSupply();

        uintparam  = Atotal;

        return (msg.sender, Atotal);

        
    }

    function addToken(address _token,uint8 decimal ) override public  {
        
        a.addToken(_token, decimal);

        
    }

    
    function setParam0(uint256 value ) override public  {
        
        a.setParam0(value);

        
    }

       
    function takeFund(uint256 value ) public  {
        
        a.setParam0(value);

        
    }

    function changebyA(string memory value) override public {
        
         ( msg.sender == address(a) ) ?  sparam = value : sparam = "not A";

    }

}
