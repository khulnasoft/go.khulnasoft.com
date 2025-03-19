package libcontainerd // import "go.khulnasoft.com/libcontainerd"

import (
	"context"

	containerd "github.com/containerd/containerd/v2/client"
	"go.khulnasoft.com/libcontainerd/remote"
	libcontainerdtypes "go.khulnasoft.com/libcontainerd/types"
)

// NewClient creates a new libcontainerd client from a containerd client
func NewClient(ctx context.Context, cli *containerd.Client, stateDir, ns string, b libcontainerdtypes.Backend) (libcontainerdtypes.Client, error) {
	return remote.NewClient(ctx, cli, stateDir, ns, b)
}
