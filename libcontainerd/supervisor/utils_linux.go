package supervisor // import "go.khulnasoft.com/libcontainerd/supervisor"

import "syscall"

// containerdSysProcAttr returns the SysProcAttr to use when exec'ing
// containerd
func containerdSysProcAttr() *syscall.SysProcAttr {
	return &syscall.SysProcAttr{
		Setsid:    true,
		Pdeathsig: syscall.SIGKILL,
	}
}
