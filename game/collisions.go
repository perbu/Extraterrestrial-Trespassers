package game

import (
	"github.com/perbu/extraterrestrial_trespassers/assets"
	"image"
)

// collides takes two assets and their positions and returns true if they collide.
func collides(aAsset assets.Asset, aPosition position, bAsset assets.Asset, bPosition position) bool {
	// Adjust the bounds of each asset by their positions.
	aBounds := image.Rect(
		aAsset.Bounds.Min.X+aPosition.x,
		aAsset.Bounds.Min.Y+aPosition.y,
		aAsset.Bounds.Max.X+aPosition.x,
		aAsset.Bounds.Max.Y+aPosition.y,
	)

	bBounds := image.Rect(
		bAsset.Bounds.Min.X+bPosition.x,
		bAsset.Bounds.Min.Y+bPosition.y,
		bAsset.Bounds.Max.X+bPosition.x,
		bAsset.Bounds.Max.Y+bPosition.y,
	)

	// Check for intersection.
	return aBounds.Overlaps(bBounds)
}
