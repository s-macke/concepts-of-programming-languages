package main

import "fmt"

// function taking an array
func modifyArray(a [3]int) {
	a[0] = 99 // modifies only the local copy
}

// function taking a slice
func modifySlice(s []int) {
	s[0] = 99 // modifies the underlying array
}

func main() {
	// Arrays
	arr := [3]int{1, 2, 3}
	modifyArray(arr)
	fmt.Println("Array after modifyArray:", arr) // [1 2 3] unchanged

	// Slices
	slice := []int{1, 2, 3}
	modifySlice(slice)
	fmt.Println("Slice after modifySlice:", slice) // [99 2 3] changed
}
