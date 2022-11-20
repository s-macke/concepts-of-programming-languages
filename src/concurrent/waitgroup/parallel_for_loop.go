package main

import (
	"fmt"
	"math/big"
	"sync"
)

func ParallelFor(n int, f func(int)) {
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(i int) {
			defer wg.Done()
			f(i)
		}(i)
	}
	wg.Wait()
}

func ProbablyPrime(value int) {
	if big.NewInt(int64(value)).ProbablyPrime(0) == true {
		fmt.Printf("%d is probably prime\n", value)
	}
}

func main() {
	fmt.Println("Start Program")
	ParallelFor(1000, ProbablyPrime)
	fmt.Println("Stop Program")
}
