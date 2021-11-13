package main

import (
	"log"
	"math/big"
	"net"
	"net/rpc"
)

type MultiplyArgs struct {
	A, B big.Int
}

type Arith struct{}

func (t *Arith) Multiply(args *MultiplyArgs, reply *big.Int) error {
	*reply = *args.A.Mul(&args.A, &args.B)
	return nil
}

func main() {
	arith := new(Arith)
	rpc.Register(arith)

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	rpc.Accept(listener)
}
