package images // import "go.khulnasoft.com/daemon/images"

import (
	"context"

	"go.khulnasoft.com/api/types/backend"
	"go.khulnasoft.com/api/types/events"
)

// LogImageEvent generates an event related to an image with only the default attributes.
func (i *ImageService) LogImageEvent(ctx context.Context, imageID, refName string, action events.Action) {
	ctx = context.WithoutCancel(ctx)
	attributes := map[string]string{}

	img, err := i.GetImage(ctx, imageID, backend.GetImageOpts{})
	if err == nil && img.Config != nil {
		// image has not been removed yet.
		// it could be missing if the event is `delete`.
		copyAttributes(attributes, img.Config.Labels)
	}
	if refName != "" {
		attributes["name"] = refName
	}
	i.eventsService.Log(action, events.ImageEventType, events.Actor{
		ID:         imageID,
		Attributes: attributes,
	})
}

// copyAttributes guarantees that labels are not mutated by event triggers.
func copyAttributes(attributes, labels map[string]string) {
	if labels == nil {
		return
	}
	for k, v := range labels {
		attributes[k] = v
	}
}
