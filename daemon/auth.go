package daemon // import "go.khulnasoft.com/daemon"

import (
	"context"

	"go.khulnasoft.com/api/types/registry"
	"go.khulnasoft.com/dockerversion"
)

// AuthenticateToRegistry checks the validity of credentials in authConfig
func (daemon *Daemon) AuthenticateToRegistry(ctx context.Context, authConfig *registry.AuthConfig) (string, string, error) {
	return daemon.registryService.Auth(ctx, authConfig, dockerversion.DockerUserAgent(ctx))
}
