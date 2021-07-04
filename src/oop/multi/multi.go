package main

import "fmt"

type Fooer interface {
	Foo()
}

type Barer interface {
	Bar()
}

type X struct{}

func (x X) Foo() {}
func (x X) Bar() {}

// END1 OMIT

type Foo struct {
	Name string
}

type Bar struct {
	Name string
}

type Y struct {
	Foo
	Bar
}

func main() {
	y := Y{
		Foo: Foo{Name: "Foo!"},
		Bar: Bar{Name: "Bar!"},
	}
	fmt.Print(y.Foo.Name)
	fmt.Print(y.Bar.Name)
	//fmt.Print(y.Name) // Ambiguous Reference
}

// EOF OMIT
