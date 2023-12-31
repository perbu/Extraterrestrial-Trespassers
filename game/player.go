package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/perbu/extraterrestrial_trespassers/assets"
	"github.com/perbu/extraterrestrial_trespassers/state"
)

type Player struct {
	Position    Position
	Asset       assets.Asset
	ShootPlayer *audio.Player
	CrashPlayer *audio.Player
	global      *state.Global
	dying       bool
}

func NewPlayer(aud *audio.Context, state *state.Global) Player {
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
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(p.Position.X), float64(p.Position.Y))
	screen.DrawImage(p.Asset.Sprite, op)
}

func (p *Player) Shoot() *Projectile {
	p.ShootPlayer.Rewind()
	p.ShootPlayer.Play()
	return &Projectile{
		Asset: assets.GetProjectile(),
		Position: Position{
			X: p.Position.X + 20,
			Y: p.Position.Y,
		},
		Speed: 5,
	}
}

func (p *Player) Crash() {
	_ = p.CrashPlayer.Rewind()
	p.CrashPlayer.Play()
	p.global.QueueAction(state.PlayerDied)
}
