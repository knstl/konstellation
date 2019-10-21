#!/usr/bin/env bash

source $(pwd)/docker/.env

NODE_NAME=$(hostname)

if [[ ! -d ${NODE_ROOT} ]]; then
  echo "Node's config DOSE NOT exist !"
  echo "" >&2
  exit 1
fi

containers=$(docker container ls | awk '{print $1}' | sed -n 2p)
if [ ! -z "$containers" ]; then
  echo -n "Remove $containers ... "
  docker rm -f ${NODE_NAME} >/dev/null
fi

echo -n "Create ${NODE_NAME} ... "
docker run -d \
  --name ${NODE_NAME} \
  -e CHAIN_ID=${CHAIN_ID} \
  -e MONIKER=NODE_NAME \
  -e NODE_TYPE=PRIVATE_TESTNET \
  -v ${NODE_ROOT}/konstellation:/root/.konstellation \
  -v ${NODE_ROOT}/konstellationcli:/root/.konstellationcli \
  ${IMAGE_OWNER}/konstellation:${CHAIN_ID}
echo "Done !"
