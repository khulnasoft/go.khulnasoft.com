//go:build !exclude_graphdriver_overlay2 && linux

package register // import "go.khulnasoft.com/daemon/graphdriver/register"

import (
	// register the overlay2 graphdriver
	_ "go.khulnasoft.com/daemon/graphdriver/overlay2"
)
