//go:build linux
// +build linux

package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("uid: ", os.Getuid())
	fmt.Println("gid: ", os.Getgid())
	fmt.Println("Pid: ", os.Getpid())
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
