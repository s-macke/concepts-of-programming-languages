package main

import (
	"bufio"
	"fmt"
	"net"
)

// use with netcat

func main() {
	l, err := net.Listen("tcp4", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go func() {
			data, _ := bufio.NewReader(c).ReadString('\n')
			_, _ = c.Write([]byte("Echo: " + data))
			c.Close()
		}()
	}
}
