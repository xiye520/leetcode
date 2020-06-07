#!/bin/bash

go fmt ./...
go vet ./...
#find ./ -name "*.go" | grep -v "/topics/" | xargs gofmt -w
#find ./ -name "*.go" | xargs gofmt -w
#find ./ | xargs dos2unix
