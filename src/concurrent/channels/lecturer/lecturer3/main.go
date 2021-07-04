package main

import (
	"fmt"
	"math/rand"
	"time"
)

func lecturer(c chan string) {
	for i := 0; i < 5; i++ {
		c <- fmt.Sprintf("%d Bla bla goroutines bla channels bla bla\n", i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

func main() {
	c := make(chan string)
	go lecturer(c)
	for i := 0; i < 5; i++ {
		fmt.Printf(<-c)
	}
}

// EOF OMIT
