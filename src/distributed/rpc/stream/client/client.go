package main

import (
	"bufio"
	"fmt"
	"log"
	"net/rpc"
	"os"
)

func Reader(client *rpc.Client) {
	var dummy struct{}
	for {
		var reply rune
		err := client.Call("Session.Read", &dummy, &reply)
		if err != nil {
			log.Fatal("read error:", err)
		}
		fmt.Print(string(reply))
	}
}

func Writer(client *rpc.Client) {
	var dummy struct{}
	reader := bufio.NewReader(os.Stdin)
	for {
		r, _, _ := reader.ReadRune()
		err := client.Call("Session.Write", &r, &dummy)
		if err != nil {
			log.Fatal("write error:", err)
		}
	}
}

func main() {
	client, err := rpc.Dial("tcp", "simulationcorner.net:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	go Reader(client)
	Writer(client)
}
