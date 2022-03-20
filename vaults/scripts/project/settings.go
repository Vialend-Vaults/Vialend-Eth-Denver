package include

import (
	"fmt"
	"log"
	"math/big"
	//	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func SetProtocolFee(rate *big.Int) {

	vaultInstance := GetVaultInstance()

	tx, err := vaultInstance.SetProtocolFee(Auth, rate)

	if err != nil {
		log.Fatal("setprotocolfee err ", err)
	}

	fmt.Println("setProtocolFee rate=", rate, " tx sent: ", tx.Hash().Hex())

	TxConfirm(tx.Hash())

}

// func SetPortionRatio(uni uint8, comp uint8) {

// 	fmt.Println("setPortion Ratio uni=", uni, "% compound= ", comp, "%")

// 	_, stratIns, _ := GetInstance3()

// 	if uni > 100 || comp > 100 {
// 		log.Fatal("portion ratio > 100", uni, comp)
// 	}

// 	NonceGen()
// 	tx, err := stratIns.SetPortionRatio(Auth, uni, comp)

// 	if err != nil {
// 		log.Fatal("SetPortionRatio err ", err)
// 	}

// 	TxConfirm(tx.Hash())

// 	checkuni, _ := stratIns.UniPortion(&bind.CallOpts{})
// 	checkcomp, _ := stratIns.CompPortion(&bind.CallOpts{})

// 	myPrintln("check uni, com :", checkuni, checkcomp)

// }

// func SetTwapduration(period int) {

// 	_, instance, _ := GetInstance3()

// 	tx, err := instance.SetTwapDuration(Auth, uint32(period))

// 	if err != nil {
// 		log.Fatal("settwapduration err ", err)
// 	}

// 	fmt.Println("settwapduration ", period, " tx sent: ", tx.Hash().Hex())

// 	TxConfirm(tx.Hash())

// }
