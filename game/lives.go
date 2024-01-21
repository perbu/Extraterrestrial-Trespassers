package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/perbu/extraterrestrial_trespassers/assets"
)

type lives struct {
	position position
	asset    assets.Asset
	game     *Game
	lives    int
}

func NewLife(x, y int, game *Game) *lives {
	return &lives{
		position: position{
			x: x,
			y: y,
		},
		asset: assets.GetPlayer(),
		game:  game,
		lives: 2,
	}
}

func (l *lives) Draw(screen *ebiten.Image) {
	lives := l.lives
	for i := 0; i < lives; i++ {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(l.position.x+i*100), float64(l.position.y))
		op.GeoM.Scale(.75, 0.75)
		screen.DrawImage(l.asset.Sprite, op)
	}
}

func (l *lives) GetLives() int {
	return l.lives
}

func (l *lives) DecrementLives() int {
	l.lives--
	return l.lives
}

func (l *lives) IncrementLives() int {
	l.lives++
	return l.lives
}
