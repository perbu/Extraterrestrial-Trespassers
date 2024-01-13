package intro

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/perbu/extraterrestrial_trespassers/state"
	"image/color"
	"math/rand"
)

type StarField struct {
	Stars []*star
	Menu  *menu
	state *state.Global
}

func NewStarField(state *state.Global) *StarField {
	sf := &StarField{
		Stars: make([]*star, 0),
		Menu:  newMenu(state),
		state: state,
	}

	for i := 0; i < 800; i++ {
		speed := rand.Intn(5) + 1
		var col color.Color
		var size float32
		switch speed {
		case 5:
			col = color.RGBA{R: 255, G: 255, B: 255, A: 255}
			size = 3
		case 4:
			col = color.RGBA{R: 150, G: 150, B: 150, A: 255}
			size = 2
		case 3:
			col = color.RGBA{R: 100, G: 100, B: 100, A: 255}
			size = 2
		case 2:
			col = color.RGBA{R: 75, G: 75, B: 75, A: 255}
			size = 1
		case 1:
			col = color.RGBA{R: 50, G: 50, B: 50, A: 255}
			size = 1
		}
		sf.Stars = append(sf.Stars, &star{
			position: position{
				X: rand.Intn(sf.state.GetWidth()),
				Y: rand.Intn(sf.state.GetHeight()),
			},
			speed: speed,
			color: col,
			size:  size,
		})
	}
	return sf
}

func (sf *StarField) Update() error {
	err := sf.Menu.Update()
	if err != nil {
		return err
	}
	for _, s := range sf.Stars {
		_ = s.Update()
		if s.position.Y > sf.state.GetHeight() {
			// Reset star
			s.position.Y = -10
		}
	}
	return nil
}

func (sf *StarField) Draw(screen *ebiten.Image) {
	for _, s := range sf.Stars {
		s.Draw(screen)
	}
	sf.Menu.Draw(screen)
}
func (sf *StarField) Layout(outsideWidth, outsideHeight int) (int, int) {
	return sf.state.GetDimensions()
}
