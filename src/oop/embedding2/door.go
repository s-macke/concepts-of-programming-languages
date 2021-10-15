package main

import "fmt"

type Door bool

func (d *Door) Open() {
	*d = true
}

func (d *Door) Close() {
	*d = false
}

func (d *Door) Set(v bool) {
	*d = Door(v)
	fmt.Println(d)
	panic("full stack message")
}

func main() {
	var d Door = false
	fmt.Println(d)
	d.Open()
	fmt.Println(d)
	d.Close()
	fmt.Println(d)
	d.Set(true)

}
