package daemon // import "go.khulnasoft.com/daemon"

import (
	"testing"

	"go.khulnasoft.com/api/types/system"
	"go.khulnasoft.com/dockerversion"
	"gotest.tools/v3/assert"
)

func TestFillLicense(t *testing.T) {
	v := &system.Info{}
	d := &Daemon{
		root: "/var/lib/docker/",
	}
	d.fillLicense(v)
	assert.Assert(t, v.ProductLicense == dockerversion.DefaultProductLicense)
}
