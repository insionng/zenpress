#!/bin/bash
export GOPATH=`pwd`
mkdir /dev/shm/archives
go build -ldflags "-s" app.go
sudo ./app
