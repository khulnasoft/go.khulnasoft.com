package client // import "go.khulnasoft.com/client"

import (
	"context"
	"net/url"

	"go.khulnasoft.com/api/types"
)

// PluginDisable disables a plugin
func (cli *Client) PluginDisable(ctx context.Context, name string, options types.PluginDisableOptions) error {
	name, err := trimID("plugin", name)
	if err != nil {
		return err
	}
	query := url.Values{}
	if options.Force {
		query.Set("force", "1")
	}
	resp, err := cli.post(ctx, "/plugins/"+name+"/disable", query, nil, nil)
	ensureReaderClosed(resp)
	return err
}
