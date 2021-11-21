//go:build linux
// +build linux

package main

import (
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("./stats")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if cmd != nil {
		panic(err)
	}
}
