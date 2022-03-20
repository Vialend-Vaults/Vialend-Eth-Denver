package main

import (
	"math/big"
	_ "time"

	project "viaroot/scripts/project"
)

type TransferStruct struct {
	AccountId    int
	Amount       *big.Int
	TokenAddress string
	ToAddress    string
}
type Switcher struct {
	ViewOnly         bool
	DeployToken      int
	TokenParam       [2]project.TokenStruct
	TransferToken    int
	TransferParam    [2]TransferStruct
	DeployFactory    int
	CreatePool       int
	InitialPool      int
	MintPool         int
	MintPoolParam    [3]int64
	DeployVault      int
	Deposit          int
	DepositParam     [3]int64
	Withdraw         int
	WithDrawParam    [2]int64
	Rebalance        int
	RebalanceParam   [3]int64
	CreatePosition   int
	IncreasePosition int
	RemovePosition   int
	Swap             int
	CollectFees      int
	Strategy1        int
	Strategy1Param   [3]int64
}

func main() {

	project.Init(-1, -1)
	project.Quiet = false

	//project.DeployViaFactory()
	project.DeployStratUniComp()

}
