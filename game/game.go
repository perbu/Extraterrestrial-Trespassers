package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/perbu/spaceinvaders/assets"
)

type Game struct {
	Player       Player
	AlienFleet   *Fleet
	Projectiles  []*Projectile
	screenHeight int
	screenWidth  int
	gameHeight   int
	gameWidth    int
}

type Position struct {
	X int
	Y int
}

func New(swidth, sheight, gwidth, gheight int) *Game {
	return &Game{
		screenHeight: sheight,
		screenWidth:  swidth,
		gameHeight:   gheight,
		gameWidth:    gwidth,
		AlienFleet:   newFleet(0, 5, 50, gwidth-50),
		Player: Player{
			Position: Position{
				X: gwidth / 2,
				Y: gheight - 50,
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
		if g.Player.Position.X < g.gameWidth-50 {
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
	g.Player.Draw(screen)
	for _, p := range g.Projectiles {
		p.Draw(screen)
	}
	g.AlienFleet.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.screenWidth, g.screenHeight
}
