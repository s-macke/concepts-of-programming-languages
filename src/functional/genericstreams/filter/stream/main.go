package main

import (
	"fmt"
	"github.com/mariomac/gostream/stream"
)

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	filteredArray := stream.OfSlice(array).
		Filter(func(n int) bool { return n%2 == 0 }).
		ToSlice()
	fmt.Println(filteredArray)
}
