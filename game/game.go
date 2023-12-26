package game

import (
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
		Lives:      NewLife(0, 0, 3),
		AlienFleet: newFleet(0, 50, gameMargin, gameWidth-gameMargin),
		Bombs:      make([]*Bomb, 0, 10),
		Player: Player{
			Position: Position{
				X: gameWidth / 2,
				Y: gameHeight - 50,
			},
			Sprite: assets.GetPlayer(),
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

	// Check for collisions
	for _, e := range g.AlienFleet.Enemies {
		for _, p := range g.Projectiles {
			if e.Collides(p) {
				// remove the projectile from the screen
				p.Position.Y = -10
				// remove the enemy from the fleet:
				e.dead = true
			}
		}
	}
	// remove dead enemies from the fleet
	g.AlienFleet.Enemies = filterEnemies(g.AlienFleet.Enemies)

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
