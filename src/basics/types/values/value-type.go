package main

import "fmt"

func foo(x int) {
	x = 2
}

func bar(s string) {
	s = "world"
}

func main() {
	x := 1
	foo(x)
	fmt.Println(x) // 1

	s := "Hello"
	bar(s)
	fmt.Println(s) // Hello
}
