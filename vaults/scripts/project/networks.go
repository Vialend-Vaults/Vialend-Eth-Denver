package include

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

/// 0: mainnet (or forked local), 1: local, 2: local , 3: gorlie, 4: gorlie,  5: goreli , 6: rinkeby, 7: ROPSTEN or ROPSTENLOCAL
var _USE_GETH = false // use geth directly without forking from ganache
var _LOAD_ENV = true  // if false, using hard coded data in network.
var Networkid = 6
var Account = 0
var ProviderSortId = 0

// var VaultFactory = "0xa8Fe7fd56DA477C0b0BB3729eE557f57a24879e1"
// var VIASTRATEGY = "0xd0f66761c2199E276721DcE320d36c9b8387c442"
// var VIAVAULT = "0x6d91365Cd1c987F8F8C2Dfb83CA90EF1B76D2E0D"

type TokenStruct struct {
	Name           string
	Symbol         string
	Decimals       uint8
	MaxTotalSupply *big.Int
}

type LendingStruct struct {
	WETH           string
	CETH           string
	DAI            string
	CDAI           string
	ETH            string
	WBTC           string
	CWBTC          string
	USDC           string
	CUSDC          string
	USDT           string
	CUSDT          string
	OSQTH          string
	OSQTHWETHPOOL  string
	ChainLinkProxy string
	SqthController string
	ATOKEN_USDC    string
	AAVE_USDC      string
	AAVE_ETH       string
}
type AccountStruct struct {
	Address    string
	PrivateKey string
}

type Params struct {
	ProviderUrl      []string
	Factory          string
	Callee           string
	PrivateKey       [11]string
	TokenA           string
	TokenB           string
	CTOKEN0          string
	CTOKEN1          string
	VaultFactory     string
	PendingTime      time.Duration
	Pool             string
	SwapRouter       string
	BonusToken       string
	Vault            string
	FeeTier          int64
	VaultBridge      string
	VaultStrat       string
	VaultHedge       string
	LendingContracts LendingStruct
	Accounts         []AccountStruct
	SwapHelper       string
}

type Indi int

const (
	Envelope Indi = iota
	Bollinger
	TMA
)

var Token [2]TokenStruct

//var Decimals0, Decimals1 uint8

var Network = Networks[Networkid]

var EthClientWS *ethclient.Client //GetEthClient(Networkid, 1)

var EthClient *ethclient.Client //GetEthClient(Networkid, 0)

var Auth *bind.TransactOpts

var FromAddress common.Address

var DEBUG = true

type Info struct {
	Name  string
	Value string
}

var InfoString []Info

