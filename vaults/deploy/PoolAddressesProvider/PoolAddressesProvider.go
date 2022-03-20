// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package api

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ApiMetaData contains all meta data concerning the Api contract.
var ApiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"marketId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"ACLAdminUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"ACLManagerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"AddressSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"proxyAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldImplementationAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newImplementationAddress\",\"type\":\"address\"}],\"name\":\"AddressSetAsProxy\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"oldMarketId\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"newMarketId\",\"type\":\"string\"}],\"name\":\"MarketIdSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"PoolConfiguratorUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"PoolDataProviderUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"PoolUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"PriceOracleSentinelUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"PriceOracleUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"proxyAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementationAddress\",\"type\":\"address\"}],\"name\":\"ProxyCreated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getACLAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getACLManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"getAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMarketId\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPoolConfigurator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPoolDataProvider\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPriceOracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPriceOracleSentinel\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAclAdmin\",\"type\":\"address\"}],\"name\":\"setACLAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAclManager\",\"type\":\"address\"}],\"name\":\"setACLManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"setAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"newImplementationAddress\",\"type\":\"address\"}],\"name\":\"setAddressAsProxy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newMarketId\",\"type\":\"string\"}],\"name\":\"setMarketId\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPoolConfiguratorImpl\",\"type\":\"address\"}],\"name\":\"setPoolConfiguratorImpl\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newDataProvider\",\"type\":\"address\"}],\"name\":\"setPoolDataProvider\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPoolImpl\",\"type\":\"address\"}],\"name\":\"setPoolImpl\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPriceOracle\",\"type\":\"address\"}],\"name\":\"setPriceOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPriceOracleSentinel\",\"type\":\"address\"}],\"name\":\"setPriceOracleSentinel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162001ffb38038062001ffb8339810160408190526200003491620003aa565b600080546001600160a01b0319163390811782556040519091829160008051602062001fdb833981519152908290a3506200006f8262000082565b6200007a816200018d565b5050620004d2565b600060018054620000939062000477565b80601f0160208091040260200160405190810160405280929190818152602001828054620000c19062000477565b8015620001125780601f10620000e65761010080835404028352916020019162000112565b820191906000526020600020905b815481529060010190602001808311620000f457829003601f168201915b5050855193945062000130936001935060208701925090506200029e565b5081604051620001419190620004b4565b604051809103902081604051620001599190620004b4565b604051908190038120907fe685c8cdecc6030c45030fd54778812cb84ed8e4467c38294403d68ba786082390600090a35050565b6000546001600160a01b03163314620001ed5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064015b60405180910390fd5b6001600160a01b038116620002545760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401620001e4565b600080546040516001600160a01b038085169392169160008051602062001fdb83398151915291a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b828054620002ac9062000477565b90600052602060002090601f016020900481019282620002d057600085556200031b565b82601f10620002eb57805160ff19168380011785556200031b565b828001600101855582156200031b579182015b828111156200031b578251825591602001919060010190620002fe565b50620003299291506200032d565b5090565b5b808211156200032957600081556001016200032e565b634e487b7160e01b600052604160045260246000fd5b60005b83811015620003775781810151838201526020016200035d565b8381111562000387576000848401525b50505050565b80516001600160a01b0381168114620003a557600080fd5b919050565b60008060408385031215620003be57600080fd5b82516001600160401b0380821115620003d657600080fd5b818501915085601f830112620003eb57600080fd5b81518181111562000400576200040062000344565b604051601f8201601f19908116603f011681019083821181831017156200042b576200042b62000344565b816040528281528860208487010111156200044557600080fd5b620004588360208301602088016200035a565b80965050505050506200046e602084016200038d565b90509250929050565b600181811c908216806200048c57607f821691505b60208210811415620004ae57634e487b7160e01b600052602260045260246000fd5b50919050565b60008251620004c88184602087016200035a565b9190910192915050565b611af980620004e26000396000f3fe608060405234801561001057600080fd5b50600436106101425760003560e01c806376d84ffc116100b8578063e4ca28b71161007c578063e4ca28b714610254578063e860accb14610267578063ed301ca91461026f578063f2fde38b14610282578063f67b184714610295578063fca513a8146102a857600080fd5b806376d84ffc146101f75780638da5cb5b1461020a578063a15644061461021b578063ca446dd91461022e578063e44e9ed11461024157600080fd5b80635dcc528c1161010a5780635dcc528c146101b15780635eb88d3d146101c4578063631adfca146101cc578063707cd716146101d4578063715018a6146101dc57806374944cec146101e457600080fd5b8063026b1d5f146101475780630e67178c1461016c57806321f8a72114610174578063530e784f14610187578063568ef4701461019c575b600080fd5b61014f6102b0565b6040516001600160a01b0390911681526020015b60405180910390f35b61014f6102c7565b61014f610182366004610fb7565b6102da565b61019a610195366004610fe5565b6102f5565b005b6101a46103b0565b6040516101639190611065565b61019a6101bf366004611078565b610442565b61014f6104e7565b61014f61050a565b61014f610529565b61019a610542565b61019a6101f2366004610fe5565b6105b6565b61019a610205366004610fe5565b610671565b6000546001600160a01b031661014f565b61019a610229366004610fe5565b610720565b61019a61023c366004611078565b6107b3565b61019a61024f366004610fe5565b61083b565b61019a610262366004610fe5565b6108ee565b61014f61099b565b61019a61027d366004610fe5565b6109b6565b61019a610290366004610fe5565b610a67565b61019a6102a33660046110be565b610b51565b61014f610b87565b60006102c2631413d3d360e21b6102da565b905090565b60006102c26820a1a62fa0a226a4a760b91b5b6000908152600260205260409020546001600160a01b031690565b6000546001600160a01b031633146103285760405162461bcd60e51b815260040161031f9061116f565b60405180910390fd5b6b50524943455f4f5241434c4560a01b600090815260026020527f740f710666bd7a12af42df98311e541e47f7fd33d382d11602457a6d540cbd6380546001600160a01b038481166001600160a01b03198316811790935560405191169283917f56b5f80d8cac1479698aa7d01605fd6111e90b15fc4d2b377417f46034876cbd9190a35050565b6060600180546103bf906111a4565b80601f01602080910402602001604051908101604052809291908181526020018280546103eb906111a4565b80156104385780601f1061040d57610100808354040283529160200191610438565b820191906000526020600020905b81548152906001019060200180831161041b57829003601f168201915b5050505050905090565b6000546001600160a01b0316331461046c5760405162461bcd60e51b815260040161031f9061116f565b6000828152600260205260408120546001600160a01b03169061048e84610ba1565b905061049a8484610c3e565b6040516001600160a01b038281168252808516919084169086907f3bbd45b5429b385e3fb37ad5cd1cd1435a3c8ec32196c7937597365a3fd3e99c9060200160405180910390a450505050565b60006102c27414149250d157d3d49050d31157d4d1539512539153605a1b6102da565b60006102c2702827a7a62fa1a7a72324a3aaa920aa27a960791b6102da565b60006102c26a20a1a62fa6a0a720a3a2a960a91b6102da565b6000546001600160a01b0316331461056c5760405162461bcd60e51b815260040161031f9061116f565b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b6000546001600160a01b031633146105e05760405162461bcd60e51b815260040161031f9061116f565b7414149250d157d3d49050d31157d4d1539512539153605a1b600090815260026020527f0d2c1bcee56447b4f46248272f34207a580a5c40f666a31f4e2fbb470ea53ab880546001600160a01b038481166001600160a01b03198316811790935560405191169283917f5326514eeca90494a14bedabcff812a0e683029ee85d1e23824d44fd14cd6ae79190a35050565b6000546001600160a01b0316331461069b5760405162461bcd60e51b815260040161031f9061116f565b6820a1a62fa0a226a4a760b91b600090815260026020527ffab167ad2009dcb80ee379700bb4bd029d97c1181ed9d961625632c8a6f051c680546001600160a01b038481166001600160a01b03198316811790935560405191169283917fe9cf53972264dc95304fd424458745019ddfca0e37ae8f703d74772c41ad115b9190a35050565b6000546001600160a01b0316331461074a5760405162461bcd60e51b815260040161031f9061116f565b600061075c631413d3d360e21b610ba1565b905061076f631413d3d360e21b83610c3e565b816001600160a01b0316816001600160a01b03167f90affc163f1a2dfedcd36aa02ed992eeeba8100a4014f0b4cdc20ea265a6662760405160405180910390a35050565b6000546001600160a01b031633146107dd5760405162461bcd60e51b815260040161031f9061116f565b60008281526002602052604080822080546001600160a01b031981166001600160a01b038681169182179093559251911692839186917f9ef0e8c8e52743bb38b83b17d9429141d494b8041ca6d616a6c77cebae9cd8b791a4505050565b6000546001600160a01b031633146108655760405162461bcd60e51b815260040161031f9061116f565b6c2220aa20afa82927ab24a222a960991b600090815260026020527fcd7944601aaa5cd7ccdae1bebec659e98c6aac8f12486b30e59db0d39698051f80546001600160a01b038481166001600160a01b03198316811790935560405191169283917fc853974cfbf81487a14a23565917bee63f527853bcb5fa54f2ae1cdf8a38356d9190a35050565b6000546001600160a01b031633146109185760405162461bcd60e51b815260040161031f9061116f565b6000610937702827a7a62fa1a7a72324a3aaa920aa27a960791b610ba1565b9050610957702827a7a62fa1a7a72324a3aaa920aa27a960791b83610c3e565b816001600160a01b0316816001600160a01b03167f8932892569eba59c8382a089d9b732d1f49272878775235761a2a6b0309cd46560405160405180910390a35050565b60006102c26c2220aa20afa82927ab24a222a960991b6102da565b6000546001600160a01b031633146109e05760405162461bcd60e51b815260040161031f9061116f565b6a20a1a62fa6a0a720a3a2a960a91b600090815260026020527f9edef266ef35fd0c6e131df0f31a330f3dd4c4d19dd31ed615c21d005c68116b80546001600160a01b038481166001600160a01b03198316811790935560405191169283917fb30efa04327bb8a537d61cc1e5c48095345ad18ef7cc04e6bacf7dfb6caaf5079190a35050565b6000546001600160a01b03163314610a915760405162461bcd60e51b815260040161031f9061116f565b6001600160a01b038116610af65760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b606482015260840161031f565b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b6000546001600160a01b03163314610b7b5760405162461bcd60e51b815260040161031f9061116f565b610b8481610e14565b50565b60006102c26b50524943455f4f5241434c4560a01b6102da565b6000818152600260205260408120546001600160a01b031680610bc75750600092915050565b6000819050806001600160a01b0316635c60da1b6040518163ffffffff1660e01b81526004016020604051808303816000875af1158015610c0c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c3091906111d9565b949350505050565b50919050565b6000828152600260205260408082205490513060248201526001600160a01b039091169190819060440160408051601f198184030181529190526020810180516001600160e01b031663189acdbd60e31b17905290506001600160a01b038316610da95730604051610caf90610f11565b6001600160a01b039091168152602001604051809103906000f080158015610cdb573d6000803e3d6000fd5b506000868152600260205260409081902080546001600160a01b0319166001600160a01b038416908117909155905163347d5e2560e21b81529194508493509063d1f5789490610d3190879085906004016111f6565b600060405180830381600087803b158015610d4b57600080fd5b505af1158015610d5f573d6000803e3d6000fd5b50505050836001600160a01b0316836001600160a01b0316867f4a465a9bd819d9662563c1e11ae958f8109e437e7f4bf1c6ef0b9a7b3f35d47860405160405180910390a4610e0d565b60405163278f794360e11b81528392506001600160a01b03831690634f1ef28690610dda90879085906004016111f6565b600060405180830381600087803b158015610df457600080fd5b505af1158015610e08573d6000803e3d6000fd5b505050505b5050505050565b600060018054610e23906111a4565b80601f0160208091040260200160405190810160405280929190818152602001828054610e4f906111a4565b8015610e9c5780601f10610e7157610100808354040283529160200191610e9c565b820191906000526020600020905b815481529060010190602001808311610e7f57829003601f168201915b50508551939450610eb893600193506020870192509050610f1e565b5081604051610ec7919061121a565b604051809103902081604051610edd919061121a565b604051908190038120907fe685c8cdecc6030c45030fd54778812cb84ed8e4467c38294403d68ba786082390600090a35050565b61088d8061123783390190565b828054610f2a906111a4565b90600052602060002090601f016020900481019282610f4c5760008555610f92565b82601f10610f6557805160ff1916838001178555610f92565b82800160010185558215610f92579182015b82811115610f92578251825591602001919060010190610f77565b50610f9e929150610fa2565b5090565b5b80821115610f9e5760008155600101610fa3565b600060208284031215610fc957600080fd5b5035919050565b6001600160a01b0381168114610b8457600080fd5b600060208284031215610ff757600080fd5b813561100281610fd0565b9392505050565b60005b8381101561102457818101518382015260200161100c565b83811115611033576000848401525b50505050565b60008151808452611051816020860160208601611009565b601f01601f19169290920160200192915050565b6020815260006110026020830184611039565b6000806040838503121561108b57600080fd5b82359150602083013561109d81610fd0565b809150509250929050565b634e487b7160e01b600052604160045260246000fd5b6000602082840312156110d057600080fd5b813567ffffffffffffffff808211156110e857600080fd5b818401915084601f8301126110fc57600080fd5b81358181111561110e5761110e6110a8565b604051601f8201601f19908116603f01168101908382118183101715611136576111366110a8565b8160405282815287602084870101111561114f57600080fd5b826020860160208301376000928101602001929092525095945050505050565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b600181811c908216806111b857607f821691505b60208210811415610c3857634e487b7160e01b600052602260045260246000fd5b6000602082840312156111eb57600080fd5b815161100281610fd0565b6001600160a01b0383168152604060208201819052600090610c3090830184611039565b6000825161122c818460208701611009565b919091019291505056fe60a060405234801561001057600080fd5b5060405161088d38038061088d83398101604081905261002f91610040565b6001600160a01b0316608052610070565b60006020828403121561005257600080fd5b81516001600160a01b038116811461006957600080fd5b9392505050565b6080516107df6100ae600039600081816101130152818161015801528181610211015281816103510152818161037a01526104a501526107df6000f3fe60806040526004361061004a5760003560e01c80633659cfe6146100545780634f1ef286146100745780635c60da1b14610087578063d1f57894146100b8578063f851a440146100cb575b6100526100e0565b005b34801561006057600080fd5b5061005261006f366004610586565b610108565b6100526100823660046105a8565b61014d565b34801561009357600080fd5b5061009c610204565b6040516001600160a01b03909116815260200160405180910390f35b6100526100c6366004610641565b610256565b3480156100d757600080fd5b5061009c610344565b6100e861039c565b61010661010160008051602061078a8339815191525490565b6103a4565b565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016141561014557610142816103c8565b50565b6101426100e0565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614156101f757610187836103c8565b6000836001600160a01b031683836040516101a3929190610703565b600060405180830381855af49150503d80600081146101de576040519150601f19603f3d011682016040523d82523d6000602084013e6101e3565b606091505b50509050806101f157600080fd5b50505050565b6101ff6100e0565b505050565b6000336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016141561024b575060008051602061078a8339815191525490565b6102536100e0565b90565b600061026e60008051602061078a8339815191525490565b6001600160a01b03161461028157600080fd5b6102ac60017f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbd610713565b60008051602061078a833981519152146102c8576102c8610738565b6102d182610408565b805115610340576000826001600160a01b0316826040516102f2919061074e565b600060405180830381855af49150503d806000811461032d576040519150601f19603f3d011682016040523d82523d6000602084013e610332565b606091505b50509050806101ff57600080fd5b5050565b6000336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016141561024b57507f000000000000000000000000000000000000000000000000000000000000000090565b61010661049a565b3660008037600080366000845af43d6000803e8080156103c3573d6000f35b3d6000fd5b6103d181610408565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b6104118161052e565b6104885760405162461bcd60e51b815260206004820152603b60248201527f43616e6e6f742073657420612070726f787920696d706c656d656e746174696f60448201527f6e20746f2061206e6f6e2d636f6e74726163742061646472657373000000000060648201526084015b60405180910390fd5b60008051602061078a83398151915255565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614156101065760405162461bcd60e51b815260206004820152603260248201527f43616e6e6f742063616c6c2066616c6c6261636b2066756e6374696f6e20667260448201527137b6903a343290383937bc3c9030b236b4b760711b606482015260840161047f565b6000813f7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a47081811480159061056257508115155b949350505050565b80356001600160a01b038116811461058157600080fd5b919050565b60006020828403121561059857600080fd5b6105a18261056a565b9392505050565b6000806000604084860312156105bd57600080fd5b6105c68461056a565b9250602084013567ffffffffffffffff808211156105e357600080fd5b818601915086601f8301126105f757600080fd5b81358181111561060657600080fd5b87602082850101111561061857600080fd5b6020830194508093505050509250925092565b634e487b7160e01b600052604160045260246000fd5b6000806040838503121561065457600080fd5b61065d8361056a565b9150602083013567ffffffffffffffff8082111561067a57600080fd5b818501915085601f83011261068e57600080fd5b8135818111156106a0576106a061062b565b604051601f8201601f19908116603f011681019083821181831017156106c8576106c861062b565b816040528281528860208487010111156106e157600080fd5b8260208601602083013760006020848301015280955050505050509250929050565b8183823760009101908152919050565b60008282101561073357634e487b7160e01b600052601160045260246000fd5b500390565b634e487b7160e01b600052600160045260246000fd5b6000825160005b8181101561076f5760208186018101518583015201610755565b8181111561077e576000828501525b50919091019291505056fe360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbca2646970667358221220945b5335f5b8f8c53d75a8973acd3c75c33f2aeb51528204a5e066bf34e7da4e64736f6c634300080a0033a264697066735822122070a421bde1c2171dfbb8efb1a623a3057dcf893a1099fdb6492ecd5343cd1b5264736f6c634300080a00338be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0",
}

