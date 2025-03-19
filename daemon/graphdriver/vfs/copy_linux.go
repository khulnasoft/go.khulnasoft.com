package vfs // import "go.khulnasoft.com/daemon/graphdriver/vfs"

import "go.khulnasoft.com/daemon/graphdriver/copy"

func dirCopy(srcDir, dstDir string) error {
	return copy.DirCopy(srcDir, dstDir, copy.Content, false)
}
