package main

import (
	"fmt"
	"math/rand"
	"time"
)

func lecturer(name string, speed int) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; i < 5; i++ {
			c <- fmt.Sprintf("%s: %d Bla bla goroutines bla channels bla bla\n", name, i)
			time.Sleep(time.Duration(rand.Intn(1e3*speed)) * time.Millisecond)
		}
	}()
	return c
}

// START OMIT
// func lecturer(name string, speed int) <-chan string { ... }

func fanIn(c1, c2 <-chan string) <-chan string {
	c := make(chan string)
	go func() { for { c <- <-c1 } }()
	go func() { for { c <- <-c2 } }()
	return c
}

func main() {
	a := lecturer("Anne", 1)
	b := lecturer("Bart", 2)
	c := fanIn(a, b)
	for i := 0; i < 10; i++ {
		fmt.Printf(<-c)
	}
}

// EOF OMIT
