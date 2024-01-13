package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/perbu/extraterrestrial_trespassers/assets"
)

type bomb struct {
	position position
	asset    assets.Asset
	speed    int
}

func newBomb(x, y, speed int) *bomb {
	return &bomb{
		position: position{
			X: x,
			Y: y,
		},
		asset: assets.GetBomb(),
		speed: speed,
	}
}
func (b *bomb) Update() {
	b.position.Y += b.speed
}

func (b *bomb) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(b.position.X), float64(b.position.Y))
	screen.DrawImage(b.asset.Sprite, op)
}

func filterBombs(bs []*bomb, maxHeight int) []*bomb {
	ret := make([]*bomb, 0)
	for _, b := range bs {
		if b.position.Y < maxHeight {
			ret = append(ret, b)
		}
	}
	return ret
}
