# Exercise 11 - Webassembly

## Exercise 11.1 - Hello world Go project in WebAssembly

* Install npm/node on your system
  - https://nodejs.org/en/download/

---
* Write a Hello World program in Go
* Compile by setting the environment variables GOARCH and GOOS

```bash
export GOARCH=wasm 
export GOOS=js 
go build -o lib.wasm main.go
```

The result will be a file called `lib.wasm`

* In case you run Windows command shell you have to execute three commands
```powershell
set GOARCH=wasm 
set GOOS=js 
go build -o lib.wasm main.go`
```

* Copy the wasm JavaScript runtime from the GOROOT directory. 
You will find the GOROOT directory via  `go env GOROOT` and copy the wasm_exec file:
 `cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .` and
 `cp "$(go env GOROOT)/misc/wasm/wasm_exec_node.js" .`

* Alternative path: `/usr/share/doc/go/misc/wasm/wasm_exec.js`
* Execute via `node wasm_exec_node.js lib.wasm`
---

* An HTTP file web server is provided in the folder `src/servers/fileserver`.
Alternatively, you can use python `python -m http.server 8080`
* Copy the index.html file and test your program inside the browser

```HTML
<html>
<body>
    <script src="wasm_exec.js"></script>
    <script>
        const go = new Go();

        async function Run() {
            let response = await fetch('lib.wasm')
            let buffer = await response.arrayBuffer();
            let result = await WebAssembly.instantiate(buffer, go.importObject);
            go.run(result.instance)
        }

        Run()
    </script>
</body>
</html>
```
---

## Exercise 11.2 - Concurrency

Javascript or WebAssembly doesn't support concurrency and are single threaded.

* Experiment with the Go concurrency in WebAssembly.
  Try to run two or more Goroutines in parallel. Does it work?
* Webassembly does not support Coroutines, e.g. blocking calls such as `time.Sleep(1 * time.Second)`.
  Test blocking calls together with concurrency. Does it work?
* Try to run the dining philosophers problem in the web browser.
  Example source code `src/concurrent/channels/philosophers`

## Exercise 11.3 - Execute a Go function inside Go

In one of our last lectures I showed a very simple reverse polish notation calculator 

```
func main() {
	var s stack
	fields := strings.Fields("2 3 5 * +")
	for _, expr := range fields {
		fmt.Println("Parse ", expr)
		switch expr {
		case "+":
			s.Push(s.Pop() + s.Pop())
		case "*":
			s.Push(s.Pop() * s.Pop())
		default:
			v, err := strconv.Atoi(expr)
			if err != nil {
				fmt.Println("Unknown expr", expr)
				return
			}
			s.Push(v)
		}
	}
	println("result:", s.Pop())
}
```

This can be programmed as a function with an input string and an output string. 
* Use a stack implementation from one of the first lectures and rewrite it, 
  so that the calculator is separated into its own function and not run inside the main routine.
* Rewrite it, so that it can be used in WebAssembly.

You can use the following code to create the necessary objects such as input field, button and output text.

```html
<html>
    <body>
        <input type="text" id="input" >
        <button type="button" onclick="execute_function()">Execute</button>
        <p id="output"></p>
    </body>
</html>
```

You can then access the elements in a Javascript `<script>` element via

```JavaScript
  alert(document.getElementById("input").value)
  document.getElementById("output").innerText = "Hello world"
```

* Try to create the necessary HTML elements via Golang.
Here is an example code.
```JavaScript
	document = js.Global().Get("document")
	p = document.Call("createElement", "p")
	p.Set("innerHTML", "Hello WASM from Go!")
	document.Get("body").Call("appendChild", p)
```