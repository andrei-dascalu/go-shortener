#!/bin/bash -e


protoc \
    --go_out=src/grpcserve --go_opt=paths=import \
    --go-grpc_out=src/grpcserve --go-grpc_opt=paths=import \
    proto/service.proto
