set -x
clang -O2 --sysroot=/tmp/wasi-libc -nostartfiles  -Wl,--no-entry -Wl,--export-all -target wasm32-unknown-wasi -o foo.wasm foo.c
#clang -O2 -nostdlib -Wl,--no-entry -Wl,--export-all --target=wasm32 -o foo.wasm foo.c
node foo.js
