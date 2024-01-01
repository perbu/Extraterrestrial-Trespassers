package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/perbu/extraterrestrial_trespassers/assets"
)

type Life struct {
	Position  Position
	Asset     assets.Asset
	NoOfLives int
}

func NewLife(x, y, noOfLives int) *Life {
	return &Life{
		Position: Position{
			X: x,
			Y: y,
		},
		Asset:     assets.GetPlayer(),
		NoOfLives: noOfLives,
	}
}

func (l *Life) Draw(screen *ebiten.Image) {
	for i := 0; i < l.NoOfLives; i++ {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(l.Position.X+i*100), float64(l.Position.Y))
		op.GeoM.Scale(.5, 0.5)
		screen.DrawImage(l.Asset.Sprite, op)
	}
}
