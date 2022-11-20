package main

import "fmt"

func main() {
	ch := make(chan int, 2) // channel with buffer size 2
	ch <- 1
	ch <- 2 // dead by now
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
