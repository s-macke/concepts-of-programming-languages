package main

import "fmt"

func main() {
	ch := make(chan int)
	ch <- 1
	ch <- 2 // dead by now
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
