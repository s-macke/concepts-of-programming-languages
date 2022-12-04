//go:build linux
// +build linux

package main

// #include <unistd.h>
import "C"
import "fmt"

func getHostname() string {
	var buf [65]C.char
	C.gethostname(&buf[0], 65)
	return C.GoString(&buf[0])
}

func main() {
	fmt.Println("The Hostname is", getHostname())
}
