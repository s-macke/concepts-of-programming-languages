package main

import (
	"fmt"
	"github.com/mariomac/gostream/stream"
)

func main() {
	array := []int{10, 5, 3, 7, 1, 0, 4, 6}
	sortedArray := stream.OfSlice(array).
		Sorted(func(a, b int) int { return a - b }).
		ToSlice()
	fmt.Println(sortedArray)
}
