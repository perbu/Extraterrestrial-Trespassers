package game

import (
	"github.com/perbu/extraterrestrial_trespassers/assets"
	"image"
)

// collides takes two assets and their positions and returns true if they collide.
func collides(aAsset assets.Asset, aPosition position, bAsset assets.Asset, bPosition position) bool {
	// Adjust the bounds of each asset by their positions.
	aBounds := image.Rect(
		aAsset.Bounds.Min.X+aPosition.X,
		aAsset.Bounds.Min.Y+aPosition.Y,
		aAsset.Bounds.Max.X+aPosition.X,
		aAsset.Bounds.Max.Y+aPosition.Y,
	)

	bBounds := image.Rect(
		bAsset.Bounds.Min.X+bPosition.X,
		bAsset.Bounds.Min.Y+bPosition.Y,
		bAsset.Bounds.Max.X+bPosition.X,
		bAsset.Bounds.Max.Y+bPosition.Y,
	)

	// Check for intersection.
	return aBounds.Overlaps(bBounds)
}
