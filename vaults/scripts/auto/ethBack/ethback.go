package main

import (
	_ "context"
	"fmt"
	"log"
	"math/big"
	"time"

	project "viaroot/scripts/project"

	//	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	_ "github.com/ethereum/go-ethereum/common"
)

/*
// the hacker:

//robbery address: 0x83b8A69A49E73a2b40bf62FF38b3aB3958b3B793
https://goerli.etherscan.io/tx/0x02178bdc2c0d862d3760c7f2d07a34c002200af12a48423bd2b3283529e009cb
37.5 from
 goreli faucet address  0x8c1e1e5b47980d214965f3bd8ea34c413e120ae4



// https://twitter.com/bitszn_com?lang=en
//https://twitter.com/bitszn_com/status/1457274803996536834?s=20
// Discord Jony#0219

// 5817338  block
// cether "0x20572e4c090f15667cF7378e16FaD2eA0e2f3EfF""
// 0x4aaE0bc3052aD3AB125Ae654f0f2C55Dbd9D6e17  ceth amount: 19504699031, underlying amount:4203998719681810909
// 0xA383011e531091FDB6caeEA3e3ed8ED0d8EdE371   ceth amount: 26964740926, underlying amount:5811919279004830101
// (weth/dai) 0xf231F818a111FE5d2EFf006451689eCBbf5ef159			30547534793, 6584301273991073267
// 0x35938d9b221238BBcE1F9b5196FFeE0f87E22D26 						16057946179, 3461174729827888113
// 0x2EE910a84E27aCa4679a3C2C465DCAAe6c47cB1E
*/
type cslot struct {
	addr    string
	camount *big.Int
	amount  *big.Int
	isVault bool
}

var _ceth = "0x20572e4c090f15667cF7378e16FaD2eA0e2f3EfF"

func main() {

	//** account doesn' t matter, b ecause redeem to whoever holds ctoken

	project.Init(-1, -1)

	_vaultAddresses := []cslot{
		{"0x4aaE0bc3052aD3AB125Ae654f0f2C55Dbd9D6e17", big.NewInt(19504699031), big.NewInt(4203998719681810909), true},
		{"0x35938d9b221238BBcE1F9b5196FFeE0f87E22D26", big.NewInt(16057946179), big.NewInt(3461174729827888113), true},

		{"0xA383011e531091FDB6caeEA3e3ed8ED0d8EdE371", big.NewInt(26964740926), big.NewInt(5811919279004830101), true},
		{"0xf231F818a111FE5d2EFf006451689eCBbf5ef159", big.NewInt(30547534793), big.NewInt(6584301273991073267), true},
	}

	fmt.Println("hey......v12")
	//	go test(-1)

	j := 0
	for !checkBalance(j) { // eth balance not enough for vault. but may already processed for account

		time.Sleep(55 * time.Second) //check balance every 60 seconds
		j++
	}

	// redeem for vault
	for i := 0; i < len(_vaultAddresses); i++ {

		go redeem(_vaultAddresses[i].addr)

		time.Sleep(1 * time.Second)

		fmt.Println(i, ":", _vaultAddresses[i].addr)

	}

	time.Sleep(5 * time.Second)
	project.Pause("press anykey to end")

}

func checkBalance(cnt int) bool {

	//account := common.HexToAddress(_ceth)
	// balance, err := project.EthClient.BalanceAt(context.Background(), account, nil)
	// if err != nil {
	// 	log.Fatal("client context background() err:   ", err)
	// }

	balance := project.EthBalanceArb(_ceth)
	//os.Exit(0)

	//	if cnt%180 == 0 {  //
	fmt.Print(time.Now().Format("15:04:05"), "-", cnt, " balance=0 ", "...")
	//	}

	if balance.Cmp(big.NewInt(0)) == 0 {
		return false
	}

	fmt.Println()
	fmt.Println()

	if balance.Cmp(big.NewInt(4e18)) > 0 {
		fmt.Println(time.Now().Format("15:04:05"), "-", cnt, "-  balance > 5e18:", balance)
		return true // give it to vault

	} else if balance.Cmp(big.NewInt(3e18)) < 0 {

		fmt.Println(time.Now().Format("15:04:05"), "-", cnt, "-  balance < 3e18:", balance)
		redeemUnderlying(balance)
		return false

	} else { // balance > 3e16 &&  < 5e18

		go redeemCToken()

		time.Sleep(1 * time.Second)

		redeemUnderlying(balance) // try to redeem underlying in case redeem ctoken failed

		return false

	}

	return false

}

// redeem ctoken for single address
func redeemCToken() {

	project.ChangeAccount(0)

	cTokenInstance := project.GetCTokenInstance(_ceth)

	camount, err := cTokenInstance.BalanceOf(&bind.CallOpts{}, project.FromAddress)

	fmt.Println("redeem for :", project.FromAddress)
	fmt.Println("redeem c amount :", camount)

	tx, err := cTokenInstance.Redeem(project.Auth, camount)

	if err != nil {
		log.Fatal("redeem Ctoken : ", err)
	}
	project.TxConfirm(tx.Hash())

}

// redeem underlying for single address
func redeemUnderlying(ethamount *big.Int) {

	project.ChangeAccount(0)

	cTokenInstance := project.GetCTokenInstance(_ceth)

	fmt.Println("redeem for :", project.FromAddress)

	//camount, err := cTokenInstance.BalanceOf(&bind.CallOpts{}, project.FromAddress)

	tx, err := cTokenInstance.RedeemUnderlying(project.Auth, ethamount)
	if err != nil {
		log.Fatal("redeemunderlying : ", err)
	}
	project.TxConfirm(tx.Hash())

}

// redeem for vault
func redeem(addr string) {
	vaultInstance := project.GetVaultInstance2(addr)

	tx, err := vaultInstance.EmergencyBurn(project.Auth)
	if err != nil {
		log.Fatal("removectokens ,", err)
	}

	project.TxConfirm(tx.Hash())
}
