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

	"go.khulnasoft.com/api/types"
	"go.khulnasoft.com/api/types/swarm"
	"go.khulnasoft.com/errdefs"
	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestConfigCreateUnsupported(t *testing.T) {
	client := &Client{
		version: "1.29",
		client:  &http.Client{},
	}
	_, err := client.ConfigCreate(context.Background(), swarm.ConfigSpec{})
	assert.Check(t, is.Error(err, `"config create" requires API version 1.30, but the Docker daemon API version is 1.29`))
}

func TestConfigCreateError(t *testing.T) {
	client := &Client{
		version: "1.30",
		client:  newMockClient(errorMock(http.StatusInternalServerError, "Server error")),
	}
	_, err := client.ConfigCreate(context.Background(), swarm.ConfigSpec{})
	assert.Check(t, is.ErrorType(err, errdefs.IsSystem))
}

func TestConfigCreate(t *testing.T) {
	expectedURL := "/v1.30/configs/create"
	client := &Client{
		version: "1.30",
		client: newMockClient(func(req *http.Request) (*http.Response, error) {
			if !strings.HasPrefix(req.URL.Path, expectedURL) {
				return nil, fmt.Errorf("Expected URL '%s', got '%s'", expectedURL, req.URL)
			}
			if req.Method != http.MethodPost {
				return nil, fmt.Errorf("expected POST method, got %s", req.Method)
			}
			b, err := json.Marshal(types.ConfigCreateResponse{
				ID: "test_config",
			})
			if err != nil {
				return nil, err
			}
			return &http.Response{
				StatusCode: http.StatusCreated,
				Body:       io.NopCloser(bytes.NewReader(b)),
			}, nil
		}),
	}

	r, err := client.ConfigCreate(context.Background(), swarm.ConfigSpec{})
	if err != nil {
		t.Fatal(err)
	}
	if r.ID != "test_config" {
		t.Fatalf("expected `test_config`, got %s", r.ID)
	}
}
