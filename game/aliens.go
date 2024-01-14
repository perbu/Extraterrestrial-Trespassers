package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/perbu/extraterrestrial_trespassers/assets"
	"github.com/perbu/extraterrestrial_trespassers/state"
	"image/color"
)

type fleet struct {
	enemies         []*enemy
	movingLeft      bool
	leftMost        int
	rightMost       int
	availableAssets []assets.Asset
	aud             *audio.Context
	game            *Game
}

type enemy struct {
	asset     assets.Asset
	position  position
	dead      bool
	enemyType enemyType
	pl        *audio.Player
	game      *Game
}

type enemyType int

const (
	enemyGreen enemyType = iota
	enemyRed
	enemyYellow
	enemyCyan
)

func (g *Game) newFleet(x, y int, global *state.Global, aud *audio.Context) *fleet {
	width, _ := global.GetDimensions()
	f := &fleet{
		aud:       aud,
		enemies:   make([]*enemy, 0, 40),
		leftMost:  global.GetMargins(),
		rightMost: width - global.GetMargins(),
		availableAssets: []assets.Asset{
			assets.GetGreen(),
			assets.GetRed(),
			assets.GetYellow(),
			assets.GetCyan(),
		},
		game: g,
	}
	for row := 0; row < 4; row++ {
		for col := 0; col < 10; col++ {
			pos := position{
				x: x + col*50,
				y: y + row*50,
			}
			e := f.spawnAlien(pos, enemyType(row), aud)
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
			if e.position.x <= f.leftMost {
				f.movingLeft = false
				f.Descend(10)
				break
			}
		}
	case false:
		for _, e := range f.enemies {
			if e.position.x >= f.rightMost {
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
		e.position.y += n
	}
}

// spawnAlien spawns an alien at the given position
func (f *fleet) spawnAlien(pos position, et enemyType, aud *audio.Context) *enemy {
	sound := assets.GetExplosion()
	pl, _ := aud.NewPlayer(sound)

	return &enemy{
		enemyType: et,
		asset:     f.availableAssets[et],
		position:  pos,
		pl:        pl,
		game:      f.game,
	}

}

func (e *enemy) Draw(screen *ebiten.Image) {

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(e.position.x), float64(e.position.y))
	screen.DrawImage(e.asset.Sprite, op)
}

func (e *enemy) Update(ml bool) {
	switch ml {
	case true:
		e.position.x -= 1
	case false:
		e.position.x += 1
	}
}

func (e *enemy) kill() {
	e.pl.Play()
	e.dead = true
	var pcolor color.Color
	switch e.enemyType {
	case enemyGreen:
		pcolor = color.RGBA{R: 0, G: 255, B: 0, A: 255}
	case enemyRed:
		pcolor = color.RGBA{R: 255, G: 0, B: 0, A: 255}
	case enemyYellow:
		pcolor = color.RGBA{R: 255, G: 255, B: 0, A: 255}
	case enemyCyan:
		pcolor = color.RGBA{R: 0, G: 255, B: 255, A: 255}
	}
	center := e.findCenter()
	for i := 0; i < 200; i++ {
		e.game.particles = append(e.game.particles, newParticle(center, pcolor))
	}

}

// findCenter returns the center position of the enemy
func (e *enemy) findCenter() position {
	bounds := e.asset.Bounds

	return position{
		x: e.position.x + bounds.Dx()/2,
		y: e.position.y + bounds.Dy()/2,
	}

}
