set -x

#go tool dist list

export GOARCH=wasm
export GOOS=wasip1
go build -o hello.wasm main.go

#wasmtime hello.wasm

wasm-dis hello.wasm | grep 'import \|export '


