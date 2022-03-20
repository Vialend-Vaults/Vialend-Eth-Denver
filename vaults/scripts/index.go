package main

import (
	"fmt"
	"math/big"
	_ "time"

	project "viaroot/scripts/project"
)

type TransferStruct struct {
	AccountId    int
	Amount       *big.Int
	TokenAddress string
	ToAddress    string
}
type Switcher struct {
	ViewOnly         bool
	DeployToken      int
	TokenParam       [2]project.TokenStruct
	TransferToken    int
	TransferParam    [2]TransferStruct
	DeployFactory    int
	CreatePool       int
	InitialPool      int
	MintPool         int
	MintPoolParam    [3]int64
	DeployVault      int
	Deposit          int
	DepositParam     [3]int64
	Withdraw         int
	WithDrawParam    [2]int64
	Rebalance        int
	RebalanceParam   [3]int64
	CreatePosition   int
	IncreasePosition int
	RemovePosition   int
	Swap             int
	CollectFees      int
	Strategy1        int
	Strategy1Param   [3]int64
}

var sw = new(Switcher)

// const (
// 	WETH  string = "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"
// 	WBTC         = "0x2260fac5e5542a773aa44fbcfedf7c193bc2c599"
// 	CWBTC        = "0xC11b1268C1A384e55C48c2391d8d480264A3A7F4"
// 	USDC         = "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"
// 	CUSDC        = "0x39AA39c021dfbaE8faC545936693aC917d5E7563"
// 	CETH         = "0x4Ddc2D193948926D02f9B1fE9e1daa0718270ED5"
// 	DAI          = "0x6B175474E89094C44Da98b954EedeAC495271d0F"
// 	CDAI         = "0x5d3a536E4D6DbD6114cc1Ead35777bAB948E3643"
// 	USDT         = "0xdAC17F958D2ee523a2206206994597C13D831ec7"
// 	CUSDT        = "0xf650C3d88D12dB855b8bf7D11Be6C55A4e07dCC9"

// 	OSQTH             = "0xf1B99e3E573A1a9C5E6B2Ce818b617F0E664E86B"
// 	UniswapV2Router02 = "0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D"
// )

