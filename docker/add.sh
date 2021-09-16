#!/bin/sh
#set -o errexit -o nounset -o pipefail

KEY_PASSWORD=${KEY_PASSWORD:-1234567890}
KEY_NAME=${KEY_NAME:-validator}
CHAIN_ID=${CHAIN_ID:-darchub}

# ------------------------------------------------------------------------------
#
# Print env variables
#
# ------------------------------------------------------------------------------
echo "Chain id" "${CHAIN_ID}"
echo "Key name" "${KEY_NAME}"
echo "Mnemonic" "${KEY_MNEMONIC}"
echo "Password" "${KEY_PASSWORD}"

if test -n "${KEY_MNEMONIC-}"
then
  echo "$KEY_MNEMONIC"
    {
      echo "${KEY_MNEMONIC}"
      echo "${KEY_PASSWORD}"
      echo "${KEY_PASSWORD}"
      echo
    } | knstld keys add "${KEY_NAME}" --recover
  # hardcode the validator account for this instance
else
    {
      echo "${KEY_PASSWORD}"
      echo "${KEY_PASSWORD}"
      echo
    } | knstld keys add "${KEY_NAME}"
fi
