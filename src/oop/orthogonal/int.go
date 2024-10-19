package main

import "fmt"

type MyInt int

func (b *MyInt) Inc() {
	*b = *b + 1
}

func main() {
	var b MyInt = 10
	fmt.Println(b)
	b.Inc()
	fmt.Println(b)
}
