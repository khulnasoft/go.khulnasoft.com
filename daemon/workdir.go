package daemon // import "go.khulnasoft.com/daemon"

// ContainerCreateWorkdir creates the working directory. This solves the
// issue arising from https://go.khulnasoft.com/issues/27545,
// which was initially fixed by https://go.khulnasoft.com/pull/27884. But that fix
// was too expensive in terms of performance on Windows. Instead,
// https://go.khulnasoft.com/pull/28514 introduces this new functionality
// where the builder calls into the backend here to create the working directory.
func (daemon *Daemon) ContainerCreateWorkdir(cID string) error {
	container, err := daemon.GetContainer(cID)
	if err != nil {
		return err
	}
	err = daemon.Mount(container)
	if err != nil {
		return err
	}
	defer daemon.Unmount(container)
	return container.SetupWorkingDirectory(daemon.idMapping.RootPair())
}
