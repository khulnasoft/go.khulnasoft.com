//go:build linux || freebsd

package daemon // import "go.khulnasoft.com/daemon"

import "go.khulnasoft.com/container"

// excludeByIsolation is a platform specific helper function to support PS
// filtering by Isolation. This is a Windows-only concept, so is a no-op on Unix.
func excludeByIsolation(container *container.Snapshot, ctx *listContext) iterationAction {
	return includeContainer
}
