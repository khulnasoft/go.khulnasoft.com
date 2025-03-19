//go:build !windows

package logging

import (
	"testing"

	"go.khulnasoft.com/api/types"
	"go.khulnasoft.com/testutil"
	"go.khulnasoft.com/testutil/daemon"
	"gotest.tools/v3/assert"
	"gotest.tools/v3/skip"
)

// Regression test for #35553
// Ensure that a daemon with a log plugin set as the default logger for containers
// does not keep the daemon from starting.
func TestDaemonStartWithLogOpt(t *testing.T) {
	skip.If(t, testEnv.IsRemoteDaemon, "cannot run daemon when remote daemon")
	skip.If(t, testEnv.DaemonInfo.OSType == "windows")
	t.Parallel()

	ctx := testutil.StartSpan(baseContext, t)

	d := daemon.New(t)
	d.Start(t, "--iptables=false", "--ip6tables=false")
	defer d.Stop(t)

	c := d.NewClientT(t)

	createPlugin(ctx, t, c, "test", "dummy", asLogDriver)
	err := c.PluginEnable(ctx, "test", types.PluginEnableOptions{Timeout: 30})
	assert.Check(t, err)
	defer c.PluginRemove(ctx, "test", types.PluginRemoveOptions{Force: true})

	d.Stop(t)
	d.Start(t, "--iptables=false", "--ip6tables=false", "--log-driver=test", "--log-opt=foo=bar")
}
