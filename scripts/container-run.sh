#!/bin/sh


set -e


POD_NAME=${1}
CONTAINER_NAME=${2}


echo "starting container ${CONTAINER_NAME}"

if [ $(podman ps -a -f pod=${POD_NAME} | grep ${CONTAINER_NAME} | wc -l) -gt 0 ]
then
	echo "stopping preexist container..."
	echo "podman stop ${CONTAINER_NAME} && podman rm ${CONTAINER_NAME}"
	eval ${CMD}
fi 

CMD=$(echo "podman run -d \
	--pod ${POD_NAME} \
	--name ${CONTAINER_NAME} \
	localhost/${CONTAINER_NAME}:latest")

echo ${CMD}

eval ${CMD}

echo "done"