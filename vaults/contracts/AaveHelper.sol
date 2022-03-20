// SPDX-License-Identifier: MIT

pragma solidity 0.8.10;


import {TransferHelper} from "@uniswap/v3-periphery/contracts/libraries/TransferHelper.sol";

//import {IPoolAddressesProvider} from "./aave-v3-core/contracts/interfaces/IPoolAddressesProvider.sol";
//import {IPool} from "./aave-v3-core/contracts/interfaces/IPool.sol";
//import {DataTypes} from "./aave-v3-core/contracts/protocol/libraries/types/DataTypes.sol";


import {IPoolAddressesProvider} from "./aave-v3-core/contracts/interfaces/IPoolAddressesProvider.sol";
import {IPool} from "./aave-v3-core/contracts/interfaces/IPool.sol";
import {DataTypes} from "./aave-v3-core/contracts/protocol/libraries/types/DataTypes.sol";

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

contract AaveHelper {

    /** 
    * Create a short using the given collateral and short addresses. _shortSize is in uinits of short
    */
    function short(
        address _aavePoolProvider,
        address _callback,
        address _collateral,
        address _shorting,
        uint256 _collateralSize,  // 1,000,000 usdc // 2x short requres: collateral / ethprice * 2 = shortsize
        uint256 _shortSize       // 666.66 eth
    ) internal returns(bool) {

        // Approve call back to transfer collateral
        TransferHelper.safeApprove(_collateral, _callback, _collateralSize);

        // set up flash loan
        bytes memory params = abi.encode(_collateral, _collateralSize);
        address[] memory assets = new address[](1);
        assets[0] = _shorting;
        uint256[] memory sizes = new uint256[](1);
        sizes[0] = _shortSize;
        uint256[] memory modes = new uint256[](1);
        modes[0] = 2;           // 0 repay, 1 stable, 2 variable

        // flash loan the amount times the size of the short, swap for more collateral and deposit.
        IPool(IPoolAddressesProvider(_aavePoolProvider).getPool()).flashLoan(
            _callback,
            assets,
            sizes,
            modes,              
            address(this),
            params,
            0
        );

        return true;
    }

    /**
    * If you have a 4x short, with a 5:4 collateral to debt ratio, this flash borrows collateral equal 
    * in value to the short (4 poritons), swaps from borrow to shorted asset, pays back short debt,
    * withdraws collateral, and repays flash loan. To be used with AaveUnwindCallback class in 
    * AaveShorter.sol file.
    */
    function unwind(
        address _aavePoolProvider,
        address _callback,
        address _collateral,
        address _shorted,
        uint256 _ethPerUsdc  // ratio or value in eth/usdc since 10^18 / 10^6 > 1
    ) internal returns(bool) {

        // Get the atoken addresses
        IPool pool = IPool(IPoolAddressesProvider(_aavePoolProvider).getPool());
        DataTypes.ReserveData memory collateralReserves = pool.getReserveData(_collateral);
        address aToken = collateralReserves.aTokenAddress;

        // Approve transfer of aTokens to callback and prep flash loan
        TransferHelper.safeApprove(aToken, _callback, IERC20(aToken).balanceOf(address(this)));
        uint256 debt = getDebt(pool, _shorted);
        bytes memory params = abi.encode(_collateral, debt / _ethPerUsdc);  // 10^18 / (10^12) -> 10^6

        address[] memory assets = new address[](1);
        assets[0] = _shorted;
        uint256[] memory sizes = new uint256[](1);
        sizes[0] = debt;
        uint256[] memory modes = new uint256[](1);
        modes[0] = 0;   // 0 repay, 1 stable, 2 variable

        // flash loan the amount times the size of the short, swap for more collateral and deposit.
        pool.flashLoan(
            _callback,
            assets,
            sizes,
            modes,              
            address(this),
            params,
            0
        );

        return true;
    }

    /**
    * Get debt amount of the given asset.
    */
    function getDebt(IPool pool, address _shorted) internal view returns(uint256) {
        DataTypes.ReserveData memory debtReserves = pool.getReserveData(_shorted);
        address debtTokenAddress = debtReserves.variableDebtTokenAddress;
        return IERC20(debtTokenAddress).balanceOf(address(this)); // denominated in underlying asset not aToken.
    }
}
