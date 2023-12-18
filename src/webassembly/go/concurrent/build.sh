set -x

export GOARCH=wasm
export GOOS=js
go build -o lib.wasm main.go
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
cp "$(go env GOROOT)/misc/wasm/wasm_exec_node.js" .
node wasm_exec_node.js lib.wasm
