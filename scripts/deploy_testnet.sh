echo 'Copying node1 home directory'
scp -r -i ~/Documents/.ssh/id_rsa.pub ./testnet/net1 root@95.216.212.119:/root/node
echo 'Copying node2 home directory'
scp -r -i ~/Documents/.ssh/id_rsa.pub ./testnet/net2 root@95.216.215.127:/root/node
echo 'Copying Node3 home directory'
scp -r -i ~/Documents/.ssh/id_rsa.pub ./testnet/net3 root@95.216.210.171:/root/node
echo 'Copying Node4 home directory'
scp -r -i ~/Documents/.ssh/id_rsa.pub ./testnet/net4 root@95.216.207.219:/root/node

#
#if [[ -f "./config/nodes.json" ]]; then
#  nets=$(jq '.[] | .name' ./config/nodes.json)
#  ips=$(jq '.[] | .ip' ./config/nodes.json)
#
##  index=0
##
##  for i in "${!AR[@]}"; do
##    printf '${AR[%s]}=%s\n' "$i" "${AR[i]}"
##  done
#  for net in $nets
#  do
#    echo $net
#    echo $index
#
#    echo "${ips[$index]}"
#    ((index++))
#  done
#
#  indexx=0
#  for ip in $ips
#  do
#    echo ip
#    echo $indexx
#
#    echo "${ips[$indexx]}"
#    ((indexx++))
#  done
#
##  for i in "${!nets[@]}"
##  do
##    net="${nets[$i]}"
##    ip=${ips[$i]}
##
##echo $i
##    echo $net
##    echo $ip
##    echo "Copying ${net} home directory"
##  done
#fi
#
