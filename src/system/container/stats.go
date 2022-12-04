//go:build linux
// +build linux

package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("User ID: ", os.Getuid())
	fmt.Println("Group ID: ", os.Getgid())
	fmt.Println("Process Id: ", os.Getpid())
	PrintHostname("Current")
	/*
		syscall.Sethostname([]byte("container"))
		PrintHostname("New")
	*/
}

func PrintHostname(pre string) {
	hostname, _ := os.Hostname()
	fmt.Println(pre, "Hostname: ", hostname)
}
