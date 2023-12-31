package game

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/perbu/spaceinvaders/assets"
	"math/rand"
)

var (
	gameWidth  = 800
	gameHeight = 600
	gameMargin = 50
)

type Game struct {
	Player      Player
	AlienFleet  *Fleet
	Projectiles []*Projectile
	Bombs       []*Bomb
	Lives       *Life
}

type Position struct {
	X int
	Y int
}

func New() *Game {
	return &Game{
		Lives:      NewLife(0, 0, 2),
		AlienFleet: newFleet(0, 30, gameMargin, gameWidth-gameMargin),
		Bombs:      make([]*Bomb, 0, 10),
		Player: Player{
			Position: Position{
				X: gameWidth / 2,
				Y: gameHeight - 50,
			},
			Asset: assets.GetPlayer(),
		},
	}
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		if g.Player.Position.X > 0 {
			g.Player.Position.X -= 5
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		if g.Player.Position.X < gameWidth-50 {
			g.Player.Position.X += 5
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		p := g.Player.Shoot()
		g.Projectiles = append(g.Projectiles, p)
	}
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
			b := newBomb(e.Position.X, e.Position.Y, 2)
			g.Bombs = append(g.Bombs, b)
		}
	}
	for _, b := range g.Bombs {
		b.Update()
	}
	g.Bombs = filterBombs(g.Bombs)

	// Check for collisions between bombs and player
	for _, b := range g.Bombs {
		if Collides(b.Asset, b.Position, g.Player.Asset, g.Player.Position) {
			fmt.Println("Player hit by bomb")
			g.Lives.NoOfLives--
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
			g.Lives.NoOfLives--
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
	return gameWidth, gameHeight
}
