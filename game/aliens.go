package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/perbu/extraterrestrial_trespassers/assets"
	"github.com/perbu/extraterrestrial_trespassers/state"
	"image/color"
	"math/rand"
)

type fleet struct {
	enemies         []*enemy
	movingLeft      bool
	leftMost        int
	rightMost       int
	availableAssets []assets.Asset
	aud             *audio.Context
	game            *Game
	level           int
}

type enemy struct {
	asset     assets.Asset
	position  position
	dead      bool
	enemyType enemyType
	pl        *audio.Player
	game      *Game
	fleet     *fleet
}

type enemyType int

const (
	enemyGreen enemyType = iota
	enemyRed
	enemyYellow
	enemyCyan
)

const (
	maxEnemySpeed = 20
)

func (g *Game) newFleet(x, y int, global *state.Global, aud *audio.Context, level int) *fleet {
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
		game:  g,
		level: level,
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
				f.Descend(20)
				break
			}
		}
	case false:
		for _, e := range f.enemies {
			if e.position.x >= f.rightMost {
				f.movingLeft = true
				f.Descend(20)
				break
			}
		}
	}
	// calculate the speed of the enemies. the fewer enemies, the faster they move

	ratio := 1.0 - (float64(f.count()) / 40.0) // is 0 initially, and 1 when there are no enemies left
	speed := int(maxEnemySpeed*ratio) + 2

	for _, e := range f.enemies {
		e.Update(f.movingLeft, speed)
	}
}

func (f *fleet) Descend(n int) {
	for _, e := range f.enemies {
		e.position.y += n
	}
}

func (f *fleet) count() int {
	return len(f.enemies)
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
		fleet:     f,
	}

}

func (e *enemy) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(e.position.x), float64(e.position.y))
	screen.DrawImage(e.asset.Sprite, op)
}

func (e *enemy) Update(ml bool, speed int) {

	switch ml {
	case true:
		e.position.x -= speed
	case false:
		e.position.x += speed
	}
	// Drop the bombs. The higher the speed the higher the chance of dropping a bomb
	if rand.Intn(1000) < speed {
		e.game.bombs = append(e.game.bombs, newBomb(e.position.x, e.position.y))
	}

}

func (e *enemy) kill() {
	e.pl.Play()
	e.dead = true
	var pcolor color.Color
	var score int
	switch e.enemyType {
	case enemyGreen:
		pcolor = color.RGBA{R: 0, G: 255, B: 0, A: 255}
		score = 5
	case enemyRed:
		pcolor = color.RGBA{R: 255, G: 0, B: 0, A: 255}
		score = 3
	case enemyYellow:
		pcolor = color.RGBA{R: 255, G: 255, B: 0, A: 255}
		score = 2
	case enemyCyan:
		pcolor = color.RGBA{R: 0, G: 255, B: 255, A: 255}
		score = 1
	}
	e.game.score.AddScore(score)
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
