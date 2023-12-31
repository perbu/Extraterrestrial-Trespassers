package game

import (
	"github.com/perbu/spaceinvaders/assets"
	"image"
	"image/color"
	"testing"
)

// TestCollides will test various scenarios for the Collides function.
func TestCollides(t *testing.T) {
	// Create dummy assets and positions for testing.
	asset1 := assets.Asset{Bounds: image.Rect(0, 0, 10, 10)}
	asset2 := assets.Asset{Bounds: image.Rect(0, 0, 5, 5)}

	// Test Case 1: Assets are overlapping
	pos1 := Position{X: 0, Y: 0}
	pos2 := Position{X: 5, Y: 5}
	if !Collides(asset1, pos1, asset2, pos2) {
		t.Errorf("Collides was incorrect, got: false, want: true.")
	}

	// Test Case 2: Assets are not overlapping
	pos2 = Position{X: 20, Y: 20}
	if Collides(asset1, pos1, asset2, pos2) {
		t.Errorf("Collides was incorrect, got: true, want: false.")
	}

	// Test Case 3: Assets are touching edges but not overlapping
	asset3 := assets.Asset{Bounds: image.Rect(0, 0, 10, 10)}
	pos3 := Position{X: 10, Y: 10}
	if Collides(asset1, pos1, asset3, pos3) {
		t.Errorf("Collides was incorrect, got: true, want: false.")
	}

	// Additional test cases can be added here.
}

// Dummy ebiten.Image to satisfy the Asset struct.
// You should replace this with a proper mock or dummy object as per your setup.
type dummyImage struct{}

func (d *dummyImage) Dispose()                    {}
func (d *dummyImage) IsInvalidated() bool         { return false }
func (d *dummyImage) Size() (int, int)            { return 0, 0 }
func (d *dummyImage) Bounds() image.Rectangle     { return image.Rect(0, 0, 0, 0) }
func (d *dummyImage) At(x, y int) color.Color     { return color.RGBA{} }
func (d *dummyImage) Set(x, y int, c color.Color) {}
