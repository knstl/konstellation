#!/usr/bin/env bash

COMMAND=$1
CHAIN_ID=$2
MONIKER=$3
KEY_NAME=$4

function usage() {
  echo "Usage:"
  echo "  ./testnet.sh [command] [chain-id] [[MONIKER] [KEY_NAME]"
  echo ""
  echo "Command:"
  echo "  connect   Connect to testnet "
  echo "  validator Become a validator "
  echo ""
}

function error() {
  echo "" >&2
  echo "Error: " >&2
  echo "    $1" >&2
  echo "" >&2
  exit 1
}

function connect() {
  konstellation unsafe-reset-all

  echo "Getting genesis"
  curl -o ~/.konstellation/config/genesis.json https://raw.githubusercontent.com/Konstellation/testnet/master/"$CHAIN_ID"/genesis.json
  echo "Getting app.toml"
  curl -o ~/.konstellation/config/app.toml https://raw.githubusercontent.com/Konstellation/testnet/master/"$CHAIN_ID"/app.toml
  echo "Getting config..."
  curl -o ~/.konstellation/config/config.toml https://raw.githubusercontent.com/Konstellation/testnet/master/"$CHAIN_ID"/config.toml

  konstellation config set moniker mmm
}

function validator() {
  if [[ -z ${MONIKER} ]]; then
    usage
    error "MONIKER must be set!"
  fi
  if [[ -z ${KEY_NAME} ]]; then
    usage
    error "KEY_NAME must be set!"
  fi

  konstellationcli tx staking create-validator --moniker "$MONIKER" --pubkey $(konstellation tendermint show-validator) --amount 100000000darc --from "$KEY_NAME" --commission-max-rate 0 --commission-rate 0 --commission-max-change-rate 0  --min-self-delegation 1 --chain-id "$CHAIN_ID"
}


if [[ -z ${COMMAND} ]]; then
  usage
  error "Command must be set!"
fi

if [[ -z ${CHAIN_ID} ]]; then
  usage
  error "CHAIN_ID must be set!"
fi

case "${COMMAND}" in
"connect")
  connect
  ;;
"validator")
  validator
  ;;
*)
  usage
  echo "" >&2
  exit 1
  ;;
esac
