const fs = require('fs');
const buf = fs.readFileSync('./foo.wasm');

WebAssembly
    .instantiate(buf)
    .then(result => {
        console.log(result.instance.exports.foo(42));
        var memory = new Uint8Array(result.instance.exports.memory.buffer);
        console.log(memory[2048]);
    });