#!/bin/sh


set -e


POD_NAME=${1}
MYSQL_ROOT_PASSWORD=${2}
DATA_DIR=${3}


echo "starting mysql"

if [ $(podman pod list | grep ${POD_NAME} | wc -l) -gt 0 ]
then
	if [ $(podman ps -a -f pod=${POD_NAME} | grep mysql | wc -l) -gt 0 ]
	then
		echo "already run."
		exit 0
	fi 

	CMD=$(echo podman run -d \
		--pod ${POD_NAME} \
		--volume ${DATA_DIR}:/var/lib/mysql:Z \
		--env MYSQL_ROOT_PASSWORD=${MYSQL_ROOT_PASSWORD} \
		--name mysql \
		mysql:5.7)
	echo ${CMD}
	eval ${CMD}
else 
	if [ $(podman ps -a | grep mysql | wc -l) -gt 0 ]
	then
		echo "already run."
		exit 0
	fi 

	CMD=$(echo podman run -d \
		-p 3306:3306 \
		--volume ${DATA_DIR}:/var/lib/mysql:Z \
		--env MYSQL_ROOT_PASSWORD=${MYSQL_ROOT_PASSWORD} \
		--name mysql \
		mysql:5.7)
	echo ${CMD}
	eval ${CMD}
fi 

echo "waiting database init"
CMD=$(echo "podman exec mysql mysql --user=root --password=${MYSQL_ROOT_PASSWORD} -e 'SELECT 1' >/dev/null 2>&1")
while ! $(eval ${CMD}); do
	echo "Waiting for database ..."
	sleep 3
done

echo "done"