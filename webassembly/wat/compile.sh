set -x
wat2wasm add.wat
wasm-dis add.wasm
node add.js
