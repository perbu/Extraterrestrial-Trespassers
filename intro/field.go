package intro

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/perbu/extraterrestrial_trespassers/game"
	"image/color"
	"math/rand"
)

type StarField struct {
	Stars []*Star
	Menu  *Menu
}

func NewStarField() *StarField {
	sf := &StarField{
		Stars: make([]*Star, 0),
		Menu:  newMenu(),
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
		sf.Stars = append(sf.Stars, &Star{
			Position: game.Position{
				X: rand.Intn(game.GameWidth),
				Y: rand.Intn(game.GameHeight),
			},
			Speed: speed,
			Color: col,
			Size:  size,
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
		if s.Position.Y > game.GameHeight {
			// Reset star
			s.Position.Y = -10
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
	return game.GameWidth, game.GameHeight
}
