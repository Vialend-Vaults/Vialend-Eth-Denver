package main

import (
	"os"
	_ "os"
	_ "time"

	p "viaroot/scripts/project"
)

func main() {

	p.Init(-1, -1)

	p.Network.ViaFactory = os.Getenv("ViaFactory")
	p.Network.VaultStrat = os.Getenv("VaultStrat")
	p.Network.Vault = os.Getenv("Vault")

	// fmt.Println(p.Network.VaultStrat, p.Network.Vault)
	// // return

	// p.ViaVaultPublicList()
	// return

	// p.GetPriceVaultCall()
	// p.GetPriceStratCall()
	//p.CalcPositionShares(0, 0)
	// // p.CheckCap(1000, 10)
	// return

	//deposit
	//p.ViaVaultPublicFunc(1, int64(1000000000000000), int64(1000000000000000000))
	// p.ChangeAccount(1)
	// p.ViaVaultPublicFunc(1, int64(2000000000000000), int64(2000000000000000000))

	//withdraw
	p.ViaVaultPublicFunc(2, uint8(100))
	// p.ChangeAccount(1)
	// p.ViaVaultPublicFunc(2, uint8(100))
	// return

}
