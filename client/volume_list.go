package client // import "go.khulnasoft.com/client"

import (
	"context"
	"encoding/json"
	"net/url"

	"go.khulnasoft.com/api/types/filters"
	"go.khulnasoft.com/api/types/volume"
)

// VolumeList returns the volumes configured in the docker host.
func (cli *Client) VolumeList(ctx context.Context, options volume.ListOptions) (volume.ListResponse, error) {
	query := url.Values{}

	if options.Filters.Len() > 0 {
		//nolint:staticcheck // ignore SA1019 for old code
		filterJSON, err := filters.ToParamWithVersion(cli.version, options.Filters)
		if err != nil {
			return volume.ListResponse{}, err
		}
		query.Set("filters", filterJSON)
	}
	resp, err := cli.get(ctx, "/volumes", query, nil)
	defer ensureReaderClosed(resp)
	if err != nil {
		return volume.ListResponse{}, err
	}

	var volumes volume.ListResponse
	err = json.NewDecoder(resp.Body).Decode(&volumes)
	return volumes, err
}
