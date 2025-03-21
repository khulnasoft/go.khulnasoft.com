//go:build linux

package bridge

import (
	"errors"

	"go.khulnasoft.com/libnetwork/iptables"
)

func (n *bridgeNetwork) setupFirewalld(config *networkConfiguration, i *bridgeInterface) error {
	d := n.driver
	d.Lock()
	driverConfig := d.config
	d.Unlock()

	// Sanity check.
	if !driverConfig.EnableIPTables {
		return errors.New("no need to register firewalld hooks, iptables is disabled")
	}

	iptables.OnReloaded(func() { n.setupIP4Tables(config, i) })
	iptables.OnReloaded(n.reapplyPerPortIptables4)
	return nil
}

func (n *bridgeNetwork) setupFirewalld6(config *networkConfiguration, i *bridgeInterface) error {
	d := n.driver
	d.Lock()
	driverConfig := d.config
	d.Unlock()

	// Sanity check.
	if !driverConfig.EnableIP6Tables {
		return errors.New("no need to register firewalld hooks, ip6tables is disabled")
	}

	iptables.OnReloaded(func() { n.setupIP6Tables(config, i) })
	iptables.OnReloaded(n.reapplyPerPortIptables6)
	return nil
}
