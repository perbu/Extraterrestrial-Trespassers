package game

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/perbu/extraterrestrial_trespassers/state"
	"math/rand"
	"time"
)

type Game struct {
	Player      Player
	AlienFleet  *Fleet
	Projectiles []*Projectile
	Bombs       []*Bomb
	Lives       *Lives
	GameOver    *GameOver
	GameIsOver  bool
	state       *state.Global
}

type Position struct {
	X int
	Y int
}

func NewGame(aud *audio.Context, global *state.Global) *Game {
	g := &Game{}
	g.Lives = NewLife(0, 0, g)
	g.AlienFleet = newFleet(0, 30, global)
	g.Bombs = make([]*Bomb, 0, 10)
	g.Player = NewPlayer(aud, global, g)
	g.state = global
	return g
}

func (g *Game) Update() error {
	if g.Player.dead {
		if g.Lives.GetLives() < 0 {
			g.state.QueueAction(state.GameOver)
			return nil
		}
		g.Player.Respawn()
		// freeze the game for 2 seconds
		g.state.FreezeUntil(time.Now().Add(2 * time.Second))
		g.Player.dead = false
	}

	g.Player.Update()
	for _, p := range g.Projectiles {
		p.Update()
	}
	g.AlienFleet.Update()
	g.Projectiles = filterProjectiles(g.Projectiles)

	// Check for collisions between projectiles and enemies
	for _, e := range g.AlienFleet.Enemies {
		for _, p := range g.Projectiles {
			if Collides(e.Asset, e.Position, p.Asset, p.Position) {
				// remove the projectile from the screen
				p.Position.Y = -10
				// remove the enemy from the fleet:
				e.dead = true
			}
		}
	}
	// remove dead enemies from the fleet
	g.AlienFleet.Enemies = filterEnemies(g.AlienFleet.Enemies)

	// Drop the bombs:
	for _, e := range g.AlienFleet.Enemies {
		// 1% chance of dropping a bomb
		if rand.Intn(1000) == 1 {
			b := newBomb(e.Position.X, e.Position.Y, 5)
			g.Bombs = append(g.Bombs, b)
		}
	}
	for _, b := range g.Bombs {
		b.Update()
	}
	g.Bombs = filterBombs(g.Bombs, g.state.GetHeight())

	// Check for collisions between bombs and player
	for _, b := range g.Bombs {
		if Collides(b.Asset, b.Position, g.Player.Asset, g.Player.Position) {
			fmt.Println("Player hit by bomb")
			g.Player.Collision()
			b.Position.Y = -10

		}
	}
	// check for collisions between bombs and projectiles
	for _, b := range g.Bombs {
		for _, p := range g.Projectiles {
			if Collides(b.Asset, b.Position, p.Asset, p.Position) {
				b.Position.Y = -10
				p.Position.Y = -10
			}
		}
	}
	// check for collisions between player and enemies:
	for _, e := range g.AlienFleet.Enemies {
		if Collides(e.Asset, e.Position, g.Player.Asset, g.Player.Position) {
			fmt.Println("Player hit by enemy")
			g.Player.Collision()
			e.dead = true
		}
	}
	return nil
}

func filterEnemies(enemies []*Enemy) []*Enemy {
	var filtered []*Enemy
	for _, e := range enemies {
		if !e.dead {
			filtered = append(filtered, e)
		}
	}
	return filtered
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Lives.Draw(screen)
	g.Player.Draw(screen)
	for _, p := range g.Projectiles {
		p.Draw(screen)
	}
	g.AlienFleet.Draw(screen)
	for _, b := range g.Bombs {
		b.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.state.GetDimensions()
}
