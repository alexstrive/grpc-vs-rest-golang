#!/bin/zsh

(cd grpc_server; go run main.go) &
(cd rest_server; go run main.go) &
wait
