#!/bin/sh


set -e


CONTAINER_NAME=${1}


echo "cleaning container ${CONTAINER_NAME}"

if [ $(podman ps -a | grep ${CONTAINER_NAME} | wc -l) -gt 0 ]
then
	CMD=$(echo "podman stop ${CONTAINER_NAME} && podman rm ${CONTAINER_NAME}")

	echo ${CMD}

	eval ${CMD}
fi 

CMD=$(echo "podman rmi ${CONTAINER_NAME}")

echo ${CMD}

eval ${CMD}

echo "done"