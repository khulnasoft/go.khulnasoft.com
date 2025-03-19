package network

import (
	"go.khulnasoft.com/api/types/container"
	"go.khulnasoft.com/api/types/network"
)

const defaultNetwork = network.NetworkNat

func isPreDefined(network string) bool {
	return !container.NetworkMode(network).IsUserDefined()
}
