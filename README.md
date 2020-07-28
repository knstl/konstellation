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
# Download genesis.json
wget -O $HOME/.konstellation/config/genesis.json https://raw.githubusercontent.com/Konstellation/testnet/master/{CHAIN_ID}/genesis.json
wget -O $HOME/.konstellation/config/config.toml https://raw.githubusercontent.com/Konstellation/testnet/master/{CHAIN_ID}/config.toml
# Alternatively enter persistant peers to config.toml provided below.
nano ~/.konstellation/config/config.toml
# Scroll down to persistant peers in `config.toml`, and add the persistant peers as a comma-separated list
# Name your node
konstellation config set moniker {MONIKER}
# Start Konstellation
konstellation start
# Check your node's status with konstellationcli
konstellationcli status
```

#### Run singlenet in docker container 
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
