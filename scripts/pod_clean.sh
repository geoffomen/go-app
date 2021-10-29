#!/bin/sh


set -e 


POD_NAME=${1}


echo "cleaning pod ${POD_NAME}"

if [ $(podman pod list | grep ${POD_NAME} | wc -l) -gt 0 ]
then
	CMD=$(echo "podman pod stop ${POD_NAME} && podman pod rm ${POD_NAME}")
	echo ${CMD}
	eval ${CMD}
fi 

echo "done"