func main() {

	// project.SendTestTokens()
	// return

	// TEST NETWORK

	project.ConfigParser()
	defer project.ConfigWrite()

	project.Init(-1, -1)
	//project.Quiet = true

	// # deploy own weth/usdc  tokens
	//0x4d43654A8669f1B89ed75dEd048173F65c4047fa
	//project.GetPool("0x5B8B635c2665791cf62fe429cB149EaB42A3cEd8", "0x98a5F1520f7F7fb1e83Fe3398f9aBd151f8C65ed", 3000)
	// project.GetPriceFromAavePool()
	// return

	//project.DeployWETH()
	//project.DeployToken("x usdc 1", "xUSDC", 6, big.NewInt(500000000000000))
	// ## account 0 must have eth to wrap for weth
	// ## now the account 0 has 7 weth and 500m usdc in wallet
	//return
	//project.TokenTransfer(0, big.NewInt(1000e6), project.Network.TokenB, "0x5D2Dd8F38a74294DD1FAF7B8c806446A3379f330")
	//project.EthTransfer(project.Network.PrivateKey[0], big.NewInt(1e18), "0x5D2Dd8F38a74294DD1FAF7B8c806446A3379f330")
	//project.Wrap(project.Network.LendingContracts.WETH, 1, big.NewInt(5e17))
	//return
	//project.Deposit(30e6, 1e16, 1)
	//project.VaultInfo3(1)
	// project.GetTVL()
	// project.MyAccountInfo(1)
	//return

	//# squeeth deploy, init, mint
	//project.DeploySqueeth()
	//project.InitSqueeth()
	//project.MintSqueeth(project.FromAddress.Hex(), project.PowX(50000, 18))
	//return

	//InitialForkEnv()
	// project.GetExpectedNormalizationFactor()
	// return

	//project.ViaVaultPublicList()
	//project.ViaStratUniCompPublicList()
	//project.FactoryPublicList()
	//project.VaultInfo()
	// project.MyAccountInfo(0)
	//return

	//# wallet swap
	// project.DeploySwapHelper()
	// project.ConfigWrite()
	// project.Init(-1, -1)
	// return
	// // note: must check if swapHelper is deployed
	//project.TokenSwap(0, project.Network.LendingContracts.WETH, project.Network.LendingContracts.USDC, 500, big.NewInt(1e17))
	project.TokenSwap(0, project.Network.LendingContracts.USDC, project.Network.LendingContracts.WETH, 500, big.NewInt(280e6))
	return

	// for i := 0; i < 5; i++ {
	// 	//eth price down
	// 	project.TokenSwap(0, project.Network.LendingContracts.WETH, project.Network.LendingContracts.USDC, 500, big.NewInt(1e17))
	// 	// 	//eth price up
	// 	project.TokenSwap(0, project.Network.LendingContracts.USDC, project.Network.LendingContracts.WETH, 500, big.NewInt(280e6))
	// 	// 	//project.Sleep(5000)
	// }
	// return
	//project.TokenSwap(0, project.Network.LendingContracts.OSQTH, project.Network.LendingContracts.WETH, 3000, project.Str2BigInt("503453817677935236749"))

	//# strat router02 swap
	// project.SwapTest(project.Network.LendingContracts.WETH, project.Network.LendingContracts.OSQTH, 3000, big.NewInt(146400239374125503))
	// project.SwapTest(project.Network.LendingContracts.OSQTH, project.Network.LendingContracts.WETH, 3000, project.Str2BigInt("397754930071024025044"))
	//return

	//# strat uniswap pool swap
	//project.Deposit(30e6, 1e16, 0)
	//	project.VaultInfo3(0)

	//	project.SwapDirectPool(project.Network.LendingContracts.WETH, project.Network.LendingContracts.OSQTH, 3000, project.Str2BigInt("100000"))
	//project.SwapDirectPool(project.Network.LendingContracts.OSQTH, project.Network.LendingContracts.WETH, 3000, project.Str2BigInt("10"))
	// project.Sleep(5000)
	//	return

	//project.VaultInfo()
	//project.GetTVL()
	// fmt.Println(project.GetBalance("0xD624c05a873B9906e5F1afD9c5d6B2dC625d36c3", project.Network.VaultStrat))
	// return

	// project.DeployCallee()

	//# deploy vaultBridge
	//project.DeployVaultBridge()
	//project.ConfigWrite()
	//project.SetVaultAddress(project.Network.VaultFactory, 0) //factory
	//project.SetVaultAddress(project.Network.Vault, 1) //vault
	// //	project.SetPermit("0x23E7A3F38a8834606bCC9F5d5485aA3EBD058Efa", 1) // daniel user1
	// //	project.SetPermit("0xE1190667976b71f1c186521e50fFdAEDF722C830", 1) // collector
	// //project.SetPermit("0x8ee95fe2DB1e3f7FAACCdEd1cbCc237267EB4a00", 1) // vadmin0
	// //project.SetPermit("0x511Ed5FC53CCCf5c4239487381fcE287B02119Fa", 1) // daniel
	// project.GetVaultAddress(0)
	// project.GetVaultAddress(1)
	// project.GetVaultAddress(2)
	// project.GetPair0(project.Network.Vault)
	// return

	// # test network
	///rinkeby
	// # aave eth and usdc pool
	//project.GetPool("0x5B8B635c2665791cf62fe429cB149EaB42A3cEd8", "0x98a5F1520f7F7fb1e83Fe3398f9aBd151f8C65ed", 500)
	// project.GetPool("0x5B8B635c2665791cf62fe429cB149EaB42A3cEd8", "0x98a5F1520f7F7fb1e83Fe3398f9aBd151f8C65ed", 3000)
	// return
	//_pool := project.GetPool(project.Network.LendingContracts.WETH, project.Network.LendingContracts.OSQTH, 3000)
	// _pool := project.GetPool(project.Network.LendingContracts.WETH, project.Network.LendingContracts.USDC, 500)
	// project.PoolInfo(_pool.Hex())
	// project.GetTVL()
	//project.VaultInfo()
	//project.GetTVL()
	//return

	// project.ViaVaultPublicList()
	// project.ViaStratUniCompPublicList()
	//return

	// project.FactoryPublicList()
	// project.Deposit(1e6, 1e11, 0)
	// project.MoveFunds()

	//# sweep other tokens from vault
	// project.Sweep(project.Network.LendingContracts.OSQTH, big.NewInt(5982302))
	// return
	//project.GetTVL()
	// //	project.MyAccountInfo(0)
	//return

	//# call funds, deploy strat2, register, move funds

	//	// project.Rebalance(1500, 0) //  1500 = $600
	// // project.MyAccountInfo(1)
	// project.GetTVL()
	// project.VaultInfo3(0)
	// return
	// project.Alloc(0)    // remove positions
	// project.CallFunds() // call funds from vault
	// project.Withdraw(100, 0)
	// return
	// project.DeployStrat2ByGoStruct()
	// project.DeployVaultByGo()
	// project.ConfigWrite()
	// project.Init(-1, -1)
	// project.Register(project.Network.VaultStrat, project.Network.Vault)
	// project.ChangeStat(project.Network.VaultStrat, project.Network.Vault, 1)
	//project.Deposit(20e6, 1e16, 1)
	//project.MoveFunds()
	//project.GetTVL()
	//project.GetTotalSupply()
	project.Rebalance(1500, 0) //  1500 = $600
	//project.Rebalance(1500, 0) //  1500 = $600
	project.VaultInfo3(0)
	return

	//## deposit & withdraw
	//project.Deposit(200e6, 1e17, 0)
	//project.Deposit(100e6, 5e16, 1)
	//project.MoveFunds()
	//project.Withdraw(100, 0)

	//project.MyAccountInfo(0)
	//project.GetTVL()
	//project.VaultInfo3(1)
	//return

	//project.GetPool(project.Network.LendingContracts.WETH, project.Network.LendingContracts.OSQTH, 3000)
	//project.SwapDirectPool(project.Network.LendingContracts.WETH, project.Network.LendingContracts.OSQTH, 3000, big.NewInt(10000))
	//project.SwapDirectPool(project.Network.LendingContracts.OSQTH, project.Network.LendingContracts.WETH, 3000, big.NewInt(598210))
	// //project.MyAccountInfo(0)
	// project.VaultInfo()
	//project.Alloc(0)

	//project.GetPool(project.Network.TokenA, project.Network.TokenB, 500)
	//project.VaultInfo()
	//project.GetTVL()
	//return

	//project.GetPool(project.Network.TokenA, project.Network.TokenB, 500)
	//project.GetPool(project.Network.LendingContracts.OSQTH, project.Network.TokenB, 3000)
	// return
	//project.Alloc(0)

	// 重启geth
	// 测试withdraw 、 rebalance

	project.Rebalance(1200, 0) //  1500 = $600
	//project.Rebalance(1500, 0) //  1500 = $600

	project.VaultInfo3(0)
	//project.VaultInfo3(1)

	return

	//## check vaultinfo
	//project.SwapTest(project.Network.LendingContracts.WETH, project.Network.LendingContracts.USDC, 500, big.NewInt(1))
	//	project.SwapTest(project.Network.LendingContracts.WETH, project.Network.LendingContracts.OSQTH, 3000, big.NewInt(1e10))
	// project.SwapTest(project.Network.LendingContracts.OSQTH, project.Network.LendingContracts.WETH, 10000, big.NewInt(125347146798))
	// //	project.SwapTest(project.Network.LendingContracts.USDC, project.Network.LendingContracts.WETH, 500, big.NewInt(100e6))
	// project.Sleep(5000)

	// project.VaultInfo()
	// return

	//## FORCE TO WITHDRAW ALL FUNDS
	// project.GetTVL()
	// project.Withdraw(100, 0)
	// //project.MoveFunds()
	// project.Rebalance(600, 0) // strategy method. call alloc/removeposition/vault.movefunds/ rebalance
	// project.Alloc(0)
	// project.CallFunds()
	// project.GetTVL()
	// project.MyAccountInfo(0)
	// return

	// project.TokenSwap(0, WETH, USDC, 500, big.NewInt(2))
	// project.MyAccountInfo(0)
	// return

	///#### Forked environment  new ganache-cli ####
	//project.Deposit(0, 2e18, 0)
	//project.MoveFunds()

	//project.SwapTest(project.Network.LendingContracts.WETH, project.Network.LendingContracts.USDC, 500, big.NewInt(1e18))
	//project.SwapTest(project.Network.LendingContracts.USDC, project.Network.LendingContracts.WETH, 500, big.NewInt(100e6))

	//	project.Sleep(5000)
	//project.MyAccountInfo(0)

	//project.SwapTest(project.Network.LendingContracts.WETH, project.Network.LendingContracts.OSQTH, 3000, big.NewInt(1e15))
	//osqth to weth worked with fee tier 3000 or  10000  with fresh forked mainnet
	//project.SwapTest(project.Network.LendingContracts.OSQTH, project.Network.LendingContracts.WETH, 10000, big.NewInt(440587665102262493))
	//project.SwapTest(project.Network.LendingContracts.OSQTH, project.Network.LendingContracts.WETH, 3000, big.NewInt(4405795610439833))

	//project.SwapTest(project.Network.LendingContracts.WETH, project.Network.LendingContracts.DAI, 3000, big.NewInt(1e12))
	//project.SwapTest(project.Network.LendingContracts.OSQTH, project.Network.LendingContracts.WETH, 3000, big.NewInt(1296748104017))

	// project.GetTVL()
	// return

	//project.VaultInfo()

	//project.Rebalance(600, 0) // strategy method. call alloc/removeposition/vault.movefunds/ rebalance

	//# test squeeth
	// project.GetExpectedNormalizationFactor()
	// return

	// InitialForkEnv()
	// return

	//InitialVaults2()
	//project.DeployVaultByGo()
	//return

	///# getPrice
	//project.GetPriceStratCall()
	// project.GetPricePer(project.Network.TokenA, 500)
	// project.GetPricePer(project.Network.LendingContracts.OSQTH, 3000)
	//project.GetPricePer(project.Network.LendingContracts.USDT)
	//project.GetPricePer(project.Network.LendingContracts.OSQTH, 3000)
	//project.GetPrice()
	//project.GetPriceFromAavePool()

	//# deploy and register strat2
	// project.DeployStrat2ByGoStruct()
	// project.ConfigWrite()
	// project.Init(-1, -1)
	// project.Register(project.Network.VaultStrat, project.Network.Vault)
	// project.ChangeStat(project.Network.VaultStrat, project.Network.Vault, 1)
	// return

	// project.SwapTest(project.Network.LendingContracts.WETH, project.Network.LendingContracts.USDC, 500, big.NewInt(1e13))
	// return

	//project.FactoryPairs("0x91642Ba015065822fc4e530C222f2492B5B97420")

	//project.Wrap(project.Network.LendingContracts.WETH, 0, 1)
	// project.Sleep(5000)
	//project.TokenSwap(0, project.Network.LendingContracts.WETH, project.Network.LendingContracts.USDC, 500, big.NewInt(1))
	//project.MyAccountInfo(0)

	// project.TokenSwap(0, WETH, OSQTH, 3000, big.NewInt(2))
	//	project.TokenSwap(0, WETH, DAI, 500, big.NewInt(2))
	// project.Sleep(5000)
	//	project.Deposit(0, 0, 0)
	//project.Withdraw(100, 0)
	// project.SetTwapduration(5)
	// return
	//project.Rebalance(500, 0)
	//return
	//project.GetTwap()

	//project.MyAccountInfo(0)
	//project.GetTVL()
	//project.Rebalance(500, 0)

	// // project.Sleep(15000)

	// fmt.Println(project.GetBalance(DAI, "0xa0df350d2637096571F7A701CBc1C5fdE30dF76A"))
	// fmt.Println(project.GetBalance(USDC, "0xa0df350d2637096571F7A701CBc1C5fdE30dF76A"))
	// fmt.Println(project.GetBalance(WETH, "0xa0df350d2637096571F7A701CBc1C5fdE30dF76A"))
	// fmt.Println(project.GetBalance(OSQTH, "0xa0df350d2637096571F7A701CBc1C5fdE30dF76A"))
	// fmt.Println(project.GetBalance(project.Network.Vault, "0xa0df350d2637096571F7A701CBc1C5fdE30dF76A"))

	//	return

	// project.PoolInfo("0x921c384F79de1BAe96d6f33E3E5b8d0B2B34cb68")
	// return

	//project.GetPool(project.Network.TokenA, project.Network.TokenB, 3000)
	//return
	//project.PoolInfo("0xee815CDC6322031952a095C6cc6FEd036Cb1F70d")
	//return
	//project.Deposit(1e14, 0, 0)
	// project.TradeSqth(100000, false)
	//return

	// project.DeployVaultFactory()
	// return

	//#reload strategy
	//// newStratAddr := project.DeployStratByGo()
	// newStratAddr := project.DeployStratByGoStruct()
	// project.ConfigWrite()
	// project.Register(newStratAddr, project.Network.Vault)
	// project.ChangeStat(newStratAddr, project.Network.Vault, 1)
	// return

	//#reload vault
	// newVaultAddr := project.DeployVaultByGo()
	// project.ConfigWrite()
	// project.Register(project.Network.VaultStrat, newVaultAddr)
	// project.ChangeStat(project.Network.VaultStrat, newVaultAddr, 1)

	//#reload strategy/strategy2 & vault
	// newStratAddr := project.DeployStratByGoStruct()
	// // //newStratAddr := project.DeployStrat2ByGoStruct()
	// newVaultAddr := project.DeployVaultByGo()
	// project.ConfigWrite()
	//project.Register(project.Network.VaultStrat, project.Network.Vault)
	// project.Register(newStratAddr, newVaultAddr)
	//project.ChangeStat(newStratAddr, newVaultAddr, 1)
	//return

	//# Just Register/Register strategy & vault
	// project.Register(project.Network.VaultStrat, project.Network.Vault)
	// project.ChangeStat(project.Network.VaultStrat, project.Network.Vault, 1)
	// return

	//# emergency
	// project.CallFunds()
	// project.EmergencyWithdraw(0)
	// // //project.EmergencyWithdraw(1)
	// project.GetTVL()
	// return

	//#CheckStatus
	// project.ViewVaults()
	// return

	// project.SetTwapduration(0)
	// project.GetTVL()
	// project.GetQuoteAtTick(project.Network.Pool, 2)
	// // //project.Rebalance(400, 2)
	// // project.GetTVL()
	// return
	project.Init(-1, -1)

	// project.IsContract("0x1F98431c8aD98523631AE4a59f267346ea31F984")
	// return

	//### Events

	// blockfrom := 6178203
	// blockend := 6178262
	// project.Events("PendingWithdraw", blockfrom, blockend)
	// project.Events("Withdraw", blockfrom, blockend)
	// project.Events("Deposit", blockfrom, blockend)
	// project.Events("MintFees", blockfrom, blockend)

	// return

	// project.DeployVaultByDeployer()
	// project.DeployStratByDeployer()
	// return

	//# WEB deploy
	// project.DeployVaultDeployer()
	// project.DeployStratDeployer()
	// // return
	// project.FactoryVault()

	// project.ConfigWrite()
	// project.Init(-1, -1)

	// // return
	// s := project.Network.VaultStrat
	// v := project.Network.Vault
	// project.ChangeStat(s, v, 1)
	// project.GetStat(s, v)

	// project.CheckActive("0x2a8179A7893d00B33D2d9DBe9F0e4bBf2Cb97DE7")
	// project.Rebalance(400, 0)
	// return

	// project.GetPair0("0xcD0E0765dBCFB556089E07781cA6E7A2e1Eab055")
	//project.ViaVaultPublicList()
	//project.ViaStratUniCompPublicList()
	// project.FactoryPublicList()
	// project.ViewVaults()
	//project.CheckStatus(project.Network.Vault, 1) // CheckVaultStatus()
	//project.ViaVaultPublicList()
	//project.ViaStratUniCompPublicList()

	//return

	//	project.SetTwapduration(0)
	// project.GetPriceStratCall()
	// project.GetTwap()
	// project.GetTickPrice()
	// return

	// project.MyAccountInfo(0)

	// project.MyAccountInfo(1)

	// //### ViaVaultInfo

	//project.GetTVL()
	// project.GetTotalAmounts()
	// project.GetCompAmounts()
	//project.LendingInfo()
	//return

	// //project.DeployVaultFactory()
	// //project.DeployVaultStrategy()
	//project.ChangeStat("0xd029FDcEB5B0E971e675D7f2766188e5F8ccEeE9", "0xD06d9CF030401a72476B9e10AA47a94CE5f3798E", 2)
	// //project.CheckVaultStatus()
	//return

	// project.Deposit(1e16, 1e6, 0)
	// //project.Withdraw(100, 0)
	// project.Rebalance(500, 0)
	// // return
	// //project.EmergencyCall() //by admin
	// project.CallFunds() // by admin, check to make sure all funds are back to vaults.
	// // // //project.ChangeStat(4)		//by admin to change the status to abandoned and only with draw allowed
	// // project.EmergencyWithdraw(0)
	// //project.EmergencyWithdraw(1)
	// project.Withdraw(100, 0)
	// // project.Withdraw(3)
	// return

	// s := "0x9a94272446f0c119E1006935c9E6D6fEB6c206f4"
	// v := "0x6E09167c444AAbe5cD49Cff5Af16B15E33096e6C"
	//project.FactoryVault()

	//project.Init(-1, -1)
	//project.DeployStratByDeployer()
	//	project.DeployVaultStrategy()
	//project.DeployViaVault()
	// //	project.DeployStratUniComp()
	//project.Sleep(5000)

	//project.SetPortionRatio(90, 100)
	//project.SetTwapduration(10)
	//return

	project.Init(-1, -1)
	project.Quiet = false

	///0
	if project.Networkid == 0 {
		//#### WETH/USDC Test
		project.Deposit(4000e6, 1e18, 0)
		//project.Withdraw(100, 0)
		//	project.Deposit(1e16, 1e5, 0)
		//	project.Deposit(2e16, 2e5, 1)

		//project.Withdraw(100, 0)
		// // project.Deposit(2e17, 2e18, 1)
		// //project.Withdraw(100, 0)
		project.GetTVL()
		// return
		// // project.GetTotalAmounts()
		// // project.GetCompAmounts()

		// //project.MoveFunds()
		project.Rebalance(400, 2) // strategy method. call alloc/removeposition/vault.movefunds/ rebalance

		project.GetTVL()
		// // project.GetTotalAmounts()
		// // project.GetCompAmounts()
		// //project.LendingInfo()

		// // //project.Rebalance(400, 2) // strategy method. call alloc/removeposition/vault.movefunds/ rebalance
		// // //project.Rebalance(600, 2) // strategy method. call alloc/removeposition/vault.movefunds/ rebalance
		// // project.Deposit(1e17, 1e6, 0)
		// // // //project.Alloc(0)
		// // //project.Rebalance(400, 0) // strategy method. call alloc/removeposition/vault.movefunds/ rebalance
		// // project.Rebalance(600, 0) // strategy method. call alloc/removeposition/vault.movefunds/ rebalance

		// // project.Deposit(2e17, 2e6, 1)
		// // project.Rebalance(600, 0) // strategy method. call alloc/removeposition/vault.movefunds/ rebalance

		// // // project.EmergencyBurn() // vault calls strategy.callFunds...alloc/removepositions/transferfunds
		// // // project.EmergencyWithdraw(0)
		// // // project.EmergencyWithdraw(1)

		project.Withdraw(100, 0)
		// project.Withdraw(100, 1)
		// project.Withdraw(100, 2)
		// project.Withdraw(100, 3)
		project.Rebalance(400, 0) // strategy method. call alloc/removeposition/vault.movefunds/ rebalance

		// // // //project.VaultInfo()
		// // // project.MyAccountInfo(0)
		// // // project.MyAccountInfo(1)
		project.GetTVL()
		// // project.GetTotalSupply()
		// // project.GetTotalAmounts()
		// // project.GetCompAmounts()

		return

	} else if project.Networkid == 3 {
		///3
		//#### WETH/USDC Test
		project.Deposit(1e14, 1e5, 0)
		//project.Withdraw(100, 0)
		//	project.Deposit(1e16, 1e5, 0)
		//	project.Deposit(2e16, 2e5, 1)

		//project.Withdraw(100, 0)
		// // project.Deposit(2e17, 2e18, 1)
		// //project.Withdraw(100, 0)
		project.GetTVL()
		// return
		// // project.GetTotalAmounts()
		// // project.GetCompAmounts()

		// //project.MoveFunds()
		project.Rebalance(400, 2) // strategy method. call alloc/removeposition/vault.movefunds/ rebalance

		project.GetTVL()
		// // project.GetTotalAmounts()
		// // project.GetCompAmounts()
		// //project.LendingInfo()

		// // //project.Rebalance(400, 2) // strategy method. call alloc/removeposition/vault.movefunds/ rebalance
		// // //project.Rebalance(600, 2) // strategy method. call alloc/removeposition/vault.movefunds/ rebalance
		// // project.Deposit(1e17, 1e6, 0)
		// // // //project.Alloc(0)
		// // //project.Rebalance(400, 0) // strategy method. call alloc/removeposition/vault.movefunds/ rebalance
		// // project.Rebalance(600, 0) // strategy method. call alloc/removeposition/vault.movefunds/ rebalance

		// // project.Deposit(2e17, 2e6, 1)
		// // project.Rebalance(600, 0) // strategy method. call alloc/removeposition/vault.movefunds/ rebalance

		// // // project.EmergencyBurn() // vault calls strategy.callFunds...alloc/removepositions/transferfunds
		// // // project.EmergencyWithdraw(0)
		// // // project.EmergencyWithdraw(1)

		project.Withdraw(100, 0)
		// project.Withdraw(100, 1)
		// project.Withdraw(100, 2)
		// project.Withdraw(100, 3)
		project.Rebalance(400, 0) // strategy method. call alloc/removeposition/vault.movefunds/ rebalance

		// // // //project.VaultInfo()
		// // // project.MyAccountInfo(0)
		// // // project.MyAccountInfo(1)
		project.GetTVL()
		// // project.GetTotalSupply()
		// // project.GetTotalAmounts()
		// // project.GetCompAmounts()

		return

	} else if project.Networkid == 4 {
		///4
		//#### WETH/DAI Test
		//project.SetTwapduration(10)
		//project.Deposit(1e17, 1e18, 0)
		//project.Deposit(2e17, 2e18, 1)
		// project.Withdraw(100, 0)
		// project.Withdraw(100, 1)
		// project.Deposit(2e17, 2e18, 1)
		// //project.Withdraw(100, 0)
		//project.GetTVL()
		// project.GetTotalAmounts()
		// project.GetCompAmounts()

		//project.MoveFunds()
		project.Rebalance(400, 2) // strategy method. call alloc/removeposition/vault.movefunds/ rebalance
		//project.Withdraw(50, 0)

		project.GetTVL()
		// project.GetTotalAmounts()
		// project.GetCompAmounts()

		//project.Rebalance(400, 2) // strategy method. call alloc/removeposition/vault.movefunds/ rebalance
		// //project.Rebalance(600, 2) // strategy method. call alloc/removeposition/vault.movefunds/ rebalance
		//project.Deposit(1e17, 1e18, 0)
		// // //project.Alloc(0)
		// //project.Rebalance(400, 0) // strategy method. call alloc/removeposition/vault.movefunds/ rebalance
		// project.Rebalance(600, 0) // strategy method. call alloc/removeposition/vault.movefunds/ rebalance

		// project.Deposit(2e17, 2e18, 1)
		// project.Rebalance(600, 0) // strategy method. call alloc/removeposition/vault.movefunds/ rebalance

		// // project.EmergencyWithdraw(0)
		// // project.EmergencyWithdraw(1)

		// project.Withdraw(100, 0)
		// project.Withdraw(100, 1)
		// project.Withdraw(100, 2)
		// project.Withdraw(100, 3)
		// project.Rebalance(400, 0) // strategy method. call alloc/removeposition/vault.movefunds/ rebalance

		// // //project.VaultInfo()
		// // project.MyAccountInfo(0)
		// // project.MyAccountInfo(1)
		//	project.GetTVL()
		// project.GetTotalSupply()
		// project.GetTotalAmounts()
		// project.GetCompAmounts()

		return
	} else if project.Networkid == 6 {
		///6 rinkeby
		//#### WETH/USDC Test
		//project.GetPriceStratCall()
		//return
		project.Deposit(3e6, 1e14, 0)
		//project.CallFunds()
		//project.GetTVL()
		//return
		//project.Withdraw(100, 0)
		//project.Rebalance(400, 0) // strategy method. call alloc/removeposition/vault.movefunds/ rebalance
		//return
		//	project.Deposit(1e16, 1e5, 0)
		//	project.Deposit(2e16, 2e5, 1)

		project.GetTVL()
		// return
		// // project.GetTotalAmounts()
		// // project.GetCompAmounts()

		project.MoveFunds()
		//project.Rebalance(400, 0) // strategy method. call alloc/removeposition/vault.movefunds/ rebalance

		//project.GetTVL()
		// // project.GetTotalAmounts()
		// // project.GetCompAmounts()
		// //project.LendingInfo()

		// // //project.Rebalance(400, 2) // strategy method. call alloc/removeposition/vault.movefunds/ rebalance
		// // //project.Rebalance(600, 2) // strategy method. call alloc/removeposition/vault.movefunds/ rebalance
		// // project.Deposit(1e17, 1e6, 0)
		// // // //project.Alloc(0)
		// // //project.Rebalance(400, 0) // strategy method. call alloc/removeposition/vault.movefunds/ rebalance
		// // project.Rebalance(600, 0) // strategy method. call alloc/removeposition/vault.movefunds/ rebalance

		// // project.Deposit(2e17, 2e6, 1)
		// // project.Rebalance(600, 0) // strategy method. call alloc/removeposition/vault.movefunds/ rebalance

		// // // project.EmergencyBurn() // vault calls strategy.callFunds...alloc/removepositions/transferfunds
		// // // project.EmergencyWithdraw(0)
		// // // project.EmergencyWithdraw(1)

		//project.Withdraw(100, 0)
		// project.Withdraw(100, 1)
		// project.Withdraw(100, 2)
		// project.Withdraw(100, 3)
		//project.Rebalance(400, 0) // strategy method. call alloc/removeposition/vault.movefunds/ rebalance

		// // // //project.VaultInfo()
		// // // project.MyAccountInfo(0)
		// // // project.MyAccountInfo(1)
		//project.GetTVL()
		// // project.GetTotalSupply()
		// // project.GetTotalAmounts()
		// // project.GetCompAmounts()

		return

	} else if project.Networkid == 7 {
		///7

		project.Init(-1, -1)
		project.Quiet = false

		//#### USDC/WETH Test
		//project.SetTwapduration(10)
		//project.Deposit(1e6, 1e7, 0)
		//project.Deposit(2e17, 2e18, 1)
		project.Withdraw(100, 0)
		// project.Withdraw(100, 1)
		// project.Deposit(2e17, 2e18, 1)
		// //project.Withdraw(100, 0)
		project.GetTVL()
		// project.GetTotalAmounts()
		// project.GetCompAmounts()

		//project.MoveFunds()
		//project.Rebalance(400, 2) // strategy method. call alloc/removeposition/vault.movefunds/ rebalance
		//project.Withdraw(50, 0)

		//project.GetTVL()
		// project.GetTotalAmounts()
		// project.GetCompAmounts()

		//project.Rebalance(400, 2) // strategy method. call alloc/removeposition/vault.movefunds/ rebalance
		// //project.Rebalance(600, 2) // strategy method. call alloc/removeposition/vault.movefunds/ rebalance
		//project.Deposit(1e17, 1e18, 0)
		// // //project.Alloc(0)
		// //project.Rebalance(400, 0) // strategy method. call alloc/removeposition/vault.movefunds/ rebalance
		// project.Rebalance(600, 0) // strategy method. call alloc/removeposition/vault.movefunds/ rebalance

		// project.Deposit(2e17, 2e18, 1)
		// project.Rebalance(600, 0) // strategy method. call alloc/removeposition/vault.movefunds/ rebalance

		// // project.EmergencyWithdraw(0)
		// // project.EmergencyWithdraw(1)

		// project.Withdraw(100, 0)
		// project.Withdraw(100, 1)
		// project.Withdraw(100, 2)
		// project.Withdraw(100, 3)
		// project.Rebalance(400, 0) // strategy method. call alloc/removeposition/vault.movefunds/ rebalance

		// // //project.VaultInfo()
		project.MyAccountInfo(0)
		// // project.MyAccountInfo(1)
		//	project.GetTVL()
		// project.GetTotalSupply()
		// project.GetTotalAmounts()
		// project.GetCompAmounts()

		return
	}
	// fmt.Println(project.Network.ViaFactory)
	// fmt.Println(project.Network.VaultStrat)
	// fmt.Println(project.Network.Vault)

	// project.Network.ViaFactory = os.Getenv("ViaFactory")
	// project.Network.VaultStrat = os.Getenv("VaultStrat")
	// project.Network.Vault = os.Getenv("Vault")
	// project.Quiet = false

	//project.GetViaVaultAddress()

	//project.DeployFactory()
	//BuildAll()

	//balance := project.EthBalance(Network.LendingContracts.CETH)
	//balance := project.EthBalanceArb("0x4Ddc2D193948926D02f9B1fE9e1daa0718270ED5")
	//project.BlockNumber()

	//project.TransferEth("1b280901929b5cd52f362b544072b66bfe29a9396db485a23da7de9f485512b0", project.X1E18(1), "0x2EE910a84E27aCa4679a3C2C465DCAAe6c47cB1E")
	// project.VaultInfo()
	// return
	// mainnet fork: weth/usdc

	//project.FindPool()
	//project.GetPool("0xc18360217d8f7ab5e7c516566761ea12ce7f9d72", "0xdac17f958d2ee523a2206206994597c13d831ec7", 3000)
	//return
	// //project.EthBalance("0xB7A41b27af9Ed23F65E36f9d92d287327c4D997d")
	// //return
	// project.DeployCallee()
	// project.DeployVialendFeemaker(-1, 0, big.NewInt(10), 50, big.NewInt(1e6), "0x5ACb5DB941E3Fc33E0c0BC80B90114b6CD0249B5")
	// project.EthBalance("0xa0df350d2637096571F7A701CBc1C5fdE30dF76A")
	// project.Wrap(project.Network.TokenB, 0, 30)
	// project.EthBalance("0xa0df350d2637096571F7A701CBc1C5fdE30dF76A")
	//return
	// project.EthBalance("0xEC2DD0d0b15D494a58653427246DC076281C377a")
	// project.Swap(0, 10000, 2, 1)
	// project.ERC20Balance(project.Network.TokenA, "0xa0df350d2637096571F7A701CBc1C5fdE30dF76A")
	// project.ERC20Balance(project.Network.TokenB, "0xa0df350d2637096571F7A701CBc1C5fdE30dF76A")
	// return

	// project.Deposit(1e17, 1e18, 0)
	// project.Deposit(2e17, 2e18, 1)
	// // project.VaultInfo()
	// project.MoveFunds()
	// project.Rebalance(400, 0)
	// //project.Alloc(0)
	// //project.EmergencyBurn()
	// project.Withdraw(100, 1)
	// project.Withdraw(100, 0)
	// //project.VaultInfo()
	// return

	// project.Deposit(1, [3]int64{4000, 10, 0}, false)
	// project.Deposit(1, [3]int64{4000, 10, 1}, false)
	// project.VaultInfo()
	//project.Strategy1(800, 0)
	//project.VaultInfo()

	// project.Withdraw(1, [2]int64{100, 0})
	// project.Withdraw(1, [2]int64{100, 1})

	// project.MyAccountInfo(1)
	// project.VaultInfo()

	// project.CheckFees()

	//project.SetProtocolFee(big.NewInt(10))
	//project.SetUniswapPortionRatio(50)

	// project.EmergencyBurn()

	// return

	// goerli test data
	// project.DeployVialendFeemaker(-1, 0, big.NewInt(10), 100, big.NewInt(1e18), "0xEa24c7256ab5c61b4dC1c5cB600A3D0bE826a440")
	// project.Sleep(1000)

	// project.Deposit(1, [3]int64{3, 3, 0}, false)
	// project.Sleep(1000)

	// project.Deposit(1, [3]int64{2, 2, 1}, false)
	// project.Sleep(1000)

	// // // //	 project.Deposit(1, [3]int64{4000, 10, 1}, false)
	// // // project.VaultInfo()
	//project.Strategy1(800, 0)
	//project.Alloc(0)
	// project.Sleep(5000)
	// project.VaultInfo()
	// project.CheckFees()
	// project.GetLendingAmounts(project.Network.Vault)
	//return
	//project.SetProtocolFee(big.NewInt(0))
	//project.WithdrawPending(100, 0)
	//project.Sleep(1000)
	// project.WithdrawPending(0, 1)
	// project.Sleep(1000)

	//project.Strategy1(800, 0)
	// project.Sleep(1000)
	// project.Withdraw(1, [2]int64{100, 1})

	//project.VaultInfo()

	// project.MyAccountInfo(0)
	// project.Sleep(1000)

	// project.MyAccountInfo(1)
	// project.Sleep(1000)

	//project.Withdraw(1, [2]int64{100, 0})
	// project.Withdraw(1, [2]int64{100, 1})

	// project.MyAccountInfo(0)
	// project.Sleep(1000)

	//project.Withdraw(1, [2]int64{100, 1})

	//	project.MyAccountInfo(1)
	//project.VaultInfo()
	// project.CheckFees()

	return

	//networkid, account, protocolfee, uniportion, team address to get fee cut
	//project.DeployVialendFeemaker(3, 1, big.NewInt(10), 90, "0xEa24c7256ab5c61b4dC1c5cB600A3D0bE826a440")

	//project.DeployArb()
	// project.DeployCallee()
	// return

	//project.SetVaultAddress("0xBC6b6e273171C428d85cDdB23D344a8400B48441", 2)
	// return

	// project.GetPool("0xC04B0d3107736C32e19F1c62b2aF67BE61d63a05", "0xdc31Ee1784292379Fbb2964b3B9C4124D8F89C60", 10000)

	// return

	//project.FindPool()
	// v := "0xb102Cd93329d7017Ae83C6E488f00EaB4844CbF2"
	// t := "0x20572e4c090f15667cF7378e16FaD2eA0e2f3EfF"
	// v := "0xb102Cd93329d7017Ae83C6E488f00EaB4844CbF2"
	// t := "0xfa5df5372c03d4968d128d624e3afeb61031a777"
	// a := big.NewInt(1e18)
	// project.Sweep(v, t, a)

	// project.EmergencyBurn()
	//project.VaultInfo2("0xb102Cd93329d7017Ae83C6E488f00EaB4844CbF2")
	//return

	//nid := int(0)
	//acc := int(0)

	//project.PrintPrice()
	//project.Wrap(project.Network.TokenA, 0, 10)

	//project.DeployVialendFeemaker(-1, 1, big.NewInt(10), 100, big.NewInt(1e8), "0xEa24c7256ab5c61b4dC1c5cB600A3D0bE826a440")
	//project.DeployVialendFeemaker(-1, 1, big.NewInt(10), 100, big.NewInt(1e18), "0xEa24c7256ab5c61b4dC1c5cB600A3D0bE826a440")
	//project.Deposit(1, [3]int64{1, 10, 0}, false)
	//project.Strategy1(800, 1)
	//project.VaultInfo()
	//project.Withdraw(1, [2]int64{100, 0})
	//project.Withdraw(1, [2]int64{100, 1})
	//project.Withdraw(1, [2]int64{100, 3})

	//project.EmergencyBurn()
	project.SetProtocolFee(big.NewInt(0))
	//project.Withdraw(1, [2]int64{100, 3})
	//project.VaultInfo2("0x4aaE0bc3052aD3AB125Ae654f0f2C55Dbd9D6e17")
	//project.MyAccountInfo(0)
	//project.PoolInfo()
	project.VaultInfo()
	return
	// // // newVault()

	// project.GetCapital(1)
	// project.GetCapital(3)
	// project.LendingInfo()
	// // project.AccountInfo()
	// project.VaultInfo()
	// // project.PoolInfo()
	// //project.FindPool()
	// // project.GetPool("0xB4FBF271143F4FBf7B91A5ded31805e42b2208d6", "0xdc31Ee1784292379Fbb2964b3B9C4124D8F89C60", 500)
	// return
	//project.Test_weth_deposit("0xB4FBF271143F4FBf7B91A5ded31805e42b2208d6", 5, 15) // weth address, accountid, amount
	//project.Test_weth_withdraw("0xB4FBF271143F4FBf7B91A5ded31805e42b2208d6", 3, 15)

	//check token decimals and info
	// fmt.Println(project.GetTokenInstance("0xC04B0d3107736C32e19F1c62b2aF67BE61d63a05"))
	//return

	//	project.Withdraw(1, [2]int64{100, 4}) // team withdraw
	//project.Deposit(1, [3]int64{1, 10, 1}, false)
	//	project.Deposit(1, [3]int64{2, 1000, 0}, false)
	//project.EmergencyBurn()

	//	project.Strategy1(1000, 1)
	// project.Strategy1(100, 1)
	//project.AccountInfo()
	// project.Withdraw(1, [2]int64{100, 0})
	//project.Withdraw(1, [2]int64{100, 1})
	//project.Withdraw(1, [2]int64{100, 3})
	//	project.Alloc(1)  // removed public as internal method
	//project.RemoveCTokens()
	project.VaultInfo()
	return

	// // // project.Strategy1( [3]int64{400, 60, 0})

	// project.AccountInfo()
	//project.Strategy1(500, 1)
	// project.Withdraw(1, [2]int64{100, 0})
	// project.Withdraw(1, [2]int64{100, 1})
	// project.Withdraw(1, [2]int64{100, 3})
	//project.VaultInfo()

	// redeemMyCtoken()
	// return

	//project.Withdraw(1, [2]int64{100, 0})
	//project.Withdraw(1, [2]int64{100, 1})
	//project.Withdraw(1, [2]int64{100, 3})
	//return

	//project.Deposit(1, [3]int64{1, 10, 0}, false)
	// project.Deposit(1, [3]int64{0, 500, 1}, false)
	// return

	// project.AccountInfo()
	// project.VaultInfo()
	// project.Withdraw(1, [2]int64{100, 0}) // team withdraw
	//project.Deposit(1, [3]int64{2, 0, 0}, true)
	//time.Sleep(10 * time.Second)
	// project.AccountInfo()

	project.Strategy1(500, 0)

	project.VaultInfo()

	project.GenFees(4, 2) // swap times, swap account, amount , sleepSeconds

	project.Strategy1(500, 0)

	project.VaultInfo()

	return

	// //project.genFees(4, 5)

	// project.Alloc(0)
	// project.AccountInfo()
	// project.VaultInfo()
	project.Strategy1(500, 1)
	project.Strategy1(200, 1)
	//project.Withdraw(1, [2]int64{100, 0})
	// project.Withdraw(1, [2]int64{100, 1})
	//project.Withdraw(1, [2]int64{100, 3})
	//project.Deposit(1, [3]int64{0, -1, 1}, false)
	project.AccountInfo()
	project.VaultInfo()
	return

}
func resetStrat() {

	project.Alloc(0)    // remove positions
	project.CallFunds() // call funds from vault
	project.GetTVL()    //
	// // //project.VaultInfo()
	//return
	project.DeployStrat2ByGoStruct()
	//	project.DeployVaultByGo()
	project.ConfigWrite()
	project.Init(-1, -1)
	project.Register(project.Network.VaultStrat, project.Network.Vault)
	project.ChangeStat(project.Network.VaultStrat, project.Network.Vault, 1)

	//project.Deposit(1000e6, 1e17, 0)
	//project.MoveFunds()
	//project.SetVaultAddress(project.Network.Vault, 1) //reg vault in bridge

}
func newVault() {

	networkId := 3
	acc := 0
	token0 := project.Network.TokenA
	token1 := project.Network.TokenB

	feetier := int64(10000) //10000, 3000, 500 //Network.FeeTier

	_protocolfee := big.NewInt(10)
	_quoteAmount := big.NewInt(1e18)
	_uniPortion := 20
	team := "0xEa24c7256ab5c61b4dC1c5cB600A3D0bE826a440"
	strategy := project.Indi(0)

	project.VaultGen(networkId, acc, token0, token1, feetier, _protocolfee, _quoteAmount, _uniPortion, team, strategy)

}

