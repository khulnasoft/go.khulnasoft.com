package tarexport // import "go.khulnasoft.com/image/tarexport"

import (
	"context"

	"github.com/containerd/platforms"
	"github.com/khulnasoft/distribution"
	"go.khulnasoft.com/api/types/events"
	"go.khulnasoft.com/image"
	"go.khulnasoft.com/layer"
	refstore "go.khulnasoft.com/reference"
	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
)

const (
	manifestFileName           = "manifest.json"
	legacyLayerFileName        = "layer.tar"
	legacyConfigFileName       = "json"
	legacyRepositoriesFileName = "repositories"
)

type manifestItem struct {
	Config       string
	RepoTags     []string
	Layers       []string
	Parent       image.ID                                 `json:",omitempty"`
	LayerSources map[layer.DiffID]distribution.Descriptor `json:",omitempty"`
}

type tarexporter struct {
	is              image.Store
	lss             layer.Store
	rs              refstore.Store
	loggerImgEvent  LogImageEvent
	platform        *platforms.Platform
	platformMatcher platforms.Matcher
}

// LogImageEvent defines interface for event generation related to image tar(load and save) operations
type LogImageEvent interface {
	// LogImageEvent generates an event related to an image operation
	LogImageEvent(ctx context.Context, imageID, refName string, action events.Action)
}

// NewTarExporter returns new Exporter for tar packages
func NewTarExporter(is image.Store, lss layer.Store, rs refstore.Store, loggerImgEvent LogImageEvent, platform *ocispec.Platform) image.Exporter {
	l := &tarexporter{
		is:             is,
		lss:            lss,
		rs:             rs,
		loggerImgEvent: loggerImgEvent,
		platform:       platform,
	}
	if platform != nil {
		l.platformMatcher = platforms.OnlyStrict(*platform)
	}
	return l
}
