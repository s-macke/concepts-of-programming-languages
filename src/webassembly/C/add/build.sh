clang -O2 -nostdlib -Wl,--no-entry -Wl,--export-all --target=wasm32 -o add.wasm add.c
wasm-dis add.wasm
node add.js
