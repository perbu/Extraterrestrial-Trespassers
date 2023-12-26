package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/perbu/spaceinvaders/assets"
)

type Life struct {
	Position  Position
	Sprite    *ebiten.Image
	NoOfLives int
}

func NewLife(x, y, noOfLives int) *Life {
	return &Life{
		Position: Position{
			X: x,
			Y: y,
		},
		Sprite:    assets.GetPlayer(),
		NoOfLives: noOfLives,
	}
}

func (l *Life) Draw(screen *ebiten.Image) {
	for i := 0; i < l.NoOfLives; i++ {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(l.Position.X+i*100), float64(l.Position.Y))
		screen.DrawImage(l.Sprite, op)
	}
}
