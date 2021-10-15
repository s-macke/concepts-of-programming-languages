package main

import "fmt"

type I interface {
	Foo()
}

type A struct {
	a int
}

type B struct {
	a int
}

type C struct {
	A
	B
	c int
}

func (c A) Foo() {
}

func main() {
	x := C{
		c: 2,
		A: A{3},
	}

	y := make([]I, 10)
	y[0] = x

	fmt.Println(x.A.a, x.A.a)
	fmt.Println(x.c)
}
