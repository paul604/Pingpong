#!/usr/bin/env bash

protoc -I . --go_out=plugins=grpc:. ./protobuf/pingpong.proto

protoc -I . --go_out=plugins=grpc:. ./protobuf/score.proto
