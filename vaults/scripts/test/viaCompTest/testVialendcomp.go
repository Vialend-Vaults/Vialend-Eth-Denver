package main

import (
	"fmt"
	"log"
	"math/big"

	ctoken "viaroot/deploy/cErc20"

	vialend "viaroot/deploy/vialendcomp"
	project "viaroot/scripts/project"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

var DAI = project.Network.LendingContracts.DAI
var cDAI = project.Network.LendingContracts.CDAI
var cETH = project.Network.LendingContracts.CETH

var vialendContract = "0x1A39EB1e350053C85e478A1AA770079bBe561f02"
var to = "0x14792757D21e54453179376c849662dE341797F2"

func main() {

	fmt.Println("Env: NetworkId=", project.Networkid, ",EthClient=", project.Network.ProviderUrl[project.ProviderSortId])

	//wrapEth()
	//unWrapEth()

	//checkEvent(5560865, 5560865)

	//checkBalance(cDAI, vialendContract)

	//mintDai(project.X1E18(100))
	// -- send DAI 100 * 1e18
	// -- receive cDai 4,709.19708954  =>  470919708954

	// check balance to get the cDai OR DAI amount
	amount := checkBalance("cDAI", cDAI, vialendContract) //big.NewInt(470919937342)

	// redeem X percentage
	x := big.NewInt(90)
	amount.Mul(amount, x).Div(amount, big.NewInt(100))
	fmt.Println(amount)

	// receive 50.000000200814806476 Dai (DAI)
	// sent 2,354.59854477 cDAIai

	// true if amount is cDAI, false: if amount is DAI
	redeemType := true

	//	redeemDai(amount,redeemType)
	//--receive Dai 100.000047814991476604
	//--sent cDai   4,709.19937342

	//mintEth(project.X1E18(1))

	//amount = checkBalance("DAI",DAI, vialendContract)
	//withDrawERC20(DAI, amount)

	Ethamount := checkETHBalance()
	cEthamount := checkBalance("cETH", cETH, vialendContract)

	//redeemEth(cEthamount, true)

	//withDrawETH(Ethamount)

	_ = amount
	_ = redeemType
	_ = cEthamount
	_ = Ethamount
}

func wrapEth() {

	instance := GetVialendInstance()
	ethAmount := (project.X1E18(1))

	project.Auth.Value = ethAmount
	tx, err := instance.Wrap(project.Auth)

	if err != nil {
		log.Fatal("unwrap instance err ", err)
	}

	fmt.Println("UNWRAP ETH tx: ", tx.Hash().Hex())

}

func unWrapEth() {
	instance := GetVialendInstance()
	ethAmount := (project.X1E18(1))

	tx, err := instance.Unwrap(project.Auth, ethAmount)

	if err != nil {
		log.Fatal("unwrap instance err ", err)
	}

	fmt.Println("UNWRAP ETH tx: ", tx.Hash().Hex())

}

func mintEth(_numEthToSupply *big.Int) {

	instance := GetVialendInstance()

	project.Auth.Value = _numEthToSupply
	tx, err := instance.SupplyEthToCompound(project.Auth,
		common.HexToAddress(cETH))

	if err != nil {
		log.Fatal("mintEth err ", err)
	}

	//get the transaction hash
	fmt.Println("mintEth tx: ", tx.Hash().Hex())

}

func mintDai(_numTokensToSupply *big.Int) {
	instance := GetVialendInstance()

	_erc20Contract := common.HexToAddress(DAI)
	_cErc20Contract := common.HexToAddress(cDAI)

	tx, err := instance.SupplyErc20ToCompound(project.Auth,
		_erc20Contract,
		_cErc20Contract,
		_numTokensToSupply,
	)

	if err != nil {
		log.Fatal(" err ", err)
	}

	//get the transaction hash
	fmt.Println("mintDai tx: ", tx.Hash().Hex())

}

func redeemDai(amount *big.Int, redeemType bool) {
	instance := GetVialendInstance()

	_cErc20Contract := common.HexToAddress(cDAI)

	tx, err := instance.RedeemCErc20Tokens(project.Auth,
		amount,
		redeemType,
		_cErc20Contract)

	if err != nil {
		log.Fatal(" err ", err)
	}

	//get the transaction hash
	fmt.Println("redeemDai tx: ", tx.Hash().Hex())

}

func redeemEth(amount *big.Int, redeemType bool) {

	instance := GetVialendInstance()

	_cEtherContract := common.HexToAddress(cETH)

	tx, err := instance.RedeemCEth(project.Auth,
		amount,
		redeemType,
		_cEtherContract)

	if err != nil {
		log.Fatal(" err ", err)
	}

	//get the transaction hash
	fmt.Println("redeemEth tx: ", tx.Hash().Hex())

}

func withDrawERC20(erc20 string, amount *big.Int) {
	instance := GetVialendInstance()

	tx, err := instance.WithdrawERC20(project.Auth,
		amount,
		common.HexToAddress(erc20),
		common.HexToAddress(to))

	if err != nil {
		log.Fatal(" err ", err)
	}

	//get the transaction hash
	fmt.Println("withdrawerc20 tx: ", tx.Hash().Hex())

}

func withDrawETH(amount *big.Int) {
	instance := GetVialendInstance()

	// withdraw eth to contract creator
	tx, err := instance.WithdrawEth(project.Auth, amount)

	if err != nil {
		log.Fatal(" err ", err)
	}

	//get the transaction hash
	fmt.Println("withdrawETH tx: ", tx.Hash().Hex())

}

func checkETHBalance() *big.Int {

	Instance := GetVialendInstance()

	bal, err := Instance.GetETHBalance(&bind.CallOpts{})

	if err != nil {
		log.Fatal("ethbalance err ", err)
	}

	fmt.Println("eth balance: ", bal)

	return (bal)

}

func checkBalance(tokenName string, cTokenAddress string, tokenHolder string) *big.Int {

	cInstance := GetCTokenInstance(cTokenAddress)

	//	bal, err := cDaiInstance.BalanceOf(&bind.CallOpts{}, FromAddress)
	bal, err := cInstance.BalanceOf(&bind.CallOpts{}, common.HexToAddress(tokenHolder))

	if err != nil {
		log.Fatal(" err ", err)
	}

	fmt.Println(tokenName, " balance: ", bal)

	return (bal)

}

func checkEvent(block0 int64, block1 int64) {
	project.VialendEvent(vialendContract, "MyLog", block0, block1)

}

func GetVialendInstance() *vialend.Api {

	instance, err := vialend.NewApi(common.HexToAddress(vialendContract), project.EthClient)
	if err != nil {
		log.Fatal("get token Instance,", err)
	}

	return instance
}

func GetCTokenInstance(Address string) *ctoken.Api {

	instance, err := ctoken.NewApi(common.HexToAddress(Address), project.EthClient)
	if err != nil {
		log.Fatal("get token Instance,", err)
	}

	return instance
}
