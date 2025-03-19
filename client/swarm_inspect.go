package client // import "go.khulnasoft.com/client"

import (
	"context"
	"encoding/json"

	"go.khulnasoft.com/api/types/swarm"
)

// SwarmInspect inspects the swarm.
func (cli *Client) SwarmInspect(ctx context.Context) (swarm.Swarm, error) {
	resp, err := cli.get(ctx, "/swarm", nil, nil)
	defer ensureReaderClosed(resp)
	if err != nil {
		return swarm.Swarm{}, err
	}

	var response swarm.Swarm
	err = json.NewDecoder(resp.Body).Decode(&response)
	return response, err
}
