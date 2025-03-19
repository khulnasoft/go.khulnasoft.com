//go:build !linux

package vfs // import "go.khulnasoft.com/daemon/graphdriver/vfs"

import (
	"go.khulnasoft.com/pkg/chrootarchive"
	"go.khulnasoft.com/pkg/idtools"
)

func dirCopy(srcDir, dstDir string) error {
	return chrootarchive.NewArchiver(idtools.IdentityMapping{}).CopyWithTar(srcDir, dstDir)
}
