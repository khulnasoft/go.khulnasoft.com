package daemon

import (
	"strings"

	"go.khulnasoft.com/libnetwork"
)

// getEndpointInNetwork returns the container's endpoint to the provided network.
func getEndpointInNetwork(name string, n *libnetwork.Network) (*libnetwork.Endpoint, error) {
	endpointName := strings.TrimPrefix(name, "/")
	return n.EndpointByName(endpointName)
}
