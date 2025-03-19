//go:build !windows

package windowsipam

import "go.khulnasoft.com/libnetwork/ipamapi"

// Register is a no-op -- windowsipam is only supported on Windows.
func Register(ipamapi.Registerer) error {
	return nil
}
