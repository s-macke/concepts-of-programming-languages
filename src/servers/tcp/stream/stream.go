package main

import (
	"io"
	"net"
	"os"
)

func main() {
	c, err := net.Dial("tcp", "towel.blinkenlights.nl:23")
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, c)
}
