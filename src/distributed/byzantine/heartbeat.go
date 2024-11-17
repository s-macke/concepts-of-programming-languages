package main

import (
	"net/http"
	"time"
)

func StartHeartbeatCalls(nodeId int) {
	for i := 0; i < len(nodeList); i++ {
		if i != nodeId {
			go RunHeartbeat(i)
		}
	}
}

func RunHeartbeat(nodeIdx int) {
	for {
		time.Sleep(1 * time.Second)
		resp, err := http.Get("http://" + nodeList[nodeIdx].domain + ":" + nodeList[nodeIdx].port + "/heartbeat")
		if err != nil {
			logger.Info("Node %d is down", nodeIdx)
			continue
		}
		logger.Info("Node %d is up", nodeIdx)
		err = resp.Body.Close()
		if err != nil {
			logger.Info("failed to close response body")
			continue
		}
	}
}
