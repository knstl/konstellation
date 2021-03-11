#!/usr/bin/env bash

COMMAND=$1

DOCKER_NETWORK=${DOCKER_NETWORK:-"konstellation-network"}
IMAGE=${IMAGE:-"knstld:latest"}
CHAIN_ID=${CHAIN_ID:-darchub}
MONIKER=${MONIKER:-dm0}


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
  if [[ "" == "$(docker network ls | grep "${DOCKER_NETWORK}")" ]]; then
    docker network create "$DOCKER_NETWORK" --subnet=172.18.0.0/16
  fi

  if [[ -d "localnet" ]]; then
    sudo rm -rdf localnet
  fi

  cosmodrome gn --chain-id "$CHAIN_ID" -n ./config/localnet.json -o ./localnet --keyring-backend test
}

function run() {
  jq -r '
    .validators[] |
    [.name, .ip] |
    @tsv' ./config/localnet.json |
    while IFS=$'\t' read -r NODE_NAME NODE_IP; do
      NODE_ROOT=$(pwd)/localnet/$NODE_NAME
      if [[ ! -d ${NODE_ROOT} ]]; then
        echo "$NODE_NAME's config DOES NOT exist !"
        echo "" >&2
        exit 1
      fi

      echo -n "Create ${NODE_NAME} on $DOCKER_NETWORK:$NODE_IP ... "
       docker run -d \
        --name "$NODE_NAME" \
        --net "$DOCKER_NETWORK" \
        --ip "$NODE_IP" \
       -v "$NODE_ROOT"/.knstld:/root/.knstld \
      "$IMAGE" /opt/run.sh

      echo "Done !"
    done
}

function start() {
  jq -r '
    .validators[] |
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
    .validators[] |
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
    .validators[] |
    [.name, .ip] |
    @tsv' ./config/localnet.json |
    while IFS=$'\t' read -r NODE_NAME NODE_IP; do
      echo -n "Remove $NODE_NAME ... "
      docker rm -f "$NODE_NAME"
      echo "Done !"
    done
}

function copy() {
  if [ ! -d "$HOME"/.knstld ]; then
    echo "Konstellation config dir does not exist"
    echo "Run konstellation init and then run this script again"
    exit
  fi

  cp -r ./localnet/config/* "$HOME"/.knstld/config/
}

if [[ -z ${COMMAND} ]]; then
  error "Command must be set !"
  usage
fi

if [[ ! -f "./config/localnet.json" ]]; then
  echo "Nodes config DOES NOT exist !"
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
