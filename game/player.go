package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/perbu/extraterrestrial_trespassers/assets"
)

type Player struct {
	Position Position
	Asset    assets.Asset
}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(p.Position.X), float64(p.Position.Y))
	screen.DrawImage(p.Asset.Sprite, op)
}

func (p *Player) Shoot() *Projectile {
	return &Projectile{
		Asset: assets.GetProjectile(),
		Position: Position{
			X: p.Position.X + 20,
			Y: p.Position.Y,
		},
		Speed: 5,
	}
}
