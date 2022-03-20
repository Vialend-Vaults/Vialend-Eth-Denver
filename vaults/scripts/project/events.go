package include

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"

	//	"time"

	//	factory "../../../../../../../uniswap/v3/deploy/UniswapV3Factory"

	//vault "viaroot/deploy/FeeMaker"
	callee "viaroot/deploy/TestUniswapV3Callee"
	vault "viaroot/deploy/vialendFeeMaker"
	vialend "viaroot/deploy/vialendcomp"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

func VaultEvent(
	eventName string,
	block_begin int64,
	block_end int64) {

	fmt.Println("----------------------------------------------")
	fmt.Println(".........Events.........  ")
	fmt.Println("----------------------------------------------")

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(block_begin),
		ToBlock:   big.NewInt(block_end),
		Addresses: []common.Address{
			common.HexToAddress(Network.Vault),
		},
	}
	logs, err := EthClient.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(vault.ApiABI)))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Event name ", eventName)

	for _, vLog := range logs {
		fmt.Println("blockHash: ", vLog.BlockHash.Hex())
		fmt.Println("block#: ", vLog.BlockNumber)
		fmt.Println("txHash: ", vLog.TxHash.Hex())

		x, _ := contractAbi.Unpack(eventName, vLog.Data)

		//Print event data
		fmt.Printf("Data:------\n", x)
		fmt.Printf("%#v\n", x)

		/*
			event := struct {
				Key   [32]byte
				Value [32]byte
			}{}

			event, err := contractAbi.Unpack("ItemSet", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("event.Key[:]: ", string(event.Key[:]))     // foo
			fmt.Println("event.Value[:]: ", string(event.Value[:])) // bar

			var topics [4]string
			for i := range vLog.Topics {
				topics[i] = vLog.Topics[i].Hex()
			}

			fmt.Println("topic[0]: ", topics[0])
		*/
	}

	//		eventSignature := []byte("ItemSet(bytes32,bytes32)")
	//	hash := crypto.Keccak256Hash(eventSignature)
	//fmt.Println("itemset hash: ", hash.Hex())

}

func VialendEvent(
	contractAddress string,
	eventName string,
	block_begin int64,
	block_end int64) {

	fmt.Println("----------------------------------------------")
	fmt.Println(".........Vialend Events.........  ")
	fmt.Println("----------------------------------------------")

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(block_begin),
		ToBlock:   big.NewInt(block_end),
		Addresses: []common.Address{
			common.HexToAddress(contractAddress),
		},
	}
	logs, err := EthClient.FilterLogs(context.Background(), query)
	if err != nil {
		log.Println(err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(vialend.ApiABI)))
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Event name ", eventName)

	for _, vLog := range logs {
		//fmt.Println("blockHash: ", vLog.BlockHash.Hex())
		fmt.Println("block#: ", vLog.BlockNumber)
		//fmt.Println("txHash: ", vLog.TxHash.Hex())

		x, _ := contractAbi.Unpack(eventName, vLog.Data)

		//Print event data
		fmt.Printf("Data:------\n", x)
		fmt.Printf("%#v\n", x)
	}

}

func TestCalleeEvent(
	eventName string,
	block_begin int64,
	block_end int64) {

	fmt.Println("----------------------------------------------")
	fmt.Println(".........callee Events.........  ")
	fmt.Println("----------------------------------------------")

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(block_begin),
		ToBlock:   big.NewInt(block_end),
		Addresses: []common.Address{
			common.HexToAddress(Network.Callee),
		},
	}
	logs, err := EthClient.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(callee.ApiABI)))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Event name ", eventName)

	for _, vLog := range logs {
		fmt.Println("blockHash: ", vLog.BlockHash.Hex())
		fmt.Println("block#: ", vLog.BlockNumber)
		fmt.Println("txHash: ", vLog.TxHash.Hex())

		x, _ := contractAbi.Unpack(eventName, vLog.Data)

		//Print event data
		fmt.Printf("Data:------\n", x)
		fmt.Printf("%#v\n", x)

		/*
			event := struct {
				Key   [32]byte
				Value [32]byte
			}{}

			event, err := contractAbi.Unpack("ItemSet", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("event.Key[:]: ", string(event.Key[:]))     // foo
			fmt.Println("event.Value[:]: ", string(event.Value[:])) // bar

			var topics [4]string
			for i := range vLog.Topics {
				topics[i] = vLog.Topics[i].Hex()
			}

			fmt.Println("topic[0]: ", topics[0])
		*/
	}

	//		eventSignature := []byte("ItemSet(bytes32,bytes32)")
	//	hash := crypto.Keccak256Hash(eventSignature)
	//fmt.Println("itemset hash: ", hash.Hex())

}
