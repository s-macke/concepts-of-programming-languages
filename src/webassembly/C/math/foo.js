const fs = require('fs');

async function Run() {
    let buffer = await fs.readFileSync('./foo.wasm');
    let result = await WebAssembly.instantiate(buffer);
    let mycos = result.instance.exports.mycos
    console.log(mycos(0.))
    console.log(mycos(3.141592/2.))
}

Run()
