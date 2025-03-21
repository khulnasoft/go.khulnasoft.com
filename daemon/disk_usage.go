package daemon // import "go.khulnasoft.com/daemon"

import (
	"context"
	"fmt"

	"go.khulnasoft.com/api/server/router/system"
	"go.khulnasoft.com/api/types"
	"go.khulnasoft.com/api/types/container"
	"go.khulnasoft.com/api/types/filters"
	"go.khulnasoft.com/api/types/image"
	"go.khulnasoft.com/api/types/volume"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

// containerDiskUsage obtains information about container data disk usage
// and makes sure that only one calculation is performed at the same time.
func (daemon *Daemon) containerDiskUsage(ctx context.Context) ([]*container.Summary, error) {
	res, _, err := daemon.usageContainers.Do(ctx, struct{}{}, func(ctx context.Context) ([]*container.Summary, error) {
		// Retrieve container list
		containers, err := daemon.Containers(ctx, &container.ListOptions{
			Size: true,
			All:  true,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to retrieve container list: %v", err)
		}

		// Remove image manifest descriptor from the result as it should not be included.
		// https://github.com/moby/moby/pull/49407#discussion_r1954396666
		for _, c := range containers {
			c.ImageManifestDescriptor = nil
		}
		return containers, nil
	})
	return res, err
}

// imageDiskUsage obtains information about image data disk usage from image service
// and makes sure that only one calculation is performed at the same time.
func (daemon *Daemon) imageDiskUsage(ctx context.Context) ([]*image.Summary, error) {
	imgs, _, err := daemon.usageImages.Do(ctx, struct{}{}, func(ctx context.Context) ([]*image.Summary, error) {
		// Get all top images with extra attributes
		imgs, err := daemon.imageService.Images(ctx, image.ListOptions{
			Filters:        filters.NewArgs(),
			SharedSize:     true,
			ContainerCount: true,
		})
		if err != nil {
			return nil, errors.Wrap(err, "failed to retrieve image list")
		}
		return imgs, nil
	})

	return imgs, err
}

// localVolumesSize obtains information about volume disk usage from volumes service
// and makes sure that only one size calculation is performed at the same time.
func (daemon *Daemon) localVolumesSize(ctx context.Context) ([]*volume.Volume, error) {
	volumes, _, err := daemon.usageVolumes.Do(ctx, struct{}{}, func(ctx context.Context) ([]*volume.Volume, error) {
		volumes, err := daemon.volumes.LocalVolumesSize(ctx)
		if err != nil {
			return nil, err
		}
		return volumes, nil
	})
	return volumes, err
}

// layerDiskUsage obtains information about layer disk usage from image service
// and makes sure that only one size calculation is performed at the same time.
func (daemon *Daemon) layerDiskUsage(ctx context.Context) (int64, error) {
	usage, _, err := daemon.usageLayer.Do(ctx, struct{}{}, func(ctx context.Context) (int64, error) {
		usage, err := daemon.imageService.LayerDiskUsage(ctx)
		if err != nil {
			return 0, err
		}
		return usage, nil
	})
	return usage, err
}

// SystemDiskUsage returns information about the daemon data disk usage.
// Callers must not mutate contents of the returned fields.
func (daemon *Daemon) SystemDiskUsage(ctx context.Context, opts system.DiskUsageOptions) (*types.DiskUsage, error) {
	eg, ctx := errgroup.WithContext(ctx)

	var containers []*container.Summary
	if opts.Containers {
		eg.Go(func() error {
			var err error
			containers, err = daemon.containerDiskUsage(ctx)
			return err
		})
	}

	var (
		images     []*image.Summary
		layersSize int64
	)
	if opts.Images {
		eg.Go(func() error {
			var err error
			images, err = daemon.imageDiskUsage(ctx)
			return err
		})
		eg.Go(func() error {
			var err error
			layersSize, err = daemon.layerDiskUsage(ctx)
			return err
		})
	}

	var volumes []*volume.Volume
	if opts.Volumes {
		eg.Go(func() error {
			var err error
			volumes, err = daemon.localVolumesSize(ctx)
			return err
		})
	}

	if err := eg.Wait(); err != nil {
		return nil, err
	}
	return &types.DiskUsage{
		LayersSize: layersSize,
		Containers: containers,
		Volumes:    volumes,
		Images:     images,
	}, nil
}
