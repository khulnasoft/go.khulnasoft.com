package rootless // import "go.khulnasoft.com/pkg/rootless"

import "os"

// RunningWithRootlessKit returns true if running under RootlessKit namespaces.
func RunningWithRootlessKit() bool {
	return os.Getenv("ROOTLESSKIT_STATE_DIR") != ""
}
