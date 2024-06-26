#!/bin/bash

PROTO_DIR=./proto
OUT_DIR=./gen

mkdir -p $OUT_DIR

protoc -I proto $PROTO_DIR/*.proto --go_out=$OUT_DIR --go_opt=paths=source_relative --go-grpc_out=$OUT_DIR --go-grpc_opt=paths=source_relative