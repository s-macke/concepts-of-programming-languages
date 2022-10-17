package main

import "fmt"

func main() {
	var a int32 = -1
	var b uint32 = 2

	if a < int32(b) {
		fmt.Printf("a is smaller than b")
	}
}
