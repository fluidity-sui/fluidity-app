# Fluidity Ethereum Worker Server

This codebase contains the offchain worker implementation for Fluidity's
Ethereum deployment.

It takes fluidity transactions from Ethereum, determines their probability of
winning, generates random numbers for them, and sends the result to the worker
client.

## Environment variables (Ethereum server)

| Name                                        | Description                                                                  |
| ------------------------------------------- | ---------------------------------------------------------------------------- |
| `FLU_WORKER_ID`                             | Worker ID used to identify the application in logging and to the AMQP queue. |
| `FLU_REDIS_ADDR`                            | Redis address pointing to Redis for state tracking                           |
| `FLU_DEBUG`                                 | Toggle debug messages produced by any application using the debug logger.    |
| `FLU_AMQP_QUEUE_ADDR`                       | AMQP queue address connected to to receive and send messages down.           |
| `FLU_ETHEREUM_CONTRACT_ADDR`                | The contract address to use based on the backend for rewarding users.        |
| `FLU_ETHEREUM_REGISTRY_ADDR`                | The registry address to query to get utility information.                    |
| `FLU_ETHEREUM_HTTP_URL`                     | HTTP address to use to connect to Geth.                                      |
| `FLU_ETHEREUM_TOKEN_BACKEND`                | Token backend to use. (compound, aave).                                      |
| `FLU_ETHEREUM_CTOKEN_ADDR`                  | Address of the corresponding CToken (ie, cUSDT)                              |
| `FLU_ETHEREUM_UNISWAP_ANCHORED_VIEW_ADDR`   | Address of the Uniswap anchored view price oracle.                           |
| `FLU_ETHEREUM_AAVE_ADDRESS_PROVIDER_ADDR`   | Address of the AAVE address provider.                                        |
| `FLU_ETHEREUM_ATOKEN_ADDR`                  | Address of the AToken to use for the underlying token.                       |
| `FLU_ETHEREUM_USD_TOKEN_ADDR`               | Token address to use to look up the price of Ethereum using AAVE.            |
| `FLU_ETHEREUM_ETH_TOKEN_ADDR`               | Token address to use to look up to get the price of Ethereum using AAVE.     |
| `FLU_ETHEREUM_UNDERLYING_TOKEN_ADDR`        | Underlying token address to use to get the APY of a token from AAVE.         |
| `FLU_ETHEREUM_UNDERLYING_TOKEN_DECIMALS`    | Underlying token decimals in place (18 for DAI, 6 for USDT and USDC, etc).   |
| `FLU_AURORA_ETH_FLUX_ADDRESS`               | Flux oracle address for looking up the price of Eth on Aurora.               |
| `FLU_AURORA_TOKEN_FLUX_ADDRESS`             | Flux oracle address for looking up the price of the current token on Aurora. |
| `FLU_ETHEREUM_ETH_DECIMAL_PLACES`           | Token decimals returned for the value of Eth by Flux oracles (usu. 8).       |
| `FLU_ETHEREUM_AMQP_QUEUE_NAME`              | Queue name to send messages down from the worker server to the client.       |
| `FLU_ETHEREUM_REDIS_APY_MOVING_AVERAGE_KEY` | Moving average key to use to track the APY with using Redis.                 |
| `FLU_ETHEREUM_APPLICATION_CONTRACTS`        | List of comma-separated contract addresses to track as applications. (e.g. `0xae461ca67b15dc8dc81ce7615e0320da1a9ab8d5,0x3041cbd36888becc7bbcbc0045e3b1f144466f5f` ) |
| `FLU_ETHEREUM_CHAINLINK_HTTP_URL`           | URL to use as the custom geth network for determining the price of ETH if the currently deployed network does not support it.                                        |
| `FLU_ETHEREUM_GLOBAL_UTILITY_REWARDS`       | (<program name>,)+ is used to enable global utility rewards for each transfer. |

## Notes

### Global program application

The global program (applications.ApplicationGlobalProgram) can be set
to reward every FLUID transfer with an additional utility reward
payout. FLU_ETHEREUM_GLOBAL_UTILITY_REWARDS is
used for this.

## Building

    make build

## Building (Docker)

    make docker
