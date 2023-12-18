set -x
wat2wasm add.wat
#wasm-dis add.wasm
#wasm2wat add.wasm
node add.js
