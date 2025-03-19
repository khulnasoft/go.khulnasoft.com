//go:build !exclude_graphdriver_fuseoverlayfs && linux

package register // import "go.khulnasoft.com/daemon/graphdriver/register"

import (
	// register the fuse-overlayfs graphdriver
	_ "go.khulnasoft.com/daemon/graphdriver/fuse-overlayfs"
)
