# Go Distributed System (Leader Election + Messaging)

A lightweight distributed system demo implemented in Go. Each node runs as a container, participates in leader election, and communicates with other nodes for message passing.

## 🚀 Features
- Implemented in Go using goroutines and channels.
- Leader election logic inspired by Raft (simplified).
- Broadcast messaging across nodes.
- Containerized and deployed to Kubernetes with multiple replicas.
- Logs show leader election events and message propagation.

## 📂 Project Structure
```aiexclude
go-distributed-system/
│── cmd/
│   └── node/
│       └── main.go         # Node entrypoint (leader election + message passing)
│
│── pkg/
│   ├── cluster/
│   │   ├── election.go     # Leader election logic
│   │   └── messaging.go    # Simple message broadcast
│   └── util/
│       └── logger.go       # Logging helper
│
│── Dockerfile              # Containerize node
│── k8s/
│   ├── deployment.yaml     # Run multiple replicas in Kubernetes
│   ├── service.yaml        # ClusterIP service for intra-node communication
│   └── configmap.yaml      # Config for cluster size / settings
│
│── README.md
```

