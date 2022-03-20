// SPDX-License-Identifier: MIT
pragma solidity ^0.7.0;

import "./@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "./@openzeppelin/contracts/token/ERC20/SafeERC20.sol";

import "./Ai.sol";
import "./Bi.sol";

contract A is Ai, ERC20 {

    
    uint256 public param0 ;
    uint256 public uintparam;

    address public immutable _B;

    struct Tokens {
        address addr;
        uint8 decimal;
    }

    Tokens[] public tokens;

    constructor (uint256 mintable, address _token0, uint8 _decimal0,  address _token1, uint8 _decimal1) ERC20("Token", "TKN") {
        
        _mint(msg.sender, mintable * (10 ** uint256(decimals())));
        tokens.push(Tokens({addr: _token0, decimal:_decimal0}));
        tokens.push(Tokens({addr: _token1, decimal:_decimal1}));
        _B = msg.sender;
    }

    function addToken(address newToken, uint8 _decimal) override public  {
        tokens.push(Tokens({addr:newToken,decimal:_decimal}));
    }

    function setParam0(uint256 value) override public {
        param0 = value;
    }

    function checkTotalSupply() public override returns( address, uint256) {
        
        uint256 Atotal = totalSupply();

        uintparam  = Atotal;

        return (msg.sender, Atotal);

        
    }

   function changebyA(string memory str) public {
       Bi(_B).changebyA(str);
   }



}