package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/perbu/spaceinvaders/assets"
)

type Player struct {
	Position Position
	Sprite   *ebiten.Image
}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(p.Position.X), float64(p.Position.Y))
	screen.DrawImage(p.Sprite, op)
}

func (p *Player) Shoot() *Projectile {
	return &Projectile{
		Sprite: assets.GetProjectile(),
		Position: Position{
			X: p.Position.X + 20,
			Y: p.Position.Y,
		},
		Speed: 5,
	}
}
