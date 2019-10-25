#!/usr/bin/env bash

COMMAND=$1
CHAIN_ID=$2

function usage() {
  echo "Usage:"
  echo "  ./testnet.sh command [chain-id]"
  echo ""
  echo "Command:"
  echo "  run       Run testnet full node on the server side. "
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

function run() {
  source $(pwd)/config/.env

  NODE_NAME=$(hostname)

  if [[ ! -d ${NODE_ROOT} ]]; then
    echo "Node's config DOSE NOT exist !"
    echo "" >&2
    exit 1
  fi

  containers=$(docker container ls | awk '{print $1}' | sed -n 2p)
  if [ ! -z "$containers" ]; then
    echo -n "Remove $containers ... "
    docker rm -f "$NODE_NAME" >/dev/null
  fi

  echo -n "Create ${NODE_NAME} ... "
  echo ""
  docker run -d \
    --name "$NODE_NAME" \
    --net=host \
    -e CHAIN_ID="$CHAIN_ID" \
    -e MONIKER="$NODE_NAME" \
    -e NODE_TYPE=PRIVATE_TESTNET \
    -p 26666:26656 \
    -p 26667:26657 \
    -p 26670:26660 \
    -v "${NODE_ROOT}"/konstellation:/root/.konstellation \
    -v "${NODE_ROOT}"/konstellationcli:/root/.konstellationcli \
    "$IMAGE_OWNER"/konstellation:"$CHAIN_ID"
  echo "Done !"
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
"run")
  run
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
