#!/bin/sh

# if mongorpc dir is not exist, create it
if [ ! -d "mongorpc" ]; then
    mkdir mongorpc
fi

# if mongorpc/v1 dir is not exist, create it
if [ ! -d "mongorpc/v1" ]; then
    mkdir mongorpc/v1
fi

protoc --proto_path=proto --go_out=mongorpc/v1 --go_opt=paths=source_relative \
    --go-grpc_out=mongorpc/v1 --go-grpc_opt=paths=source_relative \
    proto/*.proto