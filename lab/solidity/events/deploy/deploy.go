package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	api ".." // for demo
)

var provider = "ws://localhost:7545"
var contract = "0x3acA8F82e5fca1f437D71BEc742F472b843451d0"
var eventName = "Log"
var blockFrom = int64(4)
var blockEnd = int64(1351)

// api.abi needed
// var provider = "wss://goerli.infura.io/ws/v3/68070d464ba04080a428aeef1b9803c6"
// var contract = "0xE36DcbC67D9DE4A4bcF4cc496e22D505735EdB4D"
// var eventName = "Deposit"
// var blockFrom = int64(6048027)
// var blockEnd = int64(6048027)

func main() {

	// deploy() // or remix
	// return

	//listenEvent()

	queryEventFiltered()

	//queryEvent()
	// return

	//genEvent("Juliet4", 21)
	// genEvent("Paul et", 41)
	// genEvent("tom", 11)

}

func GetClient() *ethclient.Client {

	c, err := ethclient.Dial(provider)
	if err != nil {
		log.Fatal(err)
	}

	return c

}

func queryEvent() {

	client := GetClient()

	contractAddress := common.HexToAddress(contract)

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(blockFrom),
		ToBlock:   big.NewInt(blockEnd),
		Addresses: []common.Address{
			contractAddress,
		},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(api.ApiABI)))
	if err != nil {
		log.Fatal(err)
	}

	for _, vLog := range logs {
		fmt.Println("Block #:", vLog.BlockNumber) // 2394201

		event, err := contractAbi.Unpack(eventName, vLog.Data)

		if err != nil {
			log.Fatal(err)
		}

		// fmt.Println(string(event.Key[:]))   // foo
		// fmt.Println(string(event.Value[:])) // bar

		fmt.Println("event name: Instructor, len:", len(event))
		for j := 0; j < len(event); j++ {
			fmt.Println("event[", j, "]=", event[j])
		}

		fmt.Println("Topics:", len(vLog.Topics))

		for k := 0; k < len(vLog.Topics); k++ {

			fmt.Println(vLog.Topics[k])
		}

		//		var topics [4]string
		// for i := range vLog.Topics {

		// 	topics[i] = vLog.Topics[i].Hex()

		// 	//fmt.Println("topic ", i, " hash:", topics[i])
		// }

		// fmt.Println("topics0 hash:", topics[0]) // 0xe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4
	}

	// cannot replicate:
	// eventSignature := []byte("Instructor(string, uint64, string,uint64)")
	// hash := crypto.Keccak256Hash(eventSignature)
	// fmt.Println("eventSignature hash:", hash.Hex()) // 0xe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4

}

func hexToInt(hexaString string) int64 {

	// replace 0x or 0X with empty String
	numberStr := strings.Replace(hexaString, "0x", "", -1)
	numberStr = strings.Replace(numberStr, "0X", "", -1)

	output, err := strconv.ParseInt(numberStr, 16, 64)
	if err != nil {
		log.Println(err)
	}

	//fmt.Printf("Output %d", output)
	return output

}

func queryEventFiltered() {

	client := GetClient()

	contractAddress := common.HexToAddress(contract)

	fromBlock := big.NewInt(blockFrom)
	toBlock := big.NewInt(blockEnd)
	query := ethereum.FilterQuery{
		FromBlock: fromBlock,
		ToBlock:   toBlock,
		Addresses: []common.Address{
			contractAddress,
		},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(api.ApiABI)))
	if err != nil {
		log.Fatal(err)
	}

	op := ">"
	a := int64(24)
	fmt.Println("Query name ", op, " ", a, "; from ", fromBlock, " to:", toBlock)

	for _, vLog := range logs {

		age := hexToInt(vLog.Topics[2].String())

		var isOp bool
		if op == ">" {
			isOp = age > a
		}
		if op == "<" {
			isOp = age < a
		}

		if isOp {

			event, err := contractAbi.Unpack(eventName, vLog.Data)
			if err != nil {
				log.Println(err)
			}

			fmt.Println(event, " age=", age, ", block=", vLog.BlockNumber)
		}

	}

}

func listenEvent() {

	client := GetClient()

	contractAddress := common.HexToAddress(contract)

	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)

	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(api.ApiABI)))
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			//fmt.Println(vLog) // pointer to event log

			event, err := contractAbi.Unpack(eventName, vLog.Data)

			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("event name: Instructor, len:", len(event))
			for j := 0; j < len(event); j++ {
				fmt.Println("event[", j, "]=", event[j])
			}

			fmt.Println("Topics:", len(vLog.Topics))
			for k := 0; k < len(vLog.Topics); k++ {
				fmt.Println("Topics[", k, "]=", vLog.Topics[k])
			}

		}

	}

}

func deploy() {
	client := GetClient()

	privateKey, err := crypto.HexToECDSA("b8c1b5c1d81f9475fdf2e334517d29f733bdfa40682207571b12fc1142cbf329")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	address, tx, instance, err := api.DeployApi(auth, client)
	if err != nil {
		log.Fatal(err)
	}

	contract = address.Hex()

	fmt.Println(address.Hex())   // 0x147B8eb97fD247D06C4006D269c90C1908Fb5D54
	fmt.Println(tx.Hash().Hex()) // 0xdae8ba5444eefdc99f4d45cd0c4f24056cba6a02cefbf78066ef9f4188ff7dc0

	_ = instance
}

func genEvent(name string, age int64) {

	client, err := ethclient.Dial("http://localhost:7545")
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress(contract)

	instance, err := api.NewApi(contractAddress, client)

	privateKey, err := crypto.HexToECDSA("5c2313d8a6b81a83ad1df1bf12a193cbc51d5de84a000db734fd7a05aa63e5a2")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	instance.SetInfo(auth, name, big.NewInt(age))

	fname, _age, err := instance.GetInfo(&bind.CallOpts{})
	fmt.Print("fname:", fname, ", age:", _age)

}
