#!/bin/bash

echo "restart script is started"

function restart_node(){
	screen -XS cosmovisor quit
	sleep 5
	screen -dmSL cosmovisor cosmovisor start
	ulimit -n 1000000
	process_id=`/bin/ps -fu $USER| grep "cosmovisor" | grep -v "grep" | awk '{print $2}'`
	for i in $process_id
	do
		prlimit --nofile=1000000 --pid=$i
	done
	echo "restarted node"
}

export DAEMON_NAME=knstld
export DAEMON_HOME=$HOME/.knstld
export DAEMON_RESTART_AFTER_UPGRADE=true
latest_block_height=`cosmovisor status | jq -r '.SyncInfo.latest_block_height'`
sleep 10m
latest_block_height2=`cosmovisor status | jq -r '.SyncInfo.latest_block_height'`
if [ $latest_block_height -eq $latest_block_height2 ]
	then
			restart_node
fi