var Networks = [...]Params{

	{ ///0 mainnet
		//[]string{},
		//[]string{"http://192.168.0.12:8546"}, /// direct main

		[]string{"http://127.0.0.1:8545", "ws://127.0.0.1:8545"}, // fork throu geth 192.168.0.12:8546
		//[]string{"http://192.168.0.223:8545"}, // ganache-cli in centos

		"0x1F98431c8aD98523631AE4a59f267346ea31F984", //factory
		"0xEcA3eDfD09435C2C7D2583124ca9a44f82aF1e8b", //callee
		[11]string{
			"b8c1b5c1d81f9475fdf2e334517d29f733bdfa40682207571b12fc1142cbf329",
			"5c2313d8a6b81a83ad1df1bf12a193cbc51d5de84a000db734fd7a05aa63e5a2",
		}, //privatekeys

		"0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48", // usdc 	tokenB
		"0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2", // weth  tokenA
		"0x39AA39c021dfbaE8faC545936693aC917d5E7563", //cusdc
		"0x4Ddc2D193948926D02f9B1fE9e1daa0718270ED5", //cweth
		"", //VaultFactory
		10, //pendingtime  local fork
		"0x8ad599c3A0ff1De082011EFDDc58f1908eb6e6D8", // pool
		"0x68b3465833fb72A70ecDF485E0e4C7bD8665Fc45", // swap router
		"", // bonus token
		"0xE8a6f04DEF6d9fa9F02982b8E1d8Eea835fB7668", //vault address  -> ganache-cli forked main
		3000,
		"", // VaultBridge
		"", //VaultStrat
		"", //VaultHedge
		LendingStruct{
			WETH:  "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
			WBTC:  "0x2260fac5e5542a773aa44fbcfedf7c193bc2c599",
			CWBTC: "0xC11b1268C1A384e55C48c2391d8d480264A3A7F4",
			USDC:  "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
			CUSDC: "0x39AA39c021dfbaE8faC545936693aC917d5E7563",
			CETH:  "0x4Ddc2D193948926D02f9B1fE9e1daa0718270ED5",
			DAI:   "0x6B175474E89094C44Da98b954EedeAC495271d0F",
			CDAI:  "0x5d3a536E4D6DbD6114cc1Ead35777bAB948E3643",
			USDT:  "0xdAC17F958D2ee523a2206206994597C13D831ec7",
			CUSDT: "0xf650C3d88D12dB855b8bf7D11Be6C55A4e07dCC9",
			//OSQTH:          "0x921c384F79de1BAe96d6f33E3E5b8d0B2B34cb68",	// ??
			OSQTH:          "0xf1b99e3e573a1a9c5e6b2ce818b617f0e664e86b", // osqth on mainnet
			OSQTHWETHPOOL:  "0x82c427AdFDf2d245Ec51D8046b41c4ee87F0d29C", // osqth/weth uni pool
			ChainLinkProxy: "0x986b5E1e1755e3C2440e960477f25201B0a8bbD4", // usdc / eth chain link
			SqthController: "0x64187ae08781B09368e6253F9E94951243A493D5",
		},
		[]AccountStruct{},
		"", //swaphelper
	},

	{ /// 1 local usdc/usdt
		[]string{"http://127.0.0.1:7545", "ws://127.0.0.1:7545"},
		"0x0c8D15944A4f799D678029523eC1F82c84b85F32", //factory
		"0x210FA31C72D9F020D16BF948e54F108D1C688f81", //callee
		[11]string{},
		"0x59Cd9D486a8fA9b39F715915743997daA12d138e", //tokenB usdt
		"0x9D96eC63f96A4E985e227BF520dD742315AB77c7", //tokenA usdc
		"", //ctoken0
		"", //ctoken1
		"0xeBb29c07455113c30810Addc123D0D7Cd8637aea", //VaultFactory
		10,
		"0x96Dd142387281a16F72962CCb659cAc67D73d882", // pool
		"", // swap router
		"0xD0d1E195c613Cb6eea9308daB69661CAF9760eF9", // bonus token
		"0xcB0b392e747C101Ed949247730eC3aa6A75E4D3B", //vault address
		3000, // fee
		"",   // VaultBridge
		"0xdB11518974087276048364aef0596637527ddDBd", //VaultStrat
		"", //VaultHedge

		LendingStruct{DAI: "ASDF", CDAI: "ASD"},
		[]AccountStruct{},
		"", //swaphelper
	},

	{ /// 2 local weth/usdc
		[]string{"http://127.0.0.1:7545", "http://127.0.0.1:8545"},
		"0x0c8D15944A4f799D678029523eC1F82c84b85F32", //factory
		"0x210FA31C72D9F020D16BF948e54F108D1C688f81", //callee
		[11]string{},
		"0xB73A78A3C493ACdbA893da9331ff39Fe4E59bFA3", //e weth1
		"0xd8F4E5E1cE1a2961b5fB401B8c2286549607B294", //e usdc1
		"", //ctoken0
		"", //ctoken1
		"0xeBb29c07455113c30810Addc123D0D7Cd8637aea", //VaultFactory
		10,
		"0xaEbc0569A8Ad476530d765dBE6308842Bd4D699c", // pool
		"", // swap router
		"0xD0d1E195c613Cb6eea9308daB69661CAF9760eF9", // bonus token
		"0x2723f0d5F2E60D1BF686B835e630C55453307eEA", //vault address
		3000, // fee
		"",   // VaultBridge
		"",   //VaultStrat
		"",   //VaultHedge

		LendingStruct{DAI: "ASDF", CDAI: "ASD"},
		[]AccountStruct{},
		"", //swaphelper

	},
	{ ///3  goerli admin test 1
		[]string{},

		"0x1F98431c8aD98523631AE4a59f267346ea31F984", //factory
		// "0xd648DB0713965e927963182Dc44D07D122a703ed", //callee
		"0xE97f1488F053251032ef358dE5b4188cD960D413", //callee
		[11]string{}, //6 arb01 0x8a01C3E04798D0B6D7423EaFF171932943FB9A8D

		// "0x48FCb48bb7F70F399E35d9eC95fd2A614960Dcf8", //tokenA eWeth
		// "0x6f38602e142D0Bd3BC162f5912535f543D3B73d7", //tokenB  eusdc

		"0xB4FBF271143F4FBf7B91A5ded31805e42b2208d6", // tokenA Weth
		"0xD87Ba7A50B2E7E660f678A895E4B72E7CB4CCd9C", //tokenB  usdc
		"0x20572e4c090f15667cF7378e16FaD2eA0e2f3EfF", // CETH
		"0xCEC4a43eBB02f9B80916F1c718338169d6d5C1F0", //CUSDC

		"READ FROM FILE", //VaultFactory
		40,               //time pending interval
		"0x04B1560f4F58612a24cF13531F4706c817E8A5Fe", //pool
		"0xe592427a0aece92de3edee1f18e0157c05861564", // uni swap router
		"0x3C3eF6Ad37F107CDd965C4da5f007526B959532f", // team  token
		"READ FROM FILE", //vault
		3000,             // fee
		"0x601D3992aB70ffEe3858730b12d94AAB35A4a60E", // VaultBridge

		"READ FROM FILE", //VaultStrat
		"",               //VaultHedge

		LendingStruct{
			WETH:           "0xB4FBF271143F4FBf7B91A5ded31805e42b2208d6",
			CETH:           "0x20572e4c090f15667cF7378e16FaD2eA0e2f3EfF",
			USDC:           "0xD87Ba7A50B2E7E660f678A895E4B72E7CB4CCd9C",
			CUSDC:          "0xCEC4a43eBB02f9B80916F1c718338169d6d5C1F0",
			CWBTC:          "0x6CE27497A64fFFb5517AA4aeE908b1E7EB63B9fF",
			WBTC:           "0xC04B0d3107736C32e19F1c62b2aF67BE61d63a05",
			CDAI:           "0x822397d9a55d0fefd20F5c4bCaB33C5F65bd28Eb",
			DAI:            "0xdc31Ee1784292379Fbb2964b3B9C4124D8F89C60",
			ChainLinkProxy: "0x326c977e6efc84e512bb9c30f76e30c160ed06fb", // usdc/eth

		},
		[]AccountStruct{},
		"", //swaphelper

	},
	{ ///4  goerli weth / usdc fee tier 0.1%
		[]string{}, ///  provider url

		"0x1F98431c8aD98523631AE4a59f267346ea31F984", //factory
		"0x9779ECDd5E44Ab4160687CEc04a8aB8D23C53E37", //callee
		//"0xE97f1488F053251032ef358dE5b4188cD960D413", //callee

		[11]string{},

		"0xB4FBF271143F4FBf7B91A5ded31805e42b2208d6", //  Weth
		"0xdc31Ee1784292379Fbb2964b3B9C4124D8F89C60", //  DAI
		"0x20572e4c090f15667cF7378e16FaD2eA0e2f3EfF", // CETH
		"0x822397d9a55d0fefd20F5c4bCaB33C5F65bd28Eb", //CDAI

		"0x8211Ed2ed4B7aEC0084f628B63B0367F5df5f3F6", //VaultFactory

		30, //time pending interval
		"0x1738f9aAB1d370a6d0fd56a18f113DbD9e1DCd4e", //pool  (weth,dai , 500)
		"0xe592427a0aece92de3edee1f18e0157c05861564", // uni swap router  not used
		"0x3C3eF6Ad37F107CDd965C4da5f007526B959532f", // team  token  not used

		"", //ViaVault
		//"0xc62e17fcdc58b042ae55a0eb5f6efb0af5e80ee7",

		500, // fee
		"0x601D3992aB70ffEe3858730b12d94AAB35A4a60E", // VaultBridge

		"", //VaultStrategy
		"", //VaultHedge

		LendingStruct{
			WETH:           "0xB4FBF271143F4FBf7B91A5ded31805e42b2208d6", ///on remix solc0.4.12,  injected web3 deployed by 0x2ee9... test admin,
			CETH:           "0x20572e4c090f15667cF7378e16FaD2eA0e2f3EfF",
			USDC:           "0xD87Ba7A50B2E7E660f678A895E4B72E7CB4CCd9C", // REFERENCE BELOW
			CUSDC:          "0xCEC4a43eBB02f9B80916F1c718338169d6d5C1F0",
			CWBTC:          "0x6CE27497A64fFFb5517AA4aeE908b1E7EB63B9fF",
			WBTC:           "0xC04B0d3107736C32e19F1c62b2aF67BE61d63a05",
			CDAI:           "0x822397d9a55d0fefd20F5c4bCaB33C5F65bd28Eb",
			DAI:            "0xdc31Ee1784292379Fbb2964b3B9C4124D8F89C60",
			ChainLinkProxy: "0x326c977e6efc84e512bb9c30f76e30c160ed06fb", // usdc/eth
		},
		[]AccountStruct{},
		"", //swaphelper

	},
	{ ///5  goerli wbtc / usdc fee tier 0.3%
		[]string{}, ///  provider url

		"0x1F98431c8aD98523631AE4a59f267346ea31F984", //factory
		//"0xd648DB0713965e927963182Dc44D07D122a703ed", //callee
		"0x9779ECDd5E44Ab4160687CEc04a8aB8D23C53E37", //callee

		[11]string{}, //6 arb01 0x8a01C3E04798D0B6D7423EaFF171932943FB9A8D

		"0xC04B0d3107736C32e19F1c62b2aF67BE61d63a05", //  Wbtc
		"0xdc31Ee1784292379Fbb2964b3B9C4124D8F89C60", //  DAI
		"0x6CE27497A64fFFb5517AA4aeE908b1E7EB63B9fF", // CWBTC
		"0x822397d9a55d0fefd20F5c4bCaB33C5F65bd28Eb", //CDAI

		"", //VaultFactory
		30, //time pending interval
		"0x4986413d9B0CAFD0166B2fd897048dcAD7d078F8", //pool  (wbtc,dai , 1000)
		//"0xF27864FC6cCb807968aa32248da292304e575119", //pool  (wbtc,dai , 3000)
		"", // uni swap router  not used
		"", // team  token  not used

		"0xBC6b6e273171C428d85cDdB23D344a8400B48441", //vault

		10000, // fee
		"0x033F3C5eAd18496BA462783fe9396CFE751a2342", // VaultBridge
		"", //VaultStrat
		"", //VaultHedge

		LendingStruct{
			WETH:           "0xB4FBF271143F4FBf7B91A5ded31805e42b2208d6", ///on remix solc0.4.12,  injected web3 deployed by 0x2ee9... test admin,
			CETH:           "0x20572e4c090f15667cF7378e16FaD2eA0e2f3EfF",
			USDC:           "0xD87Ba7A50B2E7E660f678A895E4B72E7CB4CCd9C", // REFERENCE BELOW
			CUSDC:          "0xCEC4a43eBB02f9B80916F1c718338169d6d5C1F0",
			CWBTC:          "0x6CE27497A64fFFb5517AA4aeE908b1E7EB63B9fF",
			WBTC:           "0xC04B0d3107736C32e19F1c62b2aF67BE61d63a05",
			CDAI:           "0x822397d9a55d0fefd20F5c4bCaB33C5F65bd28Eb",
			DAI:            "0xdc31Ee1784292379Fbb2964b3B9C4124D8F89C60",
			ChainLinkProxy: "0x326c977e6efc84e512bb9c30f76e30c160ed06fb", // usdc/eth
		},
		[]AccountStruct{},
		"", //swaphelper

	},

	{ ///6  rinkeby tester admin
		[]string{"http://192.168.0.223:8546", "ws://192.168.0.223:8546"}, ///  provider url

		"0x1F98431c8aD98523631AE4a59f267346ea31F984", //factory
		"", //callee
		[11]string{
			"3ff408e9e5342c7e3f540bba9345da9b0fdaabb7a6e4eaceb7de126a1ba55391",
			"f12e87b71f16f24f192941ecc3aff77e8ebcc7be01dd8d7d1245d7644ce8c93b",
		},
		"0x6aFA9B0F44f125028573Da29c1f00A89C9C8AC99", // usdc self deployed
		"0xaab1f2c9835cA4A15f6330967aBEfBa5Edf79F7d", // weth self deployed
		//"0x4DBCdF9B62e891a7cec5A2568C3F4FAF9E8Abe2b", //tokenA  usdc
		//"0xc778417E063141139Fce010982780140Aa0cD5Ab", //tokenB Weth
		"0xd6801a1DfFCd0a410336Ef88DeF4320D6DF1883e", //ctoken0
		"0x5B281A6DdA0B271e91ae35DE655Ad301C976edb1", //ctoken1
		"", //VaultFactory
		50, //time pending interval
		"", //pool 0x3c7fADe1921Bf9D8308D76d7B09cA54839cfF033", //pool tusdc/ tweth 0xBF93aB266Cd9235DaDE543fAd2EeC884D1cCFc0c // 0x3c7fADe1921Bf9D8308D76d7B09cA54839cfF033", eweth/eusdc //pool
		"0x68b3465833fb72A70ecDF485E0e4C7bD8665Fc45", // swap router
		"",  // tto  token
		"",  //vault address
		500, // fee
		"",  // VaultBridge
		"",  //VaultStrat
		"0x1F98431c8aD98523631AE4a59f267346ea31F984", //VaultHedge   to do

		LendingStruct{
			WETH:           "0xaab1f2c9835cA4A15f6330967aBEfBa5Edf79F7d",
			USDC:           "0x6aFA9B0F44f125028573Da29c1f00A89C9C8AC99",
			CUSDC:          "0x5B281A6DdA0B271e91ae35DE655Ad301C976edb1",
			CWBTC:          "0x0014F450B8Ae7708593F4A46F8fa6E5D50620F96",
			WBTC:           "0x577D296678535e4903D59A4C929B718e1D575e0A",
			CDAI:           "0x6D7F0754FFeb405d23C51CE938289d4835bE3b14",
			DAI:            "0x5592EC0cfb4dbc12D3aB100b257153436a1f0FEa",
			CETH:           "0xd6801a1DfFCd0a410336Ef88DeF4320D6DF1883e",
			OSQTH:          "0xb5b01bb5a34bcbeccc881ee217ada81a29dd2d6d", // xSqth - deployed squeeth contract on rinkeby, xSqth
			OSQTHWETHPOOL:  "0x51B5Ac5c8df5e982B2eb4Eb3823e774485b0e2fb", // xSqth/Weth pool created through uniswap v3 UI,
			ChainLinkProxy: "0xdCA36F27cbC4E38aE16C4E9f99D39b42337F6dcf", // usdc/eth
			SqthController: "0x64187ae08781B09368e6253F9E94951243A493D5", // fake for now
			ATOKEN_USDC:    "0xD624c05a873B9906e5F1afD9c5d6B2dC625d36c3", // ATOKEN_USDC
			AAVE_USDC:      "0x5B8B635c2665791cf62fe429cB149EaB42A3cEd8", // aaveUSDC
			AAVE_ETH:       "0x98a5F1520f7F7fb1e83Fe3398f9aBd151f8C65ed", // aaveETH weth

		},
		[]AccountStruct{},
		"", //swaphelper

	},
	{ ///7  ropsten
		[]string{}, ///  provider url

		"0x1F98431c8aD98523631AE4a59f267346ea31F984", //factory
		"", //callee
		[11]string{},

		"0x07865c6e87b9f70255377e024ace6630c1eaa37f", //tokenA  usdc
		"0xc778417e063141139fce010982780140aa0cd5ab", //tokenB  Weth
		"0x2973e69b20563bcc66dC63Bde153072c33eF37fe", //ctoken0
		"0x859e9d8a4edadfEDb5A2fF311243af80F85A91b8", //ctoken1
		"", //VaultFactory
		50, //time pending interval
		"0xee815CDC6322031952a095C6cc6FEd036Cb1F70d", //pool 0x3c7fADe1921Bf9D8308D76d7B09cA54839cfF033", //pool tusdc/ tweth 0xBF93aB266Cd9235DaDE543fAd2EeC884D1cCFc0c // 0x3c7fADe1921Bf9D8308D76d7B09cA54839cfF033", eweth/eusdc //pool
		"",   // swap router
		"",   // tto  token
		"",   //vault address
		3000, // fee tier
		"",   // VaultBridge
		"",   //VaultStrat
		"",   //VaultHedge

		LendingStruct{
			WETH:           "0xc778417e063141139fce010982780140aa0cd5ab",
			USDC:           "0x07865c6e87b9f70255377e024ace6630c1eaa37f",
			CUSDC:          "0x2973e69b20563bcc66dC63Bde153072c33eF37fe",
			CWBTC:          "0x541c9cB0E97b77F142684cc33E8AC9aC17B1990F",
			WBTC:           "0x442Be68395613bDCD19778e761f03261ec46C06D",
			CDAI:           "0xbc689667C13FB2a04f09272753760E38a95B998C",
			DAI:            "0x31F42841c2db5173425b5223809CF3A38FEde360",
			CETH:           "0x859e9d8a4edadfEDb5A2fF311243af80F85A91b8",
			ChainLinkProxy: "0x0", // usdc/eth

		},
		[]AccountStruct{},
		"", //swaphelper

	},
}

