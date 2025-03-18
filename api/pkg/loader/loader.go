package loader

import (
	"go.khulnasoft.com/api/internal/loader"
	"go.khulnasoft.com/yaml/filesys"
)

// NewFileLoaderAtCwd returns a loader that loads from PWD.
// A convenience for khulnasoft edit commands.
func NewFileLoaderAtCwd(fSys filesys.FileSystem) *loader.FileLoader {
	return loader.NewLoaderOrDie(
		loader.RestrictionRootOnly, fSys, filesys.SelfDir)
}

// NewFileLoaderAtRoot returns a loader that loads from "/".
// A convenience for tests.
func NewFileLoaderAtRoot(fSys filesys.FileSystem) *loader.FileLoader {
	return loader.NewLoaderOrDie(
		loader.RestrictionRootOnly, fSys, filesys.Separator)
}
