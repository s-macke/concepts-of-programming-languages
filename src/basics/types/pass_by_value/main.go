package main

import "fmt"

type countstruct struct{ count int }

func setint(count int) {
	count = 2
}

func setstruct(st countstruct) {
	st.count = 2
}

func incarray1(count []int) {
	count = append(count, 1)
}

func incarray2(count [1]int) {
	count = [1]int{2}
}

func incmap(count map[int]int) {
	count[0]++
}

func alterstring(s string) {
	s = "world"
}

func main() {
	countint := 1
	setint(countint)
	fmt.Println("pass int: ", countint)

	countstruct := countstruct{count: 1}
	setstruct(countstruct)
	fmt.Println("pass struct: ", countstruct)

	countmap := map[int]int{
		0: 1,
	}
	incmap(countmap)
	fmt.Println("pass map:", countmap)

	countarray1 := []int{1}
	incarray1(countarray1)
	fmt.Println("pass arbitrary length array: ", countarray1)

	countarray2 := [1]int{1}
	incarray2(countarray2)
	fmt.Println("pass fixed array: ", countarray2)

	s := "hello"
	alterstring(s)
	fmt.Println("pas string", s)
}
