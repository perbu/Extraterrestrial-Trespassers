package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/perbu/spaceinvaders/assets"
)

type Bomb struct {
	Position Position
	Sprite   *ebiten.Image
	Speed    int
}

func newBomb(x, y, speed int) *Bomb {
	return &Bomb{
		Position: Position{
			X: x,
			Y: y,
		},
		Sprite: assets.GetBomb(),
		Speed:  speed,
	}
}
func (b *Bomb) Update() {
	b.Position.Y += b.Speed
}

func (b *Bomb) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(b.Position.X), float64(b.Position.Y))
	screen.DrawImage(b.Sprite, op)
}

func filterBombs(bs []*Bomb) []*Bomb {
	ret := make([]*Bomb, 0)
	for _, b := range bs {
		if b.Position.Y < 600 {
			ret = append(ret, b)
		}
	}
	return ret
}
