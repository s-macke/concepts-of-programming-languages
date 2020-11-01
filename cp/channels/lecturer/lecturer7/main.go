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

func main() {
	a := lecturer("Anne", 1)
	b := lecturer("Bart", 2)
	for i := 0; i < 10; i++ {
		select {
		case msgFromAnne := <-a:
			fmt.Printf(msgFromAnne)
		case msgFromBart := <-b:
			fmt.Printf(msgFromBart)
		}
	}
}

// EOF OMIT
