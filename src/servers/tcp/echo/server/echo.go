package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

// use with netcat

func main() {
	l, err := net.Listen("tcp4", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go func() { // asynchroneaous reply
			data, _ := bufio.NewReader(c).ReadString('\n')
			_, _ = c.Write([]byte("Echo: " + data))
			c.Close()
		}()
	}
}
