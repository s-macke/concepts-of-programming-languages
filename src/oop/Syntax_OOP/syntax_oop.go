package main

import "fmt"

type Bar struct{}

func (b *Bar) GetHello() string {
	fmt.Println(b)
	return "Hello"
}

type Foo struct {
	B *Bar
}

func main() {
	var f Foo
	fmt.Println(f.B)            // Output "nil"
	fmt.Println(f.B.GetHello()) // Null Pointer error or "Hello"?
}
