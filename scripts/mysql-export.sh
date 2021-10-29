#!/bin/sh


set -e


MYSQL_ROOT_PASSWORD=${1}
DB_NAME=${2}
DB_DUMP_FILE=${3}


echo "exporting data from mysql"

CMD=$(echo "podman exec mysql mysqldump --databases ${DB_NAME} -uroot -p${MYSQL_ROOT_PASSWORD} > ${DB_DUMP_FILE}")

echo ${CMD}

eval ${CMD}

echo "done"