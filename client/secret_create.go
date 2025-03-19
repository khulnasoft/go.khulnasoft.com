package client // import "go.khulnasoft.com/client"

import (
	"context"
	"encoding/json"

	"go.khulnasoft.com/api/types"
	"go.khulnasoft.com/api/types/swarm"
)

// SecretCreate creates a new secret.
func (cli *Client) SecretCreate(ctx context.Context, secret swarm.SecretSpec) (types.SecretCreateResponse, error) {
	if err := cli.NewVersionError(ctx, "1.25", "secret create"); err != nil {
		return types.SecretCreateResponse{}, err
	}
	resp, err := cli.post(ctx, "/secrets/create", nil, secret, nil)
	defer ensureReaderClosed(resp)
	if err != nil {
		return types.SecretCreateResponse{}, err
	}

	var response types.SecretCreateResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	return response, err
}
