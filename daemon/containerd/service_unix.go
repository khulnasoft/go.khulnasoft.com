//go:build linux || freebsd

package containerd

import (
	"go.khulnasoft.com/container"
	"go.khulnasoft.com/errdefs"
	"go.khulnasoft.com/image"
	"github.com/pkg/errors"
)

// GetLayerFolders returns the layer folders from an image RootFS.
func (i *ImageService) GetLayerFolders(img *image.Image, rwLayer container.RWLayer, containerID string) ([]string, error) {
	return nil, errdefs.NotImplemented(errors.New("not implemented"))
}
