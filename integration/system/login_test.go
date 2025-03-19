package system // import "go.khulnasoft.com/integration/system"

import (
	"fmt"
	"testing"

	"go.khulnasoft.com/api/types/registry"
	"go.khulnasoft.com/integration/internal/requirement"
	registrypkg "go.khulnasoft.com/registry"
	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
	"gotest.tools/v3/skip"
)

// Test case for GitHub 22244
func TestLoginFailsWithBadCredentials(t *testing.T) {
	skip.If(t, !requirement.HasHubConnectivity(t))

	ctx := setupTest(t)
	apiClient := testEnv.APIClient()

	_, err := apiClient.RegistryLogin(ctx, registry.AuthConfig{
		Username: "no-user",
		Password: "no-password",
	})
	assert.Assert(t, err != nil)
	assert.Check(t, is.ErrorContains(err, "unauthorized: incorrect username or password"))
	assert.Check(t, is.ErrorContains(err, fmt.Sprintf("https://%s/v2/", registrypkg.DefaultRegistryHost)))
}
