package main

func main() {
	nodeId := GetNodeId()
	node := NewNode(nodeId)
	InitLogger(node)
	node.Start()

	// wait forever
	select {}
}
