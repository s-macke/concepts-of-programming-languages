package main

import (
	"github.com/coreos/etcd/raft/raftpb"
	"math/rand"
	"sync"
	"time"
)

// a network
type network interface {
	// drop message at given rate (1.0 drops all messages)
	drop(from, to uint64, rate float64)
	// delay message for (0, d] randomly at given rate (1.0 delay all messages)
	// do we need rate here?
	delay(from, to uint64, d time.Duration, rate float64)
	disconnect(id uint64)
	connect(id uint64)
	// heal heals the network
	heal()
	nodeNetwork(id uint64) iface
}

type raftNetwork struct {
	rand         *rand.Rand
	mu           sync.Mutex
	disconnected map[uint64]bool
	dropmap      map[conn]float64
	delaymap     map[conn]delay
	recvQueues   map[uint64]chan raftpb.Message
}

type conn struct {
	from, to uint64
}

type delay struct {
	d    time.Duration
	rate float64
}

func newRaftNetwork(nodes ...uint64) *raftNetwork {
	pn := &raftNetwork{
		rand:         rand.New(rand.NewSource(1)),
		recvQueues:   make(map[uint64]chan raftpb.Message),
		dropmap:      make(map[conn]float64),
		delaymap:     make(map[conn]delay),
		disconnected: make(map[uint64]bool),
	}

	for _, n := range nodes {
		pn.recvQueues[n] = make(chan raftpb.Message, 1024)
	}
	return pn
}

func (rn *raftNetwork) nodeNetwork(id uint64) iface {
	return &nodeNetwork{id: id, raftNetwork: rn}
}

func (rn *raftNetwork) send(m raftpb.Message) {
	rn.mu.Lock()
	to := rn.recvQueues[m.To]
	if rn.disconnected[m.To] {
		to = nil
	}
	drop := rn.dropmap[conn{m.From, m.To}]
	dl := rn.delaymap[conn{m.From, m.To}]
	rn.mu.Unlock()

	if to == nil {
		return
	}
	if drop != 0 && rn.rand.Float64() < drop {
		return
	}
	// TODO: shall we dl without blocking the send call?
	if dl.d != 0 && rn.rand.Float64() < dl.rate {
		rd := rn.rand.Int63n(int64(dl.d))
		time.Sleep(time.Duration(rd))
	}

	// use marshal/unmarshal to copy message to avoid data race.
	b, err := m.Marshal()
	if err != nil {
		panic(err)
	}

	var cm raftpb.Message
	err = cm.Unmarshal(b)
	if err != nil {
		panic(err)
	}

	select {
	case to <- cm:
	default:
		// drop messages when the receiver queue is full.
	}
}

func (rn *raftNetwork) recvFrom(from uint64) chan raftpb.Message {
	rn.mu.Lock()
	fromc := rn.recvQueues[from]
	if rn.disconnected[from] {
		fromc = nil
	}
	rn.mu.Unlock()

	return fromc
}

func (rn *raftNetwork) drop(from, to uint64, rate float64) {
	rn.mu.Lock()
	defer rn.mu.Unlock()
	rn.dropmap[conn{from, to}] = rate
}

func (rn *raftNetwork) delay(from, to uint64, d time.Duration, rate float64) {
	rn.mu.Lock()
	defer rn.mu.Unlock()
	rn.delaymap[conn{from, to}] = delay{d, rate}
}

func (rn *raftNetwork) heal() {
	rn.mu.Lock()
	defer rn.mu.Unlock()
	rn.dropmap = make(map[conn]float64)
	rn.delaymap = make(map[conn]delay)
}

func (rn *raftNetwork) disconnect(id uint64) {
	rn.mu.Lock()
	defer rn.mu.Unlock()
	rn.disconnected[id] = true
}

func (rn *raftNetwork) connect(id uint64) {
	rn.mu.Lock()
	defer rn.mu.Unlock()
	rn.disconnected[id] = false
}
