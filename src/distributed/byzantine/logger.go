package main

import "fmt"

type Logger struct {
	state *Node
}

var logger Logger

func (logger *Logger) Info(format string, args ...any) {
	fmt.Printf("node=%d state=%9s: ", logger.state.nodeId, logger.state.state)
	fmt.Printf(format, args...)
	fmt.Printf("\n")
}

func InitLogger(state *Node) {
	logger = Logger{
		state: state,
	}
}
