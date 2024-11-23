package main

import (
	"time"
)

func main() {
	nodeId := GetNodeId()
	node := NewNode(nodeId)
	InitLogger(node)
	node.Start()

	time.Sleep(30 * time.Second)
	if node.state == LEADER {
		quorum := node.Attack("midnight")
		if quorum >= 1 {
			logger.Info("Attack successful")
		} else {
			logger.Info("Attack failed")
		}
	}

	// wait forever
	select {}
}
