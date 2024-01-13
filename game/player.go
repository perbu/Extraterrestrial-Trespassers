package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/perbu/extraterrestrial_trespassers/assets"
	"github.com/perbu/extraterrestrial_trespassers/state"
	"time"
)

type Player struct {
	Position    Position
	Asset       assets.Asset
	ShootPlayer *audio.Player
	CrashPlayer *audio.Player
	global      *state.Global
	game        *Game
	dead        bool
}

func NewPlayer(aud *audio.Context, state *state.Global, game *Game) Player {
	shootPlayer, _ := aud.NewPlayer(assets.GetShootSound())
	crashPlayer, _ := aud.NewPlayer(assets.GetThud())
	return Player{
		Position: Position{
			X: state.GetHeight() / 2,
			Y: state.GetHeight() - 50,
		},
		Asset:       assets.GetPlayer(),
		ShootPlayer: shootPlayer,
		CrashPlayer: crashPlayer,
		global:      state,
		game:        game,
	}
}

func (p *Player) Update() {
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		if p.Position.X > 0 {
			p.Position.X -= 5
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		if p.Position.X < p.global.GetWidth()-p.global.GetMargins() {
			p.Position.X += 5
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		p.Shoot()
	}

}
func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(p.Position.X), float64(p.Position.Y))
	// check if we are frozen, if so, we should blink the player
	if p.global.IsFrozen() {
		if time.Now().UnixMilli()/100%2 == 0 {
			return
		}
	}
	screen.DrawImage(p.Asset.Sprite, op)
}

func (p *Player) Shoot() {
	p.ShootPlayer.Rewind()
	p.ShootPlayer.Play()
	proj := &Projectile{
		Asset: assets.GetProjectile(),
		Position: Position{
			X: p.Position.X + 20,
			Y: p.Position.Y,
		},
		Speed: 5,
	}
	p.game.Projectiles = append(p.game.Projectiles, proj)
}

func (p *Player) Collision() {
	_ = p.CrashPlayer.Rewind()
	p.CrashPlayer.Play()
	p.game.Lives.DecrementLives()
	p.global.FreezeUntil(time.Now().Add(2 * time.Second))
	p.dead = true

}

func (p *Player) Respawn() {
	p.Position.X = p.global.GetWidth() / 2

}
