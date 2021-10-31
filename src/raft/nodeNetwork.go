package main

import "github.com/coreos/etcd/raft/raftpb"

// a network interface
type iface interface {
	send(m raftpb.Message)
	recv() chan raftpb.Message
	disconnect()
	connect()
}

type nodeNetwork struct {
	id uint64
	*raftNetwork
}

func (nt *nodeNetwork) connect() {
	nt.raftNetwork.connect(nt.id)
}

func (nt *nodeNetwork) disconnect() {
	nt.raftNetwork.disconnect(nt.id)
}

func (nt *nodeNetwork) send(m raftpb.Message) {
	nt.raftNetwork.send(m)
}

func (nt *nodeNetwork) recv() chan raftpb.Message {
	return nt.recvFrom(nt.id)
}
