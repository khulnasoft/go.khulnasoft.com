package layer // import "go.khulnasoft.com/layer"

import "github.com/khulnasoft/distribution"

var _ distribution.Describable = &roLayer{}

func (rl *roLayer) Descriptor() distribution.Descriptor {
	return rl.descriptor
}
