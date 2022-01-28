#!/bin/bash
# go install proto deps
go install google.golang.org/protobuf/cmd/protoc-gen-go
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
# generate code
protoc model/*.proto \
  --go_out=. \
  --go_opt=paths=source_relative \
  --proto_path=.
