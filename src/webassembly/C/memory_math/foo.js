const fs = require('fs');

async function Run() {
    let buffer = await fs.readFileSync('./foo.wasm');
    let result = await WebAssembly.instantiate(buffer);
    let foo = result.instance.exports.foo
    foo(42)
    let memory = new Uint8Array(result.instance.exports.memory.buffer);
    console.log(memory[2048]);
}

Run()
