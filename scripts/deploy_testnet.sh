if [[ -f "./config/nodes.json" ]]; then
 jq -r '
    to_entries |
    .[].value |
    @sh "scp -r -i ~/Documents/.ssh/id_rsa.pub ./testnet/\(.name) root@\(.ip):/root/node"
  ' ./config/nodes.json
fi