// ApiABI is the input ABI used to generate the binding from.
// Deprecated: Use ApiMetaData.ABI instead.
var ApiABI = ApiMetaData.ABI

// ApiBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ApiMetaData.Bin instead.
var ApiBin = ApiMetaData.Bin

// DeployApi deploys a new Ethereum contract, binding an instance of Api to it.
func DeployApi(auth *bind.TransactOpts, backend bind.ContractBackend, marketId string, owner common.Address) (common.Address, *types.Transaction, *Api, error) {
	parsed, err := ApiMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ApiBin), backend, marketId, owner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Api{ApiCaller: ApiCaller{contract: contract}, ApiTransactor: ApiTransactor{contract: contract}, ApiFilterer: ApiFilterer{contract: contract}}, nil
}

// Api is an auto generated Go binding around an Ethereum contract.
type Api struct {
	ApiCaller     // Read-only binding to the contract
	ApiTransactor // Write-only binding to the contract
	ApiFilterer   // Log filterer for contract events
}

// ApiCaller is an auto generated read-only Go binding around an Ethereum contract.
type ApiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ApiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ApiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ApiSession struct {
	Contract     *Api              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ApiCallerSession struct {
	Contract *ApiCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ApiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ApiTransactorSession struct {
	Contract     *ApiTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApiRaw is an auto generated low-level Go binding around an Ethereum contract.
type ApiRaw struct {
	Contract *Api // Generic contract binding to access the raw methods on
}

// ApiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ApiCallerRaw struct {
	Contract *ApiCaller // Generic read-only contract binding to access the raw methods on
}

// ApiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ApiTransactorRaw struct {
	Contract *ApiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewApi creates a new instance of Api, bound to a specific deployed contract.
func NewApi(address common.Address, backend bind.ContractBackend) (*Api, error) {
	contract, err := bindApi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Api{ApiCaller: ApiCaller{contract: contract}, ApiTransactor: ApiTransactor{contract: contract}, ApiFilterer: ApiFilterer{contract: contract}}, nil
}

// NewApiCaller creates a new read-only instance of Api, bound to a specific deployed contract.
func NewApiCaller(address common.Address, caller bind.ContractCaller) (*ApiCaller, error) {
	contract, err := bindApi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ApiCaller{contract: contract}, nil
}

// NewApiTransactor creates a new write-only instance of Api, bound to a specific deployed contract.
func NewApiTransactor(address common.Address, transactor bind.ContractTransactor) (*ApiTransactor, error) {
	contract, err := bindApi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ApiTransactor{contract: contract}, nil
}

// NewApiFilterer creates a new log filterer instance of Api, bound to a specific deployed contract.
func NewApiFilterer(address common.Address, filterer bind.ContractFilterer) (*ApiFilterer, error) {
	contract, err := bindApi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ApiFilterer{contract: contract}, nil
}

// bindApi binds a generic wrapper to an already deployed contract.
func bindApi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ApiABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Api *ApiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Api.Contract.ApiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Api *ApiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.Contract.ApiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Api *ApiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Api.Contract.ApiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Api *ApiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Api.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Api *ApiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Api *ApiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Api.Contract.contract.Transact(opts, method, params...)
}

// GetACLAdmin is a free data retrieval call binding the contract method 0x0e67178c.
//
// Solidity: function getACLAdmin() view returns(address)
func (_Api *ApiCaller) GetACLAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getACLAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetACLAdmin is a free data retrieval call binding the contract method 0x0e67178c.
//
// Solidity: function getACLAdmin() view returns(address)
func (_Api *ApiSession) GetACLAdmin() (common.Address, error) {
	return _Api.Contract.GetACLAdmin(&_Api.CallOpts)
}

// GetACLAdmin is a free data retrieval call binding the contract method 0x0e67178c.
//
// Solidity: function getACLAdmin() view returns(address)
func (_Api *ApiCallerSession) GetACLAdmin() (common.Address, error) {
	return _Api.Contract.GetACLAdmin(&_Api.CallOpts)
}

// GetACLManager is a free data retrieval call binding the contract method 0x707cd716.
//
// Solidity: function getACLManager() view returns(address)
func (_Api *ApiCaller) GetACLManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getACLManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetACLManager is a free data retrieval call binding the contract method 0x707cd716.
//
// Solidity: function getACLManager() view returns(address)
func (_Api *ApiSession) GetACLManager() (common.Address, error) {
	return _Api.Contract.GetACLManager(&_Api.CallOpts)
}

// GetACLManager is a free data retrieval call binding the contract method 0x707cd716.
//
// Solidity: function getACLManager() view returns(address)
func (_Api *ApiCallerSession) GetACLManager() (common.Address, error) {
	return _Api.Contract.GetACLManager(&_Api.CallOpts)
}

// GetAddress is a free data retrieval call binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(bytes32 id) view returns(address)
func (_Api *ApiCaller) GetAddress(opts *bind.CallOpts, id [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getAddress", id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddress is a free data retrieval call binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(bytes32 id) view returns(address)
func (_Api *ApiSession) GetAddress(id [32]byte) (common.Address, error) {
	return _Api.Contract.GetAddress(&_Api.CallOpts, id)
}

// GetAddress is a free data retrieval call binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(bytes32 id) view returns(address)
func (_Api *ApiCallerSession) GetAddress(id [32]byte) (common.Address, error) {
	return _Api.Contract.GetAddress(&_Api.CallOpts, id)
}

// GetMarketId is a free data retrieval call binding the contract method 0x568ef470.
//
// Solidity: function getMarketId() view returns(string)
func (_Api *ApiCaller) GetMarketId(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getMarketId")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetMarketId is a free data retrieval call binding the contract method 0x568ef470.
//
// Solidity: function getMarketId() view returns(string)
func (_Api *ApiSession) GetMarketId() (string, error) {
	return _Api.Contract.GetMarketId(&_Api.CallOpts)
}

// GetMarketId is a free data retrieval call binding the contract method 0x568ef470.
//
// Solidity: function getMarketId() view returns(string)
func (_Api *ApiCallerSession) GetMarketId() (string, error) {
	return _Api.Contract.GetMarketId(&_Api.CallOpts)
}

// GetPool is a free data retrieval call binding the contract method 0x026b1d5f.
//
// Solidity: function getPool() view returns(address)
func (_Api *ApiCaller) GetPool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getPool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPool is a free data retrieval call binding the contract method 0x026b1d5f.
//
// Solidity: function getPool() view returns(address)
func (_Api *ApiSession) GetPool() (common.Address, error) {
	return _Api.Contract.GetPool(&_Api.CallOpts)
}

// GetPool is a free data retrieval call binding the contract method 0x026b1d5f.
//
// Solidity: function getPool() view returns(address)
func (_Api *ApiCallerSession) GetPool() (common.Address, error) {
	return _Api.Contract.GetPool(&_Api.CallOpts)
}

// GetPoolConfigurator is a free data retrieval call binding the contract method 0x631adfca.
//
// Solidity: function getPoolConfigurator() view returns(address)
func (_Api *ApiCaller) GetPoolConfigurator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getPoolConfigurator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPoolConfigurator is a free data retrieval call binding the contract method 0x631adfca.
//
// Solidity: function getPoolConfigurator() view returns(address)
func (_Api *ApiSession) GetPoolConfigurator() (common.Address, error) {
	return _Api.Contract.GetPoolConfigurator(&_Api.CallOpts)
}

// GetPoolConfigurator is a free data retrieval call binding the contract method 0x631adfca.
//
// Solidity: function getPoolConfigurator() view returns(address)
func (_Api *ApiCallerSession) GetPoolConfigurator() (common.Address, error) {
	return _Api.Contract.GetPoolConfigurator(&_Api.CallOpts)
}

// GetPoolDataProvider is a free data retrieval call binding the contract method 0xe860accb.
//
// Solidity: function getPoolDataProvider() view returns(address)
func (_Api *ApiCaller) GetPoolDataProvider(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getPoolDataProvider")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPoolDataProvider is a free data retrieval call binding the contract method 0xe860accb.
//
// Solidity: function getPoolDataProvider() view returns(address)
func (_Api *ApiSession) GetPoolDataProvider() (common.Address, error) {
	return _Api.Contract.GetPoolDataProvider(&_Api.CallOpts)
}

// GetPoolDataProvider is a free data retrieval call binding the contract method 0xe860accb.
//
// Solidity: function getPoolDataProvider() view returns(address)
func (_Api *ApiCallerSession) GetPoolDataProvider() (common.Address, error) {
	return _Api.Contract.GetPoolDataProvider(&_Api.CallOpts)
}

// GetPriceOracle is a free data retrieval call binding the contract method 0xfca513a8.
//
// Solidity: function getPriceOracle() view returns(address)
func (_Api *ApiCaller) GetPriceOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getPriceOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPriceOracle is a free data retrieval call binding the contract method 0xfca513a8.
//
// Solidity: function getPriceOracle() view returns(address)
func (_Api *ApiSession) GetPriceOracle() (common.Address, error) {
	return _Api.Contract.GetPriceOracle(&_Api.CallOpts)
}

// GetPriceOracle is a free data retrieval call binding the contract method 0xfca513a8.
//
// Solidity: function getPriceOracle() view returns(address)
func (_Api *ApiCallerSession) GetPriceOracle() (common.Address, error) {
	return _Api.Contract.GetPriceOracle(&_Api.CallOpts)
}

// GetPriceOracleSentinel is a free data retrieval call binding the contract method 0x5eb88d3d.
//
// Solidity: function getPriceOracleSentinel() view returns(address)
func (_Api *ApiCaller) GetPriceOracleSentinel(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getPriceOracleSentinel")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPriceOracleSentinel is a free data retrieval call binding the contract method 0x5eb88d3d.
//
// Solidity: function getPriceOracleSentinel() view returns(address)
func (_Api *ApiSession) GetPriceOracleSentinel() (common.Address, error) {
	return _Api.Contract.GetPriceOracleSentinel(&_Api.CallOpts)
}

// GetPriceOracleSentinel is a free data retrieval call binding the contract method 0x5eb88d3d.
//
// Solidity: function getPriceOracleSentinel() view returns(address)
func (_Api *ApiCallerSession) GetPriceOracleSentinel() (common.Address, error) {
	return _Api.Contract.GetPriceOracleSentinel(&_Api.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Api *ApiCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Api *ApiSession) Owner() (common.Address, error) {
	return _Api.Contract.Owner(&_Api.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Api *ApiCallerSession) Owner() (common.Address, error) {
	return _Api.Contract.Owner(&_Api.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Api *ApiTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Api *ApiSession) RenounceOwnership() (*types.Transaction, error) {
	return _Api.Contract.RenounceOwnership(&_Api.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Api *ApiTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Api.Contract.RenounceOwnership(&_Api.TransactOpts)
}

// SetACLAdmin is a paid mutator transaction binding the contract method 0x76d84ffc.
//
// Solidity: function setACLAdmin(address newAclAdmin) returns()
func (_Api *ApiTransactor) SetACLAdmin(opts *bind.TransactOpts, newAclAdmin common.Address) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "setACLAdmin", newAclAdmin)
}

// SetACLAdmin is a paid mutator transaction binding the contract method 0x76d84ffc.
//
// Solidity: function setACLAdmin(address newAclAdmin) returns()
func (_Api *ApiSession) SetACLAdmin(newAclAdmin common.Address) (*types.Transaction, error) {
	return _Api.Contract.SetACLAdmin(&_Api.TransactOpts, newAclAdmin)
}

// SetACLAdmin is a paid mutator transaction binding the contract method 0x76d84ffc.
//
// Solidity: function setACLAdmin(address newAclAdmin) returns()
func (_Api *ApiTransactorSession) SetACLAdmin(newAclAdmin common.Address) (*types.Transaction, error) {
	return _Api.Contract.SetACLAdmin(&_Api.TransactOpts, newAclAdmin)
}

// SetACLManager is a paid mutator transaction binding the contract method 0xed301ca9.
//
// Solidity: function setACLManager(address newAclManager) returns()
func (_Api *ApiTransactor) SetACLManager(opts *bind.TransactOpts, newAclManager common.Address) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "setACLManager", newAclManager)
}

// SetACLManager is a paid mutator transaction binding the contract method 0xed301ca9.
//
// Solidity: function setACLManager(address newAclManager) returns()
func (_Api *ApiSession) SetACLManager(newAclManager common.Address) (*types.Transaction, error) {
	return _Api.Contract.SetACLManager(&_Api.TransactOpts, newAclManager)
}

// SetACLManager is a paid mutator transaction binding the contract method 0xed301ca9.
//
// Solidity: function setACLManager(address newAclManager) returns()
func (_Api *ApiTransactorSession) SetACLManager(newAclManager common.Address) (*types.Transaction, error) {
	return _Api.Contract.SetACLManager(&_Api.TransactOpts, newAclManager)
}

// SetAddress is a paid mutator transaction binding the contract method 0xca446dd9.
//
// Solidity: function setAddress(bytes32 id, address newAddress) returns()
func (_Api *ApiTransactor) SetAddress(opts *bind.TransactOpts, id [32]byte, newAddress common.Address) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "setAddress", id, newAddress)
}

// SetAddress is a paid mutator transaction binding the contract method 0xca446dd9.
//
// Solidity: function setAddress(bytes32 id, address newAddress) returns()
func (_Api *ApiSession) SetAddress(id [32]byte, newAddress common.Address) (*types.Transaction, error) {
	return _Api.Contract.SetAddress(&_Api.TransactOpts, id, newAddress)
}

// SetAddress is a paid mutator transaction binding the contract method 0xca446dd9.
//
// Solidity: function setAddress(bytes32 id, address newAddress) returns()
func (_Api *ApiTransactorSession) SetAddress(id [32]byte, newAddress common.Address) (*types.Transaction, error) {
	return _Api.Contract.SetAddress(&_Api.TransactOpts, id, newAddress)
}

// SetAddressAsProxy is a paid mutator transaction binding the contract method 0x5dcc528c.
//
// Solidity: function setAddressAsProxy(bytes32 id, address newImplementationAddress) returns()
func (_Api *ApiTransactor) SetAddressAsProxy(opts *bind.TransactOpts, id [32]byte, newImplementationAddress common.Address) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "setAddressAsProxy", id, newImplementationAddress)
}

// SetAddressAsProxy is a paid mutator transaction binding the contract method 0x5dcc528c.
//
// Solidity: function setAddressAsProxy(bytes32 id, address newImplementationAddress) returns()
func (_Api *ApiSession) SetAddressAsProxy(id [32]byte, newImplementationAddress common.Address) (*types.Transaction, error) {
	return _Api.Contract.SetAddressAsProxy(&_Api.TransactOpts, id, newImplementationAddress)
}

// SetAddressAsProxy is a paid mutator transaction binding the contract method 0x5dcc528c.
//
// Solidity: function setAddressAsProxy(bytes32 id, address newImplementationAddress) returns()
func (_Api *ApiTransactorSession) SetAddressAsProxy(id [32]byte, newImplementationAddress common.Address) (*types.Transaction, error) {
	return _Api.Contract.SetAddressAsProxy(&_Api.TransactOpts, id, newImplementationAddress)
}

// SetMarketId is a paid mutator transaction binding the contract method 0xf67b1847.
//
// Solidity: function setMarketId(string newMarketId) returns()
func (_Api *ApiTransactor) SetMarketId(opts *bind.TransactOpts, newMarketId string) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "setMarketId", newMarketId)
}

