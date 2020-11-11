package node

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"net/http"
	"time"
)

type mode int

const (
	follower mode = iota
	candidate
	leader
)

const (
	electionTimeoutMinMs = 1500                    // min. milliseconds to wait for a heartbeat (starts election when timing out)
	electionTimeoutMaxMs = 3000                    // max. milliseconds to wait for a heartbeat (starts election when timing out)
	startupWait          = 1000 * time.Millisecond // time to wait for connection to other nodes
	voteTimeout          = 1500 * time.Millisecond // time to wait for votes from other clients
	heartbeatTimeout     = 500 * time.Millisecond  // time to wait for a client to accept a heartbeat
	heartbeatTicker      = 1000 * time.Millisecond // time between two heartbeats
)

type Entry struct {
	term int
	data []byte
}

type Node struct {
	UnimplementedNodeServer
	id          int          // id of this node
	basePort    int          // port of node with id 0
	nodeCount   int          // total number of nodes
	nodeClients []NodeClient // clients to reach other nodes
	mode        mode         // current mode of this node

	currentTerm   int         // current term
	votedFor      int         // node we voted for in the current election; set to -1 if not yet voted
	log           []Entry     // log entries
	commitIndex   int         // index of highest log entry known to be committed
	lastApplied   int         // index of highest log entry applied to state machine
	electionTimer *time.Timer // when timeout is reached, a new election begins

	// leader-specific
	heartbeatTicker *time.Ticker // for periodically sending heartbeats
	nextIndex       []int        // for each server, index of the next log entry to send to that server
	matchIndex      []int        // for each server, index of highest log entry known to be replicated on server
}

func NewNode(id int, basePort int, nodeCount int) *Node {
	return &Node{
		id:          id,
		basePort:    basePort,
		nodeCount:   nodeCount,
		mode:        follower,
		votedFor:    -1,
		commitIndex: 0,
		lastApplied: 0,
		log:         []Entry{},
	}
}

// Starts an HTTP server that listens for data requests, send by e.g. curl -d 'hello' ':11000'
func (n *Node) startDataServer() {
	http.HandleFunc("/", n.acceptData)
	port := n.basePort + 1000 + n.id
	log.Printf("[%d] starting HTTP listener at port %d\n", n.id, port)
	go func() {
		err := http.ListenAndServe(fmt.Sprintf("localhost:%d", port), nil)
		if err != nil {
			log.Fatalf("[%d] failed to start HTTP server: %v", n.id, err)
		}
	}()
}

// Starts the gRPC server
func (n *Node) startGRPCServer() {
	port := n.basePort + n.id
	log.Printf("[%d] starting TCP listener on port %d", n.id, port)
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("[%d] failed to listen: %v", n.id, err)
	}

	log.Printf("[%d] starting gRPC server", n.id)
	s := grpc.NewServer()
	RegisterNodeServer(s, n)
	go func() {
		err := s.Serve(lis)
		if err != nil {
			log.Fatalf("[%d] failed to serve: %v", n.id, err)
		}
	}()
}

// Establishes gRPC connection to other nodes
func (n *Node) connectOtherNodes() {
	log.Printf("[%d] connecting to other nodes", n.id)
	n.nodeClients = make([]NodeClient, n.nodeCount)
	for i := 0; i < n.nodeCount; i++ {
		if i == n.id {
			continue
		}
		conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", n.basePort+i), grpc.WithInsecure())
		if err != nil {
			log.Fatalf("[%d] failed to connect to node %d: %v", n.id, i, err)
		}
		n.nodeClients[i] = NewNodeClient(conn)
	}
}

// Starts servers and timers and connects to other nodes
func (n *Node) Start() {
	rand.Seed(time.Now().UTC().UnixNano()) // otherwise all nodes will have the same deterministic "random" numbers

	n.startDataServer()
	n.startGRPCServer()
	time.Sleep(startupWait)
	n.connectOtherNodes()

	log.Printf("[%d] starting timers and running main loop", n.id)
	n.electionTimer = time.NewTimer(n.randomElectionTimeout())
	n.heartbeatTicker = time.NewTicker(heartbeatTicker) // we have to initialize the ticker for select {} to work
	n.heartbeatTicker.Stop()
	n.run()
}

// Main loop
func (n *Node) run() {
	for {
		select {
		case <-n.electionTimer.C:
			n.runForLeader()
		case <-n.heartbeatTicker.C:
			n.sendHeartbeat()
		}
	}
}

// Send RequestVote to all other nodes and count votes
func (n *Node) runForLeader() {
	log.Printf("[%d] election timeout reached, running for leader", n.id)
	n.mode = candidate
	n.currentTerm += 1
	n.votedFor = n.id // we vote for ourselves ...
	votes := 1        // ... and add our vote
	n.electionTimer.Reset(n.randomElectionTimeout())
	voteBox := make(chan bool, n.nodeCount-1)
	n.runForOtherNodesAsync(func(i int) {
		ctx, cancel := context.WithTimeout(context.Background(), voteTimeout)
		defer cancel()
		r, err := n.nodeClients[i].RequestVote(ctx, &RequestVoteRequest{
			Term:        int64(n.currentTerm),
			CandidateId: int64(n.id),
		})
		if err != nil {
			log.Printf("[%d] could not request vote from %d: %v", n.id, i, err)
			voteBox <- false
		} else {
			voteBox <- r.Success
		}
	})
	// wait for votes and count
	for i := 0; i < n.nodeCount-1 && n.mode == candidate; i++ {
		if <-voteBox { // all vote requests will eventually time out and respond because of the context
			votes += 1
		}
		if votes > n.nodeCount/2 {
			// we did it; let us make node land great again!
			log.Printf("[%d] voted for leader with %d votes", n.id, votes)
			n.becomeLeader()
			break
		}
	}
}

