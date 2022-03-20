package include

import (
	"context"
	"fmt"
	"log"
	"math/big"

	//	"reflect"
	"strconv"
	"strings"
	"time"

	VaultFactory "viaroot/deploy/VaultFactory"
	VaultStrat "viaroot/deploy/VaultStrategy2"
	ViaVault "viaroot/deploy/ViaVault"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

/*
VaultFactory publics

	      mapping (address => uint ) public vaultStatus;  // 1: active, 2: pending, 3: abandoned 0: empty
	      address[] public  vaultRegistry;

	   	function getAdmin() public view returns(address)
	      function getTeam() public view returns(address)
	      function setTeam(address _team) public onlyAdmin
	    	function changeVaultStat(address _vault, uint status) external onlyAdmin
	      function registerVault(address _vault) external
*/

func ViewVaults() {

	facotryInstance, _, _ := GetInstance3()

	myTitle("Check Vault Status")

	myPrintln("VaultFactory Address:", Network.VaultFactory)

	count, _ := facotryInstance.GetCount(&bind.CallOpts{})
	i := count.Int64()
	for i > 0 {
		i--
		vaults, _ := facotryInstance.Vaults(&bind.CallOpts{}, big.NewInt(i))

		myPrintln(i+1, " of ", count)
		myPrintln("strategy:", vaults.Strategy)
		myPrintln("vault:", vaults.Vault)

		stat, _ := facotryInstance.GetStat2(&bind.CallOpts{}, vaults.Strategy, vaults.Vault)
		myPrintln("stat:", stat)
	}

}

func GetStat(s string, v string) {
	facotryInstance, _, _ := GetInstance3()

	myTitle("Check Vault Status")

	stat, _ := facotryInstance.GetStat2(&bind.CallOpts{}, common.HexToAddress(s), common.HexToAddress(v))
	myPrintln("stat:", s)
	myPrintln("vault:", v)
	myPrintln("stat:", stat)

}

func ChangeStat(_strategy string, _vault string, _stat int) {

	factoryInst, _, _ := GetInstance3()

	myTitle("Change Stat")

	NonceGen()
	if tx, err := factoryInst.ChangeStat(Auth,
		common.HexToAddress(_strategy),
		common.HexToAddress(_vault),
		big.NewInt(int64(_stat))); err == nil {

		TxConfirm(tx.Hash())
	} else {
		log.Fatal("changestat", err)
	}

	GetStat(_strategy, _vault)

}

// func TradeSqth(amount int, zeroForOne bool) {

// 	_, stratInstance, _ := GetInstance4()

// 	tx, err := stratInstance.TradeSqth(Auth, big.NewInt(int64(amount)), zeroForOne)
// 	if err != nil {
// 		log.Fatal("tradesqth error", err)
// 	}

// 	TxConfirm(tx.Hash())

// }

func Register(_strategy string, _vault string) {

	// if GetPair0(_strategy).Hex() == _vault {

	// 	myPrintln("ERROR: already registered")
	// 	return
	// }

	factoryInst, _, _ := GetInstance3()

	myTitle("Register")

	NonceGen()

	if tx, err := factoryInst.Register(Auth,
		common.HexToAddress(_strategy),
		common.HexToAddress(_vault)); err == nil {

		TxConfirm(tx.Hash())
	} else {
		log.Fatal("Register", err)
	}

	GetStat(_strategy, _vault)

}

func FactoryPairs(addr string) {
	factoryInstance, _, _ := GetInstance3()

	addr2, err := factoryInstance.Pairs(&bind.CallOpts{}, common.HexToAddress(addr))
	if err != nil {
		log.Fatal("pairs[a] err-> ", err)
	}
	myPrintln("pair source:", addr)
	myPrintln("pair target:", addr2)

}
func FactoryPublicList() {

	factoryInstance, _, _ := GetInstance3()

	p1, _ := factoryInstance.GetTeam(&bind.CallOpts{})
	p2, _ := factoryInstance.GetAdmin(&bind.CallOpts{})
	p3, _ := factoryInstance.GetCount(&bind.CallOpts{})
	p4, _ := factoryInstance.GetStat(&bind.CallOpts{}, common.HexToAddress(Network.Vault))

	myTitle("factory public attributes")

	myPrintln("TEAM", p1)
	myPrintln("ADMIN", p2)
	myPrintln("VAULTS COUNT", p3)
	myPrintln("getStat", p4)

	vaultCounts := int(p3.Int64())
	for i := 0; i < vaultCounts; i++ {
		v, _ := factoryInstance.Vaults(&bind.CallOpts{}, big.NewInt(int64(i)))
		//Sleep(5000)

		myPrintln(i, ":", v)
	}

}

func GetPair0(addr string) common.Address {

	factoryInstance, _, _ := GetInstance3()

	theOther, err := factoryInstance.GetPair0(&bind.CallOpts{}, common.HexToAddress(addr))
	if err != nil {
		log.Fatal("getpair0 err-> ", err)
	}

	myPrintln("getpair0:", theOther)

	return (theOther)
}

func ViaVaultPublicList() {

	myTitle("Vault public methods")
	factoryInstance, _, viaVaultInstance := GetInstance3()

	p1, _ := factoryInstance.GetPair0(&bind.CallOpts{}, common.HexToAddress(Network.Vault))
	p3, _ := viaVaultInstance.Token0(&bind.CallOpts{})
	p4, _ := viaVaultInstance.Token1(&bind.CallOpts{})
	p5, _ := viaVaultInstance.VaultCap(&bind.CallOpts{})
	p6, _ := viaVaultInstance.IndividualCap(&bind.CallOpts{})

	myPrintln()
	myPrintln("\n", ">>>>>>via Vault  public attributes <<<<<<<<<<<", "\n")
	myPrintln("Vault:", Network.Vault)
	myPrintln("Strategy", p1)
	myPrintln("Token0", p3)
	myPrintln("Token1", p4)
	myPrintln("VaultCap", p5)
	myPrintln("IndividualCap", p6)

	myPrintln("\n", ">>>>>>via Vault  public functions <<<<<<<<<<<", "\n")
	myPrintln("1 function deposit( uint256 amountIn0, uint256 amountIn1) external")
	myPrintln("2 function withdraw( uint8 percent ) external nonReentrant returns (uint256 amount0, uint256 amount1)")
	myPrintln("3 function setVaultCap(uint256 newMax) external onlyAdmin")
	myPrintln("4 function setIndividualCap(uint256 newMax) external onlyAdmin")
	myPrintln("5 function checkCap(uint256 amount0,uint256 amount1) public view returns(uint256)")
	myPrintln("6 function sweep( address _token, uint256 amount) external  onlyAdmin")

}

func ApproveStrategy() {
	/** Things to check
	// valut factory
	// uni factory
	// admin
	// team
	// WETH
	// CETH
	// uni portion, comp portion
	//
	*/

}

func ViaStratUniCompPublicList() {
	_, viaStratInstance, _ := GetInstance3()

	myTitle("ViaStratUniCompPublicList")
	myPrintln("--Strat Uni Comp address:", Network.VaultStrat)

	p2, _ := viaStratInstance.Protocol(&bind.CallOpts{})
	p3, _ := viaStratInstance.Pool(&bind.CallOpts{})
	p4, _ := viaStratInstance.Token0(&bind.CallOpts{})
	p5, _ := viaStratInstance.Token1(&bind.CallOpts{})
	p7, _ := viaStratInstance.TickSpacing(&bind.CallOpts{})
	p8, _ := viaStratInstance.CLow(&bind.CallOpts{})
	p9, _ := viaStratInstance.CHigh(&bind.CallOpts{})
	//	p10, _ := viaStratInstance.UniPortion(&bind.CallOpts{})
	//	p11, _ := viaStratInstance.CompPortion(&bind.CallOpts{})
	p15, _ := viaStratInstance.ProtocolFeeRate(&bind.CallOpts{})
	//	p6, _ := viaStratInstance.CTOKEN(&bind.CallOpts{}, common.HexToAddress(Network.TokenA))
	//	p17, _ := viaStratInstance.CTOKEN(&bind.CallOpts{}, common.HexToAddress(Network.TokenB))
	p18, _ := viaStratInstance.MotivatorFeeRate(&bind.CallOpts{})
	//	p20, _ := viaStratInstance.GetQuoteAtTick(&bind.CallOpts{}, big.NewInt(82894), p19, common.HexToAddress(Network.TokenA), common.HexToAddress(Network.TokenB))

	//p20, _ := viaStratInstance.GetQuoteAtTick(&bind.CallOpts{}, big.NewInt(-174770), big.NewInt(1e18), common.HexToAddress(Network.TokenA), common.HexToAddress(Network.TokenB))
	p21, _ := viaStratInstance.WETH(&bind.CallOpts{})
	p22, _ := viaStratInstance.Factory(&bind.CallOpts{})

	myPrintln("\n", ">>>>>>via Strat  public attributes <<<<<<<<<<<", "\n")
	myPrintln("Factory ", p22)
	myPrintln("Team", p2)
	myPrintln("Pool", p3)
	myPrintln("Token0", p4)
	myPrintln("Token1", p5)
	myPrintln("WETH", p21)
	//	myPrintln("Ctoken0", p6)
	//	myPrintln("Ctoken1", p17)
	myPrintln("TickSpacing", p7)
	myPrintln("CLow", p8)
	myPrintln("CHigh", p9)
	//	myPrintln("UniPortion", p10)
	//	myPrintln("CompPortion", p11)
	myPrintln("ProtocolFeeRAte", p15)
	myPrintln("DecenterFeeRates", p18)
	//	myPrintln("GetQuoteAtTick", p20)

	myPrintln("\n", ">>>>>>via Strat  public functions <<<<<<<<<<<", "\n")

	myPrintln("2 function rebalance(int24 newLow, int24 newHigh) external nonReentrant")
	myPrintln("3 function getPrice() public view returns(uint256)")
	myPrintln("4 function getTwap() public view returns (int24 tick)")
	myPrintln("5 function getQuoteAtTick( int24 tick, uint128 baseAmount,  address baseToken, address quoteToken ) public pure returns (uint256 quote)")
	myPrintln("7 function setMaxTwapDeviation(int24 _maxTwapDeviation) external onlyCreator")
	myPrintln("8 function setTwapDuration(uint32 _twapDuration) external onlyCreator")
	myPrintln("9 function setUniPortionRatio(uint8 ratio) external onlyCreator")
	myPrintln("10 function setCompPortionRatio(uint8 ratio) external onlyCreator")
	myPrintln("11 function setTeam(address _team) external onlyAdmin")
	myPrintln("12 function alloc() public returns ( bool )")
	myPrintln("15 function calcShares( uint256 totalSupply, uint256 amountIn0, uint256 amountIn1) public returns(uint256 shares, uint256 amount0, uint256 amount1)")
	myPrintln("16 function getUniAmounts(int24 tickLower, int24 tickUpper)  public view returns (uint256 amount0, uint256 amount1)")

}

func ViaVaultPublicFunc(funcId int, params ...interface{}) {

	if funcId == 1 { // deposit
		doDeposit(params[0], params[1])
	}

	if funcId == 2 { // withdraw
		doWithdraw(params[0])
	}

}

func doDeposit(params ...interface{}) {

	amount0, Ok := params[0].(int64)
	amount1, Ok := params[1].(int64)
	if !Ok {
		log.Println("invalid parameter", params)
	}

	_, _, viaVaultInstance := GetInstance3()

	var maxToken0 = PowX(99999, int(Token[0].Decimals)) //new(big.Int).SetString("900000000000000000000000000000", 10)
	var maxToken1 = PowX(99999, int(Token[1].Decimals)) //new(big.Int).SetString("900000000000000000000000000000", 10)

	ApproveToken(common.HexToAddress(Network.TokenA), maxToken0, Network.Vault)
	Sleep(100)
	ApproveToken(common.HexToAddress(Network.TokenB), maxToken1, Network.Vault)
	Sleep(100)

	NonceGen()
	tx, err := viaVaultInstance.Deposit(Auth, big.NewInt(amount0), big.NewInt(amount1))
	if err != nil {
		log.Println("deposit err: ", err)
	}
	TxConfirm(tx.Hash())

}

func doWithdraw(params ...interface{}) {

	percent, Ok := params[0].(uint8)
	if !Ok {
		log.Println("invalid parameter", params)
	}

	_, _, viaVaultInstance := GetInstance3()

	NonceGen()
	tx, err := viaVaultInstance.Withdraw(Auth, uint8(percent))
	if err != nil {
		log.Println("deposit err: ", err)
	}
	TxConfirm(tx.Hash())

}

// via stratUniComp function

func ViaStratUniCompPublicFunc(funcId int, params ...interface{}) {

	if funcId == 1 { // emergency
		//doDeposit(params[0], params[1])
		myPrintln("todo")
	}

	if funcId == 2 { // rebalance
		myPrintln("todo")
	}
	if funcId == 3 {
		GetPriceStratCall()
	}

}

// func GetViaVaultAddress() {
// 	_, viaStratInstance, _ := GetInstance3()
// 	addr, _ := viaStratInstance.Vault(&bind.CallOpts{})
// 	myPrintln(addr)
// }

func GetPriceStratCall() {

	_, viaStratInstance, _ := GetInstance3()

	NonceGen()
	price, err := viaStratInstance.GetPrice(&bind.CallOpts{})
	if err != nil {
		log.Println("getprice err: ", err)
	}
	myPrintln("fun getPrice:", price)

}

// func GetTVL() {

// 	_, stratInstance, vaultInstance := GetInstance3()
// 	NonceGen()

// 	tvl, err := vaultInstance.GetTVL(&bind.CallOpts{})
// 	if err != nil {
// 		log.Println("gettvl err: ", err)
// 	}
// 	myPrintln("fun gettvl:", tvl)

// }
func GetTotalAmounts() {

	_, viaStratInstance, _ := GetInstance3()

	NonceGen()
	total0, total1, err := viaStratInstance.GetTotalAmounts(&bind.CallOpts{})
	if err != nil {
		log.Println("gettickprice err: ", err)
	}

	myPrintln("totalamounts in strat:", total0, total1)

}

// func GetCompAmounts() {

// 	_, viaStratInstance, _ := GetInstance3()

// 	NonceGen()
// 	total0, total1, err := viaStratInstance.GetAmountsInComp(&bind.CallOpts{})
// 	if err != nil {
// 		log.Println("GetCompAmount err: ", err)
// 	}

// 	myPrintln("GetCompAmount :", total0, total1)

// }

// func GetTickPrice() {

// 	_, viaStratInstance, _ := GetInstance3()

// 	NonceGen()
// 	price, err := viaStratInstance.GetTickPrice(&bind.CallOpts{})
// 	if err != nil {
// 		log.Println("gettickprice err: ", err)
// 	}
// 	myPrintln("fun gettickPrice:", price)

// }

// func GetPriceX96FromSqrtPriceX96(pool string, twapInterval int) {

// 	_, viaStratInstance, _ := GetInstance3()

// 	NonceGen()
// 	sqrtTwapx96, err := viaStratInstance.GetSqrtTwapX96(&bind.CallOpts{}, common.HexToAddress(pool), uint32((twapInterval)))
// 	if err != nil {
// 		log.Println("getsqrtTwapx96 err: ", err)
// 	}
// 	myPrintln("sqrtTwapx96:", sqrtTwapx96)

// 	priceFromSqrtPriceX96, err := viaStratInstance.GetPriceX96FromSqrtPriceX96(&bind.CallOpts{}, sqrtTwapx96)
// 	if err != nil {
// 		log.Println("getsqrtTwapx96 err: ", err)
// 	}

// 	myPrintln("price FromSqrtPriceX96:", priceFromSqrtPriceX96)

// }

// func GetTwap() {

// 	_, viaStratInstance, _ := GetInstance3()

// 	NonceGen()
// 	twap, err := viaStratInstance.GetTwap(&bind.CallOpts{})
// 	if err != nil {
// 		log.Println("getTwap err: ", err)
// 	}
// 	myPrintln("fun getTwap:", twap)

// }

func GetQuoteAtTick(pool string, twapInterval int) {

	_, viaStratInstance, _ := GetInstance3()

	NonceGen()
	twap, err := viaStratInstance.GetTwap(&bind.CallOpts{}, common.HexToAddress(pool), uint32(twapInterval))
	if err != nil {
		log.Fatal("getTwap err: ", err)
	}
	myPrintln("fun getTwap:", twap)

	quote, err := viaStratInstance.GetQuoteAtTick(&bind.CallOpts{}, twap, big.NewInt(1e18), common.HexToAddress(Network.TokenA), common.HexToAddress(Network.TokenB))
	if err != nil {
		log.Fatal("getquote err: ", err)
	}
	myPrintln("fun getquote:", quote)

}

func CalcPositionShares(a0 int, a1 int) {
	_, _, vaultInstance := GetInstance3()

	NonceGen()
	tx, err := vaultInstance.CalcPositionShares(Auth, big.NewInt(int64(a0)), big.NewInt(int64(a1)))
	if err != nil {
		log.Println("calcPositionshares err: ", err)
	}
	TxConfirm(tx.Hash())

}

// func CheckCap(a0 int, a1 int) {
// 	_, _, vaultInstance := GetInstance3()
// 	p, err := vaultInstance.CheckCap(&bind.CallOpts{}, big.NewInt(int64(a0)), big.NewInt(int64(a1)))
// 	if err != nil {
// 		log.Println("checkCap err: ", err)
// 	}
// 	myPrintln("cap:", p)

// }

func IsContract(_contract string) {
	factoryInst, _, _ := GetInstance3()
	size, err := factoryInst.IsContract(&bind.CallOpts{}, common.HexToAddress(_contract))
	if err != nil {
		log.Println("iscontract err: ", err)
	}
	myPrintln("size:", size)

}

func MoveFunds() {

	_, stratInstance, _ := GetInstance3()

	myTitle(">>>>>>>> MoveFunds to strat<<<<<<<<<<<<< ")

	if tx, err := stratInstance.MoveFunds(Auth); err == nil {
		TxConfirm(tx.Hash())
	} else {
		log.Fatal("movefunds caller=strat", err)
	}

}

func CallFunds() {

	_, _, vaultInstance := GetInstance3()

	myTitle(">>>>>>>> Call Funds from vault <<<<<<<<<<<<< ")

	if tx, err := vaultInstance.CallFunds(Auth); err == nil {
		TxConfirm(tx.Hash())
	} else {
		log.Fatal("movefunds", err)
	}

}
func ViaCommand(cmd string, params []string) {

	if cmd == "lookup" {
		fmt.Println(cmd, ", ", params)
	}

	CmdHandler()

	a1, a2, a3 := GetInstance3()

	_, _, _ = a1, a2, a3

}

func CmdHandler() {

}

func Event(contractAddr string, eventNameString string, fromBlock int, toBlock int) {

	var eventNames []string
	if eventNameString != "" {
		eventNames = strings.Split(eventNameString, ",")
	}
	if toBlock == 0 { //listen
		EventListen(contractAddr, eventNames)
	} else { // search in blocks
		EventFiltered(contractAddr, fromBlock, toBlock, eventNameString)
	}
}

func EventListen(contract string, eventNames []string) {

	contractAddress := common.HexToAddress(contract)

	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)

	sub, err := EthClientWS.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Println("event subcribe:", err)
	}

	var contractAbi *abi.ABI
	if contract == Network.Vault {
		//contractAbi, err := abi.JSON(strings.NewReader(string(ViaVault.ApiABI)))
		contractAbi, err = ViaVault.ApiMetaData.GetAbi()
	} else if contract == Network.VaultStrat {
		contractAbi, err = VaultStrat.ApiMetaData.GetAbi()
	} else if contract == Network.VaultFactory {
		contractAbi, err = VaultFactory.ApiMetaData.GetAbi()
	} else {
		log.Fatal("invalid contract. no go package")
	}
	if err != nil {
		log.Println(err)
	}

	myPrintln("listening event:", eventNames)

	for {
		select {
		case err := <-sub.Err():
			log.Println("sel1:", err)
		case vLog := <-logs:
			for _, eventName := range eventNames {

				event, err := contractAbi.Unpack(eventName, vLog.Data)
				if err != nil {
					log.Println("contractAbi.Unpack Error:", err, "-> ", eventName)
					myPrintln(vLog)
				} else {
					myPrintln(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>\n")
					myPrintln(event)
					myPrintln("<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<\n")

					eventBlockInfo(eventName, int64(vLog.BlockNumber))

					for j, v := range event {
						fmt.Println("event[", j, "]=", v)
					}
					myPrintln()

					myPrintln("Topics: (", len(vLog.Topics), ")")
					parseTopic(vLog.Topics)

				}

			}

		}

	}

}

