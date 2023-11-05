package main

import "fmt"

// Sorts an array of integers using the bubble sort algorithm.
func BubbleSort[T any](data []T, cmp func(i, j int) bool) {
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data)-1; j++ {
			if cmp(j, j+1) {
				data[j], data[j+1] = data[j+1], data[j] // Swap the values
			}
		}
	}
}

func main() {
	data := []float32{27, 15, 8, 9, 12, 4, 17, 19, 21, 23, 25}
	BubbleSort(data, func(i, j int) bool {
		return data[i] > data[j]
	})
	fmt.Println(data)
}
