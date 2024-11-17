package main

import (
	"io"
	"net/http"
	"time"
)

func Vote(node *Node) {
	for {
		if node.state != CANDIDATE {
			return
		}
		logger.Info("Start Vote")
		votes := 1
		for i := 0; i < len(nodeList); i++ {
			if i != node.nodeId {
				if CallVote(i) {
					votes++
				}
			}
		}
		if votes >= len(nodeList)/2+1 {
			logger.Info("Vote won")
			if node.state != CANDIDATE {
				return
			}
			node.SwitchState(LEADER)
			StartHeartbeatCalls(node.nodeId)
		}
		logger.Info("Vote lost")
		time.Sleep(4 * time.Second)
	}
}

func CallVote(nodeIdx int) bool {
	resp, err := http.Get("http://" + nodeList[nodeIdx].domain + ":" + nodeList[nodeIdx].port + "/vote")
	if err != nil {
		logger.Info("Node %d is down", nodeIdx)
		return false
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	logger.Info("Node %d votes %s", nodeIdx, string(body))
	return string(body) == "yes"
}
