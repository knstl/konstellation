# Konstellation network

IRIS network (a.k.a. IRISnet) aims to establish a technology foundation to facilitate construction of next-generation distributed applications. By incorporating a comprehensive service infrastructure and an enhanced IBC protocol into the Tendermint & Cosmos stack, IRISnet enables service interoperability as well as token transfers across an internet of blockchains. As the centerpiece of IRISnet, IRIS Hub (a.k.a. IRIShub) will be the first regional hub connecting to the main Cosmos Hub, thus making IRISnet an inseparable part of the whole Cosmos network.

#### Run singlenet
Run in shell from project dir
```shell script
./scripts/singlenet.sh
```

## Localnet

Run in shell from project dir
#### Create localnet
```shell script
./scripts/localnet.sh create
```
#### Run localnet
```shell script
./scripts/localnet.sh run
```
#### Copy config and genesis to konstellation dir
```shell script
./scripts/localnet.sh copy
```
#### Connect to localnet
```shell script
konstellation unsafe-reset-all
konstellation config set moniker {MONIKER}
konstellation start
```

#### Stop and remove localnet
```shell script
./scripts/localnet.sh stop
./scripts/localnet.sh rm
```

## Testnet
Run in shell from project dir
#### Create testnet
```shell script
./scripts/testnet.sh create
```
#### Deploy testnet
```shell script
./scripts/testnet.sh deploy
```
#### Run testnet nodes on the server side
```shell script
./scripts/testnet.sh run
```
#### Copy config and genesis to konstellation dir
```shell script
./scripts/testnet.sh copy
```
#### Connect to testnet
```shell script
konstellation unsafe-reset-all
konstellation config set moniker {MONIKER}
konstellation start
```

#### Run full node
```shell script
./scripts/fullnode.sh
```

## Konstellation Hub Mainnet

To join the mainnet, follow
[this guide].

## Install

See the 
[install instructions]

## Resources

* [Explorer]: https://www.konsteplorer.io/#/home
* [Demo wallet]: https://www.konstebox.io/#/home
