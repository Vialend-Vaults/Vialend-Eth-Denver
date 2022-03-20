package include

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

func BlockNumber() *big.Int {

	header, err := EthClient.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	myPrintln(header.Number.String()) // 5671744
	return header.Number

}

func EthBalance(_addr string) *big.Int {

	account := common.HexToAddress(_addr)
	myPrintln("address:", _addr)
	balance, err := EthClient.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal("eth balance via context background() err:   ", err)
	}

	myPrintln(balance)
	return balance
}
