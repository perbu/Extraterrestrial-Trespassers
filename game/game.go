package game

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/perbu/extraterrestrial_trespassers/state"
	"image/color"
	"math/rand"
	"time"
)

type Game struct {
	player      player
	alienFleet  *fleet
	projectiles []*projectile
	bombs       []*bomb
	lives       *lives
	state       *state.Global
	particles   []*particle
}

type position struct {
	X int
	Y int
}

func NewGame(aud *audio.Context, global *state.Global) *Game {
	g := &Game{}
	g.lives = NewLife(0, 0, g)
	g.alienFleet = newFleet(0, 30, global)
	g.bombs = make([]*bomb, 0, 10)
	g.player = NewPlayer(aud, global, g)
	g.state = global
	g.particles = make([]*particle, 0)
	return g
}

func (g *Game) Update() error {
	if g.player.dead {
		if g.lives.GetLives() < 0 {
			g.state.QueueAction(state.GameOver)
			return nil
		}
		g.player.Respawn()
		// freeze the game for 2 seconds
		g.state.FreezeUntil(time.Now().Add(2 * time.Second))
		g.player.dead = false
	}

	g.player.Update()
	for _, p := range g.projectiles {
		p.Update()
	}
	g.alienFleet.Update()
	g.projectiles = filterProjectiles(g.projectiles)

	// update the particles:
	for i, p := range g.particles {
		p.Update()
		// check if the particle is within bounds:
		if !g.withinBounds(p.position) {
			// remove the particle from the slice
			g.particles = removeElement(g.particles, i)
		}
		if p.age() > particleLifeTime {
			g.particles = removeElement(g.particles, i)
		}
	}

	// Check for collisions between projectiles and enemies
	for _, e := range g.alienFleet.enemies {
		for _, p := range g.projectiles {
			if collides(e.asset, e.position, p.asset, p.position) {
				// remove the projectile from the screen
				p.position.Y = -10
				// remove the enemy from the fleet:
				e.dead = true
				var pcolor color.Color
				// check what type of enemy was hit:

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
				for i := 0; i < 100; i++ {
					g.particles = append(g.particles, newParticle(e.position, pcolor))
				}
			}
		}
	}
	// remove dead enemies from the fleet
	g.alienFleet.enemies = filterEnemies(g.alienFleet.enemies)

	// Drop the bombs:
	for _, e := range g.alienFleet.enemies {
		// 1% chance of dropping a bomb
		if rand.Intn(1000) == 1 {
			b := newBomb(e.position.X, e.position.Y, 5)
			g.bombs = append(g.bombs, b)
		}
	}
	for _, b := range g.bombs {
		b.Update()
	}
	g.bombs = filterBombs(g.bombs, g.state.GetHeight())

	// Check for collisions between bombs and player
	for _, b := range g.bombs {
		if collides(b.asset, b.position, g.player.asset, g.player.position) {
			fmt.Println("player hit by bomb")
			g.player.Collision()
			b.position.Y = -10

		}
	}
	// check for collisions between bombs and projectiles
	for _, b := range g.bombs {
		for _, p := range g.projectiles {
			if collides(b.asset, b.position, p.asset, p.position) {
				b.position.Y = -10
				p.position.Y = -10
			}
		}
	}
	// check for collisions between player and enemies:
	for _, e := range g.alienFleet.enemies {
		if collides(e.asset, e.position, g.player.asset, g.player.position) {
			fmt.Println("player hit by enemy")
			g.player.Collision()
			e.dead = true
		}
	}
	return nil
}

func filterEnemies(enemies []*enemy) []*enemy {
	var filtered []*enemy
	for _, e := range enemies {
		if !e.dead {
			filtered = append(filtered, e)
		}
	}
	return filtered
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.lives.Draw(screen)
	g.player.Draw(screen)
	for _, p := range g.projectiles {
		p.Draw(screen)
	}
	g.alienFleet.Draw(screen)
	for _, b := range g.bombs {
		b.Draw(screen)
	}
	for _, p := range g.particles {
		p.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.state.GetDimensions()
}

func removeElement[T any](s []T, index int) []T {
	if index < 0 || index >= len(s) {
		// If the index is out of range, return the original slice
		return s
	}
	return append(s[:index], s[index+1:]...)
}
