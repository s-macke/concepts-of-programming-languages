
help:
	@echo "Makefile targets:"
	@echo "  install - get required dependencies"
	@echo "    proto - generate code from proto files"
	@echo "     run2 - build and start two raft nodes"
	@echo "     run3 - build and start three raft nodes"
	@echo "     run4 - build and start four raft nodes"
	@echo "     run5 - build and start five raft nodes"

install:
	go get -u google.golang.org/grpc
	go get -u google.golang.org/protobuf
	@echo "Make sure that protoc is installed for your platform"
	@echo "    macOS: brew install protoc-gen-go protoc-gen-go-grpc"
	@echo "  Windows: https://github.com/protocolbuffers/protobuf/releases/latest"

proto:
	protoc  -I node/ ./node/node.proto --go_out=plugins=grpc:node

run2:
	go build
	./raft -id 0 -count 2 & echo $$!
	./raft -id 1 -count 2 & echo $$!
	@echo "Use 'killall raft' to terminate instances"

run3:
	go build
	./raft -id 0 -count 3 & echo $$!
	./raft -id 1 -count 3 & echo $$!
	./raft -id 2 -count 3 & echo $$!
	@echo "Use 'killall raft' to terminate instances"

run4:
	go build
	./raft -id 0 -count 4 & echo $$!
	./raft -id 1 -count 4 & echo $$!
	./raft -id 2 -count 4 & echo $$!
	./raft -id 3 -count 4 & echo $$!
	@echo "Use 'killall raft' to terminate instances"

run5:
	go build
	./raft -id 0 -count 5 & echo $$!
	./raft -id 1 -count 5 & echo $$!
	./raft -id 2 -count 5 & echo $$!
	./raft -id 3 -count 5 & echo $$!
	./raft -id 4 -count 5 & echo $$!
	@echo "Use 'killall raft' to terminate instances"
