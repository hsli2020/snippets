#!/usr/bin/env bash

if [ ! -f build.sh ]; then
echo 'build.sh must be run within its container folder' 1>&2
exit 1
fi

CURDIR=`pwd`
OLDGOPATH="$GOPATH"

export GOPATH="$(cygpath -aw $CURDIR)"

#go build src/main.go
go install main

export GOPATH="$OLDGOPATH"

echo 'finished'
