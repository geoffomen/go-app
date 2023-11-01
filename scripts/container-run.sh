#!/bin/sh


set -e


echo "starting container..."
APP_NAME=${1}


TARGET_DIR=/tmp/test_$(date +'%Y-%m-%d')

if [ ! -d ${TARGET_DIR} ]; then 
    mkdir ${TARGET_DIR}
fi 

cd ${TARGET_DIR}
echo "working dir: $(pwd)"

# 初始化工作目录
CMD=$(echo "docker run --rm -v $(pwd):/workspace ${APP_NAME}")
echo ${CMD}
eval ${CMD}

# 运行
CMD=$(echo "docker compose -f deployments/container/compose.yml up -d")
echo ${CMD}
eval ${CMD}

echo "done"