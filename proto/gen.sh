#!/bin/sh 

#
# This script need protoc (protobuf-compiler) installed
#   and 2 plugins:
#       github.com/golang/protobuf/protoc-gen-go
#       google.golang.org/grpc/cmd/protoc-gen-go-grpc
#
cd "$(dirname "$0")"

protoc --go_out=. --go-grpc_out=. \
    --go_opt=paths=source_relative \
    --go-grpc_opt=paths=source_relative \
    map_reduce.proto

cd -
