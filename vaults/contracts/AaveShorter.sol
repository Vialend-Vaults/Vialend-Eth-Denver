// SPDX-License-Identifier: MIT

pragma solidity 0.8.10;

import "https://github.com/aave/aave-v3-core/blob/master/contracts/interfaces/IPool.sol";
import "https://github.com/aave/aave-v3-core/blob/master/contracts/flashloan/base/FlashLoanReceiverBase.sol";
import "https://github.com/aave/aave-v3-core/blob/master/contracts/interfaces/IPoolAddressesProvider.sol";
import "@uniswap/v3-periphery/contracts/libraries/TransferHelper.sol";
import "https://github.com/Uniswap/swap-router-contracts/blob/main/contracts/interfaces/ISwapRouter02.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "https://github.com/aave/aave-v3-core/blob/master/contracts/protocol/libraries/types/DataTypes.sol";

/**
 * @notice Contract for creating a short position using Aave for borrowing and Uniswap for swap.
 */
contract AaveShorterCallback is FlashLoanReceiverBase {
    address public constant SWAP_ROUTER = 0x68b3465833fb72A70ecDF485E0e4C7bD8665Fc45;
    address public aavePool;

    constructor(IPoolAddressesProvider _addressProvider) public FlashLoanReceiverBase(_addressProvider) {
        // noop
    }

    /**
    * @notice Executes an operation after receiving the flash-borrowed assets
    * @dev Ensure that the contract can return the debt + premium, e.g., has
    *      enough funds to repay and has approved the Pool to pull the total amount
    * @param assets The addresses of the flash-borrowed assets
    * @param amounts The amounts of the flash-borrowed assets
    * @param initiator The address of the flashloan initiator
    * @param params The byte-encoded params passed when initiating the flashloan
    * @return True if the execution of the operation succeeds, false otherwise
    */
    function executeOperation(
        address[] calldata assets,
        uint256[] calldata amounts,
        uint256[] calldata premiums,
        address initiator,
        bytes calldata params
    ) external override returns (bool) {

        require(msg.sender == address(POOL), "Only allow aave to call");
        (address _collateral, uint256 _collateralSize) = abi.decode(params, (address, uint256));

        // transfer from requires approval prior
        TransferHelper.safeTransferFrom(_collateral, initiator, address(this), _collateralSize);

        // approve
        TransferHelper.safeApprove(assets[0], SWAP_ROUTER, amounts[0]);

        // swap
        ISwapRouter02.ExactInputSingleParams memory swapParams = IV3SwapRouter.ExactInputSingleParams({
                tokenIn: assets[0],
                tokenOut: _collateral,
                fee: 3000,
                recipient: address(this),
                amountIn: amounts[0],
                amountOutMinimum: 0,   // TODO: Does this value need to be set.
                sqrtPriceLimitX96: 0
            });
        uint256 out = ISwapRouter02(SWAP_ROUTER).exactInputSingle(swapParams);

        // deposit
        TransferHelper.safeApprove(_collateral, address(POOL), _collateralSize + out);
        POOL.supply(_collateral, _collateralSize + out, initiator, 0); // USDC => AUSDC + debtTokenEth

        return true;
    }   
}


/**
 * @notice Contract for creating a short position using Aave for borrowing and Uniswap for swap.
 */
contract AaveUnwindCallback is FlashLoanReceiverBase {
    address public constant SWAP_ROUTER = 0x68b3465833fb72A70ecDF485E0e4C7bD8665Fc45;
    address public aavePool;

    constructor(IPoolAddressesProvider _addressProvider) public FlashLoanReceiverBase(_addressProvider) {
        // noop
    }

    /**
    * @notice Executes an operation after receiving the flash-borrowed assets
    * @dev Ensure that the contract can return the debt + premium, e.g., has
    *      enough funds to repay and has approved the Pool to pull the total amount
    * @param assets The addresses of the flash-borrowed assets
    * @param amounts The amounts of the flash-borrowed assets
    * @param initiator The address of the flashloan initiator
    * @param params The byte-encoded params passed when initiating the flashloan
    * @return True if the execution of the operation succeeds, false otherwise
    */
    function executeOperation(
        address[] calldata assets,
        uint256[] calldata amounts,
        uint256[] calldata premiums,
        address initiator,
        bytes calldata params
    ) external override returns (bool) {

        require(msg.sender == address(POOL), "Only allow aave POOL to call");
        (address _collateral, uint256 _collateralSizeOfDebt) = abi.decode(params, (address, uint256));
        address _aToken = POOL.getReserveData(_collateral).aTokenAddress;
        address asset = assets[0];
        uint256 amount = amounts[0];
        uint256 premium = premiums[0];

        // pay loan with flash loan
        TransferHelper.safeApprove(asset, address(POOL), amount);
        POOL.repay(
            asset,
            amount, 
            2,          // 1 stable, 2 variable
            initiator
        );

        // transfer collateral Atoken
        uint256 _collateralSize = IERC20(_aToken).balanceOf(initiator);
        TransferHelper.safeTransferFrom(_aToken, initiator, address(this), _collateralSize);

        // withdraw collateral
        POOL.withdraw(
            _collateral,
            _collateralSize,
            address(this)
        );

        // swap
        TransferHelper.safeApprove(_collateral, SWAP_ROUTER, _collateralSizeOfDebt);

        ISwapRouter02.ExactInputSingleParams memory swapParams = IV3SwapRouter.ExactInputSingleParams({
                tokenIn: _collateral,
                tokenOut: asset,
                fee: 3000,
                recipient: address(this),
                amountIn: _collateralSizeOfDebt,
                amountOutMinimum: 0,
                sqrtPriceLimitX96: 0
            });
        
        uint256 out = ISwapRouter02(SWAP_ROUTER).exactInputSingle(swapParams);
        uint256 repay = amount + premium;
        
        require(out >= repay, "Insufficient swap");

        // approve repay flash & transfer excess to inititor
        TransferHelper.safeApprove(asset, address(POOL), repay);
        IERC20(asset).transfer(initiator, out - repay);
        IERC20(_collateral).transfer(initiator, _collateralSize - _collateralSizeOfDebt);
        return true;
    }
}
