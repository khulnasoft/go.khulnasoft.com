package client // import "go.khulnasoft.com/client"

import (
	"context"
	"net/url"

	"go.khulnasoft.com/api/types"
)

// NodeRemove removes a Node.
func (cli *Client) NodeRemove(ctx context.Context, nodeID string, options types.NodeRemoveOptions) error {
	nodeID, err := trimID("node", nodeID)
	if err != nil {
		return err
	}

	query := url.Values{}
	if options.Force {
		query.Set("force", "1")
	}

	resp, err := cli.delete(ctx, "/nodes/"+nodeID, query, nil)
	defer ensureReaderClosed(resp)
	return err
}
