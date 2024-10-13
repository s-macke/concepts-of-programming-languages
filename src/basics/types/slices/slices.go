package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5} // create fixed array
	fmt.Println(arr)

	s := arr[0:2] // create a slice out of the array. Otherwise via "make([2]int)" for a new array
	fmt.Println(s)

	fmt.Println(len(s), cap(s)) // ==> 2, 5
	s[0] = 100                  // a slice is just a reference
	fmt.Println(arr)

	s = append(s, 8)
	s = append(s, 9)
	s = append(s, 10)
	s = append(s, 11)

	fmt.Println(len(s), cap(s)) // ==> 6, 10
	fmt.Println(arr)
}
