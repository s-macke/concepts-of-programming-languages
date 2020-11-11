
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
	go get -u github.com/golang/protobuf/protoc-gen-go
	@echo "Make sure that protoc is installed for your platform"
	@echo "    macOS: brew install protoc-gen-go protoc-gen-go-grpc"
	@echo "  Windows: https://github.com/protocolbuffers/protobuf/releases/latest"

proto:
	protoc --go_out=./node --go-grpc_out=./node ./node/node.proto

run2:
	go build
	./raft -id 0 -count 2 &
	./raft -id 1 -count 2 &
	@echo "Use 'killall raft' to terminate instances"

run3:
	go build
	./raft -id 0 -count 3 &
	./raft -id 1 -count 3 &
	./raft -id 2 -count 3 &
	@echo "Use 'killall raft' to terminate instances"

run4:
	go build
	./raft -id 0 -count 4 &
	./raft -id 1 -count 4 &
	./raft -id 2 -count 4 &
	./raft -id 3 -count 4 &
	@echo "Use 'killall raft' to terminate instances"

run5:
	go build
	./raft -id 0 -count 5 &
	./raft -id 1 -count 5 &
	./raft -id 2 -count 5 &
	./raft -id 3 -count 5 &
	./raft -id 4 -count 5 &
	@echo "Use 'killall raft' to terminate instances"
