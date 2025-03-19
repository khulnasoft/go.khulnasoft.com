package libnetwork

import (
	windriver "go.khulnasoft.com/libnetwork/drivers/windows"
	"go.khulnasoft.com/libnetwork/options"
	"go.khulnasoft.com/libnetwork/types"
)

const libnGWNetwork = "nat"

func getPlatformOption() EndpointOption {
	epOption := options.Generic{
		windriver.DisableICC: true,
		windriver.DisableDNS: true,
	}
	return EndpointOptionGeneric(epOption)
}

func (c *Controller) createGWNetwork() (*Network, error) {
	return nil, types.NotImplementedErrorf("default gateway functionality is not implemented in windows")
}
