// SPDX-License-Identifier: MIT

pragma solidity 0.8.10;


import {AaveHelper} from "../contracts/AaveHelper.sol";
import "@uniswap/v3-periphery/contracts/libraries/TransferHelper.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

/**
* Stub strategy uses the AaveHelper in the expected way the actual strategy will use it. 
* Currently deployed or Rinkeby at 0x42cb0f8077dA6d641c7C277f6078dfAbf77d7766
*/
contract StubStrategy is AaveHelper {

    address public aavePoolProvider = 0xA55125A90d75a95EC00130E8E8C197dB5641Eb19;
    address public shortCallback = 0xa55eC76678548b4Eff64A09fA693cC18F7c6fc8c;
    address public unwindCallback = 0x5fC0Ed674B2dA337Cf9BCfd5f17DA1F52Eb2F661;
    address public aaveUSDC = 0x5B8B635c2665791cf62fe429cB149EaB42A3cEd8;
    address public aaveEth = 0x98a5F1520f7F7fb1e83Fe3398f9aBd151f8C65ed;
    uint256 public ethPerUsdc = 250563768; //  = Eth 1 -> 1000000000000000000 / USDC 3991 -> 3991000000 (aave test net ETH value) -> 250563768 
    address public owner;

    constructor (address _owner) public {
        owner = _owner;
    }

    modifier onlyOwner() {
        require(msg.sender == owner);
        _;
    }

    function testHelperShort(uint256 usdc, uint256 leverage) public onlyOwner {
        TransferHelper.safeTransferFrom(aaveUSDC, msg.sender, address(this), usdc);
        short(aavePoolProvider, shortCallback, aaveUSDC, aaveEth, usdc, usdc * ethPerUsdc * leverage); 
    }

    function testHelperUnwind() public onlyOwner {
        unwind(aavePoolProvider, unwindCallback, aaveUSDC, aaveEth, ethPerUsdc);
        IERC20 aaveUsdcToken = IERC20(aaveUSDC);
        aaveUsdcToken.transfer(msg.sender, aaveUsdcToken.balanceOf(address(this)));
    }

    function setEthPerUsdc(uint256 _ethPerUsdc) public onlyOwner {
        ethPerUsdc = _ethPerUsdc;
    }

    function setShortCallback(address _callback) public onlyOwner {
        shortCallback = _callback;
    }

    function setUnwindCallback(address _callback) public onlyOwner {
        unwindCallback = _callback;
    }
}