func setVaultBridge() {

	// vault bridge v1 on goreli 0x033F3C5eAd18496BA462783fe9396CFE751a2342
	// vault bridge v2.x on goreli 0x428EeA0B87f8E0f5653155057f58aaaBb667A3ec

	// project.SetPermit("0x8ee95fe2DB1e3f7FAACCdEd1cbCc237267EB4a00", 1) //

	//	project.DeployVaultBridge()

	//project.DeployVaultAdmin()
	// project.AuthAdmin(Network.VaultAdmin, "0xfd8a5AE495Df1CA34F90572cb99A33B27173eDe1")
	// return

	// project.SetVaultAddress("0xBDa573F33c18c69Cda004d5035e44Dd4635f69d1", 0) //factory
	//project.SetVaultAddress("0xcD0E0765dBCFB556089E07781cA6E7A2e1Eab055", 1) // weth/usdc
	// project.SetVaultAddress("0x45aE8C6868F068d2e4AC774106aA86C8489E6E60", 2) // weth/dai

	// project.SetPermit("0xb6F0049e37D32dED0ED2FAEeE7b69930FA49A879", 1) //
	// project.SetPermit("0x8ee95fe2DB1e3f7FAACCdEd1cbCc237267EB4a00", 1) //
	project.SetPermit("0x511Ed5FC53CCCf5c4239487381fcE287B02119Fa", 1) //

	// project.GetVaultAddress(0)
	// project.GetVaultAddress(1)

}

