//go:build !linux && !freebsd && !windows

package daemon // import "go.khulnasoft.com/daemon"

import (
	"errors"

	"go.khulnasoft.com/pkg/sysinfo"
)

func checkSystem() error {
	return errors.New("the Docker daemon is not supported on this platform")
}

func setupResolvConf(_ *interface{}) {}

func getSysInfo(_ *Daemon) *sysinfo.SysInfo {
	return sysinfo.New()
}
