#!/bin/sh


set -e


TARGET=${1}


echo "cleaning ${TARGET}"

CMD=$(echo "rm build/temp/${TARGET}.o")

echo ${CMD}

eval ${CMD}

echo "done"