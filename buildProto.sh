#!/usr/bin/env bash

if [ ! -f pingpong ]; then
        mkdir pingpong
        echo "pingpong dir create"
fi
if [ ! -f score ]; then
        mkdir score
        echo "score dir create"
fi

protoc -I protobuf --go_out=plugins=grpc:pingpong/ ./protobuf/pingpong.proto

protoc -I protobuf --go_out=plugins=grpc:score/ ./protobuf/score.proto
