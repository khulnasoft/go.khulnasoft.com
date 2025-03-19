package plugins // import "go.khulnasoft.com/pkg/plugins"

import (
	"os"
	"path/filepath"
)

// specsPaths is the Windows implementation of [SpecsPaths].
func specsPaths() []string {
	return []string{filepath.Join(os.Getenv("programdata"), "docker", "plugins")}
}
