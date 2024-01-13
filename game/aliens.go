package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/perbu/extraterrestrial_trespassers/assets"
	"github.com/perbu/extraterrestrial_trespassers/state"
)

type fleet struct {
	enemies    []*enemy
	movingLeft bool
	leftMost   int
	rightMost  int
}

type enemy struct {
	asset    assets.Asset
	position position
	dead     bool
}

func newFleet(x, y int, global *state.Global) *fleet {
	a := []assets.Asset{
		assets.GetGreen(),
		assets.GetRed(),
		assets.GetYellow(),
		assets.GetBlue(),
	}
	width, _ := global.GetDimensions()
	f := &fleet{
		enemies:   make([]*enemy, 0, 40),
		leftMost:  global.GetMargins(),
		rightMost: width - global.GetMargins(),
	}
	for row := 0; row < 4; row++ {
		for col := 0; col < 10; col++ {
			e := &enemy{
				asset: a[row],
				position: position{
					X: x + col*50,
					Y: y + row*50,
				},
			}
			f.enemies = append(f.enemies, e)
		}
	}
	return f
}

func (f *fleet) Draw(screen *ebiten.Image) {
	for _, e := range f.enemies {
		e.Draw(screen)
	}
}

func (f *fleet) Update() {
	switch f.movingLeft {
	case true:
		for _, e := range f.enemies {
			if e.position.X <= f.leftMost {
				f.movingLeft = false
				f.Descend(10)
				break
			}
		}
	case false:
		for _, e := range f.enemies {
			if e.position.X >= f.rightMost {
				f.movingLeft = true
				f.Descend(10)
				break
			}
		}
	}

	for _, e := range f.enemies {
		e.Update(f.movingLeft)
	}
}

func (f *fleet) Descend(n int) {
	for _, e := range f.enemies {
		e.position.Y += n
	}
}

func (e *enemy) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(e.position.X), float64(e.position.Y))
	screen.DrawImage(e.asset.Sprite, op)
}

func (e *enemy) Update(ml bool) {
	switch ml {
	case true:
		e.position.X -= 1
	case false:
		e.position.X += 1
	}
}
