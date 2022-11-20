package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

type MultiplyArgs struct {
	A, B int
}

type Arith struct{}

func (t *Arith) Multiply(args *MultiplyArgs, reply *int) error {
	a := args.A
	b := args.B
	fmt.Println("Received arguments", a, b)
	*reply = a * b
	return nil
}

func main() {
	var arith Arith
	err := rpc.RegisterName("Math", &arith)
	if err != nil {
		log.Fatal(err)
	}

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Listening on port 1234")
	rpc.Accept(listener)
}
