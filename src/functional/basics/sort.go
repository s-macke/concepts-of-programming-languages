package main

import (
	"fmt"
	"sort"
)

func main() {
	data := []int{27, 15, 8, 9, 12, 4, 17, 19, 21, 23, 25}
	sort.SliceStable(data, func(i, j int) bool {
		return data[i] < data[j]
	})
	fmt.Println(data)
}
