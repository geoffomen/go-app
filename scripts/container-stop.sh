#!/bin/sh


set -e


echo "stopping container..."
TARGET_DIR=/tmp/test_$(date +'%Y-%m-%d')
echo "working dir: $(pwd)"
if [ -d ${TARGET_DIR} ]; then 
    cd ${TARGET_DIR}
    CMD=$(echo "docker compose -f deployments/container/compose.yml down -v")
    echo ${CMD}
    eval ${CMD}
    
    CMD=$(echo "rm -rf * && cd .. && rmdir ${TARGET_DIR}")
    echo ${CMD}
    eval ${CMD}
fi 

echo "done"