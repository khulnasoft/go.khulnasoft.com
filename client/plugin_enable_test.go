package client // import "go.khulnasoft.com/client"

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	"go.khulnasoft.com/api/types"
	"go.khulnasoft.com/errdefs"
	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestPluginEnableError(t *testing.T) {
	client := &Client{
		client: newMockClient(errorMock(http.StatusInternalServerError, "Server error")),
	}

	err := client.PluginEnable(context.Background(), "plugin_name", types.PluginEnableOptions{})
	assert.Check(t, is.ErrorType(err, errdefs.IsSystem))

	err = client.PluginEnable(context.Background(), "", types.PluginEnableOptions{})
	assert.Check(t, is.ErrorType(err, errdefs.IsInvalidParameter))
	assert.Check(t, is.ErrorContains(err, "value is empty"))

	err = client.PluginEnable(context.Background(), "    ", types.PluginEnableOptions{})
	assert.Check(t, is.ErrorType(err, errdefs.IsInvalidParameter))
	assert.Check(t, is.ErrorContains(err, "value is empty"))
}

func TestPluginEnable(t *testing.T) {
	expectedURL := "/plugins/plugin_name/enable"

	client := &Client{
		client: newMockClient(func(req *http.Request) (*http.Response, error) {
			if !strings.HasPrefix(req.URL.Path, expectedURL) {
				return nil, fmt.Errorf("Expected URL '%s', got '%s'", expectedURL, req.URL)
			}
			if req.Method != http.MethodPost {
				return nil, fmt.Errorf("expected POST method, got %s", req.Method)
			}
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewReader([]byte(""))),
			}, nil
		}),
	}

	err := client.PluginEnable(context.Background(), "plugin_name", types.PluginEnableOptions{})
	if err != nil {
		t.Fatal(err)
	}
}
