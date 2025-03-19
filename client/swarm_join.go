package client // import "go.khulnasoft.com/client"

import (
	"context"

	"go.khulnasoft.com/api/types/swarm"
)

// SwarmJoin joins the swarm.
func (cli *Client) SwarmJoin(ctx context.Context, req swarm.JoinRequest) error {
	resp, err := cli.post(ctx, "/swarm/join", nil, req, nil)
	ensureReaderClosed(resp)
	return err
}