// Switch to leader mode by stopping the election timer and starting to send heartbeats
func (n *Node) becomeLeader() {
	log.Printf("[%d] becoming leader", n.id)
	n.mode = leader
	n.electionTimer.Stop()
	n.heartbeatTicker.Reset(heartbeatTicker)
	n.sendHeartbeat() // immediately send heartbeat because the ticker waits for the first tick
}

// Switch to follower mode by stopping the heartbeat timer and starting the election timer
func (n *Node) resignFromLeader() {
	log.Printf("[%d] resigning from being leader", n.id)
	n.mode = follower
	n.electionTimer.Reset(n.randomElectionTimeout())
	n.heartbeatTicker.Stop()
}

// Send a heartbeat to all other nodes
func (n *Node) sendHeartbeat() {
	log.Printf("[%d] sending heartbeats", n.id)
	n.runForOtherNodesAsync(func(i int) {
		ctx, cancel := context.WithTimeout(context.Background(), heartbeatTimeout)
		defer cancel()
		r, err := n.nodeClients[i].AppendEntries(ctx, &AppendEntriesRequest{
			Term:         int64(n.currentTerm),
			LeaderId:     int64(n.id),
			PrevLogIndex: 0,   // TODO index of log entry immediately preceding new ones
			PrevLogTerm:  0,   // TODO term of prevLogIndex entry
			Entries:      nil, // nil means heartbeat
			LeaderCommit: int64(n.commitIndex),
		})
		if err != nil {
			log.Printf("[%d] could not send heartbeat to %d: %v", n.id, i, err)
		}
		if !r.Success && int(r.Term) > n.currentTerm {
			log.Printf("[%d] node %d is already in term %d but we are still in %d", n.id, i, r.Term, n.currentTerm)
			n.resignFromLeader()
		}
	})
}

func (n *Node) acceptData(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Printf("[%d] received %d bytes of data\n", n.id, len(bytes))
	if n.mode == leader {
		// save data and pass to followers
	} else {
		// forward data to leader
		// TODO what if we are candidate?

	}
	w.WriteHeader(http.StatusNoContent)
}

func (n *Node) RequestVote(ctx context.Context, request *RequestVoteRequest) (*RequestVoteResponse, error) {
	if int(request.Term) < n.currentTerm {
		log.Printf("[%d] received heartbeat or entries with old term, returning false", n.id)
		return &RequestVoteResponse{Success: false, Term: int64(n.currentTerm)}, nil
	}
	// TODO also check if candidate’s log is at least as up-to-date as receiver’s log
	if n.votedFor == -1 || n.votedFor == int(request.CandidateId) {
		log.Printf("[%d] voting for %d", n.id, request.CandidateId)
		n.votedFor = int(request.CandidateId)
		return &RequestVoteResponse{Success: true, Term: int64(n.currentTerm)}, nil
	} else {
		log.Printf("[%d] not voting for %d, already voted for %d", n.id, request.CandidateId, n.votedFor)
		return &RequestVoteResponse{Success: false, Term: int64(n.currentTerm)}, nil
	}
}

func (n *Node) AppendEntries(ctx context.Context, request *AppendEntriesRequest) (*AppendEntriesResponse, error) {
	if int(request.Term) < n.currentTerm {
		log.Printf("[%d] received heartbeat or entries with old term, returning false", n.id)
		return &AppendEntriesResponse{Success: false, Term: int64(n.currentTerm)}, nil
	}
	// TODO reply false if log doesn’t contain an entry at prevLogIndex whose currentTerm matches prevLogTerm
	if request.Entries == nil {
		log.Printf("[%d] received heartbeat from leader", n.id)
		n.electionTimer.Reset(n.randomElectionTimeout())
		if n.mode != follower {
			n.mode = follower // this will also stop a running election
			n.votedFor = -1
			if n.heartbeatTicker != nil {
				n.heartbeatTicker.Stop()
			}
		}
		return &AppendEntriesResponse{Success: true, Term: int64(n.currentTerm)}, nil
	} else {
		log.Printf("[%d] received new entries from leader", n.id)
		// TODO if an existing entry conflicts with a new one (same index but different terms), delete the existing entry and all that follow it
		// TODO append new entries not already in the log
		// TODO if leaderCommit > commitIndex, set commitIndex = min(leaderCommit, index of last new entry)
		return nil, status.Errorf(codes.Unimplemented, "method AppendEntries not implemented")
	}
}

// Helper function that starts a goroutine for each node client except itself
func (n *Node) runForOtherNodesAsync(f func(nodeClientIndex int)) {
	for i := range n.nodeClients {
		if i == n.id {
			continue
		}
		go f(i)
	}
}

// Generate a new random timout until starting a new election between electionTimeoutMinMs and electionTimeoutMaxMs
func (n *Node) randomElectionTimeout() time.Duration {
	ms := rand.Intn(electionTimeoutMaxMs-electionTimeoutMinMs) + electionTimeoutMinMs
	log.Printf("[%d] timeout until new election is now %dms", n.id, ms)
	return time.Duration(ms) * time.Millisecond
}
