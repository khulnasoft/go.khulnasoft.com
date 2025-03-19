//go:build !linux

package caps // import "go.khulnasoft.com/oci/caps"

func initCaps() {
	// no capabilities on Windows
}
