set -x
export WASI_SDK_PATH=/tmp/wasi-sdk-14.0
CC="${WASI_SDK_PATH}/bin/clang --sysroot=${WASI_SDK_PATH}/share/wasi-sysroot"
$CC -O2 --sysroot=/tmp/wasi-libc -nostartfiles  -Wl,--no-entry -Wl,--export=mycos -target wasm32-unknown-wasi -o foo.wasm foo.c

node foo.js
