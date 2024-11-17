package main

import (
	"time"
)

type Node struct {
	state               StateEnum
	nodeId              int
	hearbeatReceived    chan struct{}
	switchStateReceived chan StateEnum
}

func NewNode(nodeId int) *Node {
	n := &Node{
		state:  FOLLOWER,
		nodeId: nodeId,
	}
	n.hearbeatReceived = make(chan struct{})
	n.switchStateReceived = make(chan StateEnum)
	return n
}

func (node *Node) Start() {
	go node.RunServer()
	go node.StateHandler()
}

func (node *Node) SignalHeartbeat() {
	node.hearbeatReceived <- struct{}{}
}

func (node *Node) SwitchState(enum StateEnum) {
	node.switchStateReceived <- enum
}

func (node *Node) StateHandler() {
	for {
		select {
		case <-time.After(5 * time.Second):
			if node.state == LEADER {
				continue
			}
			if node.state == FOLLOWER {
				logger.Info("No heartbeat received")
				node.state.Switch(CANDIDATE)
				go Vote(node)
			}
			break

		case newState := <-node.switchStateReceived:
			node.state.Switch(newState)
			break

		case <-node.hearbeatReceived:
			logger.Info("heartbeat received")
			// voting has been done and a new leader has been elected
			// No matter of the node, you are a follower now
			// Race Condition might happen here
			node.state.Switch(FOLLOWER)
			break
		}

	}
}
