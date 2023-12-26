package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// Projectile is shot by the player at the enemy
type Projectile struct {
	Position Position
	Speed    int
	Sprite   *ebiten.Image
}

func (p *Projectile) Update() {
	p.Position.Y -= p.Speed
}

func (p *Projectile) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(p.Position.X), float64(p.Position.Y))
	screen.DrawImage(p.Sprite, op)
}

func filterProjectiles(projectiles []*Projectile) []*Projectile {
	np := make([]*Projectile, 0, len(projectiles))
	for _, p := range projectiles {
		if p.Position.Y > 0 {
			np = append(np, p)
		}
	}
	return np
}
