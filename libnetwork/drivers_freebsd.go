package libnetwork

import (
	"go.khulnasoft.com/libnetwork/driverapi"
	"go.khulnasoft.com/libnetwork/drivers/null"
)

func registerNetworkDrivers(r driverapi.Registerer, driverConfig func(string) map[string]interface{}) error {
	return null.Register(r)
}
