package network

import (
	"context"
	"fmt"

	"go.khulnasoft.com/api/types/network"
	"go.khulnasoft.com/client"
	"gotest.tools/v3/assert/cmp"
)

// IsNetworkAvailable provides a comparison to check if a docker network is available
func IsNetworkAvailable(ctx context.Context, c client.NetworkAPIClient, name string) cmp.Comparison {
	return func() cmp.Result {
		networks, err := c.NetworkList(ctx, network.ListOptions{})
		if err != nil {
			return cmp.ResultFromError(err)
		}
		for _, network := range networks {
			if network.Name == name {
				return cmp.ResultSuccess
			}
		}
		return cmp.ResultFailure(fmt.Sprintf("could not find network %s", name))
	}
}

// IsNetworkNotAvailable provides a comparison to check if a docker network is not available
func IsNetworkNotAvailable(ctx context.Context, c client.NetworkAPIClient, name string) cmp.Comparison {
	return func() cmp.Result {
		networks, err := c.NetworkList(ctx, network.ListOptions{})
		if err != nil {
			return cmp.ResultFromError(err)
		}
		for _, network := range networks {
			if network.Name == name {
				return cmp.ResultFailure(fmt.Sprintf("network %s is still present", name))
			}
		}
		return cmp.ResultSuccess
	}
}
