package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/perbu/extraterrestrial_trespassers/assets"
)

// projectile is shot by the player at the enemy
type projectile struct {
	position position
	speed    int
	asset    assets.Asset
}

func (p *projectile) Update() {
	p.position.Y -= p.speed
}

func (p *projectile) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(p.position.X), float64(p.position.Y))
	screen.DrawImage(p.asset.Sprite, op)
}

func filterProjectiles(projectiles []*projectile) []*projectile {
	np := make([]*projectile, 0, len(projectiles))
	for _, p := range projectiles {
		if p.position.Y > 0 {
			np = append(np, p)
		}
	}
	return np
}
