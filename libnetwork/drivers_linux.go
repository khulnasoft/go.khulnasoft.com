package libnetwork

import (
	"fmt"

	"go.khulnasoft.com/libnetwork/datastore"
	"go.khulnasoft.com/libnetwork/driverapi"
	"go.khulnasoft.com/libnetwork/drivers/bridge"
	"go.khulnasoft.com/libnetwork/drivers/host"
	"go.khulnasoft.com/libnetwork/drivers/ipvlan"
	"go.khulnasoft.com/libnetwork/drivers/macvlan"
	"go.khulnasoft.com/libnetwork/drivers/null"
	"go.khulnasoft.com/libnetwork/drivers/overlay"
)

func registerNetworkDrivers(r driverapi.Registerer, store *datastore.Store, driverConfig func(string) map[string]interface{}) error {
	for _, nr := range []struct {
		ntype    string
		register func(driverapi.Registerer, *datastore.Store, map[string]interface{}) error
	}{
		{ntype: bridge.NetworkType, register: bridge.Register},
		{ntype: host.NetworkType, register: func(r driverapi.Registerer, _ *datastore.Store, _ map[string]interface{}) error {
			return host.Register(r)
		}},
		{ntype: ipvlan.NetworkType, register: ipvlan.Register},
		{ntype: macvlan.NetworkType, register: macvlan.Register},
		{ntype: null.NetworkType, register: func(r driverapi.Registerer, _ *datastore.Store, _ map[string]interface{}) error {
			return null.Register(r)
		}},
		{ntype: overlay.NetworkType, register: func(r driverapi.Registerer, _ *datastore.Store, config map[string]interface{}) error {
			return overlay.Register(r, config)
		}},
	} {
		if err := nr.register(r, store, driverConfig(nr.ntype)); err != nil {
			return fmt.Errorf("failed to register %q driver: %w", nr.ntype, err)
		}
	}

	return nil
}
