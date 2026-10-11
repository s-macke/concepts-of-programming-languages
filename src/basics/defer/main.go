package main

import "fmt"

func printNumbers() {
	for i := 0; i < 3; i++ {
		defer fmt.Println(i)
	}
}

func main() {
	printNumbers()
}
