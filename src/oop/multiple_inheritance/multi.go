package main

import "fmt"

// END1 OMIT

type Foo struct {
	Name string
}

type Bar struct {
	Name string
}

type X struct {
	Foo
	Bar
}

func main() {
	y := X{
		Foo: Foo{Name: "Foo!"},
		Bar: Bar{Name: "Bar!"},
	}
	fmt.Print(y.Foo.Name)
	fmt.Print(y.Bar.Name)
	//fmt.Print(y.Name) // compile error, Ambiguous Reference
}

// EOF OMIT
