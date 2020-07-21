#!/usr/bin/env bash

# Meaning of options:
#    -I ../proto                             -- tell protoc where to find google/api/annotations.proto imported in user_manager.proto
#    --go_out=.                              -- generate go code (protobuf messages) in current directory
#    --go_opt=paths=source_relative \        -- generate go code next to the source files (otherwise creates files in go_package path)
#    --go-grpc_out=. \                       -- generate go-grpc code (server interface) in current directory
#    --go-grpc_opt=paths=source_relative \   -- generate go-grpc code next to the source files (otherwise creates files in go_package path)
#    --go-http_out=. \                       -- generate go-http code (http.Handler)
#    --go-http_opt=paths=source_relative \   -- generate go-http code next to the source files (otherwise creates files in go_package path)

protoc -I ../proto -I ./ \
    --go_out=. \
    --go_opt=paths=source_relative \
    --go-grpc_out=. \
    --go-grpc_opt=paths=source_relative \
    --go-http_out=. \
    --go-http_opt=paths=source_relative \
    user_manager.proto
