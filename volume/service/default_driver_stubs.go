//go:build !linux && !windows

package service // import "go.khulnasoft.com/volume/service"

import (
	"go.khulnasoft.com/pkg/idtools"
	"go.khulnasoft.com/volume/drivers"
)

func setupDefaultDriver(_ *drivers.Store, _ string, _ idtools.Identity) error { return nil }
