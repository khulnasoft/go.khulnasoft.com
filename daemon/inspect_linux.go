package daemon // import "go.khulnasoft.com/daemon"

import (
	"go.khulnasoft.com/api/types/backend"
	"go.khulnasoft.com/api/types/container"
	containerpkg "go.khulnasoft.com/container"
)

// This sets platform-specific fields
func setPlatformSpecificContainerFields(container *containerpkg.Container, contJSONBase *container.ContainerJSONBase) *container.ContainerJSONBase {
	contJSONBase.AppArmorProfile = container.AppArmorProfile
	contJSONBase.ResolvConfPath = container.ResolvConfPath
	contJSONBase.HostnamePath = container.HostnamePath
	contJSONBase.HostsPath = container.HostsPath

	return contJSONBase
}

func inspectExecProcessConfig(e *containerpkg.ExecConfig) *backend.ExecProcessConfig {
	return &backend.ExecProcessConfig{
		Tty:        e.Tty,
		Entrypoint: e.Entrypoint,
		Arguments:  e.Args,
		Privileged: &e.Privileged,
		User:       e.User,
	}
}
