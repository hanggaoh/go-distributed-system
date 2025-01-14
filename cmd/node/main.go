package main

import (
	"go-distributed-system/pkg/cluster"
	"go-distributed-system/pkg/util"
	"os"
	"strconv"
)

func main() {
	logger := util.NewLogger("[node] ")

	nodeID, err := strconv.Atoi(os.Getenv("NODE_ID"))
	if err != nil {
		logger.Fatalf("Invalid NODE_ID: %v", err)
	}

	node := cluster.NewNode(nodeID, logger)
	go node.Run()

	// Wait indefinitely
	select {}
}
