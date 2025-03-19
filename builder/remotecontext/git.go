package remotecontext // import "go.khulnasoft.com/builder/remotecontext"

import (
	"context"
	"os"

	"github.com/containerd/log"
	"go.khulnasoft.com/builder"
	"go.khulnasoft.com/builder/remotecontext/git"
	"go.khulnasoft.com/pkg/archive"
)

// MakeGitContext returns a Context from gitURL that is cloned in a temporary directory.
func MakeGitContext(gitURL string) (builder.Source, error) {
	root, err := git.Clone(gitURL, git.WithIsolatedConfig(true))
	if err != nil {
		return nil, err
	}

	c, err := archive.Tar(root, archive.Uncompressed)
	if err != nil {
		return nil, err
	}

	defer func() {
		err := c.Close()
		if err != nil {
			log.G(context.TODO()).WithField("action", "MakeGitContext").WithField("module", "builder").WithField("url", gitURL).WithError(err).Error("error while closing git context")
		}
		err = os.RemoveAll(root)
		if err != nil {
			log.G(context.TODO()).WithField("action", "MakeGitContext").WithField("module", "builder").WithField("url", gitURL).WithError(err).Error("error while removing path and children of root")
		}
	}()
	return FromArchive(c)
}
