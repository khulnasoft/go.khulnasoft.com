//go:build linux || windows

package service // import "go.khulnasoft.com/volume/service"
import (
	"go.khulnasoft.com/pkg/idtools"
	"go.khulnasoft.com/volume"
	"go.khulnasoft.com/volume/drivers"
	"go.khulnasoft.com/volume/local"
	"github.com/pkg/errors"
)

func setupDefaultDriver(store *drivers.Store, root string, rootIDs idtools.Identity) error {
	d, err := local.New(root, rootIDs)
	if err != nil {
		return errors.Wrap(err, "error setting up default driver")
	}
	if !store.Register(d, volume.DefaultDriverName) {
		return errors.New("local volume driver could not be registered")
	}
	return nil
}
