#!/usr/bin/env bash
GOPATH=`pwd`
export GOPATH=$GOPATH
SCRIPT=$GOPATH"/Test_"$1".go"
go run $SCRIPT