func ChangeAccount(account int) {

	Auth = GetSignature(Networkid, account)
	//fmt.Println("fromAddress changed: ", FromAddress)

}

func GetAddress(accId int) common.Address {

	///get fromAddress
	privateKey, err := crypto.HexToECDSA(Network.PrivateKey[accId])

	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	return crypto.PubkeyToAddress(*publicKeyECDSA)
}

func AddSettingString(name string, value string) {

	InfoString = append(InfoString, Info{name, value})

}

// 0: https, 1: wss
func GetEthClient(nid int, sortId int) *ethclient.Client {

	Networkid = nid
	Network = Networks[Networkid]

	c, err := ethclient.Dial(Network.ProviderUrl[sortId])
	if err != nil {
		myPrintln("networkid:", Networkid)
		myPrintln("Network.ProviderUrl[sortId]:", Network.ProviderUrl[sortId], sortId)

		//log.Fatal("getEthClient err:", err)
		log.Fatal("getEthClient err:", err)
	}

	return c

}

func GetSignature(nid int, accId int) *bind.TransactOpts {

	Account = accId

	privateKey, err := crypto.HexToECDSA(Network.PrivateKey[accId])

	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	FromAddress = crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Println("signed by ", FromAddress)

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Value = big.NewInt(0)          // in wei
	auth.GasLimit = uint64(8721975)     // in gwei
	auth.GasPrice = big.NewInt(4 * 1e9) // 2*1e9 = 2 gwei.  1 gwei = 1e9 wei

	return auth

}

