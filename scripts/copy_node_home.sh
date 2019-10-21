echo 'Copying node1 home directory'
scp -r -i ~/Documents/.ssh/id_rsa.pub ./testnet/node0 root@95.216.212.119:/root/node
echo 'Copying node2 home directory'
scp -r -i ~/Documents/.ssh/id_rsa.pub ./testnet/node1 root@95.216.215.127:/root/node
echo 'Copying Node3 home directory'
scp -r -i ~/Documents/.ssh/id_rsa.pub ./testnet/node2 root@95.216.210.171:/root/node
echo 'Copying Node4 home directory'
scp -r -i ~/Documents/.ssh/id_rsa.pub ./testnet/node3 root@95.216.207.219:/root/node
