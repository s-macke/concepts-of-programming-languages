package main

import (
	"context"
	"fmt"
	"github.com/s-macke/concepts-of-programming-languages/src/distributed/greeter/proto"
	"google.golang.org/grpc"
	"log"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	ctx := context.Background()
	c := greeter.NewGreeterClient(conn)
	reply, err := c.SayHello(ctx, &greeter.HelloRequest{Name: "world"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	fmt.Printf("Received '%s'\n", reply.Message)
}
