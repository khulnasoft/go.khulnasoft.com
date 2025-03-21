package client // import "go.khulnasoft.com/client"

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"go.khulnasoft.com/api/types/image"
	"go.khulnasoft.com/errdefs"
	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestImageInspectError(t *testing.T) {
	client := &Client{
		client: newMockClient(errorMock(http.StatusInternalServerError, "Server error")),
	}

	_, err := client.ImageInspect(context.Background(), "nothing")
	assert.Check(t, is.ErrorType(err, errdefs.IsSystem))
}

func TestImageInspectImageNotFound(t *testing.T) {
	client := &Client{
		client: newMockClient(errorMock(http.StatusNotFound, "Server error")),
	}

	_, err := client.ImageInspect(context.Background(), "unknown")
	assert.Check(t, is.ErrorType(err, errdefs.IsNotFound))
}

func TestImageInspectWithEmptyID(t *testing.T) {
	client := &Client{
		client: newMockClient(func(req *http.Request) (*http.Response, error) {
			return nil, errors.New("should not make request")
		}),
	}
	_, err := client.ImageInspect(context.Background(), "")
	assert.Check(t, is.ErrorType(err, errdefs.IsNotFound))
}

func TestImageInspect(t *testing.T) {
	expectedURL := "/images/image_id/json"
	expectedTags := []string{"tag1", "tag2"}
	client := &Client{
		client: newMockClient(func(req *http.Request) (*http.Response, error) {
			if !strings.HasPrefix(req.URL.Path, expectedURL) {
				return nil, fmt.Errorf("Expected URL '%s', got '%s'", expectedURL, req.URL)
			}
			content, err := json.Marshal(image.InspectResponse{
				ID:       "image_id",
				RepoTags: expectedTags,
			})
			if err != nil {
				return nil, err
			}
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewReader(content)),
			}, nil
		}),
	}

	imageInspect, err := client.ImageInspect(context.Background(), "image_id")
	if err != nil {
		t.Fatal(err)
	}
	if imageInspect.ID != "image_id" {
		t.Fatalf("expected `image_id`, got %s", imageInspect.ID)
	}
	if !reflect.DeepEqual(imageInspect.RepoTags, expectedTags) {
		t.Fatalf("expected `%v`, got %v", expectedTags, imageInspect.RepoTags)
	}
}
