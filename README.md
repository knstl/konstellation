# Konstellation network

IRIS network (a.k.a. IRISnet) aims to establish a technology foundation to facilitate construction of next-generation distributed applications. By incorporating a comprehensive service infrastructure and an enhanced IBC protocol into the Tendermint & Cosmos stack, IRISnet enables service interoperability as well as token transfers across an internet of blockchains. As the centerpiece of IRISnet, IRIS Hub (a.k.a. IRIShub) will be the first regional hub connecting to the main Cosmos Hub, thus making IRISnet an inseparable part of the whole Cosmos network.

#### Run singlenet
Run in shell from project dir
```shell script
./scripts/singlenet.sh
```

#### Connect to network
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
