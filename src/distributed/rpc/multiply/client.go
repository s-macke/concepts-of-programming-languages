package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type MultiplyArgs struct {
	A, B int
}

func Multiply(a, b int) int {
	args := MultiplyArgs{a, b}
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	// Synchronous call
	var reply int
	err = client.Call("Math.Multiply", &args, &reply)
	if err != nil {
		log.Fatal("Mutliply error:", err)
	}
	return reply
}

func main() {
	a := 7
	b := 8
	fmt.Printf("Multiply: %d * %d = %d\n", a, b, Multiply(a, b))
}
