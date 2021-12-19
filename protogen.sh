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


# if mongorpc-swift/proto dir is not exist, create it
if [ ! -d "mongorpc-swift/proto" ]; then
    mkdir mongorpc-swift/proto
fi

# protogen swift
protoc -I=proto --proto_path=proto --swift_out=mongorpc-swift/proto \
    --grpc-swift_out=Client=true,Server=false:mongorpc-swift/proto \
    proto/*.proto