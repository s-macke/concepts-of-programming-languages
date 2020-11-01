package main

import (
	"fmt"
	"math/rand"
	"time"
)

func lecturer() {
	for i := 0; i < 5; i++ {
		fmt.Printf("%d Bla bla goroutines bla channels bla bla\n", i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

func main() {
	go lecturer()
}

// EOF OMIT
