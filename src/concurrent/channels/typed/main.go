package main

import (
	"fmt"
	"time"
)

func main() {
	timerChan := make(chan time.Time)

	go func() {
		time.Sleep(3 * time.Second)
		timerChan <- time.Now() // send time on timerChan
	}()

	fmt.Println("Started at", time.Now())
	// Do something else; when ready, receive.
	// Receive will block until timerChan delivers.
	// Value sent is other goroutine's completion time.
	completedAt := <-timerChan
	fmt.Println("Completed at", completedAt)
}
