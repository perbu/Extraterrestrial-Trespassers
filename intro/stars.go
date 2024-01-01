package intro

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/perbu/extraterrestrial_trespassers/game"
	"image/color"
)

type Star struct {
	Position game.Position
	Speed    int
	Color    color.Color
	Size     float32
}

func (s *Star) Update() error {
	s.Position.Y += s.Speed
	return nil
}

func (s *Star) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, float32(s.Position.X), float32(s.Position.Y), s.Size, s.Size, s.Color, false)
	// ebitenutil.DrawRect(screen, float64(s.Position.X), float64(s.Position.Y), 3, 3, s.Color)
	// screen.Set(s.Position.X, s.Position.Y, s.Color)
}
