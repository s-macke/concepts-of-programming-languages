const fs = require('fs');
const buf = fs.readFileSync('./add.wasm');

WebAssembly.instantiate(buf)
.then(result => {
    console.log(result.instance.exports.add(5, 2))
});

