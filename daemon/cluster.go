package daemon // import "go.khulnasoft.com/daemon"

import (
	"go.khulnasoft.com/api/types/filters"
	"go.khulnasoft.com/api/types/network"
	lncluster "go.khulnasoft.com/libnetwork/cluster"
)

// Cluster is the interface for go.khulnasoft.com/daemon/cluster.(*Cluster).
type Cluster interface {
	ClusterStatus
	NetworkManager
	SendClusterEvent(event lncluster.ConfigEventType)
}

// ClusterStatus interface provides information about the Swarm status of the Cluster
type ClusterStatus interface {
	IsAgent() bool
	IsManager() bool
}

// NetworkManager provides methods to manage networks
type NetworkManager interface {
	GetNetwork(input string) (network.Inspect, error)
	GetNetworks(filters.Args) ([]network.Inspect, error)
	RemoveNetwork(input string) error
}
