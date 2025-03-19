package httputils // import "go.khulnasoft.com/api/server/httputils"

import (
	"io"

	"go.khulnasoft.com/api/types/container"
	"go.khulnasoft.com/api/types/network"
)

// ContainerDecoder specifies how
// to translate an io.Reader into
// container configuration.
type ContainerDecoder interface {
	DecodeConfig(src io.Reader) (*container.Config, *container.HostConfig, *network.NetworkingConfig, error)
}
