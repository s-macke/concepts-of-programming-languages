package main

import "fmt"

type A struct {
}

func (a A) Foo() {
	a.Bar()
}

func (a A) Bar() {
	fmt.Println("a.bar")
}

type B struct {
	A
}

func (b B) Bar() {
	fmt.Println("b.bar")
}

func (b B) Foo() {
	b.Bar()
}

func main() {
	b := B{}
	b.Foo() // "a.bar" or "b.bar"?
}

// EOF OMIT
