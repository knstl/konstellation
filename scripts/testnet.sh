#!/usr/bin/env bash

COMMAND=$1
CHAIN_ID=$2

function usage() {
  echo "Usage:"
  echo "  ./testnet.sh command [chain-id]"
  echo ""
  echo "Command:"
  echo "  deploy    Deploy testnet to testnodes. "
  echo "  copy      Copy config and genesis to yout konstellation dir. "
  echo ""
}

function create() {
  if [[ -d "testnet" ]]; then
    sudo rm -rdf testnet
  fi

  konstellation testnet --chain-id "$CHAIN_ID"
}

function deploy() {
 if [[ -f "./config/testnet.json" ]]; then
 jq -r '
    to_entries |
    .[].value |
    @sh "scp -r -i ~/Documents/.ssh/id_rsa.pub ./testnet/\(.name) root@\(.ip):/root/node"
  ' ./config/testnet.json
fi
}

function copy() {
  if [ ! -d $HOME/.konstellation ]; then
    echo "Konstellation config dir does not exist"
    echo "Run konstellation init and then run this script again"
    exit
  fi

  cp -r ./testnet/config/* $HOME/.konstellation/config/
}

if [[ -z ${COMMAND} ]]; then
  error "Command must be set !"
  usage
fi

if [[ -z ${CHAIN_ID} ]]; then
  CHAIN_ID="darchub"
fi

if [[ ! -f "./config/testnet.json" ]]; then
  echo "Nodes config DOSE NOT exist !"
  echo "" >&2
  exit 1
fi

case "${COMMAND}" in
"create")
  create
  ;;
"deploy")
  deploy
  ;;
"copy")
  copy
  ;;
*)
  usage
  echo "" >&2
  exit 1
  ;;
esac
