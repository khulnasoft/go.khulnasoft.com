package volume // import "go.khulnasoft.com/api/types/volume"

import "go.khulnasoft.com/api/types/filters"

// ListOptions holds parameters to list volumes.
type ListOptions struct {
	Filters filters.Args
}

// PruneReport contains the response for Engine API:
// POST "/volumes/prune"
type PruneReport struct {
	VolumesDeleted []string
	SpaceReclaimed uint64
}
