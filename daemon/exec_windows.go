package daemon // import "go.khulnasoft.com/daemon"

import (
	"context"

	"go.khulnasoft.com/container"
	"go.khulnasoft.com/daemon/config"
	"github.com/opencontainers/runtime-spec/specs-go"
)

func (daemon *Daemon) execSetPlatformOpt(ctx context.Context, daemonCfg *config.Config, ec *container.ExecConfig, p *specs.Process) error {
	if ec.Container.ImagePlatform.OS == "windows" {
		p.User.Username = ec.User
	}
	return nil
}
