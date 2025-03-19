//go:build !linux

package daemon // import "go.khulnasoft.com/daemon"

import (
	"context"

	"go.khulnasoft.com/container"
	"go.khulnasoft.com/daemon/config"
	"go.khulnasoft.com/libcontainerd/types"
	"github.com/opencontainers/runtime-spec/specs-go"
)

// initializeCreatedTask performs any initialization that needs to be done to
// prepare a freshly-created task to be started.
func (daemon *Daemon) initializeCreatedTask(ctx context.Context, cfg *config.Config, tsk types.Task, container *container.Container, spec *specs.Spec) error {
	return nil
}
