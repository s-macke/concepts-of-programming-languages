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

func main() {
	ParallelFor(1000, func(value int) {
		if big.NewInt(int64(value)).ProbablyPrime(0) == true {
			fmt.Printf("%d is probably prime\n", value)
		}
	})
}
