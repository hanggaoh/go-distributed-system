package cluster

import (
	"go-distributed-system/pkg/util"
)

// Message is a simple message struct.
type Message struct {
	Content string
}

// Broadcast sends a message to all nodes in the cluster.
func Broadcast(logger *util.Logger, message Message) {
	logger.Printf("Broadcasting message: %s", message.Content)
	// TODO: Implement actual broadcast logic
}
