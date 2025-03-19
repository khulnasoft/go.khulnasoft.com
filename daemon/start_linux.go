package daemon // import "go.khulnasoft.com/daemon"

import (
	"context"
	"fmt"

	"go.khulnasoft.com/container"
	"go.khulnasoft.com/daemon/config"
	"go.khulnasoft.com/errdefs"
	"go.khulnasoft.com/libcontainerd/types"
	"go.khulnasoft.com/oci"
	"github.com/opencontainers/runtime-spec/specs-go"
)

// initializeCreatedTask performs any initialization that needs to be done to
// prepare a freshly-created task to be started.
func (daemon *Daemon) initializeCreatedTask(
	ctx context.Context,
	cfg *config.Config,
	tsk types.Task,
	ctr *container.Container,
	spec *specs.Spec,
) error {
	if ctr.Config.NetworkDisabled {
		return nil
	}
	nspath, ok := oci.NamespacePath(spec, specs.NetworkNamespace)
	if ok && nspath == "" { // the runtime has been instructed to create a new network namespace for tsk.
		sb, err := daemon.netController.GetSandbox(ctr.ID)
		if err != nil {
			return errdefs.System(err)
		}
		if err := sb.SetKey(ctx, fmt.Sprintf("/proc/%d/ns/net", tsk.Pid())); err != nil {
			return errdefs.System(err)
		}
	}
	if err := daemon.allocateNetwork(ctx, cfg, ctr); err != nil {
		return fmt.Errorf("%s: %w", errSetupNetworking, err)
	}
	return nil
}
