package main

import (
	"constraints"
	"fmt"
)

func min[T constraints.Ordered](x, y T) T {
	if x < y {
		return x
	} else {
		return y
	}
}

func main() {
	fmt.Println(min(42, 24))
}
