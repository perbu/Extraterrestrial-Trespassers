package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/perbu/extraterrestrial_trespassers/assets"
	"github.com/perbu/extraterrestrial_trespassers/state"
	"time"
)

type player struct {
	position    position
	asset       assets.Asset
	shootPlayer *audio.Player
	crashPlayer *audio.Player
	global      *state.Global
	game        *Game
	dead        bool
}

func NewPlayer(aud *audio.Context, state *state.Global, game *Game) player {
	shootPlayer, _ := aud.NewPlayer(assets.GetShootSound())
	crashPlayer, _ := aud.NewPlayer(assets.GetThud())
	return player{
		position: position{
			x: state.GetHeight() / 2,
			y: state.GetHeight() - 50,
		},
		asset:       assets.GetPlayer(),
		shootPlayer: shootPlayer,
		crashPlayer: crashPlayer,
		global:      state,
		game:        game,
	}
}

func (p *player) Update() {
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		if p.position.x > 0 {
			p.position.x -= 5
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		if p.position.x < p.global.GetWidth()-p.global.GetMargins() {
			p.position.x += 5
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		p.Shoot()
	}

}
func (p *player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(p.position.x), float64(p.position.y))
	// check if we are frozen, if so, we should blink the player
	if p.global.IsFrozen() {
		if time.Now().UnixMilli()/100%2 == 0 {
			return
		}
	}
	screen.DrawImage(p.asset.Sprite, op)
}

func (p *player) Shoot() {
	_ = p.shootPlayer.Rewind()
	p.shootPlayer.Play()
	proj := &projectile{
		asset: assets.GetProjectile(),
		position: position{
			x: p.position.x + 20,
			y: p.position.y,
		},
		speed: projectileSpeed,
	}
	p.game.projectiles = append(p.game.projectiles, proj)
}

func (p *player) Collision() {
	_ = p.crashPlayer.Rewind()
	p.crashPlayer.Play()
	p.game.lives.DecrementLives()
	p.global.FreezeUntil(time.Now().Add(2 * time.Second))
	p.dead = true

}

func (p *player) Respawn() {
	p.position.x = p.global.GetWidth() / 2

}
