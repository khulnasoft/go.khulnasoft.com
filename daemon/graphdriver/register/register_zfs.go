//go:build (!exclude_graphdriver_zfs && linux) || (!exclude_graphdriver_zfs && freebsd)

package register // import "go.khulnasoft.com/daemon/graphdriver/register"

import (
	// register the zfs driver
	_ "go.khulnasoft.com/daemon/graphdriver/zfs"
)
