#!/bin/bash
# microtick and bitcanna contributed significantly here.
set -uxe

# set environment variables
export GOPATH=~/go
export PATH=$PATH:~/go/bin
export RPC=http://node1.konstellation.tech:26657
export RPCN=http://node1.konstellation.tech:26657
export APPNAME=KNSTLD

# Install Gaia
go install ./...

# MAKE HOME FOLDER AND GET GENESIS
knstld init notional-konstellation-relays
wget -O ~/.knstld/config/genesis.json https://raw.githubusercontent.com/Konstellation/konstellation/master/config/genesis.json


INTERVAL=1000

# GET TRUST HASH AND TRUST HEIGHT

LATEST_HEIGHT=$(curl -s $RPC/block | jq -r .result.block.header.height);
BLOCK_HEIGHT=$(($LATEST_HEIGHT-INTERVAL))
TRUST_HASH=$(curl -s "$RPC/block?height=$BLOCK_HEIGHT" | jq -r .result.block_id.hash)


# TELL USER WHAT WE ARE DOING
echo "TRUST HEIGHT: $BLOCK_HEIGHT"
echo "TRUST HASH: $TRUST_HASH"


# export state sync vars
export $(echo $APPNAME)_STATESYNC_ENABLE=true
export $(echo $APPNAME)_P2P_MAX_NUM_OUTBOUND_PEERS=500
export $(echo $APPNAME)_STATESYNC_RPC_SERVERS="$RPC,$RPCN"
export $(echo $APPNAME)_STATESYNC_TRUST_HEIGHT=$BLOCK_HEIGHT
export $(echo $APPNAME)_STATESYNC_TRUST_HASH=$TRUST_HASH
export $(echo $APPNAME)_P2P_SEEDS="1bd4b89e05e5d7ea5d2dba89c799c2e624cb35d7@node1.konstellation.tech:26656"


knstld start 
