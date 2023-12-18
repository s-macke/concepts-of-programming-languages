set -x
clang -O2 -nostdlib -Wl,--no-entry -Wl,--export-all --target=wasm32 -o foo.wasm foo.c
node foo.js
