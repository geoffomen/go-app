#!/bin/sh


set -e


APP=$1
SRC="cmd/${APP}/main.go"
OUT="build/temp/${APP}.o"

COMMIT_VAR="commitId"
BRANCH_VAR="branchName"
BUILDTIME_VAR="buildTime" 

COMMIT_ID=$(git rev-parse HEAD)
BRANCH_NAME=$(git rev-parse --abbrev-ref HEAD)
BUILD_DATE=$(date +'%Y-%m-%d_%T')


echo "working directory $(pwd)"
echo "building ${APP}"

CMD=`echo docker run --rm \
	-v "$PWD":/usr/src/${APP}:Z \
	-w /usr/src/${APP} \
	-e GOOS=linux \
	-e GOARCH=amd64 \
	golang:1.20.1 \
	go build \
	-ldflags \"-X main.${BRANCH_VAR}=${BRANCH_NAME} -X main.${COMMIT_VAR}=${COMMIT_ID} -X main.${BUILDTIME_VAR}=${BUILD_DATE}\" \
	-mod=vendor \
	-o ${OUT} \
	-v ${SRC}`

echo ${CMD}
eval ${CMD}

echo "done"