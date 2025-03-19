package daemon // import "go.khulnasoft.com/daemon"

import (
	"go.khulnasoft.com/api/types/container"
	libcontainerdtypes "go.khulnasoft.com/libcontainerd/types"
)

func toContainerdResources(resources container.Resources) *libcontainerdtypes.Resources {
	// We don't support update, so do nothing
	return nil
}
