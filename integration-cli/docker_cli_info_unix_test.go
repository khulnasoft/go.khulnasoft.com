//go:build !windows

package main

import (
	"testing"

	"go.khulnasoft.com/client"
	"go.khulnasoft.com/daemon/config"
	"go.khulnasoft.com/testutil"
	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func (s *DockerCLIInfoSuite) TestInfoSecurityOptions(c *testing.T) {
	testRequires(c, testEnv.IsLocalDaemon, DaemonIsLinux)
	if !seccompEnabled() && !Apparmor() {
		c.Skip("test requires Seccomp and/or AppArmor")
	}

	apiClient, err := client.NewClientWithOpts(client.FromEnv)
	assert.NilError(c, err)
	defer apiClient.Close()
	info, err := apiClient.Info(testutil.GetContext(c))
	assert.NilError(c, err)

	if Apparmor() {
		assert.Check(c, is.Contains(info.SecurityOptions, "name=apparmor"))
	}
	if seccompEnabled() {
		assert.Check(c, is.Contains(info.SecurityOptions, "name=seccomp,profile="+config.SeccompProfileDefault))
	}
}
