package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/perbu/extraterrestrial_trespassers/assets"
	"math/rand"
)

type bomb struct {
	position position
	asset    assets.Asset
	speed    int
}

func newBomb(x, y, maxSpeed int) *bomb {
	return &bomb{
		position: position{
			x: x,
			y: y,
		},
		asset: assets.GetBomb(),
		speed: rand.Intn(maxSpeed) + 2,
	}
}
func (b *bomb) Update() {
	b.position.y += b.speed
}

func (b *bomb) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(b.position.x), float64(b.position.y))
	screen.DrawImage(b.asset.Sprite, op)
}

func filterBombs(bs []*bomb, maxHeight int) []*bomb {
	ret := make([]*bomb, 0)
	for _, b := range bs {
		if b.position.y < maxHeight {
			ret = append(ret, b)
		}
	}
	return ret
}
