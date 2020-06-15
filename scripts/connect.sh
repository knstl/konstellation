#!/usr/bin/env bash

COMMAND=$1
CHAIN_ID=$2
MONIKER=$3
KEY=$4

function usage() {
  echo "Usage:"
  echo "  ./testnet.sh [command] [chain-id] [params]"
  echo ""
  echo "Command:"
  echo "  connect   Connect to testnet "
  echo "  validator Become validator "
  echo ""
}

#konstellation unsafe-reset-all

#konstellation config set moniker mmm

function connect() {
  konstellation unsafe-reset-all

  echo "Getting genesis"
  curl -o ~/.konstellation/config/genesis.json https://raw.githubusercontent.com/Konstellation/testnet/master/"$CHAIN_ID"/genesis.json
  echo "Getting app.toml"
  curl -o ~/.konstellation/config/app.toml https://raw.githubusercontent.com/Konstellation/testnet/master/darchub/app.toml
  echo "Getting config..."
  curl -o ~/.konstellation/config/config.toml https://raw.githubusercontent.com/Konstellation/testnet/master/darchub/config.toml

  konstellation config set moniker mmm

}

function validator() {
  if [[ -z ${MONIKER} ]]; then
    error "Moniker must be set !"
    usage
  fi
  if [[ -z ${KEY} ]]; then
    error "Key must be set !"
    usage
  fi

  konstellationcli tx staking create-validator --moniker "$MONIKER" --pubkey $(konstellation tendermint show-validator) --amount 100000000darc --from "$KEY" --commission-max-rate 0 --commission-rate 0 --commission-max-change-rate 0  --min-self-delegation 1 --chain-id "$CHAIN_ID"
}


if [[ -z ${COMMAND} ]]; then
  error "Command must be set !"
  usage
fi

if [[ -z ${CHAIN_ID} ]]; then
  CHAIN_ID="darchub"
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
