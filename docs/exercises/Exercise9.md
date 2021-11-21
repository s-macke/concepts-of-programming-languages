# Exercise 8 - Webassembly

## Exercise 8.1 - Hello world Go project in WebAssembly

* Install npm/node on your system
  - https://nodejs.org/en/download/
  - Alternative: Use the provided docker file in the `webassembly/docker` folder

---
* Write a Hello World program in Go
* Compile with `GOARCH=wasm GOOS=js go build -o lib.wasm main.go` The result will be a file called `lib.wasm`
* In case you run Windows command shell you have to execute three commands
```
set GOARCH=wasm 
set GOOS=js 
go build -o lib.wasm main.go`
```
* Copy the wasm javascript runtime from the GOROOT directory. You will find the GOROOT directory via  `go env GOROOT` and copy the wasm_exec file: `cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .`. Alternative path: `/usr/share/doc/go/misc/wasm/wasm_exec.js`-
* Execute via `node wasm_exec.js lib.wasm`
---
* An HTTP file web server is provided in the folder `wp/fileserver`. Compile and start it with root path `webassembly`.    
* Copy the index.html file from `webassembly/hello/index.html`

## Exercise 8.2 - Execute a Go function inside Go

* The folder `mandelbrot/main.go` contains a program, which calculates the so called Mandelbrot set.
* Compile it and ensure that it produces an image `mandelbrot.png` as result
* Add a function, which returns for the default parameters the PNG as Base64 encoded string.
```Go
func GetImageAsBase64() string {

}
```
  - Use `bytes.Buffer` as io.Writer object and encode via the `base64.StdEncoding.EncodeToString` function. 

* Rewrite the function header, so that it can be executed via a WebAssembly call and alter the main function accordingly.

* Create a website with following content and ensure, that it shows a circle.

```HTML
<html>
<head>
    <meta charset="utf-8"/>
    <title>Go Mandelbrot</title>
</head>

<body>

<img id="image" src="">

<script>
    document.getElementById("image").src = " data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAADIAAAAyAQAAAAA2RLUcAAAABGdBTUEAALGPC/xhBQAAAAJiS0dEAAHdihOkAAAAB3RJTUUH5AsPECUqQdxoMgAAAGpJREFUGNOt0LENgCAQheF3obBkBEdhNBmNURiBksKA9zSnCQYrab6Sez/6+RLmZrmMCLQBnlbA0QIIzQBoUjc1qkFVsCY06h93ujxW6t4WKv87+2+8a7zb9tz7bK/ttx7Wx3pZv17ko/cB/caa7HEsGrQAAAAldEVYdGRhdGU6Y3JlYXRlADIwMjAtMTEtMTVUMTU6Mzc6NDIrMDE6MDATkZwSAAAAJXRFWHRkYXRlOm1vZGlmeQAyMDIwLTExLTE1VDE1OjM3OjQyKzAxOjAwYswkrgAAAABJRU5ErkJggg=="
</script>

</body>
</html>
```

* Load the wasm file and, execute the function GetImageAsBase64() and add the resulting string to the image.
* Add coordinates and zoom level as parameters to your function.
* Don't return the string, but use Go to call a JavaScript functions to create the image object inside the Go function. 

## Exercise 8.3 - Concurrency

* Try to run two or more goroutines in parallel. Does it work?
