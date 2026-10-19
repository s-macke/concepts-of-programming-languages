package main

import "fmt"

type Counter struct {
	n int
}

func (c Counter) Inc() {
	c.n++
}

func main() {
	c := Counter{}
	c.Inc()
	c.Inc()
	fmt.Println(c.n)
}
