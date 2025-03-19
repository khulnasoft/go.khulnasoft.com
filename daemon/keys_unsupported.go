//go:build !linux

package daemon // import "go.khulnasoft.com/daemon"

// modifyRootKeyLimit is a noop on unsupported platforms.
func modifyRootKeyLimit() error {
	return nil
}