func GetSignatureByKey(key string) (*bind.TransactOpts, string) {

	privateKey, err := crypto.HexToECDSA(key)

	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	publicAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Println("signed by ", FromAddress)

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Value = big.NewInt(0)          // in wei
	auth.GasLimit = uint64(8721975)     // in gwei
	auth.GasPrice = big.NewInt(4 * 1e9) // 2*1e9 = 2 gwei.  1 gwei = 1e9 wei

	return auth, publicAddress.Hex()

}
func NonceGen() {

	return

	Sleep(17000)
	nonce, err := EthClient.PendingNonceAt(context.Background(), FromAddress)

	if err != nil {
		log.Fatal("NonceGen ", err)
	}

	Auth.Nonce = big.NewInt(int64(nonce))

}

func PrivateToPublic(_privateKey string) common.Address {

	privateKey, err := crypto.HexToECDSA(_privateKey)
	if err != nil {
		log.Fatal("HexToECDSA ", err)
	}

	publicKey := privateKey.Public()

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	return crypto.PubkeyToAddress(*publicKeyECDSA)
}

func Init(nid int, acc int) {

	if nid == -1 {
		nid = Networkid
	}

	if acc == -1 {
		acc = Account

	}

	myTitle(" Current Network ")
	if _LOAD_ENV {
		myPrintln("loading from local env....")
		loadEnv(_USE_GETH)
	} else {

		myPrintln("loading from code....")
		fmt.Println("Env: NetworkId=", Networkid, ",EthClient=", Network.ProviderUrl)
	}

	myTitle(" initial ethereum clients")
	EthClient = GetEthClient(nid, 0)
	EthClientWS = GetEthClient(nid, 1)
	myPrintln("DONE")

	myTitle("initial default signature")
	Auth = GetSignature(nid, acc)

	_, Token[0].Name, Token[0].Symbol, Token[0].Decimals, Token[0].MaxTotalSupply = GetTokenInstance(Network.TokenA)
	_, Token[1].Name, Token[1].Symbol, Token[1].Decimals, Token[1].MaxTotalSupply = GetTokenInstance(Network.TokenB)

	myTitle("initial token pairs")
	myPrintln(Token[0].Symbol, ":", common.HexToAddress(Network.TokenA), ", Decimals:", Token[0].Decimals)
	myPrintln(Token[1].Symbol, ":", common.HexToAddress(Network.TokenB), ", Decimals:", Token[1].Decimals)

	fmt.Println(Cfg.Description)

	Network.VaultFactory = Cfg.Contracts.VAULT_FACTORY
	Network.VaultStrat = Cfg.Contracts.VAULT_STRATEGY
	Network.Vault = Cfg.Contracts.VAULT
	Network.VaultBridge = Cfg.Contracts.VAULT_BRIDGE
	Network.Callee = Cfg.Contracts.CALLEE
	Network.SwapHelper = Cfg.Contracts.SWAPHELPER

	myTitle("Loading contracts ")
	myPrintln("VaultFactory:", Network.VaultFactory)
	myPrintln("Vault Strategy:", Network.VaultStrat)
	myPrintln("Vault: ", Network.Vault)
	myPrintln("Vault bridge:", Network.VaultBridge)

	myPrintln("-")

}

