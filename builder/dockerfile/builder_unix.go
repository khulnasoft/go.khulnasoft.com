//go:build !windows

package dockerfile // import "go.khulnasoft.com/builder/dockerfile"

func defaultShellForOS(os string) []string {
	return []string{"/bin/sh", "-c"}
}
