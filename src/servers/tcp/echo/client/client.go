package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8081")
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(conn, "Hello world\n")
	io.Copy(os.Stdout, conn)
}
