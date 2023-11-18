package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var myYoutubevideo struct {
	likes int32
}

func Viewer() {
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	atomic.AddInt32(&myYoutubevideo.likes, 1)
}

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			Viewer()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(myYoutubevideo.likes, "likes")
}
