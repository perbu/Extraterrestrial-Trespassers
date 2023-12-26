package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/perbu/spaceinvaders/assets"
)

type Fleet struct {
	Enemies    []*Enemy
	MovingLeft bool
	Leftmost   int
	Rightmost  int
}

type Enemy struct {
	Sprite   *ebiten.Image
	Position Position
	dead     bool
}

func newFleet(x, y, leftmost, rightmost int) *Fleet {
	sprites := []*ebiten.Image{
		assets.GetGreen(),
		assets.GetRed(),
		assets.GetYellow(),
		assets.GetBlue(),
	}
	f := &Fleet{
		Enemies:   make([]*Enemy, 0, 40),
		Leftmost:  leftmost,
		Rightmost: rightmost,
	}
	for row := 0; row < 4; row++ {
		for col := 0; col < 10; col++ {
			e := &Enemy{
				Sprite: sprites[row],
				Position: Position{
					X: x + col*50,
					Y: y + row*50,
				},
			}
			f.Enemies = append(f.Enemies, e)
		}
	}
	return f
}

func (f *Fleet) Draw(screen *ebiten.Image) {
	for _, e := range f.Enemies {
		e.Draw(screen)
	}
}

func (f *Fleet) Update() {
	switch f.MovingLeft {
	case true:
		for _, e := range f.Enemies {
			if e.Position.X <= f.Leftmost {
				f.MovingLeft = false
				f.Descend(10)
				break
			}
		}
	case false:
		for _, e := range f.Enemies {
			if e.Position.X >= f.Rightmost {
				f.MovingLeft = true
				f.Descend(10)
				break
			}
		}
	}

	for _, e := range f.Enemies {
		e.Update(f.MovingLeft)
	}
}

func (f *Fleet) Descend(n int) {
	for _, e := range f.Enemies {
		e.Position.Y += n
	}
}

func (e *Enemy) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(e.Position.X), float64(e.Position.Y))
	screen.DrawImage(e.Sprite, op)
}

func (e *Enemy) Update(ml bool) {
	switch ml {
	case true:
		e.Position.X -= 1
	case false:
		e.Position.X += 1
	}
}

func (e *Enemy) Collides(p *Projectile) bool {
	x1 := e.Position.X
	y1 := e.Position.Y
	b := e.Sprite.Bounds()
	x2 := e.Position.X + b.Max.X
	y2 := e.Position.Y + b.Max.Y
	res := p.Position.X >= x1 && p.Position.X <= x2 && p.Position.Y >= y1 && p.Position.Y <= y2
	return res
}
