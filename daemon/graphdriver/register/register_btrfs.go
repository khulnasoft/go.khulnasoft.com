//go:build !exclude_graphdriver_btrfs && linux

package register // import "go.khulnasoft.com/daemon/graphdriver/register"

import (
	// register the btrfs graphdriver
	_ "go.khulnasoft.com/daemon/graphdriver/btrfs"
)
