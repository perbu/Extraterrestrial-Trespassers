package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/perbu/extraterrestrial_trespassers/assets"
	"github.com/perbu/extraterrestrial_trespassers/state"
)

type Life struct {
	Position  Position
	Asset     assets.Asset
	NoOfLives int
	global    *state.Global
}

func NewLife(x, y int, global *state.Global) *Life {
	return &Life{
		Position: Position{
			X: x,
			Y: y,
		},
		Asset:     assets.GetPlayer(),
		NoOfLives: 2,
		global:    global,
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

func (l *Life) Die() {
	if l.NoOfLives == 0 {
		l.global.QueueAction(state.GameOver)
	}
	l.NoOfLives--
}
