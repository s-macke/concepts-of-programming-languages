package main

import (
	"fmt"
	"math/rand"
	"time"
)

var myYoutubevideo struct {
	likes int32
}

func Viewer() {
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	myYoutubevideo.likes = myYoutubevideo.likes + 1
}

func main() {
	for i := 0; i < 10000; i++ {
		go Viewer()
	}

	time.Sleep(3 * time.Second)
	fmt.Println(myYoutubevideo.likes)
}
