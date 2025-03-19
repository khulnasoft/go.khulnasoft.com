package client // import "go.khulnasoft.com/client"

import (
	"context"
	"encoding/json"
	"net/url"

	"go.khulnasoft.com/api/types/image"
)

// ImageRemove removes an image from the docker host.
func (cli *Client) ImageRemove(ctx context.Context, imageID string, options image.RemoveOptions) ([]image.DeleteResponse, error) {
	query := url.Values{}

	if options.Force {
		query.Set("force", "1")
	}
	if !options.PruneChildren {
		query.Set("noprune", "1")
	}

	resp, err := cli.delete(ctx, "/images/"+imageID, query, nil)
	defer ensureReaderClosed(resp)
	if err != nil {
		return nil, err
	}

	var dels []image.DeleteResponse
	err = json.NewDecoder(resp.Body).Decode(&dels)
	return dels, err
}
