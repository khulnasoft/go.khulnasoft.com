package client // import "go.khulnasoft.com/client"

import (
	"context"
	"encoding/json"

	"go.khulnasoft.com/api/types/container"
)

// ContainerUpdate updates the resources of a container.
func (cli *Client) ContainerUpdate(ctx context.Context, containerID string, updateConfig container.UpdateConfig) (container.UpdateResponse, error) {
	containerID, err := trimID("container", containerID)
	if err != nil {
		return container.UpdateResponse{}, err
	}

	resp, err := cli.post(ctx, "/containers/"+containerID+"/update", nil, updateConfig, nil)
	defer ensureReaderClosed(resp)
	if err != nil {
		return container.UpdateResponse{}, err
	}

	var response container.UpdateResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	return response, err
}
