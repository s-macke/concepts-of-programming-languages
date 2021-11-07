package main

import (
	"log"
	"math/big"
	"net/http"
	"net/rpc"
)

type MultiplyArgs struct {
	A, B big.Int
}

type Arith int

func (t *Arith) Multiply(args *MultiplyArgs, reply *big.Int) error {
	*reply = *args.A.Mul(&args.A, &args.B)
	return nil
}

func main() {
	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	log.Fatal(http.ListenAndServe(":1234", nil))
}
