package daemon // import "go.khulnasoft.com/daemon"

import (
	"testing"

	"go.khulnasoft.com/api/types/container"
)

func TestToContainerdResources_Defaults(t *testing.T) {
	checkResourcesAreUnset(t, toContainerdResources(container.Resources{}))
}
