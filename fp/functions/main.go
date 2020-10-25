package main

import "fmt"

// START OMIT
func aBlock(i int) {
	fmt.Printf("Entering block: i=%v\n", i)
}

func do(f func(int), loops int) {
	for i := 0; i < loops; i++ {
		f(i)
	}
}

func main() {
	do(aBlock, 5)
}

// EOF OMIT
