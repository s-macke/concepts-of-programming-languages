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
	time.Sleep(time.Duration(rand.Intn(300)) * time.Millisecond)
	myYoutubevideo.likes = myYoutubevideo.likes + 1
}

func main() {
	for i := 0; i < 10000; i++ {
		go Viewer()
	}
	for i := 0; i < 8; i++ {
		fmt.Println(myYoutubevideo.likes, "likes")
		time.Sleep(200 * time.Millisecond)
	}
}
