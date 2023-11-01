#!/bin/sh


# 容器启动时，会检查容器内的`/workspace`目录，发现该目录为空时，自动将运行需要的文件拷贝至该目录。
# 初始化脚本运行完后，即可在`/workspace`启动整个项目。
# 初始化时需要挂载一个外部的工作目录，影射到容器中的`/workspace`中。

export APP_NAME=$(ls -A /opt)


echo "初始化脚本正在运行..."
echo "需要挂载一个工作目录到容器中的workspace目录，如果没有这样做，请重新运行容器，eg: "
echo "docker run --rm -v $(pwd):/workspace ${APP_NAME}"

if [ ! -d /workspace ]; then
    mkdir -p /workspace
fi

if [ ! -d /workspace/configs ]; then
    mkdir -p /workspace/configs/
fi

if [ ! -f /workspace/configs/container-${APP_NAME}.yml ]; then
	cp -r /opt/${APP_NAME}/configs/container-${APP_NAME}.yml /workspace/configs/container-${APP_NAME}.yml
fi

if [ ! -d /workspace/website ]; then
    cp -r /opt/${APP_NAME}/website /workspace/website
fi

if [ ! -d /workspace/deployments ]; then
    cp -r /opt/${APP_NAME}/deployments /workspace/deployments
fi

if [ ! -f /workspace/deployments/container/compose.yml ]; then
	cp /workspace/deployments/container/docker-compose-${APP_NAME}.yml /workspace/deployments/container/compose.yml
fi

if [ ! -f /workspace/start.sh ]; then
	cp -r /opt/${APP_NAME}/deployments/container/start.sh /workspace/start.sh
fi

if [ ! -f /workspace/stop.sh ]; then
	cp -r /opt/${APP_NAME}/deployments/container/stop.sh /workspace/stop.sh
fi


echo "初始化完成，在工作目录中执行'sh start.sh'即可启动项目，执行'sh stop.sh'则可停止项目。"