#!/usr/bin/env bash

PARAM_COUNT=2
COMMAND=$1
CHAIN_ID=$2
NODE_PREFIX="testnode-"
DOCKER_NETWORK="konstellation-network"
TEMP_CONTAINER="temp-localnet"

function usage() {
  echo "Usage:"
  echo "  ./localnet.sh command chain-id"
  echo ""
  echo "Command:"
  echo "  create   Create network. "
  echo "  run      Create new container for each node. "
  echo "  start    Start exist containers. "
  echo "  stop     Stop exist containers. "
  echo "  rm       Remove exist containers. "
  echo ""
}

function node_count() {
  # shellcheck disable=SC2010
  ls -1 localnet | grep node -c
}

function create_localnet() {
  params="-e CHAIN_ID=${CHAIN_ID} -e NODE_TYPE=PRIVATE_TESTNET "
  if [[ -d "localnet" ]]; then
    sudo rm -rdf localnet
  fi
  if [[ "" != "$(docker ps | grep ${TEMP_CONTAINER})" ]]; then
    docker rm -f ${TEMP_CONTAINER} >/dev/null
  fi
  docker run -d --name ${TEMP_CONTAINER} ${params} konstellation:${CHAIN_ID}
  #    docker run -d --name ${TEMP_CONTAINER} ${params} konstellation:${CHAIN_ID} > /dev/null
  docker exec ${TEMP_CONTAINER} sh -c "konstellation localnet --chain-id ${CHAIN_ID} --output-dir /localnet"
  docker cp ${TEMP_CONTAINER}:/localnet ./
  docker rm -f ${TEMP_CONTAINER} >/dev/null
}

function create() {
  # Create a network for connections between nodes
  if [[ "" == "$(docker network ls | grep ${DOCKER_NETWORK})" ]]; then
    docker network create ${DOCKER_NETWORK}
  fi

  create_localnet
}

function run() {
  # Create new container for each node
  for ((i = 0; i < $(node_count); i++)); do
    NODE_ROOT=$(pwd)/localnet/node${i}
    if [[ ! -d ${NODE_ROOT} ]]; then
      echo "Node${i}'s config DOSE NOT exist !"
      echo "" >&2
      exit 1
    fi

    NODE_NAME=${NODE_PREFIX}${i}
    echo -n "Create ${NODE_NAME} ... "
    docker run -d \
      --name ${NODE_NAME} \
      --net ${DOCKER_NETWORK} \
      -e CHAIN_ID=${CHAIN_ID} \
      -e MONIKER=testnode-${i} \
      -e NODE_TYPE=PRIVATE_TESTNET \
      -v ${NODE_ROOT}/konstellation:/root/.konstellation \
      -v ${NODE_ROOT}/konstellationcli:/root/.konstellationcli \
      konstellation:${CHAIN_ID}
    echo "Done !"
  done
}

function start() {
  for ((i = 0; i < $(node_count); i++)); do
    NODE_NAME=${NODE_PREFIX}${i}
    echo -n "Start ${NODE_NAME} ... "
    docker start ${NODE_NAME}
    echo "Done !"
  done
}

function stop() {
  for ((i = 0; i < $(node_count); i++)); do
    NODE_NAME=${NODE_PREFIX}${i}
    echo -n "Stop ${NODE_NAME} ... "
    docker stop ${NODE_NAME}
    echo "Done !"
  done
}

function rm() {
  for ((i = 0; i < $(node_count); i++)); do
    NODE_NAME=${NODE_PREFIX}${i}
    echo -n "Remove ${NODE_NAME} ... "
    docker rm -f ${NODE_NAME} >/dev/null
    echo "Done !"
  done
}

if [[ $# != "${PARAM_COUNT}" ]]; then
  usage
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
*)
  usage
  echo "" >&2
  exit 1
  ;;
esac
