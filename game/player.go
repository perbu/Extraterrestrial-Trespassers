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
	gun         *gun
}

func newPlayer(aud *audio.Context, state *state.Global, game *Game) player {
	shootPlayer, _ := aud.NewPlayer(assets.GetShootSound())
	crashPlayer, _ := aud.NewPlayer(assets.GetThud())
	p := player{
		position: position{
			x: state.GetWidth() / 2,
			y: state.GetHeight() - 50,
		},
		asset:       assets.GetPlayer(),
		shootPlayer: shootPlayer,
		crashPlayer: crashPlayer,
		global:      state,
		game:        game,
	}
	p.gun = p.newGun() // attach a gun.
	return p
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
	p.gun.Draw(screen)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(p.position.x), float64(p.position.y))

	if p.global.IsFrozen() {
		if time.Now().UnixMilli()/100%2 == 0 {
			return
		}
	}
	screen.DrawImage(p.asset.Sprite, op)

}

func (p *player) Shoot() {
	vol := p.game.state.GetVolume()
	p.shootPlayer.SetVolume(vol)
	_ = p.shootPlayer.Rewind()
	p.shootPlayer.Play()
	gunpos := p.getGunPlacement()
	proj := &projectile{
		asset:    assets.GetProjectile(),
		position: gunpos,
		speed:    p.gun.fire(),
	}
	p.game.projectiles = append(p.game.projectiles, proj)
}

func (p *player) getGunPlacement() position {
	return position{
		x: p.position.x + p.asset.Bounds.Max.X/2,
		y: p.position.y,
	}
}

func (p *player) Collision() {
	vol := p.game.state.GetVolume()
	p.crashPlayer.SetVolume(vol)
	_ = p.crashPlayer.Rewind()
	p.crashPlayer.Play()
	p.game.lives.DecrementLives()
	p.global.FreezeUntil(time.Now().Add(2 * time.Second))
	p.dead = true

}

func (p *player) Respawn() {
	centerX := p.global.GetWidth() / 2
	p.position.x = centerX
	rect := p.asset.Bounds
	// leave one whole player on each side of the center
	xMin := centerX - rect.Max.X
	xMax := centerX + rect.Max.X
	yMin := 300
	// now clear the bombs that are 200 pixels above the player
	for i, bomb := range p.game.bombs {
		if bomb.position.x > xMin && bomb.position.x < xMax && bomb.position.y > yMin {
			p.game.bombs = append(p.game.bombs[:i], p.game.bombs[i+1:]...)
		}
	}
}
