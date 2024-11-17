package main

import (
	"fmt"
	"os"
	"strconv"
)

func PrintUsage() {
	fmt.Printf("Usage: %s <nodeId>\n", os.Args[0])
	os.Exit(0)
}

func GetNodeId() int {
	if len(os.Args) != 2 {
		PrintUsage()
	}
	nodeId, err := strconv.Atoi(os.Args[1])
	if err != nil {
		PrintUsage()
	}
	return nodeId
}
