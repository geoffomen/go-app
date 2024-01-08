#!/bin/sh

GO=`which go`
OLD_MODULE=ibingli.com
NEW_MODULE=ibingli.com

# edit module name
${GO} mod edit -module ${NEW_MODULE}

# rename all imported module
find . \
	-type f \
	-name '*.go' \
	-exec sed -i -e "s,${OLD_MODULE},${NEW_MODULE},g" {} +
