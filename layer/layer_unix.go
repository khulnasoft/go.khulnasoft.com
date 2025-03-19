//go:build linux || freebsd || darwin || openbsd

package layer // import "go.khulnasoft.com/layer"

import "go.khulnasoft.com/pkg/stringid"

func (ls *layerStore) mountID(name string) string {
	return stringid.GenerateRandomID()
}
