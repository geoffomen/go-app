#!/bin/sh


set -e


IMAGE_NAME=${1}
DOCKER_FILE=build/container/${IMAGE_NAME}/Dockerfile


echo "building image ${IMAGE_NAME}"

if [ $(docker images -a | grep ${IMAGE_NAME} | wc -l) -gt 0 ]
then
	echo "already exist. delete it"
	CMD=$(echo "docker rmi ${IMAGE_NAME}")
	echo ${CMD}
	eval ${CMD}
fi 

CMD=$(echo "docker build --build-arg APP_NAME=${IMAGE_NAME} -t ${IMAGE_NAME} -f ${DOCKER_FILE} .")
echo ${CMD}
eval ${CMD}

echo "done"