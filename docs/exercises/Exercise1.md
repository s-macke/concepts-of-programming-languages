# Exercise 1 - Getting Started with Go

If you do not finish during the lecture period, please finish it as homework.

## Setup

- Install Go from <https://go.dev> or install it via a package manager.
- Check that Go works on the command line: `go version`
- Look at the Go environment via `go env`, especially the variable `GOPATH`. Downloaded dependencies are stored in the module cache below it (`$GOPATH/pkg/mod`). Your own code can live anywhere.
- Create a project directory, e.g. `hello`, and change into it.
- Initialize a module via `go mod init hello`. Check the content of the created file `go.mod`.
- Create a file `main.go` with a Hello World program.
- Run the program via `go run .`
- Compile the program via `go build` and run the created binary.
- Optional, but desirable: Install Visual Studio Code, IntelliJ, GoLand or any other editor with Go support and get familiar with it.
- Try to run the Hello World example inside <https://goplay.tools/>
  Alternative: <https://go.dev/play>
  Alternative: <https://go-playground-wasm.vercel.app/>
- Optional: Get familiar with the community. Look at what others have done: <https://github.com/avelino/awesome-go>

## Learn the basics

Work through the first seven examples at <https://gobyexample.com/>: Hello World, Values, Variables, Constants, For, If/Else and Switch. Type them in and run them yourself instead of only reading them.

Alternative: the interactive Tour of Go at <https://go.dev/tour>, which runs directly in the browser.

## Check Go's classification yourself

In the lecture, Go was classified as *compiled*, *native machine code* and *cross compiler*. Check it yourself. The following commands work the same on Windows, macOS and Linux.

- Which target platforms does Go support? Run `go tool dist list`. How many are there? Can you find WebAssembly?
- Build your Hello World for your own system via `go build`. How large is the binary? Why is a simple Hello World that large?
- Cross-compile the same program for Linux on ARM:

  ```
  go env -w GOOS=linux GOARCH=arm64
  go build -o hello-linux
  go env -u GOOS GOARCH
  ```

  Important: The last line resets the target to your own system. Otherwise all following builds are also built for Linux.
- Inspect both binaries via `go version -m <binary>`. For which system and architecture was each one built? Can you run the Linux binary on your computer?

## What does this print in Go?

In the lecture we saw some surprising results in JavaScript. Now it is Go's turn.

For each snippet:

1. Guess the output **before** you run it.
2. Run it in your `main.go` or in the playground.
3. Explain the result. Which concept of the language is behind it?

```go
fmt.Println(0.1 + 0.2)
```

```go
a, b := 0.1, 0.2
fmt.Println(a + b)
```

```go
var x int8 = 127
x++
fmt.Println(x)
```

```go
s := "Hello, 世界"
fmt.Println(len(s), len([]rune(s)))
```

Compare the first two snippets: Why do they print different results? And how does JavaScript behave here?

## If you finish early

- Pick an unfamiliar language from the concepts board in Miro, e.g. Zig, Elixir or Forth. Write and run a Hello World in an online playground. Which concepts do you notice that differ from Go? This might also be an inspiration for the topic of your semester work.

## Answer the following question
- In your opinion, what are the characteristics of a successful language? Add your answer to the Miro board.

## After this Exercise

- You should have a working Go installation on your computer
- You should know how to compile and run Go code
