package filesys

import "go.khulnasoft.com/yaml/filesys"

const (
	// Separator is deprecated, use go.khulnasoft.com/yaml/filesys.Separator.
	Separator = filesys.Separator
	// SelfDir is deprecated, use go.khulnasoft.com/yaml/filesys.SelfDir.
	SelfDir = filesys.SelfDir
	// ParentDir is deprecated, use go.khulnasoft.com/yaml/filesys.ParentDir.
	ParentDir = filesys.ParentDir
)

type (
	// FileSystem is deprecated, use go.khulnasoft.com/yaml/filesys.FileSystem.
	FileSystem = filesys.FileSystem
	// FileSystemOrOnDisk is deprecated, use go.khulnasoft.com/yaml/filesys.FileSystemOrOnDisk.
	FileSystemOrOnDisk = filesys.FileSystemOrOnDisk
	// ConfirmedDir is deprecated, use go.khulnasoft.com/yaml/filesys.ConfirmedDir.
	ConfirmedDir = filesys.ConfirmedDir
)

// MakeEmptyDirInMemory is deprecated, use go.khulnasoft.com/yaml/filesys.MakeEmptyDirInMemory.
func MakeEmptyDirInMemory() FileSystem { return filesys.MakeEmptyDirInMemory() }

// MakeFsInMemory is deprecated, use go.khulnasoft.com/yaml/filesys.MakeFsInMemory.
func MakeFsInMemory() FileSystem { return filesys.MakeFsInMemory() }

// MakeFsOnDisk is deprecated, use go.khulnasoft.com/yaml/filesys.MakeFsOnDisk.
func MakeFsOnDisk() FileSystem { return filesys.MakeFsOnDisk() }

// NewTmpConfirmedDir is deprecated, use go.khulnasoft.com/yaml/filesys.NewTmpConfirmedDir.
func NewTmpConfirmedDir() (filesys.ConfirmedDir, error) { return filesys.NewTmpConfirmedDir() }

// RootedPath is deprecated, use go.khulnasoft.com/yaml/filesys.RootedPath.
func RootedPath(elem ...string) string { return filesys.RootedPath(elem...) }

// StripTrailingSeps is deprecated, use go.khulnasoft.com/yaml/filesys.StripTrailingSeps.
func StripTrailingSeps(s string) string { return filesys.StripTrailingSeps(s) }

// StripLeadingSeps is deprecated, use go.khulnasoft.com/yaml/filesys.StripLeadingSeps.
func StripLeadingSeps(s string) string { return filesys.StripLeadingSeps(s) }

// PathSplit is deprecated, use go.khulnasoft.com/yaml/filesys.PathSplit.
func PathSplit(incoming string) []string { return filesys.PathSplit(incoming) }

// PathJoin is deprecated, use go.khulnasoft.com/yaml/filesys.PathJoin.
func PathJoin(incoming []string) string { return filesys.PathJoin(incoming) }

// InsertPathPart is deprecated, use go.khulnasoft.com/yaml/filesys.InsertPathPart.
func InsertPathPart(path string, pos int, part string) string {
	return filesys.InsertPathPart(path, pos, part)
}
