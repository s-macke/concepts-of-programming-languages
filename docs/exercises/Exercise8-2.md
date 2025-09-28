# Exercise 8 - Distributed Consensus and the Raft Algorithm

If you do not finish during the lecture period, please finish it as homework.

## Exercise 8.1 - Read the Raft Paper

Read the [Raft Paper](https://raft.github.io/raft.pdf):

* Section 1
* Section 2
* Section 4
* Section 5 (until including 5.3)

## Exercise 7.2 - Coordinate an attack on a fortified city with Raft

Three armies are attacking a fortified city. The armies are commanded by generals. 
The generals need to agree on the time to attack the city.
Only if at least two generals agree on the time, the attack will be successful.
They agree to organize via a simplified protocol based on the Raft algorithm.

### Read arguments

On start of the program the node id should be read as argument from the console.
```raft --node 0```

Use the `flag` package to read the arguments.

### Define list of servers

Define a list of servers which correspond to the three armies.
The list can be hardcoded in the program and should at least
contain the ports of the servers.

```
var nodePortList = []string{
	"10000", "10001", "10002",
}
```

### Heartbeat Endpoint

Implement the heartbeat mechanism. 

1. Open a server on the port defined by the node id and 
implement a heartbeat endpoint.
Test via curl, httpie, wget or a browser 


### State

Define an [enum](https://gobyexample.com/enums] containin) for the node state

* Follower
* Candidate
* Leader

Implement the [stringer](https://go.dev/tour/methods/17) interface for the enum.
Write a function to switch the state of the node.
It should apply the following rules, otherwise stop the program.

* A leader can never change
* A candidate can only change to leader or follower
* A follower can only change to candidate

### Heartbeat client
Define node 0 as the leader, the others as followers.
The leader should send heartbeats to the other nodes in 
intervals of 1 second.



**Examples:**

* Testable key value store using interfaces and goroutines:  
  https://github.com/0xqab/concepts-of-programming-languages/tree/master/dp/kvstore/core/raft
* Networkable implementation using gRPC (voting only):  
  https://github.com/0xqab/concepts-of-programming-languages/tree/master/dp/raft
* Highly abstracted from the paper:  
  https://github.com/hashicorp/raft
* More implementations:  
  github.com/search?q=go-raft
