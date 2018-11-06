#!/bin/bash

workdir=`dirname $0`
cd ${workdir}
workdir=`pwd`

export GOPATH=${GOPATH}:${workdir}

go run ${workdir}/src/worm.go
