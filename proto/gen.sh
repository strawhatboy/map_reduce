#!/bin/sh 
cd "$(dirname "$0")"

protoc --go_out=. --go-grpc_out=. \
    --go_opt=paths=source_relative \
    --go-grpc_opt=paths=source_relative \
    map_reduce.proto

cd -
