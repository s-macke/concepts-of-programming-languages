package main

import (
	"fmt"
	"runtime"
)
import "time"
import "math"

func IsPrime(value int) bool {
	for i := 2; i <= int(math.Floor(float64(value)/2)); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}

func ListPrimes() {
	n := 10000001
	for {
		if IsPrime(n) {
			fmt.Printf("%d\n", n)
		}
		n += 2
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