// SetMarketId is a paid mutator transaction binding the contract method 0xf67b1847.
//
// Solidity: function setMarketId(string newMarketId) returns()
func (_Api *ApiSession) SetMarketId(newMarketId string) (*types.Transaction, error) {
	return _Api.Contract.SetMarketId(&_Api.TransactOpts, newMarketId)
}

// SetMarketId is a paid mutator transaction binding the contract method 0xf67b1847.
//
// Solidity: function setMarketId(string newMarketId) returns()
func (_Api *ApiTransactorSession) SetMarketId(newMarketId string) (*types.Transaction, error) {
	return _Api.Contract.SetMarketId(&_Api.TransactOpts, newMarketId)
}

// SetPoolConfiguratorImpl is a paid mutator transaction binding the contract method 0xe4ca28b7.
//
// Solidity: function setPoolConfiguratorImpl(address newPoolConfiguratorImpl) returns()
func (_Api *ApiTransactor) SetPoolConfiguratorImpl(opts *bind.TransactOpts, newPoolConfiguratorImpl common.Address) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "setPoolConfiguratorImpl", newPoolConfiguratorImpl)
}

// SetPoolConfiguratorImpl is a paid mutator transaction binding the contract method 0xe4ca28b7.
//
// Solidity: function setPoolConfiguratorImpl(address newPoolConfiguratorImpl) returns()
func (_Api *ApiSession) SetPoolConfiguratorImpl(newPoolConfiguratorImpl common.Address) (*types.Transaction, error) {
	return _Api.Contract.SetPoolConfiguratorImpl(&_Api.TransactOpts, newPoolConfiguratorImpl)
}

