package include

import (
	"log"
	"math/big"

	//	"time"

	//	factory "../../../../../../../uniswap/v3/deploy/UniswapV3Factory"
	Squeeth "viaroot/deploy/Squeeth"
	token "viaroot/deploy/Tokens/erc20/deploy/Token"
	pool "viaroot/deploy/uniswap/v3/deploy/UniswapV3Pool"

	//vault "viaroot/deploy/FeeMaker"
	callee "viaroot/deploy/TestUniswapV3Callee"
	VaultFactory "viaroot/deploy/VaultFactory"
	VaultStrategy "viaroot/deploy/VaultStrategy2"

	VaultStrategy2 "viaroot/deploy/VaultStrategy2"

	ViaVault "viaroot/deploy/ViaVault"
	cErc20 "viaroot/deploy/cErc20"
	vault "viaroot/deploy/vialendFeeMaker"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func GetInstance3() (*VaultFactory.Api, *VaultStrategy.Api, *ViaVault.Api) {

	A1, err := VaultFactory.NewApi(common.HexToAddress(Network.VaultFactory), EthClient)
	if err != nil {
		log.Println("VaultFactory Instance err:", err)
	}

	A2, err := VaultStrategy.NewApi(common.HexToAddress(Network.VaultStrat), EthClient)
	if err != nil {
		log.Println("VaultStrat Instance err:", err)
	}

	A3, err := ViaVault.NewApi(common.HexToAddress(Network.Vault), EthClient)
	if err != nil {
		log.Println("ViaVault Instance err:", err)
	}

	return A1, A2, A3
}

func GetInstance4() (*VaultFactory.Api, *VaultStrategy2.Api, *ViaVault.Api) {

	A1, err := VaultFactory.NewApi(common.HexToAddress(Network.VaultFactory), EthClient)
	if err != nil {
		log.Println("VaultFactory Instance err:", err)
	}

	A2, err := VaultStrategy2.NewApi(common.HexToAddress(Network.VaultStrat), EthClient)
	if err != nil {
		log.Println("VaultStrat Instance err:", err)
	}

	A3, err := ViaVault.NewApi(common.HexToAddress(Network.Vault), EthClient)
	if err != nil {
		log.Println("ViaVault Instance err:", err)
	}

	return A1, A2, A3
}

func GetVaultInstance() *vault.Api {

	instance, err := vault.NewApi(common.HexToAddress(Network.Vault), EthClient)
	if err != nil {
		log.Fatal("vaultInstance err:", err)
	}
	return instance
}

func GetVaultInstance2(_addr string) *vault.Api {

	instance, err := vault.NewApi(common.HexToAddress(_addr), EthClient)
	if err != nil {
		log.Fatal("vaultInstance err:", err)
	}
	return instance
}

func GetCalleeInstance() *callee.Api {

	instance, err := callee.NewApi(common.HexToAddress(Network.Callee), EthClient)
	if err != nil {
		log.Fatal("CalleeInstance err:", err)
	}
	return instance
}

func GetTokenInstance(TokenAddress string) (*token.Api, string, string, uint8, *big.Int) {

	instance, err := token.NewApi(common.HexToAddress(TokenAddress), EthClient)
	if err != nil {
		log.Fatal("get token Instance,", err)
	}

	name, err := instance.Name(&bind.CallOpts{})
	if err != nil {
		log.Println("token address:", TokenAddress)
		log.Fatal("get token name,", err)
	}
	symbol, err := instance.Symbol(&bind.CallOpts{})
	if err != nil {
		log.Fatal("get token symbol,", err)
	}
	decimals, err := instance.Decimals(&bind.CallOpts{})
	if err != nil {
		log.Fatal("get token decimals,", err)
	}
	maxsupply, err := instance.TotalSupply(&bind.CallOpts{})
	if err != nil {
		log.Fatal("get token maxsupply,", err)
	}
	return instance, name, symbol, decimals, maxsupply
}

func GetCTokenInstance(Address string) *cErc20.Api {

	instance, err := cErc20.NewApi(common.HexToAddress(Address), EthClient)
	if err != nil {
		log.Fatal("get token Instance,", err)
	}

	return instance
}

func GetSqueethInstance(Address string) *Squeeth.Api {

	instance, err := Squeeth.NewApi(common.HexToAddress(Address), EthClient)
	if err != nil {
		log.Fatal("get Squeeth Instance,", err)
	}

	return instance
}

func GetPoolInstance() *pool.Api {

	_pool := GetPool(Network.TokenA, Network.TokenB, Network.FeeTier)
	//instance, err := pool.NewApi(common.HexToAddress(Network.Pool), EthClient)
	instance, err := pool.NewApi(_pool, EthClient)

	if err != nil {
		log.Fatal("poolInstance err:", err)
	}
	return instance

}

func GetPoolInstance1(_pool string) *pool.Api {

	//instance, err := pool.NewApi(common.HexToAddress(Network.Pool), EthClient)
	instance, err := pool.NewApi(common.HexToAddress(_pool), EthClient)

	if err != nil {
		log.Fatal("poolInstance err:", err)
	}
	return instance

}
