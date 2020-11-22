set -x
clang -O2 --sysroot=/tmp/wasi-libc -nostartfiles  -Wl,--no-entry -Wl,--export-all -target wasm32-unknown-wasi -o foo.wasm foo.c
node foo.js

