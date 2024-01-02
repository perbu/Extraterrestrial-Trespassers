package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/perbu/extraterrestrial_trespassers/assets"
)

type Bomb struct {
	Position Position
	Asset    assets.Asset
	Speed    int
}

func newBomb(x, y, speed int) *Bomb {
	return &Bomb{
		Position: Position{
			X: x,
			Y: y,
		},
		Asset: assets.GetBomb(),
		Speed: speed,
	}
}
func (b *Bomb) Update() {
	b.Position.Y += b.Speed
}

func (b *Bomb) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(b.Position.X), float64(b.Position.Y))
	screen.DrawImage(b.Asset.Sprite, op)
}

func filterBombs(bs []*Bomb, maxHeight int) []*Bomb {
	ret := make([]*Bomb, 0)
	for _, b := range bs {
		if b.Position.Y < maxHeight {
			ret = append(ret, b)
		}
	}
	return ret
}

func (b *Bomb) Collides(p *Player) bool {
	x1 := b.Position.X
	y1 := b.Position.Y
	bounds := b.Asset.Sprite.Bounds()
	x2 := b.Position.X + bounds.Max.X
	y2 := b.Position.Y + bounds.Max.Y
	res := p.Position.X >= x1 && p.Position.X <= x2 && p.Position.Y >= y1 && p.Position.Y <= y2
	return res
}
