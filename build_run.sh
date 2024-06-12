#!/bin/bash

PWD=`pwd`
MODULE_FILE=$PWD/go.mod
DIR_NAME=$(basename ${PWD})
MODULE_NAME=${DIR_NAME,,}

if [ ! -f $MODULE_FILE ]; then
    go mod init $MODULE_NAME
fi

go work use .

BUILD_NAME=build

if [ ! -d ${BUILD_NAME} ]; then
    mkdir ${BUILD_NAME}
fi

go build -o ${BUILD_NAME}/

exec ${PWD}/${BUILD_NAME}/${MODULE_NAME} $*