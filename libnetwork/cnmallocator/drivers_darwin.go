package cnmallocator

import (
	"go.khulnasoft.com/libnetwork/driverapi"
	"go.khulnasoft.com/libnetwork/drivers/overlay/ovmanager"
	"github.com/moby/swarmkit/v2/manager/allocator/networkallocator"
)

var initializers = map[string]func(driverapi.Registerer) error{
	"overlay": ovmanager.Register,
}

// PredefinedNetworks returns the list of predefined network structures
func (*Provider) PredefinedNetworks() []networkallocator.PredefinedNetworkData {
	return nil
}
