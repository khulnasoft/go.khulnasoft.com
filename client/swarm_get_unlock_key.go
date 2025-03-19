package client // import "go.khulnasoft.com/client"

import (
	"context"
	"encoding/json"

	"go.khulnasoft.com/api/types"
)

// SwarmGetUnlockKey retrieves the swarm's unlock key.
func (cli *Client) SwarmGetUnlockKey(ctx context.Context) (types.SwarmUnlockKeyResponse, error) {
	resp, err := cli.get(ctx, "/swarm/unlockkey", nil, nil)
	defer ensureReaderClosed(resp)
	if err != nil {
		return types.SwarmUnlockKeyResponse{}, err
	}

	var response types.SwarmUnlockKeyResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	return response, err
}
