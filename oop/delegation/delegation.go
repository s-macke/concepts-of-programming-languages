package main

import "fmt"

type A struct {
}

func (a A) Foo() {
	a.Bar()
}

func (a A) Bar() {
	fmt.Print("a.bar")
}

type B struct {
	A
}

func (b B) Bar() {
	fmt.Print("b.bar")
}

func main() {
	b := B{}
	b.Foo() // "a.bar" or "b.bar"?
}

// EOF OMIT
