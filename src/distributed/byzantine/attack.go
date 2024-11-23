package main

import (
	"net/http"
	"strings"
)

func (node *Node) Attack(when string) int {
	quorum := 0
	for nodeIdx := 0; nodeIdx < len(nodeList); nodeIdx++ {
		if nodeIdx == node.nodeId {
			continue
		}
		resp, err := http.Post("http://"+nodeList[nodeIdx].domain+":"+nodeList[nodeIdx].port+"/attack", "text/plain", strings.NewReader(when))
		if err == nil && resp.StatusCode == http.StatusOK {
			quorum++
		}
	}
	return quorum
}
