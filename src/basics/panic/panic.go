package main

import "fmt"

func getElement(array []int, index int) int {
	if index >= len(array) {
		panic("Out of bounds")
	}
	return array[index]
}

func getElementWithRecover(array []int, index int) (value int) {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("Recovered with message '", r, "'")
		}
		value = -1
	}()
	return getElement(array, index)
}

func main() {
	array := []int{3, 4, 5}
	ret := getElementWithRecover(array, 3)
	fmt.Println("return value: ", ret)
}

// EOF OMIT
