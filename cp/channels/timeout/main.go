package main

import (
	"fmt"
	"time"
)

func setTimeout(task func() int, timeout time.Duration) (int, error) {
	taskChan := make(chan int)
	go func() {
		taskChan <- task()
	}()
	timerChan := time.After(timeout)
	select {
	case result := <-taskChan:
		return result, nil
	case <-timerChan:
		return 0, fmt.Errorf("operation timed out")
	}
}

// MAIN OMIT
func main() {
	res, err := setTimeout(func() int {
		time.Sleep(2000 * time.Millisecond)
		return 1
	}, 1*time.Second)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("operation returned %d", res)
	}
}
// EOF OMIT
