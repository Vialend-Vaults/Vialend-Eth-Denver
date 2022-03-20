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
