//go:build !linux && !darwin && !freebsd && !windows

package daemon // import "go.khulnasoft.com/daemon"

func (daemon *Daemon) setupDumpStackTrap(_ string) {
	return
}
