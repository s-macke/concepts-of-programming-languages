package main

import (
	"fmt"
	"os"
	"plugin"
	"reflect"
)

func main() {
	p, err := plugin.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	v, err := p.Lookup("V")
	if err != nil {
		panic(err)
	}
	f, err := p.Lookup("F")
	if err != nil {
		panic(err)
	}
	fmt.Println("type of v is ", reflect.TypeOf(v).String())
	fmt.Println("type of f is ", reflect.TypeOf(f).String())
	*v.(*int) = 7
	f.(func())() // prints "Hello, number 7"
}
