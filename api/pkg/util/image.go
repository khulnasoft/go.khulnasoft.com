package util

import (
	"go.khulnasoft.com/api/internal/image"
)

// Splits image string name into name, tag and digest
func SplitImageName(imageName string) (name string, tag string, digest string) {
	return image.Split(imageName)
}
