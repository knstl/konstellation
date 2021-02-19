#!/usr/bin/env bash

COMMAND=$1
SUBCOMMAND=$2

IMAGE=${IMAGE:-"knstld:latest"}
VOLUME=${VOLUME:-knstld_data}
CHAIN_ID=${CHAIN_ID:-darchub}
MONIKER=${MONIKER:-dm0}

function usage() {
  echo "Usage:"
  echo "  ./start.sh [command] [subcommand]"
  echo ""
  echo "Command:"
  echo "  setup   Setup docker container [vol - volume, b - bind mount]"
  echo "  run     Run docker container [vol - volume, b - bind mount]"
  echo ""
}

function error() {
  echo "" >&2
  echo "Error: " >&2
  echo "    $1" >&2
  echo "" >&2
  exit 1
}

function setup_volume() {
  docker run --rm -it -e KEY_PASSWORD="$KEY_PASSWORD" -e KEY_NAME="$KEY_NAME" -e KEY_MNEMONIC="$KEY_MNEMONIC" \
   --mount type=volume,source="$VOLUME",target=/root "$IMAGE" /opt/setup.sh
}

function setup_bind() {
  docker run --rm -it -e KEY_PASSWORD="KEY_PASSWORD" -e KEY_NAME="$KEY_NAME" -e KEY_MNEMONIC="KEY_MNEMONIC" \
    -v ~/.knstld:/root/.knstld \
    "$IMAGE" /opt/setup.sh
}

function run_volume() {
  docker run --rm -it -p 26657:26657 -p 26656:26656 -p 1317:1317 \
      --mount type=volume,source="$VOLUME",target=/root \
      "$IMAGE" /opt/run.sh
}

function run_bind() {
  docker run --rm -it -p 26657:26657 -p 26656:26656 -p 1317:1317 \
      -v ~/.knstld:/root/.knstld \
      "$IMAGE" /opt/run.sh
}

if [[ -z ${COMMAND} ]]; then
  usage
  error "Command must be set!"
fi

if [[ -z ${SUBCOMMAND} ]]; then
  usage
  error "Subcommand must be set!"
fi

function run() {
  case "${SUBCOMMAND}" in
  "vol")
    run_volume
    ;;
  "b")
    run_bind
    ;;
  *)
    usage
    echo "" >&2
    exit 1
    ;;
  esac
}

function setup() {
  case "${SUBCOMMAND}" in
  "vol")
    setup_volume
    ;;
  "b")
    setup_bind
    ;;
  *)
    usage
    echo "" >&2
    exit 1
    ;;
  esac
}

case "${COMMAND}" in
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
