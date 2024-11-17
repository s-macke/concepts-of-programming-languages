package main

type StateEnum int

const (
	// FOLLOWER is the start state.
	FOLLOWER StateEnum = iota
	// CANDIDATE is the candidate state.
	CANDIDATE
	// LEADER is the leader state.
	LEADER
)

// Stringer interface.
func (s StateEnum) String() string {
	switch s {
	case FOLLOWER:
		return "FOLLOWER"
	case LEADER:
		return "LEADER"
	case CANDIDATE:
		return "CANDIDATE"
	default:
		panic("Unknown State! Should never happen.")
	}
}

func (s *StateEnum) Switch(enum StateEnum) {
	if *s == enum {
		return
	}
	logger.Info("Node switches to state %s", enum)
	if *s == LEADER {
		panic("Cannot switch from LEADER to another state")
	}
	if *s == FOLLOWER && enum == LEADER {
		panic("Cannot switch from FOLLOWER to LEADER state")
	}
	*s = enum
}
