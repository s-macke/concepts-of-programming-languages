package main

import (
	"encoding/json"
	"fmt"
	"runtime"
)

func main() {
	println("OS: ", runtime.GOOS)
	println("ARCH: ", runtime.GOARCH)
	println("ROOT: ", runtime.GOROOT())
	println("Number of CPUs: ", runtime.NumCPU())
	println("Current number of goroutines: ", runtime.NumGoroutine())
	var rtm runtime.MemStats
	runtime.ReadMemStats(&rtm)
	rtmAsBytes, _ := json.MarshalIndent(rtm, "", "  ")
	fmt.Println(string(rtmAsBytes))
}
