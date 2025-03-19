//go:build !freebsd && !linux && !windows

package libnetwork

import "go.khulnasoft.com/libnetwork/driverapi"

func registerNetworkDrivers(r driverapi.Registerer, driverConfig func(string) map[string]interface{}) error {
	return nil
}
