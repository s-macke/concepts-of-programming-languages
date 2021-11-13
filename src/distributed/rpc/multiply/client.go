package main

import (
	"fmt"
	"log"
	"math/big"
	"net/rpc"
)

type MultiplyArgs struct {
	A, B big.Int
}

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	// Synchronous call
	args := MultiplyArgs{*big.NewInt(7), *big.NewInt(8)}
	var reply big.Int
	err = client.Call("Arith.Multiply", &args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %s*%s=%s\n", args.A.String(), args.B.String(), reply.String())

}
