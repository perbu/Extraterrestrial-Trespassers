package game

import (
	"github.com/perbu/extraterrestrial_trespassers/assets"
	"image"
)

// Collides takes two assets and their positions and returns true if they collide.
func Collides(aAsset assets.Asset, aPosition Position, bAsset assets.Asset, bPosition Position) bool {
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
