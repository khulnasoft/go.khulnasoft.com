package checkpoint // import "go.khulnasoft.com/api/server/router/checkpoint"

import "go.khulnasoft.com/api/types/checkpoint"

// Backend for Checkpoint
type Backend interface {
	CheckpointCreate(container string, config checkpoint.CreateOptions) error
	CheckpointDelete(container string, config checkpoint.DeleteOptions) error
	CheckpointList(container string, config checkpoint.ListOptions) ([]checkpoint.Summary, error)
}
