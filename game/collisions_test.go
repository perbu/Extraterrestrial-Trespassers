package game

import (
	"github.com/perbu/spaceinvaders/assets"
	"image"
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
