package main

import "fmt"

// START OMIT
func do(f func(int), loops int) {
	for i := 0; i < loops; i++ {
		f(i)
	}
}

func main() {
	printvalue := func(i int) { fmt.Println(i) }
	do(printvalue, 5)
}

// EOF OMIT
