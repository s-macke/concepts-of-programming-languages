package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	filteredArray := []int{}
	for _, v := range array {
		if v%2 == 0 {
			filteredArray = append(filteredArray, v)
		}
	}
	fmt.Println(filteredArray)
}