func BuildAll() {

	sw.ViewOnly = false

	sw.DeployFactory = 0
	//...manual step... update the new factory addres in networks.go

	sw.DeployToken = 1
	sw.TokenParam[0] = project.TokenStruct{"f weth 1", "fWETH1", 18, big.NewInt(50000000000000)}
	sw.TokenParam[1] = project.TokenStruct{"f usdc 1", "fUSDC1", 6, big.NewInt(500000000000000000)}
	//...manual step... update the new token addres in networks.go

	sw.TransferToken = 0
	sw.TransferParam[0] = TransferStruct{0, project.X1E18(900), project.Network.TokenA, "0xeBb29c07455113c30810Addc123D0D7Cd8637aea"}
	sw.TransferParam[1] = TransferStruct{0, project.X1E18(900), project.Network.TokenB, "0xeBb29c07455113c30810Addc123D0D7Cd8637aea"}

	sw.CreatePool = 1 // *Note: if token0+token1+fee = pool exists ERROR: createPool VM Exception while processing transaction: revert
	sw.InitialPool = 1
	//...manual step... update the new pool addres in networks.go

	sw.MintPool = 1
	sw.MintPoolParam = [3]int64{1000, 1, 1} // currently hardcoded 1000 * 1e18 as the liquidity,
	// .... manual step.... new vault may apply

	sw.DeployVault = 1
	//...manual step... update the new vault addres in networks.go

	sw.Deposit = 0
	sw.DepositParam = [3]int64{1, 1, 1} // {amount0, amount1 , account}

	sw.Strategy1 = 0
	sw.Strategy1Param = [3]int64{600, 60, 3} // {tick range, tickspacing, account}

	sw.Withdraw = 0
	sw.WithDrawParam = [2]int64{100, 1} // { percent, account}
	// accountid,  amount of shares in percentage %

	sw.Rebalance = 0
	sw.RebalanceParam = [3]int64{10, 60, 3} //[2]int64{22000, 60} // 12000,60   {tick range , tickspacing, account}

	sw.Swap = 0

	// 1: single swap, 2: multiple swaps
	// swapAmount, _ := new(big.Int).SetString("85175185371092425157", 10) // 85 * 1e18
	// zeroForOne := false
	//swapAmount, _ := new(big.Int).SetString("139190665697301284354", 10) // 139 * 1e18
	//zeroForOne := true

	sw.CollectFees = 0

	sw.CreatePosition = -1
	sw.IncreasePosition = -1
	sw.RemovePosition = -1

	if sw.DeployFactory > 0 {
		project.DeployFactory() ///new factory  for crating new pools if old pool already exists
	}

	if sw.DeployToken > 0 {

		token0 := project.DeployToken(sw.TokenParam[0].Name, sw.TokenParam[0].Symbol, sw.TokenParam[0].Decimals, sw.TokenParam[0].MaxTotalSupply)
		token1 := project.DeployToken(sw.TokenParam[1].Name, sw.TokenParam[1].Symbol, sw.TokenParam[1].Decimals, sw.TokenParam[1].MaxTotalSupply)

		//always make token0 = weth = tokenA
		project.Network.TokenA = token0
		project.Network.TokenB = token1

	}

	if sw.TransferToken > 0 {

		project.TokenTransfer(
			sw.TransferParam[0].AccountId,
			sw.TransferParam[0].Amount,
			sw.TransferParam[0].TokenAddress,
			sw.TransferParam[0].ToAddress)

		project.TokenTransfer(
			sw.TransferParam[1].AccountId,
			sw.TransferParam[1].Amount,
			sw.TransferParam[1].TokenAddress,
			sw.TransferParam[1].ToAddress)

	}

	if sw.CreatePool > 0 {
		project.CreatePool(sw.CreatePool) /// need to edit networks to setup address of token0 and token1,  fee tier
	}
	if sw.InitialPool > 0 {
		project.InitialPool(sw.InitialPool)
	}

	if sw.MintPool > 0 {
		project.MintPool(sw.MintPoolParam[0], sw.MintPoolParam[1], sw.MintPoolParam[2])
		//os.Exit(0)
	}

	if sw.DeployVault > 0 {
		project.DeployVault() /// deployed by test admin 2, edit networks. token0, token1, fee to get the pool address
	}

	//	project.Deposit(sw.DepositParam, false) /// deposit token0 amount * 1e18, token1 amount * 1e6

	//	project.Rebalance(sw.Rebalance, sw.RebalanceParam) /// make sure Account = 0

	//	project.Withdraw(sw.Withdraw, sw.WithDrawParam) /// withdraw shares, input number in percentage %

	// print all deployed addresses
	for _, i := range project.InfoString {
		fmt.Println(i)
	}

	project.CheckPrice(false, 3)
	project.Equation(false, false)

}

