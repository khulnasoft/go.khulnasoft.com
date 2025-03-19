package layer // import "go.khulnasoft.com/layer"

import (
	"io"

	"github.com/khulnasoft/distribution"
)

func (ls *layerStore) RegisterWithDescriptor(ts io.Reader, parent ChainID, descriptor distribution.Descriptor) (Layer, error) {
	return ls.registerWithDescriptor(ts, parent, descriptor)
}
