package client // import "go.khulnasoft.com/client"

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	"go.khulnasoft.com/api/types/container"
	"go.khulnasoft.com/errdefs"
	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestContainerStartError(t *testing.T) {
	client := &Client{
		client: newMockClient(errorMock(http.StatusInternalServerError, "Server error")),
	}
	err := client.ContainerStart(context.Background(), "nothing", container.StartOptions{})
	assert.Check(t, is.ErrorType(err, errdefs.IsSystem))

	err = client.ContainerStart(context.Background(), "", container.StartOptions{})
	assert.Check(t, is.ErrorType(err, errdefs.IsInvalidParameter))
	assert.Check(t, is.ErrorContains(err, "value is empty"))

	err = client.ContainerStart(context.Background(), "    ", container.StartOptions{})
	assert.Check(t, is.ErrorType(err, errdefs.IsInvalidParameter))
	assert.Check(t, is.ErrorContains(err, "value is empty"))
}

func TestContainerStart(t *testing.T) {
	expectedURL := "/containers/container_id/start"
	client := &Client{
		client: newMockClient(func(req *http.Request) (*http.Response, error) {
			if !strings.HasPrefix(req.URL.Path, expectedURL) {
				return nil, fmt.Errorf("Expected URL '%s', got '%s'", expectedURL, req.URL)
			}
			// we're not expecting any payload, but if one is supplied, check it is valid.
			if req.Header.Get("Content-Type") == "application/json" {
				var startConfig interface{}
				if err := json.NewDecoder(req.Body).Decode(&startConfig); err != nil {
					return nil, fmt.Errorf("Unable to parse json: %s", err)
				}
			}

			checkpoint := req.URL.Query().Get("checkpoint")
			if checkpoint != "checkpoint_id" {
				return nil, fmt.Errorf("checkpoint not set in URL query properly. Expected 'checkpoint_id', got %s", checkpoint)
			}

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewReader([]byte(""))),
			}, nil
		}),
	}

	err := client.ContainerStart(context.Background(), "container_id", container.StartOptions{CheckpointID: "checkpoint_id"})
	if err != nil {
		t.Fatal(err)
	}
}
