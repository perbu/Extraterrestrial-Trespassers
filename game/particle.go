package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"math"
	"math/rand"
)

type particle struct {
	position  position
	speed     int
	direction int
	size      int
	color     color.Color
}

func newParticle(pos position) *particle {
	return &particle{
		position:  pos,
		speed:     rand.Intn(10) + 3,
		direction: rand.Intn(360),
		size:      rand.Intn(3) + 1,
		color:     randomParticleColor(),
	}
}

func (p *particle) Update() bool {
	// update the position of the particle, use the direction and speed
	// to calculate the new position
	p.position.X += int(float64(p.speed) * math.Cos(float64(p.direction)))
	p.position.Y += int(float64(p.speed) * math.Sin(float64(p.direction)))
	return true
}

func (p *particle) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen,
		float32(p.position.X),
		float32(p.position.Y),
		float32(p.size),
		float32(p.size), p.color, false)
}

func randomParticleColor() color.Color {
	return color.RGBA{
		R: 255,
		G: 255,
		B: uint8(rand.Intn(128) + 128),
		A: uint8(rand.Intn(255)),
	}
}
