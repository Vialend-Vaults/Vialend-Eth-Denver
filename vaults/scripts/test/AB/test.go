package main

import (
	"fmt"
	"log"
	_ "time"

	AB "viaroot/deploy/AB"
	project "viaroot/scripts/project"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func main() {

	// Bins := project.DeployAB()

	// _AA, _ := Bins.A(&bind.CallOpts{})

	// fmt.Println("AA address:", _AA)
	// return

	_B := common.HexToAddress("0x402a3dBDB42f246a9Fb7c70F7b89F19265d5F54B")
	_A := common.HexToAddress("0x9d98133037dDef780d69C4B498b51e1f7cDe4c53")

	Ainstance, err := AB.NewA(_A, project.EthClient)
	if err != nil {
		log.Fatal("AInstance err:", err)
	}

	//result, _ := Ainstance.CheckTotalSupply(project.Auth)
	//	fmt.Println("checktotalsupply by A:", result)

	fmt.Println("call changebyA  from A:")
	tx, _ := Ainstance.ChangebyA(project.Auth, "hey its A")

	project.TxConfirm(tx.Hash())

	Binstance, err := AB.NewB(_B, project.EthClient)
	if err != nil {
		log.Fatal("BInstance err:", err)
	}

	sparam, _ := Binstance.Sparam(&bind.CallOpts{})

	fmt.Println("sparam should be 'hey its A';  got:'", sparam, "'")

	// call changebyA from B
	tx, _ = Binstance.ChangebyA(project.Auth, "hello its B")

	project.TxConfirm(tx.Hash())

	sparamB, _ := Binstance.Sparam(&bind.CallOpts{})
	fmt.Println("sparamB should be 'hello its B';  got:'", sparamB, "'")

}
