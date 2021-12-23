[![CircleCI](https://circleci.com/gh/circleci/circleci-docs.svg?style=shield)](https://circleci.com/gh/Konstellation/konstellation)

Konstellation is the blockchain built using the [Cosmos SDK](https://github.com/cosmos/cosmos-sdk). Konstellation will be interact with other sovereign blockchains using a protocol called [IBC](https://github.com/cosmos/ics/tree/master/ibc) that enables Inter-Blockchain Communication.

# Konstellation network

## Mainnet Full Node Quick Start
With each version of the Konstellation Hub, the chain is restarted from a new Genesis state. We are currently on darchub.

Get mainnet config [here](https://github.com/Konstellation/konstellation/tree/master/config)

### Build from code

This assumes that you're running Linux or MacOS and have installed [Go 1.17+](https://golang.org/dl/).  This guide helps you:

Build, Install, and Name your Node

Current latest release is `v0.4.3`
```bash
# Clone Konstellation from the latest release found here: https://github.com/konstellation/konstellation/releases
git clone -b <latest_release> https://github.com/konstellation/konstellation
# Enter the folder Konstellation was cloned into
cd konstellation
# Change git release branch `v0.4.3`
git checkout v0.4.3
# Compile and install Konstellation
make build
# Check konstellation version
build/knstld version
```

### To join mainnet follow this steps

#### Initialize data and folders
```bash
build/knstld unsafe-reset-all
```

#### Genesis & Seeds
Download [genesis.json](https://raw.githubusercontent.com/Konstellation/konstellation/master/config/genesis.json)
```
wget -O $HOME/.knstld/config/genesis.json https://raw.githubusercontent.com/Konstellation/konstellation/master/config/genesis.json
```
Download [config.toml](https://raw.githubusercontent.com/Konstellation/konstellation/master/config/config.toml) with predefined seeds and persistent peers
```
wget -O $HOME/.knstld/config/config.toml https://raw.githubusercontent.com/Konstellation/konstellation/master/config/config.toml
```

Alternatively enter persistent peers to config.toml provided [here](https://github.com/Konstellation/konstellation/tree/master/config)

1) Open ~/.knstld/config/config.toml with text editor. Alternatively you can use cli editor, like nano ``` nano ~/.knstld/config/config.toml ```
2) Scroll down to persistant peers in `config.toml`, and add the persistant peers as a comma-separated list

#### Setting Up a New Node
You can edit this moniker, in the ~/.knstld/config/config.toml file:
```bash
# A custom human readable name for this node
moniker = "<your_custom_moniker>"
```

You can edit the ~/.knstld/config/app.toml file in order to enable the anti spam mechanism and reject incoming transactions with less than the minimum gas prices:
```
# This is a TOML config file.
# For more information, see https://github.com/toml-lang/toml

##### main base config options #####

# The minimum gas prices a validator is willing to accept for processing a
# transaction. A transaction's fees must meet the minimum of any denomination
# specified in this config (e.g. 10udarc).

minimum-gas-prices = ""
```
Your full node has been initialized!

#### Run a full node
```
# Start Konstellation
build/knstld start

# to run process in background run
screen -dmSL knstld build/knstld start

# Check your node's status with konstellation cli
build/knstld status
```

Wait for the konstellation block synchroniztion complete

### Create a key
Add new
``` bash
build/knstld keys add <key_name>
```

Or import via mnemonic
```bash
build/knstld keys add <key_name> -i
```

As a result, you got
```bash
- name: <key_name>
  type: local
  address: <key_address>
  pubkey: <key_pubkey>
  mnemonic: ""
  threshold: 0
  pubkeys: []


**Important** write this mnemonic phrase in a safe place.
It is the only way to recover your account if you ever forget your password.

<key_mnemonic>
```

### To become a validator follow this steps
Before setting up your validator node, make sure you've already gone through the [Full Node Setup](https://github.com/Konstellation/konstellation#to-join-mainnet-follow-this-steps)

#### What is a Validator?
[Validators](https://docs.cosmos.network/v0.44/modules/staking/01_state.html#validator) are responsible for committing new blocks to the blockchain through voting. A validator's stake is slashed if they become unavailable or sign blocks at the same height.
Please read about [Sentry Node Architecture](https://hub.cosmos.network/main/validators/security.html#sentry-nodes-ddos-protection) to protect your node from DDOS attacks and to ensure high-availability.

#### Create Your Validator

Your `darcvalconspub` can be used to create a new validator by staking tokens. You can find your validator pubkey by running:

```bash
build/knstld tendermint show-validator
```

To create your validator, just use the following command:

Check if your key(address) has enough balance:

```bash
build/knstld query bank balances <key address>
```

For test nodes, `chain-id` is `darchub`.\
You need transction fee `2udarc` to make your transaction for creating validator.\
Don't use more `udarc` than you have! 

```bash
build/knstld tx staking create-validator \
  --amount=1000000udarc \
  --pubkey=$(build/knstld tendermint show-validator) \
  --moniker=<choose a moniker> \
  --chain-id=<chain_id> \
  --commission-rate="0.10" \
  --commission-max-rate="0.20" \
  --commission-max-change-rate="0.01" \
  --min-self-delegation="1" \
  --from=<key_name> \
  --fees=2udarc
```

* NOTE: If you have troubles with \'\\\' symbol, run the command in a single line like `build/knstld tx staking create-validator --amount=1000000udarc --pubkey=$(build/knstld tendermint show-validator) ...`

When specifying commission parameters, the `commission-max-change-rate` is used to measure % _point_ change over the `commission-rate`. E.g. 1% to 2% is a 100% rate increase, but only 1 percentage point.

`Min-self-delegation` is a strictly positive integer that represents the minimum amount of self-delegated voting power your validator must always have. A `min-self-delegation` of 1 means your validator will never have a self-delegation lower than `1000000udarc`

You can check that you are in the validator set by using a third party explorer or using cli tool
```bash
build/knstld query staking validators --chain-id=<chain_id>
```

* Note: You can edit the params after, by running command `build/knstld tx staking edit-validator ... —from <key_name> --chain-id=<chain_id> --fees=2udarc` with the necessary options

## How to init chain

Add key to your keyring
```build/knstld keys add <key name>```

Initialize genesis and config files.
```build/knstld init <moniker> --chain-id <chain id>```

Replace all denoms `stake` to `udarc` in `genesis.json`

Add genesis account
```build/knstld add-genesis-account <key name> 200000000000udarc``` - 200000darc

Create genesis transaction
```build/knstld gentx <key name> 100000000000udarc --chain-id <chain id>``` - create CreateValidator transaction

Collect all of gentxs
```build/knstld collect-gentxs```

Run network
```build/knstld start```

## Dockerized

We provide a docker image to help with test setups. There are two modes to use it

Build: ```docker build -t knstld:latest .```

### Dev server
Bring up a local node with a test account containing tokens

This is just designed for local testing/CI - do not use these scripts in production. Very likely you will assign tokens to accounts whose mnemonics are public on github.
Prepend `VOLTYPE=vol|b` if you want to bind mount or volume into container as storage

#### Set IMAGE env variable
```shell script
export IMAGE=knstld:latest
```
#### Set other env variable
```shell script
export CHAIN_ID=darchub
export MONIKER=<YOUR_MONIKER>
```

**You can git clone and run script or call docker commands directly****
**In a case of using volume (not bind mount) as containe volume, change `-v ~/.knstld:/root/.knstld` to `--mount type=volume,source="$VOLUME",target=/root`**

#### Init
Initialize blockchain folder
```shell script
# Using script
./docker/start.sh init

# Using docker cmd
docker run --rm -it \
  -e MONIKER="$MONIKER" \
  -e CHAIN_ID="$CHAIN_ID" \
  -v ~/.knstld:/root/.knstld "$IMAGE" /opt/init.sh
```

#### Moniker
Change moniker
```shell script
export MONIKER=moniker 

# Using script
./docker/start.sh config

# Using docker cmd
docker run --rm -it \
  -e MONIKER="$MONIKER" \
  -v ~/.knstld:/root/.knstld "$IMAGE" /opt/config.sh
```

#### Setup
**Omit `KEY_NAME`, `KEY_PASSWORD`, `KEY_MNEMONIC` if you want to create a new identity.**
Setup genaccs, gentxs, collectGentxs

```shell script
export KEY_PASSWORD="..."
export KEY_NAME="..."
export KEY_MNEMONIC="..."

# Using script
./docker/start.sh setup

# Using docker cmd
docker run --rm -it \
  -e KEY_PASSWORD="$KEY_PASSWORD" \
  -e KEY_NAME="$KEY_NAME" \
  -e KEY_MNEMONIC="$KEY_MNEMONIC" \
  -v ~/.knstld:/root/.knstld "$IMAGE" /opt/setup.sh
```

#### Run 
Run blockchain node in container
```shell script
# Using script
./docker/start.sh run

# Using docker cmd
docker run --rm -it \
  -p 26657:26657 \
  -p 26656:26656 \
  -p 1317:1317 \
  -p 9090:9090 \
  -v ~/.knstld:/root/.knstld  "$IMAGE" /opt/run.sh
```

### Localnet
See `Localnet` in [Cosmodrome](https://github.com/konstellation/cosmodrome) project repo

## Network generation
To generate network see `Testnet` in [Cosmodrome](https://github.com/konstellation/cosmodrome) project repo
