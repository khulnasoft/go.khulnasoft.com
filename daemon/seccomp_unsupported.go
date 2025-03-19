//go:build !linux

package daemon // import "go.khulnasoft.com/daemon"

import (
	"context"

	"github.com/containerd/containerd/v2/core/containers"
	coci "github.com/containerd/containerd/v2/pkg/oci"
	"go.khulnasoft.com/container"
)

const supportsSeccomp = false

// WithSeccomp sets the seccomp profile
func WithSeccomp(daemon *Daemon, c *container.Container) coci.SpecOpts {
	return func(ctx context.Context, _ coci.Client, _ *containers.Container, s *coci.Spec) error {
		return nil
	}
}
