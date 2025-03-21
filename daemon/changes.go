package daemon // import "go.khulnasoft.com/daemon"

import (
	"context"
	"errors"
	"time"

	"go.khulnasoft.com/internal/metrics"
	"go.khulnasoft.com/pkg/archive"
)

// ContainerChanges returns a list of container fs changes
func (daemon *Daemon) ContainerChanges(ctx context.Context, name string) ([]archive.Change, error) {
	start := time.Now()

	container, err := daemon.GetContainer(name)
	if err != nil {
		return nil, err
	}

	if isWindows && container.IsRunning() {
		return nil, errors.New("Windows does not support diff of a running container")
	}

	c, err := daemon.imageService.Changes(ctx, container)
	if err != nil {
		return nil, err
	}
	metrics.ContainerActions.WithValues("changes").UpdateSince(start)
	return c, nil
}
