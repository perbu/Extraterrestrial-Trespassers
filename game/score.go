package game

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/perbu/extraterrestrial_trespassers/assets"
	"golang.org/x/image/font"
	"image/color"
)

type score struct {
	score    int
	position position
	face     font.Face
	game     *Game
}

func (g *Game) newScore() *score {
	face, err := assets.GetFont(16)
	if err != nil {
		panic(err)
	}
	g.state.GetWidth()
	return &score{
		position: position{
			g.state.GetWidth() - 300,
			20,
		},
		face: face,
		game: g,
	}
}

func (s *score) AddScore(amount int) {
	s.score += amount
}

func (s *score) Draw(screen *ebiten.Image) {
	mess := fmt.Sprintf("Level %d Score: %d", s.game.alienFleet.level, s.score)
	text.Draw(screen, mess, s.face, s.position.x, s.position.y, color.White)
}
