package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/perbu/extraterrestrial_trespassers/assets"
	"golang.org/x/image/font"
	"image/color"
	"strconv"
)

type score struct {
	score    int
	position position
	face     font.Face
}

func (g *Game) newScore() *score {
	face, err := assets.GetFont(16)
	if err != nil {
		panic(err)
	}
	g.state.GetWidth()
	return &score{
		position: position{
			g.state.GetWidth() - 100,
			20,
		},
		face: face,
	}
}

func (s *score) AddScore(amount int) {
	s.score += amount
}

func (s *score) Draw(screen *ebiten.Image) {
	text.Draw(screen, strconv.Itoa(s.score), s.face, s.position.x, s.position.y, color.White)
}
