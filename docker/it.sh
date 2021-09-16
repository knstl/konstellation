#!/bin/sh
#set -o errexit -o nounset -o pipefail

KEY_PASSWORD=${KEY_PASSWORD:-1234567890}
KEY_NAME=${KEY_NAME:-validator}
CHAIN_ID=${CHAIN_ID:-darchub}
args="$*"
CMD=${CMD:-version}
if [ -n "$args" ]
then
  CMD=$args
fi

knstld $CMD

