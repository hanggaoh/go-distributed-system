package cluster

import (
	"go-distributed-system/pkg/util"
	"math/rand"
	"time"
)

// NodeState represents the state of a node in the cluster.
type NodeState int

const (
	Follower NodeState = iota
	Candidate
	Leader
)

// Node represents a node in the cluster.
type Node struct {
	ID      int
	State   NodeState
	logger  *util.Logger
	Timeout *time.Timer
}

// NewNode creates a new node.
func NewNode(id int, logger *util.Logger) *Node {
	return &Node{
		ID:     id,
		State:  Follower,
		logger: logger,
	}
}

// StartElection starts the leader election process.
func (n *Node) StartElection() {
	n.State = Candidate
	n.logger.Printf("Node %d starting election", n.ID)
	// TODO: Implement election logic
	n.ResetTimeout()
}

// Run starts the node's main loop.
func (n *Node) Run() {
	n.logger.Printf("Node %d starting up", n.ID)
	n.ResetTimeout()

	for {
		select {
		case <-n.Timeout.C:
			n.StartElection()
		}
	}
}

// ResetTimeout resets the node's election timeout.
func (n *Node) ResetTimeout() {
	duration := time.Duration(150+rand.Intn(150)) * time.Millisecond
	if n.Timeout == nil {
		n.Timeout = time.NewTimer(duration)
	} else {
		n.Timeout.Reset(duration)
	}
}
