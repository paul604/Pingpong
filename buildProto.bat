@echo off
cls

protoc -I . --go_out=plugins=grpc:. ./protobuf/pingpong.proto

protoc -I . --go_out=plugins=grpc:. ./protobuf/score.proto
