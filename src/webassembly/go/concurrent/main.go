package main

import (
	"fmt"
	"math/big"
	"runtime"
)
import "time"

func ListPrimes() {
	for n := 100000000001; ; n++ {
		if big.NewInt(int64(n)).ProbablyPrime(100000) {
			fmt.Printf("%d\n", n)
		}
		runtime.Gosched()
	}
}

func main() {
	go ListPrimes()
	for {
		fmt.Printf("Hello, WebAssembly!\n")
		time.Sleep(1 * time.Second)
	}
}
