# Exercise 8 - Distributed Programming

If you do not finish during the lecture period, please finish it as homework.

## Exercise 8.1 - Call Rest API

Call the GitHub API to receive the repositories with the most stars and the query "awesome"

https://api.github.com/search/repositories?sort=stars&order=desc&q=awesome

Use [json-to-go](https://mholt.github.io/json-to-go/) to instantly create a Go structure from an arbitrary JSON.

## Exercise 8.2 - Coordinate an attack on a fortified city with Raft

Three armies, each led by a general, plan to attack a fortified city. The armies may arrive at different times to surround the city, and none of the generals can be certain whether all the armies will arrive at all. Consequently, they have not appointed a single leader.
For the attack to succeed, at least two armies must strike simultaneously. Therefore, the generals need to coordinate and agree on a precise time to launch the attack.
To achieve this, they decide to use a simplified version of the Raft algorithm.

In the following, we will represent the armies as nodes and facilitate communication between them using REST APIs.

### Define list of servers (nodes)

Define a list of servers which correspond to the three armies.
The list can be hardcoded in the program and must at least
contain the ports of the servers.

```
var nodePortList = []string{
	"10000", "10001", "10002",
}
```

### Start program, read node id argument from console

On start of the program the node id should be read as argument from the console.
```raft --node 0```

Use the `flag` [packag](https://pkg.go.dev/flag) to read the arguments.

### State

Define an [enum](https://gobyexample.com/enums) for the node state

* Follower
* Candidate
* Leader

Write a function to switch the state of the node.
It should apply the following rules, otherwise stop the program.

* A leader can never change
* A candidate can only change to a leader or follower
* A follower can only change to a candidate

### Attack Endpoint

For now, node 0 starts as a leader, the others as followers. Hence, the leader election is skipped for now. Start with the attack coordination.

Implement an "attack" endpoint and start the server on the port defined by the port list.
The "attack" endpoint uses the REST Method POST, receives the time of the attack as string, prints the attack time on screen and returns a 200 status code.
You can test it via curl, httpie or wget.

```
curl -X POST http://localhost:10000/attack -d "midnight"
```

If the node is the leader: After a few seconds (e.g. 30 seconds) the leader should send an attack request to the other nodes to attack at the same time.
Only when the leader receives an OK response from at least one other node, the attack is successful.

Optional task: You can send a commit message to the followers to confirm the attack. You just have to implement a commit endpoint.

### Simplified Heartbeat Endpoint

Implement the heartbeat mechanism. Implement a heartbeat endpoint via a GET method. It should just return a 200 status code and turn the node into a follower.
The leader node should send a heartbeat to the other nodes in intervals of 1 second.

### Simplified Vote Endpoint

Implement a vote endpoint. 
The vote endpoint uses the REST Method GET.
If the node is the leader, it should return "no" to the other nodes, otherwise "yes" and turn into a "follower".

### Simplified Leader Election

From now on all nodes start from the follower state. No one is the leader.

- The nodes wait for heartbeats from the leader. E.g. for two seconds.(Example via the time.After function: [main.go](../../src/concurrent/channels/timeout/main.go)
- Because there is no leader, the nodes become candidates and start an election. 
- The nodes wait for a random time between 1 and 3 seconds and start a voting process.
- For the voting they send a request to all other nodes to the voting endpoint.
- If they receive a majority of "yes" votes, they become the leader and start the heartbeats.

**Examples:**

* Testable key value store using interfaces and goroutines:  
  https://github.com/0xqab/concepts-of-programming-languages/tree/master/dp/kvstore/core/raft
* Networkable implementation using gRPC (voting only):  
  https://github.com/0xqab/concepts-of-programming-languages/tree/master/dp/raft
* Highly abstracted from the paper:  
  https://github.com/hashicorp/raft
* More implementations:  
  github.com/search?q=go-raft