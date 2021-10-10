package main

import "fmt"

type countstruct struct {count int}

func inc(count int) {
	count++
}

func incstruct(st countstruct) {
	st.count++
}

func incarray1(count []int) {
	count[0]++
}

func incarray2(count [1]int) {
	count[0]++
}

func incmap(count map[int]int) {
	count[0]++
}

func alterstring(s string) {
	s = "world"
}

func main() {
	countint := 1
	inc(countint)
	fmt.Println(countint)

	countarray1 := []int{1}
	incarray1(countarray1)
	fmt.Println(countarray1)

	countarray2 := [1]int{1}
	incarray2(countarray2)
	fmt.Println(countarray2)

	countstruct := countstruct{count: 1}
	incstruct(countstruct)
	fmt.Println(countstruct)

	countmap := map[int]int{
		0: 1,
	}
	incmap(countmap)
	fmt.Println(countmap)

	s := "hello"
	alterstring(s)
	fmt.Println(s)
}

