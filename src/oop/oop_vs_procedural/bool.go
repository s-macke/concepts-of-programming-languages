package main

//Compile with *go run -gcflags '-N -l' door.go* to disable optimization and code inlining

import "fmt"

type MyBoolean bool

func (d *MyBoolean) Set(v bool) {
	fmt.Println(d)              // print pointer to d
	panic("full stack message") // Force panic. Shows used parameters for Set after compilation
	*d = MyBoolean(v)
}

func main() {
	var d MyBoolean = false
	d.Set(true)
}