// fileter e.g.
func EventFiltered(contract string, _from int, _to int, eventName string) {

	contractAddress := common.HexToAddress(contract)

	fromBlock := big.NewInt(int64(_from))
	toBlock := big.NewInt(int64(_to))
	query := ethereum.FilterQuery{
		FromBlock: fromBlock,
		ToBlock:   toBlock,
		Addresses: []common.Address{
			contractAddress,
		},
	}

	logs, err := EthClientWS.FilterLogs(context.Background(), query)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(" Block from ", fromBlock, " to:", toBlock)

	//var contractAbi abi.ABI
	var contractAbi *abi.ABI
	if contract == Network.Vault {
		//	contractAbi, err = abi.JSON(strings.NewReader(string(ViaVault.ApiABI)))
		contractAbi, err = ViaVault.ApiMetaData.GetAbi()
	} else if contract == Network.VaultStrat {
		//contractAbi, err = abi.JSON(strings.NewReader(string(VaultStrat.ApiABI)))
		contractAbi, err = VaultStrat.ApiMetaData.GetAbi()
	} else if contract == Network.VaultFactory {
		//contractAbi, err = abi.JSON(strings.NewReader(string(VaultFactory.ApiABI)))
		contractAbi, err = VaultFactory.ApiMetaData.GetAbi()
	} else {
		log.Fatal("invalid contract. no go package")
	}

	if err != nil {
		log.Println("contractAbi->", err)
	}

	myPrintln(">>>>>>> logs>>>>>>>>>\n", logs)

	for _, vLog := range logs {

		//event, err := contractAbi.Unpack(eventName, vLog.Data)
		event := struct {
			Shares  *big.Int
			Amount0 *big.Int
			Amount1 *big.Int
		}{}
		err = contractAbi.UnpackIntoInterface(&event, eventName, vLog.Data)
		if err != nil {
			log.Println("contractAbi.Unpack error:", err, "-> ", eventName)
		}

		myPrintln(">>>>>>>>>the log>>>>>>>\n", event)

		//eventBlockInfo(eventName, int64(vLog.BlockNumber))

		// for j, v := range event {
		// 	fmt.Println("event[", j, "]=", v)
		// }
		// fmt.Println()

		myPrintln("Topics: (", len(vLog.Topics), ")")
		parseTopic(vLog.Topics)
		//		typeof(vLog.Data)
		//		typeof(event)

	}

}

