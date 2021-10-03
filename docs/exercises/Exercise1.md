# Exercise 1 - Getting Started with Go

If you do not finish during the lecture period, please finish it as homework.

## Setup

- Install Go from <http://golang.org> inside a virtual disk
  (.vhdx on Windows or .sparseimage on Mac) in a /software/go subdirectory.
  To have a really independent image choose the archives and not the installers.

  - Win: <https://www.windowscentral.com/how-create-and-set-vhdx-or-vhd-windows-10>
  - Mac: <https://support.apple.com/de-de/guide/disk-utility/dskutl11888/mac>

- Create a Go workspace on the disk in the <<DISK>>/codebase/go directory
  - <https://www.youtube.com/watch?v=XCsL89YtqCs>
- Create a shell script (.sh / .cmd) to make your changes to the `GOPATH` and PATH environment variables persistent.
  - Check the configuration with `go env`
- Create a `HelloWorld.go` program inside the `GOPATH`/src/hello folder
- Test the `HelloWorld` program with "go run HelloWorld"
- Compile the `HelloWorld` program with "go build" and "go install". You need to init the module first via `go mod init`.
- What is the difference between go build and go install?
- Optional, but desirable: Install Visual Studio Code, IntelliJ or any other Editor with Go support inside your virtual disk.
- Try to run the hello world example code inside <https://goplay.tools/>
  Alternative: https://play.golang.org/
  Alternative: https://go-playground-wasm.vercel.app/
- Optional: Get familiar with the community. Look what others have done: <https://github.com/avelino/awesome-go>

## After this Exercise

- You should know how to compile and run Go code
- You should know about the meaning of the `GOPATH` and PATH variables
- You should have a portable Go installation inside a separate disk or directory on your computer
