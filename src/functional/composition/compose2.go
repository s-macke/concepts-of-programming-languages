package main

import "fmt"

// Type any makes the code readable
type any interface{}
type function func(any) any

func main() {
	compose := func(g, f function) function {
		return func(x any) any {
			return g(f(x))
		}
	}

	square := func(x any) any { return x.(int) * x.(int) }

	fmt.Printf("%v\n", compose(square, square)(2)) // --> 4*4 = 16
	fmt.Printf("%v\n", compose(compose(square, square), square)(2))
}
