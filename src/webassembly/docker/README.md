# WebAssembly development environment for WAST, C, and GO

Start with
```
docker build -t wasm_env .
```

Run with
```
docker run -it --name wasm -v $PWD:/root/wasm wasm_env:latest
```
where $PWD is the path to your `webassembly` folder

