package main

import (
	"fmt"
	"sort"
)

func main() {
	array := []int{10, 5, 3, 7, 1, 0, 4, 6}
	sort.SliceStable(array, func(i, j int) bool {
		return array[i] < array[j]
	})
	fmt.Println(array)
}
