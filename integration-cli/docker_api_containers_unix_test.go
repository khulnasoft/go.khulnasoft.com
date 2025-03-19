//go:build !windows

package main

import (
	"testing"

	"go.khulnasoft.com/sys/mount"
)

func mountWrapper(t *testing.T, device, target, mType, options string) error {
	t.Helper()
	err := mount.Mount(device, target, mType, options)
	if err != nil {
		return err
	}
	t.Cleanup(func() { _ = mount.Unmount(target) })
	return nil
}
