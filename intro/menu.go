package intro

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/perbu/extraterrestrial_trespassers/assets"
	"github.com/perbu/extraterrestrial_trespassers/state"
	"golang.org/x/image/font"
	"image/color"
)

type intro struct {
	menuOptions      []string
	currentSelection int
	face             font.Face
	state            *state.Global
	mode             mode
}

type mode int

const (
	menu mode = iota
	credits
)

func newMenu(state *state.Global) *intro {
	face, err := assets.GetFont(24)
	if err != nil {
		panic(err)
	}
	return &intro{
		menuOptions: []string{"Start Game", "Credits", "Quit"},
		face:        face,
		state:       state,
	}
}

func (g *intro) Update() error {
	if g.mode == credits {
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			g.mode = menu
		}
		return nil
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		switch g.currentSelection {
		case 0:
			g.state.QueueAction(state.NewGame)
		case 1:
			g.mode = credits
			return nil
		case 2:
			g.state.QueueAction(state.Quit)
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		if g.currentSelection > 0 {
			g.currentSelection--
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		if g.currentSelection < len(g.menuOptions)-1 {
			g.currentSelection++
		}
	}
	return nil
}

func (g *intro) Draw(screen *ebiten.Image) {
	if g.mode == credits {
		text.Draw(screen, "Credits", g.face, 100, 100, color.White)
		text.Draw(screen, "perbu", g.face, 100, 200, color.White)
		return
	}
	for i, option := range g.menuOptions {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(g.state.GetWidth()/2-len(option)*4), float64(100+i*20))
		if i == g.currentSelection {
			text.Draw(screen, option, g.face,
				g.state.GetWidth()/2-len(option)*8,
				300+i*40, color.RGBA{0, 255, 0, 255}) // Highlight selected option
		} else {
			text.Draw(screen, option, g.face,
				g.state.GetWidth()/2-len(option)*8,
				300+i*40, color.White)
		}
	}
}

func (g *intro) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.state.GetDimensions()
}
