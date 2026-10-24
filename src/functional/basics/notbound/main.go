package main

import "fmt"

func main() {
	fmt.Println(func(x, y int) int { return x + y }(3, 4)) // -> returns 7
}