func loadEnv(useGETH bool) {

	// NOTE: add new environment,need to restart liteide

	networkNames := map[int]string{0: "MAINNET", 1: "LOCAL", 2: "LOCAL", 3: "GOERLI", 4: "GOERLI", 5: "GOERLI", 6: "RINKEBY", 7: "ROPSTEN"}

	if useGETH {
		networkNames = map[int]string{0: "MAINNET", 1: "LOCAL", 2: "LOCAL", 3: "GOERLILOCAL", 4: "GOERLILOCAL", 5: "GOERLILOCAL", 6: "RINKEBYLOCAL", 7: "ROPSTENLOCAL"}
	}

	accountNames := map[int]string{0: "MAINNET", 1: "LOCAL", 2: "LOCAL", 3: "GOERLI", 4: "GOERLI", 5: "GOERLI", 6: "RINKEBY", 7: "ROPSTEN"}

	i := 0

	for i = 0; i < 8; i++ {
		Networks[i].ProviderUrl = strings.Split(os.Getenv("VIALEND_NETWORK_"+networkNames[i]+"_PROVIDER"), ";")
		//		fmt.Println(Networks[i].ProviderUrl[0])
		//		fmt.Println(Networks[i].ProviderUrl[1])

		j := 0

		//# print chosen network
		if i == Networkid {
			myPrintln(networkNames[i], Networks[i].ProviderUrl)
		}

		accounts := os.Getenv("VIALEND_" + accountNames[i] + "_ACCOUNT" + strconv.Itoa(j))

		for accounts != "" {

			v := os.Getenv("VIALEND_" + accountNames[i] + "_ACCOUNT" + strconv.Itoa(j))

			if len(strings.TrimSpace(v)) != 0 {
				k := strings.Split(v, ":")
				Networks[i].PrivateKey[j] = k[1]
				//fmt.Println(Networks[i].PrivateKey[j])
				j++

			} else {
				break
			}

		}
	}

}

func TxConfirm(tx common.Hash) {

	myPrintln("tx: ", tx.Hex())

	tr, err := EthClient.TransactionReceipt(context.Background(), tx)
	if err != nil {
		log.Println("TransactionReceipt:  ", err)
	}

	delay := time.Duration(5)

	for i := 0; i < 100; i++ {
		if err != nil {
			//log.Println("TxConfirm transactionReceipt Failed :  ", err)
			log.Println("tx receipt", err, ".. trying in ", delay, " seconds")
			time.Sleep(delay * time.Second)
		} else {
			break
		}
		tr, err = EthClient.TransactionReceipt(context.Background(), tx)
	}

	if tr.Status == 0 {
		log.Fatal("!!!!!!!! Failed tx ")
	}

	myPrintln("BlockNumber:", tr.BlockNumber)
	myPrintln("Status:", Stat2Str(tr.Status))
	myPrintln("CumulativeGasUsed:", tr.CumulativeGasUsed)
	myPrintln("GasUsed:", tr.GasUsed)

}

func setEnv() {
	// Set Environment Variables
	//
}