// SetPoolConfiguratorImpl is a paid mutator transaction binding the contract method 0xe4ca28b7.
//
// Solidity: function setPoolConfiguratorImpl(address newPoolConfiguratorImpl) returns()
func (_Api *ApiTransactorSession) SetPoolConfiguratorImpl(newPoolConfiguratorImpl common.Address) (*types.Transaction, error) {
	return _Api.Contract.SetPoolConfiguratorImpl(&_Api.TransactOpts, newPoolConfiguratorImpl)
}

// SetPoolDataProvider is a paid mutator transaction binding the contract method 0xe44e9ed1.
//
// Solidity: function setPoolDataProvider(address newDataProvider) returns()
func (_Api *ApiTransactor) SetPoolDataProvider(opts *bind.TransactOpts, newDataProvider common.Address) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "setPoolDataProvider", newDataProvider)
}

// SetPoolDataProvider is a paid mutator transaction binding the contract method 0xe44e9ed1.
//
// Solidity: function setPoolDataProvider(address newDataProvider) returns()
func (_Api *ApiSession) SetPoolDataProvider(newDataProvider common.Address) (*types.Transaction, error) {
	return _Api.Contract.SetPoolDataProvider(&_Api.TransactOpts, newDataProvider)
}

// SetPoolDataProvider is a paid mutator transaction binding the contract method 0xe44e9ed1.
//
// Solidity: function setPoolDataProvider(address newDataProvider) returns()
func (_Api *ApiTransactorSession) SetPoolDataProvider(newDataProvider common.Address) (*types.Transaction, error) {
	return _Api.Contract.SetPoolDataProvider(&_Api.TransactOpts, newDataProvider)
}

