//go:build !windows

package main

import (
	"os"
	"strings"

	"go.khulnasoft.com/pkg/sysinfo"
)

var sysInfo *sysinfo.SysInfo

func setupLocalInfo() {
	sysInfo = sysinfo.New()
}

func cpuCfsPeriod() bool {
	return testEnv.DaemonInfo.CPUCfsPeriod
}

func cpuCfsQuota() bool {
	return testEnv.DaemonInfo.CPUCfsQuota
}

func cpuShare() bool {
	return testEnv.DaemonInfo.CPUShares
}

func oomControl() bool {
	return testEnv.DaemonInfo.OomKillDisable
}

func pidsLimit() bool {
	return sysInfo.PidsLimit
}

func memoryLimitSupport() bool {
	return testEnv.DaemonInfo.MemoryLimit
}

func memoryReservationSupport() bool {
	return sysInfo.MemoryReservation
}

func swapMemorySupport() bool {
	return testEnv.DaemonInfo.SwapLimit
}

func memorySwappinessSupport() bool {
	return testEnv.IsLocalDaemon() && sysInfo.MemorySwappiness
}

func blkioWeight() bool {
	return testEnv.IsLocalDaemon() && sysInfo.BlkioWeight
}

func cgroupCpuset() bool {
	return testEnv.DaemonInfo.CPUSet
}

func seccompEnabled() bool {
	return sysInfo.Seccomp
}

func bridgeNfIptables() bool {
	content, err := os.ReadFile("/proc/sys/net/bridge/bridge-nf-call-iptables")
	return err == nil && strings.TrimSpace(string(content)) == "1"
}

func unprivilegedUsernsClone() bool {
	content, err := os.ReadFile("/proc/sys/kernel/unprivileged_userns_clone")
	return err != nil || !strings.Contains(string(content), "0")
}
