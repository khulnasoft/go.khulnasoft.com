//go:build !windows

package container // import "go.khulnasoft.com/daemon/cluster/executor/container"

const (
	testAbsPath        = "/foo"
	testAbsNonExistent = "/some-non-existing-host-path/"
)
