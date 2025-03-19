//go:build !linux && !freebsd

package zfs // import "go.khulnasoft.com/daemon/graphdriver/zfs"

func checkRootdirFs(rootdir string) error {
	return nil
}

func getMountpoint(id string) string {
	return id
}
