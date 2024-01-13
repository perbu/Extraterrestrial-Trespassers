package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/perbu/extraterrestrial_trespassers/assets"
	"sync"
)

type Lives struct {
	position Position
	asset    assets.Asset
	game     *Game
	lives    int
	mu       sync.Mutex
}

func NewLife(x, y int, game *Game) *Lives {
	return &Lives{
		position: Position{
			X: x,
			Y: y,
		},
		asset: assets.GetPlayer(),
		game:  game,
		lives: 2,
	}
}

func (l *Lives) Draw(screen *ebiten.Image) {
	lives := l.lives
	for i := 0; i < lives; i++ {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(l.position.X+i*100), float64(l.position.Y))
		op.GeoM.Scale(.5, 0.5)
		screen.DrawImage(l.asset.Sprite, op)
	}
}

func (l *Lives) GetLives() int {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.lives
}

func (l *Lives) SetLives(lives int) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.lives = lives
}

func (l *Lives) DecrementLives() int {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.lives--
	return l.lives
}
