package client // import "go.khulnasoft.com/client"

import (
	"context"
	"encoding/json"
	"net/url"

	"go.khulnasoft.com/api/types"
	"go.khulnasoft.com/api/types/filters"
	"go.khulnasoft.com/api/types/swarm"
)

// ConfigList returns the list of configs.
func (cli *Client) ConfigList(ctx context.Context, options types.ConfigListOptions) ([]swarm.Config, error) {
	if err := cli.NewVersionError(ctx, "1.30", "config list"); err != nil {
		return nil, err
	}
	query := url.Values{}

	if options.Filters.Len() > 0 {
		filterJSON, err := filters.ToJSON(options.Filters)
		if err != nil {
			return nil, err
		}

		query.Set("filters", filterJSON)
	}

	resp, err := cli.get(ctx, "/configs", query, nil)
	defer ensureReaderClosed(resp)
	if err != nil {
		return nil, err
	}

	var configs []swarm.Config
	err = json.NewDecoder(resp.Body).Decode(&configs)
	return configs, err
}
