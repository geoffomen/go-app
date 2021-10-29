#!/bin/sh


set -e


POD_NAME=${1}


echo "creating pod ${POD_NAME}"

if [ $(podman pod list | grep ${POD_NAME} | wc -l) -gt 0 ]
then
	echo "already exist."
	echo "done"
	exit 0
fi 

CMD=$(echo "podman pod create \
	--publish 8000:8000 \
	--publish 3306:3306 \
	--name ${POD_NAME}")
echo ${CMD}
eval ${CMD}

echo "done"