set -x
GOARCH=wasm GOOS=js go build -o lib.wasm main.go
#cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
cp "/usr/share/doc/go/misc/wasm/wasm_exec.js" .

