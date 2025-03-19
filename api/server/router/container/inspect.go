// FIXME(thaJeztah): remove once we are a module; the go:build directive prevents go from downgrading language version to go1.16:
//go:build go1.22

package container // import "go.khulnasoft.com/api/server/router/container"

import (
	"context"
	"net/http"

	"go.khulnasoft.com/api/server/httputils"
	"go.khulnasoft.com/api/types/backend"
	"go.khulnasoft.com/api/types/container"
	"go.khulnasoft.com/api/types/versions"
	"go.khulnasoft.com/internal/sliceutil"
	"go.khulnasoft.com/pkg/stringid"
)

// getContainersByName inspects container's configuration and serializes it as json.
func (c *containerRouter) getContainersByName(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error {
	ctr, err := c.backend.ContainerInspect(ctx, vars["name"], backend.ContainerInspectOptions{
		Size: httputils.BoolValue(r, "size"),
	})
	if err != nil {
		return err
	}

	version := httputils.VersionFromContext(ctx)
	if versions.LessThan(version, "1.45") {
		shortCID := stringid.TruncateID(ctr.ID)
		for nwName, ep := range ctr.NetworkSettings.Networks {
			if container.NetworkMode(nwName).IsUserDefined() {
				ep.Aliases = sliceutil.Dedup(append(ep.Aliases, shortCID, ctr.Config.Hostname))
			}
		}
	}
	if versions.LessThan(version, "1.48") {
		ctr.ImageManifestDescriptor = nil
	}

	return httputils.WriteJSON(w, http.StatusOK, ctr)
}
