package main

type NodeConnection struct {
	domain string
	port   string
}

var nodeList = []NodeConnection{
	{
		domain: "localhost",
		port:   "10000",
	},
	{
		domain: "localhost",
		port:   "10001",
	},
	{
		domain: "localhost",
		port:   "10002",
	},
}
