package include

import (
	"log"
	"math/big"

	//	"reflect"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func GetPrice() {

	myPrintln("----------------------------------------------")
	myPrintln(".........GetPrice strat2 .........  ")
	myPrintln("----------------------------------------------")

	_, stratInstance, _ := GetInstance4()

	price, err := stratInstance.GetPrice(&bind.CallOpts{})

	if err != nil {
		log.Fatal("getprice err:", err)
	}

	myPrintln("price:", price)

}

func GetPricePer(_token string, fee int) *big.Int {

	myPrintln("----------------------------------------------")
	myPrintln(".........GetPricePer strat2 .........  ")
	myPrintln("----------------------------------------------")

	_, stratInstance, _ := GetInstance4()

	price, err := stratInstance.GetEthPriceFromPool(&bind.CallOpts{}, common.HexToAddress(_token), big.NewInt(int64(fee)))

	if err != nil {
		bal, symb := TokenInfo(_token, Network.VaultStrat)
		log.Println(symb, ",", bal, ", ", _token, ",", fee)

		log.Fatal("getpricefrompool err:", err)

	}

	//myPrintln("price per :", price)

	return price

}

func GetExpectedNormalizationFactor() {

	_, stratInstance, _ := GetInstance4()

	norm, err := stratInstance.GetSqueethNorm(&bind.CallOpts{}, common.HexToAddress(Network.LendingContracts.SqthController))
	if err != nil {
		log.Println("sqthController address:", Network.LendingContracts.SqthController)
		log.Fatal("getsqueethnorm ", err)
	}
	myPrintln("normalization factor:", norm)

}

func SwapTest(tokenIn string, tokenOut string, fee int, amount *big.Int) {

	myPrintln("----------------------------------------------")
	myPrintln(".........Swap throu uniswap router 02 .........  ")
	myPrintln("----------------------------------------------")

	myPrintln("tokenIn:", tokenIn)
	myPrintln("tokenOut:", tokenOut)
	myPrintln("amount:", amount)

	_, stratInstance, _ := GetInstance4()

	tx, err := stratInstance.SwapTest(Auth,
		common.HexToAddress(tokenIn),
		common.HexToAddress(tokenOut),
		big.NewInt(int64(fee)),
		amount,
	)

	if err != nil {
		log.Fatal("swap err:", err)
	}

	myPrintln("tx :", tx.Hash())

	TxConfirm(tx.Hash())

}
