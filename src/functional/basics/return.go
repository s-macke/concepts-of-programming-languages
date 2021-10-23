package main

import "fmt"

type areaFunc func(int, int) int

func getAreaFunc() areaFunc {
	return func(x, y int) int {
		return x * y
	}
}

func main() {
	areaF := getAreaFunc()
	res := areaF(2, 4)
	fmt.Println(res)
}
