package daemon // import "go.khulnasoft.com/daemon"

import (
	"github.com/Microsoft/hcsshim/cmd/containerd-shim-runhcs-v1/options"
	"go.khulnasoft.com/container"
	"go.khulnasoft.com/daemon/config"
	"go.khulnasoft.com/pkg/system"
)

func (daemon *Daemon) getLibcontainerdCreateOptions(*configStore, *container.Container) (string, interface{}, error) {
	if system.ContainerdRuntimeSupported() {
		opts := &options.Options{}
		return config.WindowsV2RuntimeName, opts, nil
	}
	return "", nil, nil
}
