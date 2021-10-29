#!/bin/sh

set -e

echo "working directory $(pwd)"
echo "building ${1}"

CC="go"
SRC="cmd/${1}/main.go"
OUT="build/temp/${1}.o"

COMMIT_VAR="commitId"
BRANCH_VAR="branchName"
BUILDTIME_VAR="buildTime" 

COMMIT_ID=$(git rev-parse HEAD)
BRANCH_NAME=$(git rev-parse --abbrev-ref HEAD)
BUILD_DATE=$(date +'%Y-%m-%d_%T')

CMD=`echo ${CC} build \
	-mod=vendor \
	-ldflags \"-X main.${BRANCH_VAR}=${BRANCH_NAME} -X main.${COMMIT_VAR}=${COMMIT_ID} -X main.${BUILDTIME_VAR}=${BUILD_DATE}\" \
	-o ${OUT} \
	-v \
	${SRC}`

echo "going to execute command: ${CMD}"
eval ${CMD}

echo "done"