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

type menu struct {
	menuOptions      []string
	currentSelection int
	face             font.Face
	state            *state.Global
}

func newMenu(state *state.Global) *menu {
	face, err := assets.GetFont(24)
	if err != nil {
		panic(err)
	}
	return &menu{
		menuOptions: []string{"Start Game", "Credits", "Quit"},
		face:        face,
		state:       state,
	}
}

func (g *menu) Update() error {
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
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		switch g.currentSelection {
		case 0:
			g.state.QueueAction(state.NewGame)
		case 1:
			g.state.QueueAction(state.ShowCredits)
		case 2:
			g.state.QueueAction(state.Quit)
		}
	}
	// Add logic here for when an option is selected (e.g., pressing Enter)

	return nil
}

func (g *menu) Draw(screen *ebiten.Image) {
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

func (g *menu) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.state.GetDimensions()
}