func Events(eventName string, _from int, _to int) {

	fromBlock := big.NewInt(int64(_from))
	toBlock := big.NewInt(int64(_to))
	query := ethereum.FilterQuery{
		FromBlock: fromBlock,
		ToBlock:   toBlock,
		Addresses: []common.Address{
			common.HexToAddress(Network.VaultFactory),
			common.HexToAddress(Network.Vault),
			common.HexToAddress(Network.VaultStrat),
		},
	}

	logs, err := EthClientWS.FilterLogs(context.Background(), query)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(" Block from ", fromBlock, " to:", toBlock)

	//var contractAbi abi.ABI
	var contractAbi *abi.ABI
	contractAbi, err = ViaVault.ApiMetaData.GetAbi()

	if err != nil {
		log.Println(" ViaVault.ApiMetaData.GetAbi error ", err)
	}

	Events := map[string]interface{}{}
	for _, vLog := range logs {
		err = contractAbi.UnpackIntoMap(Events, eventName, vLog.Data)
		if err != nil {
			//log.Println(": UnpackIntoMap error:", err, "-> ", eventName)
			continue
		}

		if len(Events) > 0 {
			eventBlockInfo(eventName, int64(vLog.BlockNumber))

			myPrintln("Topics: Array(", len(vLog.Topics), ")")

			parseTopic(vLog.Topics)

			myPrintln("Logs: ")

			myPrintln(Events)
		}

	}

}

func typeof(v interface{}) {
	fmt.Println(fmt.Sprintf("%T", v))
}

func eventBlockInfo(eventName string, blockNumber int64) {
	block, err := EthClient.BlockByNumber(context.Background(), big.NewInt(int64(blockNumber)))
	if err != nil {
		log.Fatal("block ", err)
	}
	myPrintln("## event: ", eventName, " | event time:", time.Unix(int64(block.Time()), 0).Format("2006.01.02 15:04:05"), " | blocknumber:", blockNumber)

}

func parseTopic(topic []common.Hash) {

	for k := 0; k < len(topic); k++ {

		//hint := hexToInt(topic[k].String())
		fmt.Println(k, ": ", topic[k])
		//fmt.Println("hexToInt:", hint)
	}

}

func hexToInt(hexaString string) int64 {

	// replace 0x or 0X with empty String
	numberStr := strings.Replace(hexaString, "0x", "", -1)
	numberStr = strings.Replace(numberStr, "0X", "", -1)

	output, err := strconv.ParseInt(numberStr, 16, 64)
	if err != nil {
		log.Println("WARNING: hexToInt error. can be ignored")
	}

	//fmt.Printf("Output %d", output)
	return output

}
