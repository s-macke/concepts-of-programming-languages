package main

import "fmt"

var V int

func F() {
	fmt.Printf("Hello, number %d\n", V)
}

func main() {
	fmt.Println("plugin loaded")
}
