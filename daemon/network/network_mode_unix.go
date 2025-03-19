//go:build !windows

package network

import (
	"go.khulnasoft.com/api/types/container"
	"go.khulnasoft.com/api/types/network"
)

const defaultNetwork = network.NetworkBridge

func isPreDefined(network string) bool {
	n := container.NetworkMode(network)
	return n.IsBridge() || n.IsHost() || n.IsNone() || n.IsDefault()
}
