set +x

export GOARCH=wasm
export GOOS=js
#go build -o lib.wasm *.go

~/goinstall/go/bin/go build -o lib.wasm *.go

node wasm_exec.js lib.wasm

