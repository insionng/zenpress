#!/bin/bash
export GOPATH=`pwd`
go build -ldflags "-s" app.go
sudo ./app
