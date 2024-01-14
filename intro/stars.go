package intro

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
)

type position struct {
	X int
	Y int
}

type star struct {
	position position
	speed    int
	color    color.Color
	size     float32
}

func (s *star) Update() error {
	s.position.Y += s.speed
	return nil
}

func (s *star) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, float32(s.position.X), float32(s.position.Y), s.size, s.size, s.color, false)
	// ebitenutil.DrawRect(screen, float64(s.position.x), float64(s.position.y), 3, 3, s.color)
	// screen.Set(s.position.x, s.position.y, s.color)
}
