package main

import "fmt"

type A struct{}

func (a A) Foo() {
	fmt.Print("a.foo ")
	a.Bar()
}

func (a A) Bar() {
	fmt.Println("a.bar")
}

type B struct {
	A
}

func (b B) Foo() {
	fmt.Print("b.foo ")
	b.A.Foo()
}

func (b B) Bar() {
	fmt.Println("b.bar")
}

func main() {
	a := A{}
	a.Foo()

	b := B{}
	b.Foo() // "b.foo a.foo a.bar" or "b.foo a.foo b.bar"?
}
