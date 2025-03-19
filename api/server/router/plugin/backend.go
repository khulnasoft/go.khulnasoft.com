package plugin // import "go.khulnasoft.com/api/server/router/plugin"

import (
	"context"
	"io"
	"net/http"

	"github.com/distribution/reference"
	"go.khulnasoft.com/api/types"
	"go.khulnasoft.com/api/types/backend"
	"go.khulnasoft.com/api/types/filters"
	"go.khulnasoft.com/api/types/registry"
	"go.khulnasoft.com/plugin"
)

// Backend for Plugin
type Backend interface {
	Disable(name string, config *backend.PluginDisableConfig) error
	Enable(name string, config *backend.PluginEnableConfig) error
	List(filters.Args) ([]types.Plugin, error)
	Inspect(name string) (*types.Plugin, error)
	Remove(name string, config *backend.PluginRmConfig) error
	Set(name string, args []string) error
	Privileges(ctx context.Context, ref reference.Named, metaHeaders http.Header, authConfig *registry.AuthConfig) (types.PluginPrivileges, error)
	Pull(ctx context.Context, ref reference.Named, name string, metaHeaders http.Header, authConfig *registry.AuthConfig, privileges types.PluginPrivileges, outStream io.Writer, opts ...plugin.CreateOpt) error
	Push(ctx context.Context, name string, metaHeaders http.Header, authConfig *registry.AuthConfig, outStream io.Writer) error
	Upgrade(ctx context.Context, ref reference.Named, name string, metaHeaders http.Header, authConfig *registry.AuthConfig, privileges types.PluginPrivileges, outStream io.Writer) error
	CreateFromContext(ctx context.Context, tarCtx io.ReadCloser, options *types.PluginCreateOptions) error
}
