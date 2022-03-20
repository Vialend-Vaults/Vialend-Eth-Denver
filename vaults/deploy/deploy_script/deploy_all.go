package main

import (
	"context"
	"crypto/ecdsa"

	"fmt"
	"math/big"

	//api "../FundKeeper"
	config "../../test/scripts/config"
	callee "../TestUniswapV3Callee"
	tester "../tester"
	unwrap "../unwrap"
	vialendcomp "../vialendcomp"
	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type nClient struct {
	clientUrl  string
	privateKey string
	pool       string
}

func main() {

	nid := 1 // 0: local, 1: goerli, 2: goerli, 3: goerli user 4: rinkeby
	deploy_tester := false
	deploy_TestUniswapV3Callee := false
	deploy_vialendcomp := true
	deploy_unwrap := false

	networks := [...]nClient{

		{ //0 local
			"http://127.0.0.1:7545",
			"e8ef3a782d9002408f2ca6649b5f95b3e5772364a5abe203f1678817b6093ff0",
			"0x606B7f7093aa4df2282763F4e9f714706221838b"},

		{ //1 test admin
			"https://goerli.infura.io/v3/68070d464ba04080a428aeef1b9803c6",
			"284b65567176c10bc010345042b1d9852fcc1c42ae4b76317e6da040318fbe7f",
			"0xBF93aB266Cd9235DaDE543fAd2EeC884D1cCFc0c"},

		{ //2 test admin
			"https://goerli.infura.io/v3/68070d464ba04080a428aeef1b9803c6",
			"2b200539ce93eab329be1bd7c199860782e547eb7f95a43702c1b0641c0486a7",
			"0xBF93aB266Cd9235DaDE543fAd2EeC884D1cCFc0c"},

		{ //3 test user 2
			"https://goerli.infura.io/v3/68070d464ba04080a428aeef1b9803c6",
			"01e8c8df56230b8b6e4ce6371bed124f4f9950c51d64adc581938239724ed5e6",
			"0xBF93aB266Cd9235DaDE543fAd2EeC884D1cCFc0c"},

		{ //4
			"https://rinkeby.infura.io/v3/68070d464ba04080a428aeef1b9803c6",
			"2b200539ce93eab329be1bd7c199860782e547eb7f95a43702c1b0641c0486a7", // 0x2EE910a84E27aCa4679a3C2C465DCAAe6c47cB1E
			"pool"},
	}

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
	/*auth, err := config.Auth(client)
	if err != nil {
		panic(err)
	}
	*/

	if deploy_TestUniswapV3Callee {
		address, tx, instance, err := callee.DeployApi(auth, client)
		if err != nil {
			panic(err)
		}
		fmt.Println("contract address:", address.Hex())

		_, _ = instance, tx
	}

	if deploy_tester {
		address, tx, instance, err := tester.DeployApi(auth, client)
		if err != nil {
			panic(err)
		}
		fmt.Println("contract address:", address.Hex())

		_, _ = instance, tx
	}

	if deploy_vialendcomp {
		address, tx, instance, err := vialendcomp.DeployApi(auth, client, common.HexToAddress(config.Network.LendingContracts.WETH))
		if err != nil {
			panic(err)
		}
		fmt.Println("vialendcomp contract address:", address.Hex())

		_, _ = instance, tx
	}

	if deploy_unwrap {

		address, tx, instance, err := unwrap.DeployApi(auth, client,
			common.HexToAddress(config.Network.LendingContracts.WETH))
		if err != nil {
			panic(err)
		}
		fmt.Println("unwrap contract address:", address.Hex())

		_, _ = instance, tx
	}

}
