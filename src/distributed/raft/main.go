package main

import (
	"flag"
	"github.com/s-macke/concepts-of-programming-languages/dp/raft/node"
)

var (
	id       = flag.Int("id", 0, "Node identifier [0...n]")
	count    = flag.Int("count", 5, "Number of started nodes")
	basePort = flag.Int("baseport", 10000, "Port number of node with id 0")
)

func main() {
	flag.Parse()
	n := node.NewNode(*id, *basePort, *count)
	n.Start()
}