// SetPoolImpl is a paid mutator transaction binding the contract method 0xa1564406.
//
// Solidity: function setPoolImpl(address newPoolImpl) returns()
func (_Api *ApiTransactor) SetPoolImpl(opts *bind.TransactOpts, newPoolImpl common.Address) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "setPoolImpl", newPoolImpl)
}

// SetPoolImpl is a paid mutator transaction binding the contract method 0xa1564406.
//
// Solidity: function setPoolImpl(address newPoolImpl) returns()
func (_Api *ApiSession) SetPoolImpl(newPoolImpl common.Address) (*types.Transaction, error) {
	return _Api.Contract.SetPoolImpl(&_Api.TransactOpts, newPoolImpl)
}

// SetPoolImpl is a paid mutator transaction binding the contract method 0xa1564406.
//
// Solidity: function setPoolImpl(address newPoolImpl) returns()
func (_Api *ApiTransactorSession) SetPoolImpl(newPoolImpl common.Address) (*types.Transaction, error) {
	return _Api.Contract.SetPoolImpl(&_Api.TransactOpts, newPoolImpl)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x530e784f.
//
// Solidity: function setPriceOracle(address newPriceOracle) returns()
func (_Api *ApiTransactor) SetPriceOracle(opts *bind.TransactOpts, newPriceOracle common.Address) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "setPriceOracle", newPriceOracle)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x530e784f.
//
// Solidity: function setPriceOracle(address newPriceOracle) returns()
func (_Api *ApiSession) SetPriceOracle(newPriceOracle common.Address) (*types.Transaction, error) {
	return _Api.Contract.SetPriceOracle(&_Api.TransactOpts, newPriceOracle)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x530e784f.
//
// Solidity: function setPriceOracle(address newPriceOracle) returns()
func (_Api *ApiTransactorSession) SetPriceOracle(newPriceOracle common.Address) (*types.Transaction, error) {
	return _Api.Contract.SetPriceOracle(&_Api.TransactOpts, newPriceOracle)
}

// SetPriceOracleSentinel is a paid mutator transaction binding the contract method 0x74944cec.
//
// Solidity: function setPriceOracleSentinel(address newPriceOracleSentinel) returns()
func (_Api *ApiTransactor) SetPriceOracleSentinel(opts *bind.TransactOpts, newPriceOracleSentinel common.Address) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "setPriceOracleSentinel", newPriceOracleSentinel)
}

// SetPriceOracleSentinel is a paid mutator transaction binding the contract method 0x74944cec.
//
// Solidity: function setPriceOracleSentinel(address newPriceOracleSentinel) returns()
func (_Api *ApiSession) SetPriceOracleSentinel(newPriceOracleSentinel common.Address) (*types.Transaction, error) {
	return _Api.Contract.SetPriceOracleSentinel(&_Api.TransactOpts, newPriceOracleSentinel)
}

