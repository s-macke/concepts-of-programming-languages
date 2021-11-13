package main

import (
	"bufio"
	"fmt"
	"net"
)

// use with netcat

var broker = NewBroker()

func handleConnection(c net.Conn) {
	user := c.RemoteAddr().String() + ": "
	fmt.Printf("Connection from %s\n", c.RemoteAddr().String())
	defer c.Close()
	_, err := c.Write([]byte("Connected to chat server\n"))
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = c.Write([]byte("Your username " + user + "\n"))
	if err != nil {
		fmt.Println(err)
		return
	}

	go func() {
		var read = broker.Subscribe()
		for {
			s := <-read
			_, err := c.Write([]byte(s.(string)))
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}()

	for {
		data, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		broker.Publish(user + data)
	}
}

func main() {
	go broker.Start()
	defer broker.Stop()

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
		go handleConnection(c)
	}
}
