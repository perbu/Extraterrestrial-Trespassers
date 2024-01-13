package game

import (
	"github.com/perbu/extraterrestrial_trespassers/assets"
	"github.com/perbu/extraterrestrial_trespassers/state"
	"golang.org/x/image/font"
)

type GameOver struct {
	Position position
	Text     string
	Face     font.Face
	Done     bool
	state    *state.Global
}

func (g *GameOver) Update() error {
	if g.Position.Y < g.state.GetHeight()/2 {
		g.Position.Y += 1
	} else {
		g.Done = true
	}
	return nil
}

func NewGameOver(state *state.Global) *GameOver {
	face, err := assets.GetFont(24)
	if err != nil {
		panic(err)
	}
	return &GameOver{
		Position: position{X: state.GetWidth()/2 - 100, Y: 0},
		Text:     "Game Over",
		Face:     face,
		state:    state,
	}
}
