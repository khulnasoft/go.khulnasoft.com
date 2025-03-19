package ipams

import (
	"go.khulnasoft.com/libnetwork/ipamapi"
	"go.khulnasoft.com/libnetwork/ipams/defaultipam"
	"go.khulnasoft.com/libnetwork/ipams/null"
	remoteIpam "go.khulnasoft.com/libnetwork/ipams/remote"
	"go.khulnasoft.com/libnetwork/ipams/windowsipam"
	"go.khulnasoft.com/libnetwork/ipamutils"
	"go.khulnasoft.com/pkg/plugingetter"
)

// Register registers all the builtin drivers (ie. default, windowsipam, null
// and remote). If 'pg' is nil, the remote driver won't be registered.
func Register(r ipamapi.Registerer, pg plugingetter.PluginGetter, lAddrPools, gAddrPools []*ipamutils.NetworkToSplit) error {
	if err := defaultipam.Register(r, lAddrPools, gAddrPools); err != nil {
		return err
	}
	if err := windowsipam.Register(r); err != nil {
		return err
	}
	if err := null.Register(r); err != nil {
		return err
	}
	if pg != nil {
		if err := remoteIpam.Register(r, pg); err != nil {
			return err
		}
	}
	return nil
}
