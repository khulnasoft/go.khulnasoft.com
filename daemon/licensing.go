package daemon // import "go.khulnasoft.com/daemon"

import (
	"go.khulnasoft.com/api/types/system"
	"go.khulnasoft.com/dockerversion"
)

func (daemon *Daemon) fillLicense(v *system.Info) {
	v.ProductLicense = dockerversion.DefaultProductLicense
}
