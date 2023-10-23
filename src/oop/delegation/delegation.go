package main

import "fmt"

type A struct {}
func (a A) Foo() {
	fmt.Print("a.foo ")
	a.Bar()
}

type B struct {
	A
}
func (b B) Foo() {
	fmt.Print("b.foo ")
	b.Bar()
}

func main() {
	b := B{}
	b.Foo() // What happens?
}
