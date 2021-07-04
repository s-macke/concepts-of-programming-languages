// +build linux

// Simple "Docker"/Container implementation by Liz Rice.

package main

import (
	"log"
	"os"
	"os/exec"
)

// go run main.go run <cmd> <args>
func main() {
	switch os.Args[1] {
	case "run":
		run()
	default:
		panic("help")
	}
}

func run() {
	log.Printf("Running %v \n", os.Args[1:])

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	//cmd.SysProcAttr = &syscall.SysProcAttr{
	//	Cloneflags:   syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
	//	Unshareflags: syscall.CLONE_NEWNS,
	//}

	must(cmd.Run())
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
