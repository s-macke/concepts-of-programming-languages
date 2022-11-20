
help:
	@echo "Makefile targets:"
	@echo "  install - get required dependencies"
	@echo "    proto - generate code from proto files"
	@echo "   server - start server"
	@echo "   client - start client"

install:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go get google.golang.org/grpc

proto:
	protoc  -I ./proto ./proto/greeter.proto --go_out=./proto --go-grpc_out=./proto

server:
	go run server/server.go

client:
	go run client/client.go

.PHONY: help install proto server client