// SetPriceOracleSentinel is a paid mutator transaction binding the contract method 0x74944cec.
//
// Solidity: function setPriceOracleSentinel(address newPriceOracleSentinel) returns()
func (_Api *ApiTransactorSession) SetPriceOracleSentinel(newPriceOracleSentinel common.Address) (*types.Transaction, error) {
	return _Api.Contract.SetPriceOracleSentinel(&_Api.TransactOpts, newPriceOracleSentinel)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Api *ApiTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Api *ApiSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Api.Contract.TransferOwnership(&_Api.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Api *ApiTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Api.Contract.TransferOwnership(&_Api.TransactOpts, newOwner)
}

// ApiACLAdminUpdatedIterator is returned from FilterACLAdminUpdated and is used to iterate over the raw logs and unpacked data for ACLAdminUpdated events raised by the Api contract.
type ApiACLAdminUpdatedIterator struct {
	Event *ApiACLAdminUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ApiACLAdminUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiACLAdminUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ApiACLAdminUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ApiACLAdminUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiACLAdminUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiACLAdminUpdated represents a ACLAdminUpdated event raised by the Api contract.
type ApiACLAdminUpdated struct {
	OldAddress common.Address
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterACLAdminUpdated is a free log retrieval operation binding the contract event 0xe9cf53972264dc95304fd424458745019ddfca0e37ae8f703d74772c41ad115b.
//
// Solidity: event ACLAdminUpdated(address indexed oldAddress, address indexed newAddress)
func (_Api *ApiFilterer) FilterACLAdminUpdated(opts *bind.FilterOpts, oldAddress []common.Address, newAddress []common.Address) (*ApiACLAdminUpdatedIterator, error) {

	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _Api.contract.FilterLogs(opts, "ACLAdminUpdated", oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return &ApiACLAdminUpdatedIterator{contract: _Api.contract, event: "ACLAdminUpdated", logs: logs, sub: sub}, nil
}

// WatchACLAdminUpdated is a free log subscription operation binding the contract event 0xe9cf53972264dc95304fd424458745019ddfca0e37ae8f703d74772c41ad115b.
//
// Solidity: event ACLAdminUpdated(address indexed oldAddress, address indexed newAddress)
func (_Api *ApiFilterer) WatchACLAdminUpdated(opts *bind.WatchOpts, sink chan<- *ApiACLAdminUpdated, oldAddress []common.Address, newAddress []common.Address) (event.Subscription, error) {

	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _Api.contract.WatchLogs(opts, "ACLAdminUpdated", oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiACLAdminUpdated)
				if err := _Api.contract.UnpackLog(event, "ACLAdminUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseACLAdminUpdated is a log parse operation binding the contract event 0xe9cf53972264dc95304fd424458745019ddfca0e37ae8f703d74772c41ad115b.
//
// Solidity: event ACLAdminUpdated(address indexed oldAddress, address indexed newAddress)
func (_Api *ApiFilterer) ParseACLAdminUpdated(log types.Log) (*ApiACLAdminUpdated, error) {
	event := new(ApiACLAdminUpdated)
	if err := _Api.contract.UnpackLog(event, "ACLAdminUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiACLManagerUpdatedIterator is returned from FilterACLManagerUpdated and is used to iterate over the raw logs and unpacked data for ACLManagerUpdated events raised by the Api contract.
type ApiACLManagerUpdatedIterator struct {
	Event *ApiACLManagerUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ApiACLManagerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiACLManagerUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ApiACLManagerUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ApiACLManagerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiACLManagerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiACLManagerUpdated represents a ACLManagerUpdated event raised by the Api contract.
type ApiACLManagerUpdated struct {
	OldAddress common.Address
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterACLManagerUpdated is a free log retrieval operation binding the contract event 0xb30efa04327bb8a537d61cc1e5c48095345ad18ef7cc04e6bacf7dfb6caaf507.
//
// Solidity: event ACLManagerUpdated(address indexed oldAddress, address indexed newAddress)
func (_Api *ApiFilterer) FilterACLManagerUpdated(opts *bind.FilterOpts, oldAddress []common.Address, newAddress []common.Address) (*ApiACLManagerUpdatedIterator, error) {

	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _Api.contract.FilterLogs(opts, "ACLManagerUpdated", oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return &ApiACLManagerUpdatedIterator{contract: _Api.contract, event: "ACLManagerUpdated", logs: logs, sub: sub}, nil
}

// WatchACLManagerUpdated is a free log subscription operation binding the contract event 0xb30efa04327bb8a537d61cc1e5c48095345ad18ef7cc04e6bacf7dfb6caaf507.
//
// Solidity: event ACLManagerUpdated(address indexed oldAddress, address indexed newAddress)
func (_Api *ApiFilterer) WatchACLManagerUpdated(opts *bind.WatchOpts, sink chan<- *ApiACLManagerUpdated, oldAddress []common.Address, newAddress []common.Address) (event.Subscription, error) {

	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _Api.contract.WatchLogs(opts, "ACLManagerUpdated", oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiACLManagerUpdated)
				if err := _Api.contract.UnpackLog(event, "ACLManagerUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseACLManagerUpdated is a log parse operation binding the contract event 0xb30efa04327bb8a537d61cc1e5c48095345ad18ef7cc04e6bacf7dfb6caaf507.
//
// Solidity: event ACLManagerUpdated(address indexed oldAddress, address indexed newAddress)
func (_Api *ApiFilterer) ParseACLManagerUpdated(log types.Log) (*ApiACLManagerUpdated, error) {
	event := new(ApiACLManagerUpdated)
	if err := _Api.contract.UnpackLog(event, "ACLManagerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiAddressSetIterator is returned from FilterAddressSet and is used to iterate over the raw logs and unpacked data for AddressSet events raised by the Api contract.
type ApiAddressSetIterator struct {
	Event *ApiAddressSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ApiAddressSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiAddressSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ApiAddressSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ApiAddressSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiAddressSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiAddressSet represents a AddressSet event raised by the Api contract.
type ApiAddressSet struct {
	Id         [32]byte
	OldAddress common.Address
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAddressSet is a free log retrieval operation binding the contract event 0x9ef0e8c8e52743bb38b83b17d9429141d494b8041ca6d616a6c77cebae9cd8b7.
//
// Solidity: event AddressSet(bytes32 indexed id, address indexed oldAddress, address indexed newAddress)
func (_Api *ApiFilterer) FilterAddressSet(opts *bind.FilterOpts, id [][32]byte, oldAddress []common.Address, newAddress []common.Address) (*ApiAddressSetIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _Api.contract.FilterLogs(opts, "AddressSet", idRule, oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return &ApiAddressSetIterator{contract: _Api.contract, event: "AddressSet", logs: logs, sub: sub}, nil
}

// WatchAddressSet is a free log subscription operation binding the contract event 0x9ef0e8c8e52743bb38b83b17d9429141d494b8041ca6d616a6c77cebae9cd8b7.
//
// Solidity: event AddressSet(bytes32 indexed id, address indexed oldAddress, address indexed newAddress)
func (_Api *ApiFilterer) WatchAddressSet(opts *bind.WatchOpts, sink chan<- *ApiAddressSet, id [][32]byte, oldAddress []common.Address, newAddress []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _Api.contract.WatchLogs(opts, "AddressSet", idRule, oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiAddressSet)
				if err := _Api.contract.UnpackLog(event, "AddressSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAddressSet is a log parse operation binding the contract event 0x9ef0e8c8e52743bb38b83b17d9429141d494b8041ca6d616a6c77cebae9cd8b7.
//
// Solidity: event AddressSet(bytes32 indexed id, address indexed oldAddress, address indexed newAddress)
func (_Api *ApiFilterer) ParseAddressSet(log types.Log) (*ApiAddressSet, error) {
	event := new(ApiAddressSet)
	if err := _Api.contract.UnpackLog(event, "AddressSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiAddressSetAsProxyIterator is returned from FilterAddressSetAsProxy and is used to iterate over the raw logs and unpacked data for AddressSetAsProxy events raised by the Api contract.
type ApiAddressSetAsProxyIterator struct {
	Event *ApiAddressSetAsProxy // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ApiAddressSetAsProxyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiAddressSetAsProxy)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ApiAddressSetAsProxy)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ApiAddressSetAsProxyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiAddressSetAsProxyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiAddressSetAsProxy represents a AddressSetAsProxy event raised by the Api contract.
type ApiAddressSetAsProxy struct {
	Id                       [32]byte
	ProxyAddress             common.Address
	OldImplementationAddress common.Address
	NewImplementationAddress common.Address
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterAddressSetAsProxy is a free log retrieval operation binding the contract event 0x3bbd45b5429b385e3fb37ad5cd1cd1435a3c8ec32196c7937597365a3fd3e99c.
//
// Solidity: event AddressSetAsProxy(bytes32 indexed id, address indexed proxyAddress, address oldImplementationAddress, address indexed newImplementationAddress)
func (_Api *ApiFilterer) FilterAddressSetAsProxy(opts *bind.FilterOpts, id [][32]byte, proxyAddress []common.Address, newImplementationAddress []common.Address) (*ApiAddressSetAsProxyIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var proxyAddressRule []interface{}
	for _, proxyAddressItem := range proxyAddress {
		proxyAddressRule = append(proxyAddressRule, proxyAddressItem)
	}

	var newImplementationAddressRule []interface{}
	for _, newImplementationAddressItem := range newImplementationAddress {
		newImplementationAddressRule = append(newImplementationAddressRule, newImplementationAddressItem)
	}

	logs, sub, err := _Api.contract.FilterLogs(opts, "AddressSetAsProxy", idRule, proxyAddressRule, newImplementationAddressRule)
	if err != nil {
		return nil, err
	}
	return &ApiAddressSetAsProxyIterator{contract: _Api.contract, event: "AddressSetAsProxy", logs: logs, sub: sub}, nil
}

// WatchAddressSetAsProxy is a free log subscription operation binding the contract event 0x3bbd45b5429b385e3fb37ad5cd1cd1435a3c8ec32196c7937597365a3fd3e99c.
//
// Solidity: event AddressSetAsProxy(bytes32 indexed id, address indexed proxyAddress, address oldImplementationAddress, address indexed newImplementationAddress)
func (_Api *ApiFilterer) WatchAddressSetAsProxy(opts *bind.WatchOpts, sink chan<- *ApiAddressSetAsProxy, id [][32]byte, proxyAddress []common.Address, newImplementationAddress []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var proxyAddressRule []interface{}
	for _, proxyAddressItem := range proxyAddress {
		proxyAddressRule = append(proxyAddressRule, proxyAddressItem)
	}

	var newImplementationAddressRule []interface{}
	for _, newImplementationAddressItem := range newImplementationAddress {
		newImplementationAddressRule = append(newImplementationAddressRule, newImplementationAddressItem)
	}

	logs, sub, err := _Api.contract.WatchLogs(opts, "AddressSetAsProxy", idRule, proxyAddressRule, newImplementationAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiAddressSetAsProxy)
				if err := _Api.contract.UnpackLog(event, "AddressSetAsProxy", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAddressSetAsProxy is a log parse operation binding the contract event 0x3bbd45b5429b385e3fb37ad5cd1cd1435a3c8ec32196c7937597365a3fd3e99c.
//
// Solidity: event AddressSetAsProxy(bytes32 indexed id, address indexed proxyAddress, address oldImplementationAddress, address indexed newImplementationAddress)
func (_Api *ApiFilterer) ParseAddressSetAsProxy(log types.Log) (*ApiAddressSetAsProxy, error) {
	event := new(ApiAddressSetAsProxy)
	if err := _Api.contract.UnpackLog(event, "AddressSetAsProxy", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiMarketIdSetIterator is returned from FilterMarketIdSet and is used to iterate over the raw logs and unpacked data for MarketIdSet events raised by the Api contract.
type ApiMarketIdSetIterator struct {
	Event *ApiMarketIdSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ApiMarketIdSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiMarketIdSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ApiMarketIdSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ApiMarketIdSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiMarketIdSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiMarketIdSet represents a MarketIdSet event raised by the Api contract.
type ApiMarketIdSet struct {
	OldMarketId common.Hash
	NewMarketId common.Hash
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMarketIdSet is a free log retrieval operation binding the contract event 0xe685c8cdecc6030c45030fd54778812cb84ed8e4467c38294403d68ba7860823.
//
// Solidity: event MarketIdSet(string indexed oldMarketId, string indexed newMarketId)
func (_Api *ApiFilterer) FilterMarketIdSet(opts *bind.FilterOpts, oldMarketId []string, newMarketId []string) (*ApiMarketIdSetIterator, error) {

	var oldMarketIdRule []interface{}
	for _, oldMarketIdItem := range oldMarketId {
		oldMarketIdRule = append(oldMarketIdRule, oldMarketIdItem)
	}
	var newMarketIdRule []interface{}
	for _, newMarketIdItem := range newMarketId {
		newMarketIdRule = append(newMarketIdRule, newMarketIdItem)
	}

	logs, sub, err := _Api.contract.FilterLogs(opts, "MarketIdSet", oldMarketIdRule, newMarketIdRule)
	if err != nil {
		return nil, err
	}
	return &ApiMarketIdSetIterator{contract: _Api.contract, event: "MarketIdSet", logs: logs, sub: sub}, nil
}

// WatchMarketIdSet is a free log subscription operation binding the contract event 0xe685c8cdecc6030c45030fd54778812cb84ed8e4467c38294403d68ba7860823.
//
// Solidity: event MarketIdSet(string indexed oldMarketId, string indexed newMarketId)
func (_Api *ApiFilterer) WatchMarketIdSet(opts *bind.WatchOpts, sink chan<- *ApiMarketIdSet, oldMarketId []string, newMarketId []string) (event.Subscription, error) {

	var oldMarketIdRule []interface{}
	for _, oldMarketIdItem := range oldMarketId {
		oldMarketIdRule = append(oldMarketIdRule, oldMarketIdItem)
	}
	var newMarketIdRule []interface{}
	for _, newMarketIdItem := range newMarketId {
		newMarketIdRule = append(newMarketIdRule, newMarketIdItem)
	}

	logs, sub, err := _Api.contract.WatchLogs(opts, "MarketIdSet", oldMarketIdRule, newMarketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiMarketIdSet)
				if err := _Api.contract.UnpackLog(event, "MarketIdSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMarketIdSet is a log parse operation binding the contract event 0xe685c8cdecc6030c45030fd54778812cb84ed8e4467c38294403d68ba7860823.
//
// Solidity: event MarketIdSet(string indexed oldMarketId, string indexed newMarketId)
func (_Api *ApiFilterer) ParseMarketIdSet(log types.Log) (*ApiMarketIdSet, error) {
	event := new(ApiMarketIdSet)
	if err := _Api.contract.UnpackLog(event, "MarketIdSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Api contract.
type ApiOwnershipTransferredIterator struct {
	Event *ApiOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ApiOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ApiOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ApiOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiOwnershipTransferred represents a OwnershipTransferred event raised by the Api contract.
type ApiOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Api *ApiFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ApiOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Api.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ApiOwnershipTransferredIterator{contract: _Api.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Api *ApiFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ApiOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Api.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiOwnershipTransferred)
				if err := _Api.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Api *ApiFilterer) ParseOwnershipTransferred(log types.Log) (*ApiOwnershipTransferred, error) {
	event := new(ApiOwnershipTransferred)
	if err := _Api.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiPoolConfiguratorUpdatedIterator is returned from FilterPoolConfiguratorUpdated and is used to iterate over the raw logs and unpacked data for PoolConfiguratorUpdated events raised by the Api contract.
type ApiPoolConfiguratorUpdatedIterator struct {
	Event *ApiPoolConfiguratorUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ApiPoolConfiguratorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiPoolConfiguratorUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ApiPoolConfiguratorUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ApiPoolConfiguratorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiPoolConfiguratorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiPoolConfiguratorUpdated represents a PoolConfiguratorUpdated event raised by the Api contract.
type ApiPoolConfiguratorUpdated struct {
	OldAddress common.Address
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPoolConfiguratorUpdated is a free log retrieval operation binding the contract event 0x8932892569eba59c8382a089d9b732d1f49272878775235761a2a6b0309cd465.
//
// Solidity: event PoolConfiguratorUpdated(address indexed oldAddress, address indexed newAddress)
func (_Api *ApiFilterer) FilterPoolConfiguratorUpdated(opts *bind.FilterOpts, oldAddress []common.Address, newAddress []common.Address) (*ApiPoolConfiguratorUpdatedIterator, error) {

	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _Api.contract.FilterLogs(opts, "PoolConfiguratorUpdated", oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return &ApiPoolConfiguratorUpdatedIterator{contract: _Api.contract, event: "PoolConfiguratorUpdated", logs: logs, sub: sub}, nil
}

// WatchPoolConfiguratorUpdated is a free log subscription operation binding the contract event 0x8932892569eba59c8382a089d9b732d1f49272878775235761a2a6b0309cd465.
//
// Solidity: event PoolConfiguratorUpdated(address indexed oldAddress, address indexed newAddress)
func (_Api *ApiFilterer) WatchPoolConfiguratorUpdated(opts *bind.WatchOpts, sink chan<- *ApiPoolConfiguratorUpdated, oldAddress []common.Address, newAddress []common.Address) (event.Subscription, error) {

	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _Api.contract.WatchLogs(opts, "PoolConfiguratorUpdated", oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiPoolConfiguratorUpdated)
				if err := _Api.contract.UnpackLog(event, "PoolConfiguratorUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePoolConfiguratorUpdated is a log parse operation binding the contract event 0x8932892569eba59c8382a089d9b732d1f49272878775235761a2a6b0309cd465.
//
// Solidity: event PoolConfiguratorUpdated(address indexed oldAddress, address indexed newAddress)
func (_Api *ApiFilterer) ParsePoolConfiguratorUpdated(log types.Log) (*ApiPoolConfiguratorUpdated, error) {
	event := new(ApiPoolConfiguratorUpdated)
	if err := _Api.contract.UnpackLog(event, "PoolConfiguratorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiPoolDataProviderUpdatedIterator is returned from FilterPoolDataProviderUpdated and is used to iterate over the raw logs and unpacked data for PoolDataProviderUpdated events raised by the Api contract.
type ApiPoolDataProviderUpdatedIterator struct {
	Event *ApiPoolDataProviderUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ApiPoolDataProviderUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiPoolDataProviderUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ApiPoolDataProviderUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ApiPoolDataProviderUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiPoolDataProviderUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiPoolDataProviderUpdated represents a PoolDataProviderUpdated event raised by the Api contract.
type ApiPoolDataProviderUpdated struct {
	OldAddress common.Address
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPoolDataProviderUpdated is a free log retrieval operation binding the contract event 0xc853974cfbf81487a14a23565917bee63f527853bcb5fa54f2ae1cdf8a38356d.
//
// Solidity: event PoolDataProviderUpdated(address indexed oldAddress, address indexed newAddress)
func (_Api *ApiFilterer) FilterPoolDataProviderUpdated(opts *bind.FilterOpts, oldAddress []common.Address, newAddress []common.Address) (*ApiPoolDataProviderUpdatedIterator, error) {

	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _Api.contract.FilterLogs(opts, "PoolDataProviderUpdated", oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return &ApiPoolDataProviderUpdatedIterator{contract: _Api.contract, event: "PoolDataProviderUpdated", logs: logs, sub: sub}, nil
}

// WatchPoolDataProviderUpdated is a free log subscription operation binding the contract event 0xc853974cfbf81487a14a23565917bee63f527853bcb5fa54f2ae1cdf8a38356d.
//
// Solidity: event PoolDataProviderUpdated(address indexed oldAddress, address indexed newAddress)
func (_Api *ApiFilterer) WatchPoolDataProviderUpdated(opts *bind.WatchOpts, sink chan<- *ApiPoolDataProviderUpdated, oldAddress []common.Address, newAddress []common.Address) (event.Subscription, error) {

	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _Api.contract.WatchLogs(opts, "PoolDataProviderUpdated", oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiPoolDataProviderUpdated)
				if err := _Api.contract.UnpackLog(event, "PoolDataProviderUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePoolDataProviderUpdated is a log parse operation binding the contract event 0xc853974cfbf81487a14a23565917bee63f527853bcb5fa54f2ae1cdf8a38356d.
//
// Solidity: event PoolDataProviderUpdated(address indexed oldAddress, address indexed newAddress)
func (_Api *ApiFilterer) ParsePoolDataProviderUpdated(log types.Log) (*ApiPoolDataProviderUpdated, error) {
	event := new(ApiPoolDataProviderUpdated)
	if err := _Api.contract.UnpackLog(event, "PoolDataProviderUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiPoolUpdatedIterator is returned from FilterPoolUpdated and is used to iterate over the raw logs and unpacked data for PoolUpdated events raised by the Api contract.
type ApiPoolUpdatedIterator struct {
	Event *ApiPoolUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ApiPoolUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiPoolUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ApiPoolUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ApiPoolUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiPoolUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiPoolUpdated represents a PoolUpdated event raised by the Api contract.
type ApiPoolUpdated struct {
	OldAddress common.Address
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPoolUpdated is a free log retrieval operation binding the contract event 0x90affc163f1a2dfedcd36aa02ed992eeeba8100a4014f0b4cdc20ea265a66627.
//
// Solidity: event PoolUpdated(address indexed oldAddress, address indexed newAddress)
func (_Api *ApiFilterer) FilterPoolUpdated(opts *bind.FilterOpts, oldAddress []common.Address, newAddress []common.Address) (*ApiPoolUpdatedIterator, error) {

	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _Api.contract.FilterLogs(opts, "PoolUpdated", oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return &ApiPoolUpdatedIterator{contract: _Api.contract, event: "PoolUpdated", logs: logs, sub: sub}, nil
}

// WatchPoolUpdated is a free log subscription operation binding the contract event 0x90affc163f1a2dfedcd36aa02ed992eeeba8100a4014f0b4cdc20ea265a66627.
//
// Solidity: event PoolUpdated(address indexed oldAddress, address indexed newAddress)
func (_Api *ApiFilterer) WatchPoolUpdated(opts *bind.WatchOpts, sink chan<- *ApiPoolUpdated, oldAddress []common.Address, newAddress []common.Address) (event.Subscription, error) {

	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _Api.contract.WatchLogs(opts, "PoolUpdated", oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiPoolUpdated)
				if err := _Api.contract.UnpackLog(event, "PoolUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePoolUpdated is a log parse operation binding the contract event 0x90affc163f1a2dfedcd36aa02ed992eeeba8100a4014f0b4cdc20ea265a66627.
//
// Solidity: event PoolUpdated(address indexed oldAddress, address indexed newAddress)
func (_Api *ApiFilterer) ParsePoolUpdated(log types.Log) (*ApiPoolUpdated, error) {
	event := new(ApiPoolUpdated)
	if err := _Api.contract.UnpackLog(event, "PoolUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiPriceOracleSentinelUpdatedIterator is returned from FilterPriceOracleSentinelUpdated and is used to iterate over the raw logs and unpacked data for PriceOracleSentinelUpdated events raised by the Api contract.
type ApiPriceOracleSentinelUpdatedIterator struct {
	Event *ApiPriceOracleSentinelUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ApiPriceOracleSentinelUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiPriceOracleSentinelUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ApiPriceOracleSentinelUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ApiPriceOracleSentinelUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiPriceOracleSentinelUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiPriceOracleSentinelUpdated represents a PriceOracleSentinelUpdated event raised by the Api contract.
type ApiPriceOracleSentinelUpdated struct {
	OldAddress common.Address
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPriceOracleSentinelUpdated is a free log retrieval operation binding the contract event 0x5326514eeca90494a14bedabcff812a0e683029ee85d1e23824d44fd14cd6ae7.
//
// Solidity: event PriceOracleSentinelUpdated(address indexed oldAddress, address indexed newAddress)
func (_Api *ApiFilterer) FilterPriceOracleSentinelUpdated(opts *bind.FilterOpts, oldAddress []common.Address, newAddress []common.Address) (*ApiPriceOracleSentinelUpdatedIterator, error) {

	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _Api.contract.FilterLogs(opts, "PriceOracleSentinelUpdated", oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return &ApiPriceOracleSentinelUpdatedIterator{contract: _Api.contract, event: "PriceOracleSentinelUpdated", logs: logs, sub: sub}, nil
}

// WatchPriceOracleSentinelUpdated is a free log subscription operation binding the contract event 0x5326514eeca90494a14bedabcff812a0e683029ee85d1e23824d44fd14cd6ae7.
//
// Solidity: event PriceOracleSentinelUpdated(address indexed oldAddress, address indexed newAddress)
func (_Api *ApiFilterer) WatchPriceOracleSentinelUpdated(opts *bind.WatchOpts, sink chan<- *ApiPriceOracleSentinelUpdated, oldAddress []common.Address, newAddress []common.Address) (event.Subscription, error) {

	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _Api.contract.WatchLogs(opts, "PriceOracleSentinelUpdated", oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiPriceOracleSentinelUpdated)
				if err := _Api.contract.UnpackLog(event, "PriceOracleSentinelUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePriceOracleSentinelUpdated is a log parse operation binding the contract event 0x5326514eeca90494a14bedabcff812a0e683029ee85d1e23824d44fd14cd6ae7.
//
// Solidity: event PriceOracleSentinelUpdated(address indexed oldAddress, address indexed newAddress)
func (_Api *ApiFilterer) ParsePriceOracleSentinelUpdated(log types.Log) (*ApiPriceOracleSentinelUpdated, error) {
	event := new(ApiPriceOracleSentinelUpdated)
	if err := _Api.contract.UnpackLog(event, "PriceOracleSentinelUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiPriceOracleUpdatedIterator is returned from FilterPriceOracleUpdated and is used to iterate over the raw logs and unpacked data for PriceOracleUpdated events raised by the Api contract.
type ApiPriceOracleUpdatedIterator struct {
	Event *ApiPriceOracleUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ApiPriceOracleUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiPriceOracleUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ApiPriceOracleUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ApiPriceOracleUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiPriceOracleUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiPriceOracleUpdated represents a PriceOracleUpdated event raised by the Api contract.
type ApiPriceOracleUpdated struct {
	OldAddress common.Address
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPriceOracleUpdated is a free log retrieval operation binding the contract event 0x56b5f80d8cac1479698aa7d01605fd6111e90b15fc4d2b377417f46034876cbd.
//
// Solidity: event PriceOracleUpdated(address indexed oldAddress, address indexed newAddress)
func (_Api *ApiFilterer) FilterPriceOracleUpdated(opts *bind.FilterOpts, oldAddress []common.Address, newAddress []common.Address) (*ApiPriceOracleUpdatedIterator, error) {

	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _Api.contract.FilterLogs(opts, "PriceOracleUpdated", oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return &ApiPriceOracleUpdatedIterator{contract: _Api.contract, event: "PriceOracleUpdated", logs: logs, sub: sub}, nil
}

// WatchPriceOracleUpdated is a free log subscription operation binding the contract event 0x56b5f80d8cac1479698aa7d01605fd6111e90b15fc4d2b377417f46034876cbd.
//
// Solidity: event PriceOracleUpdated(address indexed oldAddress, address indexed newAddress)
func (_Api *ApiFilterer) WatchPriceOracleUpdated(opts *bind.WatchOpts, sink chan<- *ApiPriceOracleUpdated, oldAddress []common.Address, newAddress []common.Address) (event.Subscription, error) {

	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _Api.contract.WatchLogs(opts, "PriceOracleUpdated", oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiPriceOracleUpdated)
				if err := _Api.contract.UnpackLog(event, "PriceOracleUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePriceOracleUpdated is a log parse operation binding the contract event 0x56b5f80d8cac1479698aa7d01605fd6111e90b15fc4d2b377417f46034876cbd.
//
// Solidity: event PriceOracleUpdated(address indexed oldAddress, address indexed newAddress)
func (_Api *ApiFilterer) ParsePriceOracleUpdated(log types.Log) (*ApiPriceOracleUpdated, error) {
	event := new(ApiPriceOracleUpdated)
	if err := _Api.contract.UnpackLog(event, "PriceOracleUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiProxyCreatedIterator is returned from FilterProxyCreated and is used to iterate over the raw logs and unpacked data for ProxyCreated events raised by the Api contract.
type ApiProxyCreatedIterator struct {
	Event *ApiProxyCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ApiProxyCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiProxyCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ApiProxyCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ApiProxyCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiProxyCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiProxyCreated represents a ProxyCreated event raised by the Api contract.
type ApiProxyCreated struct {
	Id                    [32]byte
	ProxyAddress          common.Address
	ImplementationAddress common.Address
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterProxyCreated is a free log retrieval operation binding the contract event 0x4a465a9bd819d9662563c1e11ae958f8109e437e7f4bf1c6ef0b9a7b3f35d478.
//
// Solidity: event ProxyCreated(bytes32 indexed id, address indexed proxyAddress, address indexed implementationAddress)
func (_Api *ApiFilterer) FilterProxyCreated(opts *bind.FilterOpts, id [][32]byte, proxyAddress []common.Address, implementationAddress []common.Address) (*ApiProxyCreatedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var proxyAddressRule []interface{}
	for _, proxyAddressItem := range proxyAddress {
		proxyAddressRule = append(proxyAddressRule, proxyAddressItem)
	}
	var implementationAddressRule []interface{}
	for _, implementationAddressItem := range implementationAddress {
		implementationAddressRule = append(implementationAddressRule, implementationAddressItem)
	}

	logs, sub, err := _Api.contract.FilterLogs(opts, "ProxyCreated", idRule, proxyAddressRule, implementationAddressRule)
	if err != nil {
		return nil, err
	}
	return &ApiProxyCreatedIterator{contract: _Api.contract, event: "ProxyCreated", logs: logs, sub: sub}, nil
}

// WatchProxyCreated is a free log subscription operation binding the contract event 0x4a465a9bd819d9662563c1e11ae958f8109e437e7f4bf1c6ef0b9a7b3f35d478.
//
// Solidity: event ProxyCreated(bytes32 indexed id, address indexed proxyAddress, address indexed implementationAddress)
func (_Api *ApiFilterer) WatchProxyCreated(opts *bind.WatchOpts, sink chan<- *ApiProxyCreated, id [][32]byte, proxyAddress []common.Address, implementationAddress []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var proxyAddressRule []interface{}
	for _, proxyAddressItem := range proxyAddress {
		proxyAddressRule = append(proxyAddressRule, proxyAddressItem)
	}
	var implementationAddressRule []interface{}
	for _, implementationAddressItem := range implementationAddress {
		implementationAddressRule = append(implementationAddressRule, implementationAddressItem)
	}

	logs, sub, err := _Api.contract.WatchLogs(opts, "ProxyCreated", idRule, proxyAddressRule, implementationAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiProxyCreated)
				if err := _Api.contract.UnpackLog(event, "ProxyCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProxyCreated is a log parse operation binding the contract event 0x4a465a9bd819d9662563c1e11ae958f8109e437e7f4bf1c6ef0b9a7b3f35d478.
//
// Solidity: event ProxyCreated(bytes32 indexed id, address indexed proxyAddress, address indexed implementationAddress)
func (_Api *ApiFilterer) ParseProxyCreated(log types.Log) (*ApiProxyCreated, error) {
	event := new(ApiProxyCreated)
	if err := _Api.contract.UnpackLog(event, "ProxyCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
