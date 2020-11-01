package main

import (
	"fmt"
	"math/rand"
	"time"
)

func lecturer() <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; i < 5; i++ {
			c <- fmt.Sprintf("%d Bla bla goroutines bla channels bla bla\n", i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

func main() {
	c := lecturer()
	for i := 0; i < 5; i++ {
		fmt.Printf(<-c)
	}
}

// EOF OMIT
