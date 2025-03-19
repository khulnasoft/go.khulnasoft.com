package cnmallocator

import (
	"go.khulnasoft.com/libnetwork/driverapi"
	"go.khulnasoft.com/libnetwork/drivers/bridge/brmanager"
	"go.khulnasoft.com/libnetwork/drivers/host"
	"go.khulnasoft.com/libnetwork/drivers/ipvlan/ivmanager"
	"go.khulnasoft.com/libnetwork/drivers/macvlan/mvmanager"
	"go.khulnasoft.com/libnetwork/drivers/overlay/ovmanager"
	"github.com/moby/swarmkit/v2/manager/allocator/networkallocator"
)

var initializers = map[string]func(driverapi.Registerer) error{
	"overlay": ovmanager.Register,
	"macvlan": mvmanager.Register,
	"bridge":  brmanager.Register,
	"ipvlan":  ivmanager.Register,
	"host":    host.Register,
}

// PredefinedNetworks returns the list of predefined network structures
func (*Provider) PredefinedNetworks() []networkallocator.PredefinedNetworkData {
	return []networkallocator.PredefinedNetworkData{
		{Name: "bridge", Driver: "bridge"},
		{Name: "host", Driver: "host"},
	}
}
