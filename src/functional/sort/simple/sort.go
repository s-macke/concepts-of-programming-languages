package main

import "fmt"

// Sorts an array of integers using the bubble sort algorithm.
func BubbleSort(data []int) {
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data)-1; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j] // Swap the values

			}
		}
	}
}

func main() {
	data := []int{27, 15, 8, 9, 12, 4, 17, 19, 21, 23, 25}
	BubbleSort(data)
	fmt.Println(data)
}
