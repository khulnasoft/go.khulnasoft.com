package libnetwork

import (
	"fmt"

	"go.khulnasoft.com/libnetwork/datastore"
	"go.khulnasoft.com/libnetwork/driverapi"
	"go.khulnasoft.com/libnetwork/drivers/null"
	"go.khulnasoft.com/libnetwork/drivers/windows"
	"go.khulnasoft.com/libnetwork/drivers/windows/overlay"
)

func registerNetworkDrivers(r driverapi.Registerer, store *datastore.Store, _ func(string) map[string]interface{}) error {
	for _, nr := range []struct {
		ntype    string
		register func(driverapi.Registerer) error
	}{
		{ntype: null.NetworkType, register: null.Register},
		{ntype: overlay.NetworkType, register: overlay.Register},
	} {
		if err := nr.register(r); err != nil {
			return fmt.Errorf("failed to register %q driver: %w", nr.ntype, err)
		}
	}

	return windows.RegisterBuiltinLocalDrivers(r, store)
}
