#!/bin/sh


set -e


CONTAINER_NAME=${1}


echo "cleaning container ${CONTAINER_NAME}"

CMD=$(echo "docker rmi ${CONTAINER_NAME}")
echo ${CMD}
eval ${CMD}

echo "done"