
help:
	@echo "Makefile targets:"
	@echo "  install - get required dependencies"
	@echo "    proto - generate code from proto files"
	@echo "   server - start server"
	@echo "   client - start client"

install:
	go get google.golang.org/grpc

proto:
	protoc  -I ./proto ./proto/greeter.proto --go_out=plugins=grpc:proto

server:
	go run server/server.go

client:
	go run client/client.go

.PHONY: help install proto server client