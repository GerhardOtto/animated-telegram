#!/bin/bash
set -e

protoc \
  --go_out=client \
  --go_opt=module=techops.technical.challenge.com \
  --go-grpc_out=client \
  --go-grpc_opt=module=techops.technical.challenge.com \
  exampleservice.proto
