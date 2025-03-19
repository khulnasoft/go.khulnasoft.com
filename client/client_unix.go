//go:build !windows

package client // import "go.khulnasoft.com/client"

// DefaultDockerHost defines OS-specific default host if the DOCKER_HOST
// (EnvOverrideHost) environment variable is unset or empty.
const DefaultDockerHost = "unix:///var/run/docker.sock"
