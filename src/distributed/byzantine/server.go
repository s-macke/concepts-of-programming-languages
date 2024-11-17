package main

import (
	"fmt"
	"net/http"
)

func (node *Node) RunServer() {
	http.HandleFunc("/heartbeat", func(w http.ResponseWriter, r *http.Request) {
		node.SignalHeartbeat()
		_, _ = fmt.Fprintf(w, "Node %d\n", node.nodeId)
	})
	http.HandleFunc("/vote", func(w http.ResponseWriter, r *http.Request) {
		if node.state != LEADER {
			_, _ = fmt.Fprintf(w, "yes")
			node.SwitchState(FOLLOWER)
		} else {
			_, _ = fmt.Fprintf(w, "no")
		}
	})
	port := nodeList[node.nodeId].port
	logger.Info("Start on port %s ...", port)
	err := http.ListenAndServe("localhost:"+port, nil)
	if err != nil {
		panic(err)
	}
	panic("Stop server")
}
