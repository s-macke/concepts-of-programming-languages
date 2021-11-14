//go:build linux || darwin
// +build linux darwin

package main

import (
	"golang.org/x/sys/unix"
	"log"
	"syscall"
	"unsafe"
)

func main() {
	path := "mydir"
	dir, _ := syscall.BytePtrFromString(path)
	mode := 0777
	_, _, e := unix.Syscall(unix.SYS_MKDIR, uintptr(unsafe.Pointer(dir)), uintptr(mode), 0)
	if e != 0 {
		log.Fatal("Error code", e)
	}
}
