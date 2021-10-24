//go:build linux || darwin
// +build linux darwin

// Copyright 2018 Johannes Weigend, Johannes  Siedersleben
// Licensed under the Apache License, Version 2.0
package main

import (
	"golang.org/x/sys/unix"
	"log"
	"os"
	"syscall"
	"unsafe"
)

func main() {
	os.Mkdir()
	path := "mydir"
	dir, _ := syscall.BytePtrFromString(path)
	mode := 0777
	_, _, e := unix.Syscall(unix.SYS_MKDIR, uintptr(unsafe.Pointer(dir)), uintptr(mode), 0)
	if e != 0 {
		log.Fatal("Error code", e)
	}
}
