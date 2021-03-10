#!/bin/sh
#set -o errexit -o nounset -o pipefail

MONIKER=${MONIKER:-dm0}
echo "Moniker" "${MONIKER}"

sed -i "s/^moniker.*/moniker = \"$MONIKER\"/" /root/.knstld/config/config.toml
