package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func Sleep() {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
}

func main() {
	fmt.Println("Number of Goroutines running: ", runtime.NumGoroutine())
	for i := 0; i < 1000000; i++ {
		go Sleep()
	}
	for i := 0; i < 15; i++ {
		fmt.Println("Number of Goroutines running: ", runtime.NumGoroutine())
		time.Sleep(100 * time.Millisecond)
	}

}
