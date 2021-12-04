const fs = require('fs');

async function Run() {
    let buffer = await fs.readFileSync('./add.wasm');
    let result = await WebAssembly.instantiate(buffer);
    let add = result.instance.exports.add
    console.log(add(5, 2))
}

Run()


