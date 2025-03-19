package client // import "go.khulnasoft.com/client"

import (
	"context"
	"net/url"

	"go.khulnasoft.com/api/types/checkpoint"
)

// CheckpointDelete deletes the checkpoint with the given name from the given container
func (cli *Client) CheckpointDelete(ctx context.Context, containerID string, options checkpoint.DeleteOptions) error {
	containerID, err := trimID("container", containerID)
	if err != nil {
		return err
	}

	query := url.Values{}
	if options.CheckpointDir != "" {
		query.Set("dir", options.CheckpointDir)
	}

	resp, err := cli.delete(ctx, "/containers/"+containerID+"/checkpoints/"+options.CheckpointID, query, nil)
	ensureReaderClosed(resp)
	return err
}
