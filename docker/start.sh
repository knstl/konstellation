#!/usr/bin/env bash

# ------------------------------------------------------------------------------
#
# Parameters which might be changed while upgrading.
#
# ------------------------------------------------------------------------------
INITIALIZED_FLAG="/initialized.flag"

# ------------------------------------------------------------------------------
#
# Initial Konstellation block chain full node
#
#   1. Init server's local configuration
#   2. Download preset configuration from github
#   3. Add flag to avoid reinitialization
#
# ------------------------------------------------------------------------------
function init_full_node() {
  if [[ -z ${MONIKER} ]]; then
    echo "Environment MONIKER must be set !" >&2
    exit 1
  fi

  if [ -n "${KEY_NAME}" ]
  then
    echo "${KEY_PASSWORD}"
    echo "${KEY_MNEMONIC}"
    {
      echo "${KEY_PASSWORD}"
      echo "${KEY_MNEMONIC}"
      echo
    } | konstellationcli keys add "${KEY_NAME}" --interactive
  fi

  konstellation init "${MONIKER}" --chain-id "${CHAIN_ID}"
  # shellcheck disable=SC2164
  cd /root/.konstellation/config
  rm -f config.toml genesis.json
  cp /root/.konstellation/config.toml /root/.konstellation/config/config.toml
  cp /root/.konstellation/genesis.json /root/.konstellation/config/genesis.json
  konstellation unsafe-reset-all
  #    wget https://raw.githubusercontent.com/hashgard/testnets/master/sif/${CHAIN_ID}/config/config.toml
  #    wget https://raw.githubusercontent.com/hashgard/testnets/master/sif/${CHAIN_ID}/config/genesis.json
  sed -i "s|moniker.*|moniker = \"${MONIKER}\"|g" config.toml
  sed -i "s|seeds.*|seeds = \"${SEEDS}\"|g" config.toml

  ip=$(ifconfig eth0 | head -2 | awk '{print $2}' | sed -n 2p)
  sed -i "s|external_address.*|external_address = \"${ip}\"|g" config.toml

  if [ -n "$SEED" ]
  then
    sed -i "s|seed_mode.*|seed_mode = true|g" config.toml
  fi
}

# ------------------------------------------------------------------------------
#
# Initial Konstellation block chain private single node
#
#   1. Generate account
#   2. Init server's local configuration
#   3. Assign coins to the account
#   4. Add account to genesis
#
# ------------------------------------------------------------------------------
function init_private_single() {
  # Create private key for first delegation
  echo "${KEY_PASSWORD}"
  echo "${KEY_MNEMONIC}"
  {
    echo "${KEY_PASSWORD}"
    echo "${KEY_MNEMONIC}"
    echo
  } | konstellationcli keys add "${KEY_NAME}" --interactive

  # Init konstellation chain
  konstellation init "${MONIKER}" --chain-id "${CHAIN_ID}"
  konstellation add-genesis-account "${KEY_NAME}" "${COIN_GENESIS}"
  echo "${KEY_PASSWORD}" | konstellation gentx --name "${KEY_NAME}" --amount "${COIN_DELEGATE}"
  konstellation collect-gentxs
}

# ------------------------------------------------------------------------------
#
# Initial Konstellation block chain private multiple nodes
#
#   Folder .konstellation and .konstellationcli had been created by 'konstellation testnet'
#   and mount to /root/.konstellation and /root/.konstellationcli. So nothing to do.
#
# ------------------------------------------------------------------------------
function init_private_testnet() {
  echo "nothing to do !" >/dev/null
}

# ------------------------------------------------------------------------------
#
# Initial Konstellation block chain private multiple nodes
#
#   Folder .konstellation and .konstellationcli had been created by 'konstellation testnet'
#   and mount to /root/.konstellation and /root/.konstellationcli. So nothing to do.
#
# ------------------------------------------------------------------------------
function konstellation_start() {
  konstellation start

  # Hold the container for debugging
  while true; do
    sleep 1
  done
}

# ------------------------------------------------------------------------------
#
# Error prompt and exit abnormally
#
# ------------------------------------------------------------------------------
function error() {
  echo "" >&2
  echo "Error: " >&2
  echo "    $1" >&2
  echo "" >&2
  exit 1
}

# ------------------------------------------------------------------------------
#
# Chain id must be set to environment
#
# ------------------------------------------------------------------------------
if [[ -z ${CHAIN_ID} ]]; then
  error "Environment CHAIN_ID must be set !"
fi

# ------------------------------------------------------------------------------
#
# For container restart
#
# ------------------------------------------------------------------------------
if [[ -e ${INITIALIZED_FLAG} ]]; then
  konstellation_start
fi

# ------------------------------------------------------------------------------
#
# Print env variables
#
# ------------------------------------------------------------------------------
echo "Chain id" "${CHAIN_ID}"
echo "Moniker" "${MONIKER}"

# ------------------------------------------------------------------------------
#
# Config client global settings
#
# ------------------------------------------------------------------------------
konstellationcli config chain-id "${CHAIN_ID}"
konstellationcli config trust-node true
konstellationcli config output json
konstellationcli config indent true

# ------------------------------------------------------------------------------
#
# Node type:
#
#     FULL_NODE       - Run Konstellation block chain full node which can connect to
#                       Konstellation testnet and change to validator.
#     PRIVATE_SINGLE  - Run Konstellation block chain with single node.(Default)
#     PRIVATE_TESTNET - Run Konstellation block chain with multiple nodes created by
#                       command 'konstellation testnet'
#
# ------------------------------------------------------------------------------
if [[ -z ${NODE_TYPE} ]]; then
  NODE_TYPE="PRIVATE_SINGLE"
fi
case "${NODE_TYPE}" in
"FULL_NODE")
  init_full_node
  ;;
"PRIVATE_SINGLE")
  init_private_single$()
  ;;
"PRIVATE_TESTNET")
  init_private_testnet
  ;;
*)
  error "Environment NODE_TYPE must be one of FULL_NODE/PRIVATE_SINGLE/PRIVATE_TESTNET !"
  ;;
esac
echo "Node type "${NODE_TYPE}

# Mark initial successfully
touch ${INITIALIZED_FLAG}

# Start service
konstellation_start
