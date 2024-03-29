package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/perbu/extraterrestrial_trespassers/assets"
	"github.com/perbu/extraterrestrial_trespassers/state"
	"image/color"
	"math/rand"
	"time"
)

type fleet struct {
	enemies         []*enemy
	movingLeft      bool
	leftMost        int
	rightMost       int
	availableAssets []assets.Asset
	audio           *audio.Context
	game            *Game
	level           int
	bombFactor      int
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
	maxEnemySpeed = 12 // level * 3 is added
)

func (g *Game) newFleet(global *state.Global, audio *audio.Context, level int) *fleet {
	width, _ := global.GetDimensions()
	f := &fleet{
		audio:     audio,
		enemies:   make([]*enemy, 0, 40),
		leftMost:  global.GetMargins(),
		rightMost: width - global.GetMargins(),
		availableAssets: []assets.Asset{
			assets.GetGreen(),
			assets.GetRed(),
			assets.GetYellow(),
			assets.GetCyan(),
		},
		game:       g,
		level:      level,
		bombFactor: 1300,
	}
	f.populate()
	return f
}
func (f *fleet) populate() {
	for row := 0; row < 4; row++ {
		for col := 0; col < 10; col++ {
			pos := position{
				x: 0 + col*70,
				y: 30 + row*50,
			}
			e := f.spawnAlien(pos, enemyType(row), f.audio)
			f.enemies = append(f.enemies, e)
		}
	}

}

func (f *fleet) Draw(screen *ebiten.Image) {
	for _, e := range f.enemies {
		e.Draw(screen)
	}
}

func (f *fleet) Update() {
	// check if there are enemies left:
	if f.count() == 0 {
		f.level++
		f.populate()
		f.game.state.FreezeUntil(time.Now().Add(2 * time.Second))
		f.game.lives.IncrementLives()
	}

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
	speed := int(maxEnemySpeed*ratio) + f.level*3

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
	if rand.Intn(e.fleet.bombFactor) < speed {
		e.game.bombs = append(e.game.bombs, newBomb(e.position.x, e.position.y, speed/2))
	}

}

func (e *enemy) kill() {
	vol := e.fleet.game.state.GetVolume()
	e.pl.SetVolume(vol)
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
