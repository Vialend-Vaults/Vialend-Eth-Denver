/*

	this deploy file is deprecated. check vaults/test/scripts/deploy.go

	go file still be referenced
*/

package main

import (
	"context"
	"crypto/ecdsa"

	"fmt"
	"math/big"

	//api "../deploy/token/tusdc"
	//api "../deploy/token/tweth"
	api "viaroot/deploy/Tokens/erc20/deploy/ERC20fixedSupply"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type nClient struct {
	clientUrl  string
	privateKey string
	pool       string //not in use
}

func main() {

	networks := [...]nClient{

		{ //0
			"http://127.0.0.1:7545",
			"e8ef3a782d9002408f2ca6649b5f95b3e5772364a5abe203f1678817b6093ff0",
			"n/a"},

		{ //1 Tester Admin 0x2EE910a84E27aCa4679a3C2C465DCAAe6c47cB1E
			"https://goerli.infura.io/v3/68070d464ba04080a428aeef1b9803c6",
			"2b200539ce93eab329be1bd7c199860782e547eb7f95a43702c1b0641c0486a7",
			""},
		{ //2 Tester Admi 2 0x6dd19aEB91d1f43C46f0DD74C9E8A92BFe2a3Cd0
			"https://goerli.infura.io/v3/68070d464ba04080a428aeef1b9803c6",
			"284b65567176c10bc010345042b1d9852fcc1c42ae4b76317e6da040318fbe7f",
			""},

		{ //3  Tester Admin 3 0xD8Dbe65b64428464fFa14DEAbe288b83C240e713
			"https://goerli.infura.io/v3/68070d464ba04080a428aeef1b9803c6",
			"d8cda34b6928af75aff58c60fe9ed3339896b57a13fa88695aa6da7b775cda2a",
			""},

		{"https://rinkeby.infura.io/v3/68070d464ba04080a428aeef1b9803c6",
			"2b200539ce93eab329be1bd7c199860782e547eb7f95a43702c1b0641c0486a7",
			"pool"},
	}

	//	total0, _ := new(big.Int).SetString("26512386160", 10) // totalSupply , 6

	//name, symbol, decimals, totalSupply := "ttt Bonus", "tBonus", uint8(18), big.NewInt(1e9)

	//name, symbol, decimals, totalSupply := "tto USDC", "tUSDC6", uint8(6), big.NewInt(26512386160)
	//name, symbol, decimals, totalSupply := "tto USDC", "tUSDC", uint8(18), big.NewInt(26512386160)
	//(name,symbol,decimals,totalSupply) := ("tto WETH","tWETH", uint8(18), big.NewInt(6718113))

	//name, symbol, decimals, totalSupply := "ddd USDC", "dUSDC", uint8(6), big.NewInt(26512386160)
	//	name, symbol, decimals, totalSupply := "ddd WETH", "dWETH", uint8(18), big.NewInt(6718113)

	//name, symbol, decimals, totalSupply := "eee USDC", "eUSDC", uint8(6), big.NewInt(26512386160)
	//name, symbol, decimals, totalSupply := "eee WETH", "eWETH", uint8(18), big.NewInt(6718113)

	//name, symbol, decimals, totalSupply := "local USDC", "USDC", uint8(18), big.NewInt(50000000000)
	name, symbol, decimals, totalSupply := "local USDT", "USDT", uint8(18), big.NewInt(50000000000)

	nid := 0
	client, err := ethclient.Dial(networks[nid].clientUrl)
	privateKey, err := crypto.HexToECDSA(networks[nid].privateKey)

	if err != nil {
		panic(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("invalid key")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		panic(err)
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		panic(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		panic(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)                 // in wei
	auth.GasLimit = uint64(6721975)            // in units
	auth.GasPrice = big.NewInt(8 * 1000000000) // 1 gwei = 1000000000

	address, tx, instance, err := api.DeployApi(auth, client, name, symbol, decimals, totalSupply)
	if err != nil {
		panic(err)
	}

	fmt.Println("contract address:", address.Hex())

	_, _ = instance, tx
}