func InitialVaults2() {

	//project.DeployCallee()
	//	project.Sleep(9000)
	// // // // return

	project.DeployVaultFactory()
	// // return
	//project.Sleep(9000)
	//newStratAddr := project.DeployStratByGoStruct()
	newStratAddr := project.DeployStrat2ByGoStruct()
	//project.Sleep(9000)

	newVaultAddr := project.DeployVaultByGo()
	// //		project.DeployVaultByGo()
	//project.Sleep(9000)

	project.ConfigWrite()
	// project.Sleep(25000)
	_, _ = newStratAddr, newVaultAddr
	project.Register(newStratAddr, newVaultAddr)
	//project.Sleep(5000)
	project.ChangeStat(newStratAddr, newVaultAddr, 1)
	//project.Sleep(5000)
	// // project.Register(project.Network.VaultStrat, project.Network.Vault)
	// // return
	// // project.ChangeStat(project.Network.VaultStrat, project.Network.Vault, 1)
	// // return
}

func InitialForkEnv() {

	project.DeployCallee()
	project.Sleep(9000)
	// // // // return
	project.DeployVaultBridge()
	project.Sleep(9000)

	project.DeployVaultFactory()
	// // return
	project.Sleep(9000)
	// project.DeployStratByGoStruct()
	project.DeployStrat2ByGoStruct()
	project.Sleep(9000)

	project.DeployVaultByGo()
	project.Sleep(9000)

	project.ConfigWrite()
	project.Init(-1, -1)

	project.Register(project.Network.VaultStrat, project.Network.Vault)
	project.Sleep(5000)
	project.ChangeStat(project.Network.VaultStrat, project.Network.Vault, 1)
	project.Sleep(5000)

	//project.ChangeStat(project.Network.VaultStrat, project.Network.Vault, 1)

	//	InitVaultBridge()
	project.Sleep(5000)

	project.Wrap(project.Network.LendingContracts.WETH, 0, big.NewInt(2e18))
	project.Sleep(5000)
	//return

	// swap within wallet , error when forking infura mainnet,  ERROR pool instance:Post "http://127.0.0.1:8545": EOF
	// project.TokenSwap(0, project.Network.LendingContracts.WETH, project.Network.LendingContracts.OSQTH, 3000, big.NewInt(2))
	// project.Sleep(5000)
	// return

	project.Deposit(0, 2e18, 0)
	project.Sleep(5000)

	project.MoveFunds() // move funds to stratgy
	project.Sleep(9000)

	project.SwapTest(project.Network.LendingContracts.WETH, project.Network.LendingContracts.USDC, 500, big.NewInt(1e17))
	project.Sleep(9000)
	project.SwapTest(project.Network.LendingContracts.WETH, project.Network.LendingContracts.OSQTH, 3000, big.NewInt(1e17))
	project.Sleep(9000)
	// project.SwapTest(project.Network.LendingContracts.USDC, project.Network.LendingContracts.WETH, 500, big.NewInt(1000e6))
	// project.Sleep(9000)
	project.MyAccountInfo(0)

	//	return

	// //project.TokenSwap(0, WETH, DAI, 500, big.NewInt(2))
	// project.Sleep(5000)
	// //project.TokenSwap(0, WETH, OSQTH, 3000, big.NewInt(2))
	//	project.Sleep(5000)

	//FOR USDC/WETH
	//project.Deposit(4000e6, 1e18, 0)

	//FOR DAI/WETH
	//project.Deposit(1e18, 1e18, 0)

	//project.Sleep(5000)

	//project.MoveFunds() // move funds to strategy
	project.Rebalance(400, 0) // strategy method. call alloc/removeposition/vault.movefunds/ rebalance

	project.Sleep(5000)

	project.GetTVL()

	// ////project.TokenSwap(0, project.Network.LendingContracts.WETH, project.Network.LendingContracts.USDC, big.NewInt(1))
}
