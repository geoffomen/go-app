#!/bin/sh


set -e


make container args_appname=example

echo "容器创建成功"
echo "使用帮助: "
echo '先cd到一个应用的工作目录，然后执行命令docker run --rm -v $(pwd):/workspace example来初化始该工作目录;'
echo '初始化完成，在工作目录中执行sh start.sh即可启动应用，执行sh stop.sh则可停止应用。'