//go:build linux || darwin
// +build linux darwin

// Copyright 2018 Johannes Weigend, Johannes  Siedersleben
// Licensed under the Apache License, Version 2.0
package main

import (
	"fmt"
	"golang.org/x/sys/unix"
)

func main() {
	pid, _, _ := unix.Syscall(unix.SYS_GETPID, 0, 0, 0)
	fmt.Println("process id: ", pid)
}
