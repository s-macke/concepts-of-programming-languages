package main

import "fmt"

type A struct {
}

func (a A) Foo() {
	fmt.Println("a.foo")
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

/*
func (b B) Foo() {
	fmt.Println("b.foo")
	b.Bar()
}
*/
func main() {
	b := B{}
	b.Foo() // "a.bar" or "b.bar"?
}

// EOF OMIT
