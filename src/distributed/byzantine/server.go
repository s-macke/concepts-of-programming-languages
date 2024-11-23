package main

import (
	"fmt"
	"io"
	"net/http"
)

func (node *Node) RunServer() {
	http.HandleFunc("/attack", func(w http.ResponseWriter, r *http.Request) {
		b, _ := io.ReadAll(r.Body)
		fmt.Println("Attack at " + string(b))
		w.WriteHeader(http.StatusOK)
	})
	http.HandleFunc("/commit", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Attack ")
		w.WriteHeader(http.StatusOK)
	})

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
