#!/bin/bash
rm -rf ./bin/*
cp ./config.toml ./bin/
mkdir -p ./bin/assets
cp -a ./assets/* ./bin/assets/
mkdir -p ./bin/templates
cp -a ./templates/* ./bin/templates/
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ./bin/example.exe ./