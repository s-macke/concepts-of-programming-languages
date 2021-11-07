package main

import (
	"fmt"
	"log"
	"math/big"
	"net/rpc"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	// Synchronous call
	args := &MultiplyArgs{*big.NewInt(7), *big.NewInt(8)}
	var reply big.Int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %s*%s=%s", args.A.String(), args.B.String(), reply.String())

}
