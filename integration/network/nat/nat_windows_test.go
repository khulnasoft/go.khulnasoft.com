package nat // import "go.khulnasoft.com/integration/network/nat"

import (
	"testing"

	"go.khulnasoft.com/integration/internal/network"
	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestWindowsNoDisableIPv4(t *testing.T) {
	ctx := setupTest(t)
	c := testEnv.APIClient()

	_, err := network.Create(ctx, c, "ipv6only",
		network.WithDriver("nat"),
		network.WithIPv4(false),
	)
	assert.Check(t, is.ErrorContains(err, "IPv4 cannot be disabled on Windows"))
}
