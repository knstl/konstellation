[![CircleCI](https://circleci.com/gh/circleci/circleci-docs.svg?style=shield)](https://circleci.com/gh/Konstellation/konstellation)

Konstellation is the blockchain built using the [Cosmos SDK](https://github.com/cosmos/cosmos-sdk). Konstellation will be interact with other sovereign blockchains using a protocol called [IBC](https://github.com/cosmos/ics/tree/master/ibc) that enables Inter-Blockchain Communication.

# Konstellation network

## Testnet Full Node Quick Start
With each version of the Konstellation Hub, the chain is restarted from a new Genesis state. We are currently on knstlhub-1.

Get testnet config [here](https://github.com/Konstellation/testnet)

### Build from code

This assumes that you're running Linux or MacOS and have installed [Go 1.14+](https://golang.org/dl/).  This guide helps you:

* build and install Konstellation
* allow you to name your node
* download config file of add seeds to your config file
* download genesis state
* start your node
* use konstellation to check the status of your node.

Build, Install, and Name your Node
```bash
# Clone Konstellation from the latest release found here: https://github.com/konstellation/konstellation/releases
git clone -b <latest_release> https://github.com/konstellation/konstellation
# Enter the folder Konstellation was cloned into
cd konstellation
# Compile and install Konstellation
make install
```

### Using binaries
```bash
# linux
wget https://github.com/Konstellation/konstellation/releases/download/{KONSTELLATION_VERSION}/linux_amd64.tar.gz
tar -xvzf linux_amd64.tar.gz
sudo cp ./linux_amd64/* /usr/local/bin
# macos
wget https://gist.github.com/Konstellation/b9168ec665bf8991a1cd20fd999452fa/raw/2c53c4c2fa0d90e7a10a6b7f2b5e28c35bec73d2/darwin_amd64.tar.gz

# win
wget https://gist.github.com/Konstellation/b9168ec665bf8991a1cd20fd999452fa/raw/2c53c4c2fa0d90e7a10a6b7f2b5e28c35bec73d2/windows_amd64.tar.gz

```

### To join testnet follow this steps
Download Genesis, Start your Node, Check your Node Status
```bash
# Initialize data and folders
# konstellation init {MONIKER} --chain-id {CHAIN_ID}
konstellation unsafe-reset-all
```

#### Genesis & Seeds
```
# Download genesis.json
wget -O $HOME/.konstellation/config/genesis.json https://raw.githubusercontent.com/Konstellation/testnet/master/{CHAIN_ID}/genesis.json
wget -O $HOME/.konstellation/config/config.toml https://raw.githubusercontent.com/Konstellation/testnet/master/{CHAIN_ID}/config.toml
# Alternatively enter persistant peers to config.toml provided below.
nano ~/.konstellation/config/config.toml
# Scroll down to persistant peers in `config.toml`, and add the persistant peers as a comma-separated list
```

#### Setting Up a New Node
Name your node
```
konstellation config set moniker {MONIKER}
```

You can edit this moniker later, in the ~/.gaiad/config/config.toml file:
```bash
# A custom human readable name for this node
moniker = "<your_custom_moniker>"
```

You can edit the ~/.gaiad/config/app.toml file in order to enable the anti spam mechanism and reject incoming transactions with less than the minimum gas prices:
```
# This is a TOML config file.
# For more information, see https://github.com/toml-lang/toml

##### main base config options #####

# The minimum gas prices a validator is willing to accept for processing a
# transaction. A transaction's fees must meet the minimum of any denomination
# specified in this config (e.g. 10uatom).

minimum-gas-prices = ""
```
Your full node has been initialized!

#### Run a full node
```
# Start Konstellation
konstellation start
# Check your node's status with konstellationcli
konstellationcli status
```

### Create a key
Add new
``` bash
konstellationcli keys add <key_name>
```

Or import via mnemonic
```bash
konstellationcli keys add <key_name> -i
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
Before setting up your validator node, make sure you've already gone through the [Full Node Setup](https://github.com/Konstellation/konstellation#to-join-testnet-follow-this-steps)

#### What is a Validator?
[Validators](https://hub.cosmos.network/master/validators/overview.html) are responsible for committing new blocks to the blockchain through voting. A validator's stake is slashed if they become unavailable or sign blocks at the same height.
Please read about [Sentry Node Architecture](https://hub.cosmos.network/master/validators/validator-faq.html#how-can-validators-protect-themselves-from-denial-of-service-attacks) to protect your node from DDOS attacks and to ensure high-availability.

#### Create Your Validator

Your `darcvalconspub` can be used to create a new validator by staking tokens. You can find your validator pubkey by running:

```bash
konstellation tendermint show-validator
```

To create your validator, just use the following command:
 
Don't use more `udarc` than you have! 

```bash
konstellationcli tx staking create-validator \
  --amount=100000000000udarc \
  --pubkey=$(konstellation tendermint show-validator) \
  --moniker="choose a moniker" \
  --chain-id=<chain_id> \
  --commission-rate="0.10" \
  --commission-max-rate="0.20" \
  --commission-max-change-rate="0.01" \
  --min-self-delegation="1" \
  --from=<key_name>
```

When specifying commission parameters, the `commission-max-change-rate` is used to measure % _point_ change over the `commission-rate`. E.g. 1% to 2% is a 100% rate increase, but only 1 percentage point.

`Min-self-delegation` is a strictly positive integer that represents the minimum amount of self-delegated voting power your validator must always have. A `min-self-delegation` of 1 means your validator will never have a self-delegation lower than `1000000darc`

You can confirm that you are in the validator set by using a third party explorer or using cli tool
```bash
konstellationcli q staking validator $(konstlelation tendermint show-validator)
```

### Run singlenet in docker container 
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
