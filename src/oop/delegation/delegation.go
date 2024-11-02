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
	//b.Bar() //b.A.Bar()
	b.A.Bar()
}

func (b B) Bar() {
	fmt.Println("b.bar")
}

func main() {
	a := A{}
	a.Foo()

	b := B{}
	b.Foo() // "a.bar" or "b.bar"?

}
