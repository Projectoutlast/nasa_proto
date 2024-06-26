#!/bin/bash

PROTO_DIR=./proto
OUT_DIR=./gen

mkdir -p $OUT_DIR

do protoc -I proto proto/*.proto --go_out=./gen --go_opt=paths=source_relative --go-grpc_out=./gen --go-grpc_opt=paths=source_relative
done