#!/bin/sh


set -e


MYSQL_ROOT_PASSWORD=${1}
DB_DUMP_FILE=${2}


echo "importing data into mysql"

CMD=$(echo "podman exec -i mysql mysql -uroot -p${MYSQL_ROOT_PASSWORD} < ${DB_DUMP_FILE}")

echo ${CMD}

eval ${CMD}

echo "done"