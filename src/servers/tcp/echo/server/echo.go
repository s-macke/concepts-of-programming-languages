package main

import (
	"bufio"
	"fmt"
	"net"
)

// use with netcat or browser

func main() {
	l, err := net.Listen("tcp4", ":8081")
	if err != nil {
		panic(err)
	}
	defer l.Close()
	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("Connection established from", c.RemoteAddr())
		go func() { // asynchronous reply
			data, _ := bufio.NewReader(c).ReadString('\n')
			fmt.Println("Received string", string(data))
			_, _ = c.Write([]byte("Echo: " + data))
			c.Close()
		}()
	}
}
