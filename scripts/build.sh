#!/bin/sh


set -e


APP=$1
CC=$(which go)
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

CMD=`echo ${CC} build \
	-mod=vendor \
	-ldflags \"-X main.${BRANCH_VAR}=${BRANCH_NAME} -X main.${COMMIT_VAR}=${COMMIT_ID} -X main.${BUILDTIME_VAR}=${BUILD_DATE}\" \
	-o ${OUT} \
	-v \
	${SRC}`

echo ${CMD}
eval ${CMD}

echo "done"