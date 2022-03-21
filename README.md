# ViaLend
## Sustainable Liquidity Management

Automate and connect proven DeFi protocols, for everyday investors to simply:
* earn more
* risk less 
* save time and money. 

This project builds upon our EthGlobal 2021 submission that had us
automate a liquidity provisioning vault for Uniswap V3. Using a 50/50
allocation that leverages concentrated liquidity but rebalances with no
fees or permanent loss. In the previous version, we solved this vault
design by integrating with Compound and using the lending pool as a
reserve in order to rebalance our position in V3 and avoid swap fees. In
EthDenver 2022 we want to see if we could achieve and solve three
challenges:

1. Use Squeeth to hedge impermanent loss of automated market making.
2. Use Aave flash loans to create a ETH short for delta neutrality.
3. Improve UI landing page

### Rinkeby Deployed Contracts 
# Strategy Contact
0x510c44CAFACFb4e19D9eBb34d7DC5b3B7b3340be

# Test rebalance transactions
https://rinkeby.etherscan.io/tx/0x28cbde6134ddbf2f763a3df01fe7578ff5f0d69046a3ead7af4670eeb3654287

# Test Aave short and unwind transactions
https://rinkeby.etherscan.io/tx/0xe9a572f939510e6aba4ea1f8ff042efced91f2c9478af9f0db1a7b60dc3d2c23
https://rinkeby.etherscan.io/tx/0x9dbf3c6f0aebfd6a12e4ddc3984d023a28cd552285a95622a9e3f5133048eeac

