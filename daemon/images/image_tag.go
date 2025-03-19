package images // import "go.khulnasoft.com/daemon/images"

import (
	"context"

	"github.com/distribution/reference"
	"go.khulnasoft.com/api/types/events"
	"go.khulnasoft.com/image"
)

// TagImage adds the given reference to the image ID provided.
func (i *ImageService) TagImage(ctx context.Context, imageID image.ID, newTag reference.Named) error {
	if err := i.referenceStore.AddTag(newTag, imageID.Digest(), true); err != nil {
		return err
	}

	if err := i.imageStore.SetLastUpdated(imageID); err != nil {
		return err
	}
	i.LogImageEvent(ctx, imageID.String(), reference.FamiliarString(newTag), events.ActionTag)
	return nil
}
