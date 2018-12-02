#!/bin/bash
rm -rf ./bin/*
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/example ./