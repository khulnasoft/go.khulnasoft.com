package client // import "go.khulnasoft.com/client"

import (
	"context"
	"encoding/json"
	"net/url"

	"go.khulnasoft.com/api/types"
	"go.khulnasoft.com/api/types/filters"
	"go.khulnasoft.com/api/types/swarm"
)

// TaskList returns the list of tasks.
func (cli *Client) TaskList(ctx context.Context, options types.TaskListOptions) ([]swarm.Task, error) {
	query := url.Values{}

	if options.Filters.Len() > 0 {
		filterJSON, err := filters.ToJSON(options.Filters)
		if err != nil {
			return nil, err
		}

		query.Set("filters", filterJSON)
	}

	resp, err := cli.get(ctx, "/tasks", query, nil)
	defer ensureReaderClosed(resp)
	if err != nil {
		return nil, err
	}

	var tasks []swarm.Task
	err = json.NewDecoder(resp.Body).Decode(&tasks)
	return tasks, err
}
