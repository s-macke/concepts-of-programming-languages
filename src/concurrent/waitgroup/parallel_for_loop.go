package main

import (
	"fmt"
	"math/big"
	"sync"
)

var wg sync.WaitGroup

func ParallelFor(n int, f func(int)) {
	wg.Add(n)
	for i := 0; i < n; i++ {
		go f(i)
	}
	wg.Wait()
}

func ProbablyPrime(value int) {
	if big.NewInt(int64(value)).ProbablyPrime(0) == true {
		fmt.Printf("%d is probably prime\n", value)
	}
	wg.Done()
}

func main() {
	fmt.Println("Start Program")
	ParallelFor(1000, ProbablyPrime)
	fmt.Println("Stop Program")
}
