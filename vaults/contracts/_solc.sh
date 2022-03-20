Warning: Contract code size is 24696 bytes and exceeds 24576 bytes (a limit introduced in Spurious Dragon). This contract may not be deployable on mainnet. Consider enabling the optimizer (with a low "runs" value!), turning off revert strings, or using libraries.


#start wsl
#bash "/mnt/c/Users/xdotk/torukmakto/vialend/contracts/vaults/v2.0/contracts/_solcScripts.sh"
#chmod +x _solcScripts.sh
 #./_solcScripts.sh


solc --optimize --overwrite --abi ./aave-v3-core/contracts/protocol/pool/Pool.sol -o ../build/aavePool.abi
solc --optimize --overwrite --bin ./aave-v3-core/contracts/protocol/pool/Pool.sol -o ../build/aavePool.abi
/usr/bin/abigen --abi=../build/aavePool.abi/pool.abi --bin=../build/aavePool.abi/pool.bin --pkg=api --out=../deploy/aavePool/Pool.go

solc --optimize --overwrite --abi ./aave-v3-core/contracts/protocol/configuration/PoolAddressesProvider.sol -o ../build
solc --optimize --overwrite --bin ./aave-v3-core/contracts/protocol/configuration/PoolAddressesProvider.sol -o ../build
/usr/bin/abigen --abi=../build/PoolAddressesProvider.abi --bin=../build/PoolAddressesProvider.bin --pkg=api --out=../deploy/PoolAddressesProvider/PoolAddressesProvider.go

solc --optimize --overwrite --abi SwapHelper.sol -o ../build
solc --optimize --overwrite --bin SwapHelper.sol -o ../build
/usr/bin/abigen --abi=../build/SwapHelper.abi --bin=../build/SwapHelper.bin --pkg=api --out=../deploy/SwapHelper/SwapHelper.go


solc --optimize --overwrite --abi VaultFactory.sol -o ../build
solc --optimize --overwrite --bin VaultFactory.sol -o ../build
/usr/bin/abigen --abi=../build/VaultFactory.abi --bin=../build/VaultFactory.bin --pkg=api --out=../deploy/VaultFactory/VaultFactory.go


solc --optimize --overwrite --abi VaultStrategy2.sol -o ../build
solc --optimize --overwrite --bin VaultStrategy2.sol -o ../build
/usr/bin/abigen --abi=../build/VaultStrategy2.abi --bin=../build/VaultStrategy2.bin --pkg=api --out=../deploy/VaultStrategy2/VaultStrategy2.go

solc --optimize --overwrite --abi ViaVault.sol -o ../build
solc --optimize --overwrite --bin ViaVault.sol -o ../build
/usr/bin/abigen --abi=../build/ViaVault.abi --bin=../build/ViaVault.bin --pkg=api --out=../deploy/ViaVault/ViaVault.go


solc --optimize --overwrite --abi StratDeployer.sol -o ../build
solc --optimize --overwrite --bin StratDeployer.sol -o ../build
abigen --abi=../build/StratDeployer.abi --bin=../build/StratDeployer.bin --pkg=api --out=../deploy/StratDeployer/StratDeployer.go
solc --optimize --overwrite --abi VaultDeployer.sol -o ../build
solc --optimize --overwrite --bin VaultDeployer.sol -o ../build
abigen --abi=../build/VaultDeployer.abi --bin=../build/VaultDeployer.bin --pkg=api --out=../deploy/VaultDeployer/VaultDeployer.go


solc --optimize --overwrite --abi VaultFactory.sol -o ../build
solc --optimize --overwrite --bin VaultFactory.sol -o ../build
/usr/bin/abigen --abi=../build/VaultFactory.abi --bin=../build/VaultFactory.bin --pkg=api --out=../deploy/VaultFactory/VaultFactory.go
solc --optimize --overwrite --abi VaultStrategy.sol -o ../build
solc --optimize --overwrite --bin VaultStrategy.sol -o ../build
/usr/bin/abigen --abi=../build/VaultStrategy.abi --bin=../build/VaultStrategy.bin --pkg=api --out=../deploy/VaultStrategy/VaultStrategy.go
solc --optimize --overwrite --abi ViaVault.sol -o ../build
solc --optimize --overwrite --bin ViaVault.sol -o ../build
/usr/bin/abigen --abi=../build/ViaVault.abi --bin=../build/ViaVault.bin --pkg=api --out=../deploy/ViaVault/ViaVault.go

solc --optimize --overwrite --abi test.sol -o ../build
solc --optimize --overwrite --bin test.sol -o ../build
/usr/bin/abigen --abi=../build/test.abi --bin=../build/test.bin --pkg=api --out=../deploy/test/test.go


/usr/bin/abigen --abi=../build/VaultFactory.abi --bin=../build/VaultFactory.bin --pkg=api --out=../deploy/VaultFactory/VaultFactory.go
/usr/bin/abigen --abi=../build/ViaVault.abi --bin=../build/ViaVault.bin --pkg=api --out=../deploy/ViaVault/ViaVault.go
/usr/bin/abigen --abi=../build/VaultStrategy.abi --bin=../build/VaultStrategy.bin --pkg=api --out=../deploy/VaultStrategy/VaultStrategy.go


#"C:\Program Files\Geth\abigen"  --abi=../build/ViaVault.abi --bin=../build/ViaVault.bin --pkg=api --out=../deploy/ViaVault/ViaVault.go


#abigen --abi=../build/StratUniComp.abi --bin=../build/StratUniComp.bin --pkg=api --out=../deploy/StratUniComp/StratUniComp.go
#abigen --abi=../build/ViaStrategy.abi --bin=../build/ViaStrategy.bin --pkg=api --out=../deploy/ViaStrategy/ViaStrategy.go
#abigen --abi=../build/ViaFactory.abi --bin=../build/ViaFactory.bin --pkg=api --out=../deploy/ViaFactory/ViaFactory.go


// WARNING *** suspect this to be run on linux caused problem that deposit and withdraw display name on etherscan changed to transfer*
//go run "/mnt/c/Users/xdotk/torukmakto/vialend/contracts/vaults/v2.0/scripts/index.go"
//go run ../scripts/auto/event/main.event.go -l Deposit


solc --optimize --overwrite --abi WPowerPerp.sol -o ../../build
solc --optimize --overwrite --bin WPowerPerp.sol -o ../../build
/usr/bin/abigen --abi=../../build/WPowerPerp.abi --bin=../../build/WPowerPerp.bin --pkg=api --out=../../deploy/Squeeth/WPowerPerp.go


solc --optimize --overwrite --abi TestUniswapV3Callee.sol -o ../../build
