#!/usr/bin/env bash
set -o xtrace

COMMAND=$1

IMAGE=${IMAGE:-"knstld:latest"}
VOLUME=${VOLUME:-knstld_data}
CHAIN_ID=${CHAIN_ID:-darchub}
MONIKER=${MONIKER:-dm0}
VOLTYPE=${VOLTYPE:-b}
TEST=${TEST:-}

function usage() {
  echo "Usage:"
  echo " [IMAGE=] [VOLUME=vol|b] ./start.sh [command]"
  echo " vol - volume, b - bind mount"
  echo "Command:"
  echo "  init    Init docker container "
  echo "  setup   Setup docker container "
  echo "  run     Run docker container "
  echo ""
}

function error() {
  echo "" >&2
  echo "Error: " >&2
  echo "    $1" >&2
  echo "" >&2
  exit 1
}


volumes=()

if [[ -n $TEST ]]
then
  volumes+=( -v ~/pj/konstellation/docker/:/opt/ )
fi

if [[ $VOLTYPE = "b" ]]
then
  volumes+=( -v ~/.knstld:/root/.knstld )
else
  volumes+=( --mount type=volume,source="$VOLUME",target=/root )
fi

if [[ -z ${COMMAND} ]]; then
  usage
  error "Command must be set!"
fi

function init() {
  docker run --rm -it \
  -e MONIKER="$MONIKER" \
  -e CHAIN_ID="$CHAIN_ID" \
    "${volumes[@]}" "$IMAGE" /opt/init.sh
}

function setup() {
  docker run --rm -it \
  -e KEY_PASSWORD="$KEY_PASSWORD" \
  -e KEY_NAME="$KEY_NAME" \
  -e KEY_MNEMONIC="$KEY_MNEMONIC" \
    "${volumes[@]}" "$IMAGE" /opt/setup.sh
}

function run() {
  docker run --rm -it \
   -p 26657:26657 \
   -p 26656:26656 \
   -p 1317:1317 \
   -p 9090:9090 \
    "${volumes[@]}" "$IMAGE" /opt/run.sh
}

function config() {
  docker run --rm -it \
  -e MONIKER="$MONIKER" \
    "${volumes[@]}" "$IMAGE" /opt/config.sh
}

case "${COMMAND}" in
"init")
  init
  ;;
"config")
  config
  ;;
"setup")
  setup
  ;;
"run")
  run
  ;;
*)
  usage
  echo "" >&2
  exit 1
  ;;
esac
