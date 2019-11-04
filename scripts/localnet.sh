#!/usr/bin/env bash

COMMAND=$1
CHAIN_ID=$2
DOCKER_NETWORK="konstellation-network"

function usage() {
  echo "Usage:"
  echo "  ./localnet.sh command [chain-id]"
  echo ""
  echo "Command:"
  echo "  create   Create network. "
  echo "  run      Create new container for each node. "
  echo "  start    Start exist containers. "
  echo "  stop     Stop exist containers. "
  echo "  rm       Remove exist containers. "
  echo "  copy     Copy config and genesis to yout konstellation dir. "
  echo ""
}

function create() {
  # Create a network for connections between nodes
  if [[ "" == "$(docker network ls | grep ${DOCKER_NETWORK})" ]]; then
    docker network create ${DOCKER_NETWORK}
  fi

  if [[ -d "localnet" ]]; then
    sudo rm -rdf localnet
  fi

  konstellation localnet --chain-id "$CHAIN_ID"
}

function run() {
  jq -r '
    .[] |
    [.name, .ip] |
    @tsv' ./config/localnet.json |
    while IFS=$'\t' read -r NODE_NAME NODE_IP; do
      NODE_ROOT=$(pwd)/localnet/$NODE_NAME
      if [[ ! -d ${NODE_ROOT} ]]; then
        echo "$NODE_NAME's config DOSE NOT exist !"
        echo "" >&2
        exit 1
      fi

      echo -n "Create ${NODE_NAME} ... "
      docker run -d \
        --name "$NODE_NAME" \
        --net "$DOCKER_NETWORK" \
        --ip "$NODE_IP" \
        -e CHAIN_ID="$CHAIN_ID" \
        -e MONIKER="$NODE_NAME" \
        -e NODE_TYPE=PRIVATE_TESTNET \
        -v "$NODE_ROOT"/konstellation:/root/.konstellation \
        -v "$NODE_ROOT"/konstellationcli:/root/.konstellationcli \
        konstellation:"$CHAIN_ID"
      echo "Done !"
    done
}

function start() {
  jq -r '
    .[] |
    [.name, .ip] |
    @tsv' ./config/localnet.json |
    while IFS=$'\t' read -r NODE_NAME _; do
      echo -n "Start $NODE_NAME ... "
      docker start "$NODE_NAME"
      echo "Done !"
    done
}

function stop() {
  jq -r '
    .[] |
    [.name, .ip] |
    @tsv' ./config/localnet.json |
    while IFS=$'\t' read -r NODE_NAME _; do
      echo -n "Stop $NODE_NAME ... "
      docker stop "$NODE_NAME"
      echo "Done !"
    done
}

function rm() {
  jq -r '
    .[] |
    [.name, .ip] |
    @tsv' ./config/localnet.json |
    while IFS=$'\t' read -r NODE_NAME NODE_IP; do
      echo -n "Remove $NODE_NAME ... "
      docker rm -f "$NODE_NAME"
      echo "Done !"
    done
}

function copy() {
  if [ ! -d $HOME/.konstellation ]; then
    echo "Konstellation config dir does not exist"
    echo "Run konstellation init and then run this script again"
    exit
  fi

  cp -r ./localnet/config/* $HOME/.konstellation/config/
}

if [[ -z ${COMMAND} ]]; then
  error "Command must be set !"
  usage
fi

if [[ -z ${CHAIN_ID} ]]; then
  source ./config/.env
fi

if [[ ! -f "./config/localnet.json" ]]; then
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
"start")
  start
  ;;
"stop")
  stop
  ;;
"rm")
  rm
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
