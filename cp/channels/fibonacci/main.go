package main

import "fmt"

func fib() <-chan int {
	c := make(chan int)
	go func() {
		for i, j := 0, 1; ; i, j = i+j, i {
			c <- i
		}
	}()
	return c
}

// MAIN OMIT
func main() {
	fibChan := fib() // <- write func fib
	for n := 1; n <= 10; n++ {
		fmt.Printf("The %dth Fibonacci number is %d\n", n, <-fibChan)
	}
}
// EOF OMIT
