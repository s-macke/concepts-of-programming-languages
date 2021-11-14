//go:build linux
// +build linux

package main

// #include <stdlib.h>
import "C"
import "fmt"

func Seed(i int) {
	C.srandom(C.uint(i))
}

func Random() int {
	var r C.long = C.random() // calls "int rand(void)" from the C std library
	return int(r)
}

func main() {
	Seed(1)
	fmt.Printf("random int from C is %d\n", Random())
}
