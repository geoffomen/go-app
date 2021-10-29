#!/bin/sh


set -e


CONTAIN_NAME=${1}
DOCKER_FILE=${2}


echo "building container ${CONTAIN_NAME}"

if [ $(podman images -a | grep localhost/${CONTAIN_NAME} | wc -l) -gt 0 ]
then
	echo "already exist."
	exit 0
fi 

CMD=$(echo "podman build -t ${CONTAIN_NAME} -f ${DOCKER_FILE} .")

echo ${CMD}

eval ${CMD}

echo "done"