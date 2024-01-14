package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"math"
	"math/rand"
	"time"
)

type particle struct {
	position  position
	speed     float64
	direction float64
	size      int
	color     color.Color
	created   time.Time
}

func newParticle(pos position, col color.Color) *particle {

	return &particle{
		position:  pos,
		speed:     float64(rand.Intn(10) + 3),
		direction: generateDirection(),
		size:      2,
		color:     col,
		created:   time.Now(),
	}
}

func (p *particle) Update() bool {
	// update the position of the particle, use the direction and speed
	// to calculate the new position
	p.position.X += int(p.speed * math.Cos(p.direction))
	p.position.Y += int(p.speed * math.Sin(p.direction))
	return true
}

const particleLifeTime = 1500

func (p *particle) Draw(screen *ebiten.Image) {
	age := p.age()
	if age > particleLifeTime {
		return
	}
	normalizedAgeFactor := float64(age) / float64(particleLifeTime) // Normalize to [0, 1]
	// newAlpha := uint8(float64(alpha) * (1 - normalizedAgeFactor))   // Fade alpha based on age
	newColor := multiplyAlpha(p.color, 1-normalizedAgeFactor)
	// fmt.Printf("age: %d, normalizedAgeFactor: %f, newAlpha: %d\n", age, normalizedAgeFactor, newAlpha)
	vector.DrawFilledRect(screen,
		float32(p.position.X),
		float32(p.position.Y),
		float32(p.size),
		float32(p.size), newColor, false)
}
func (p *particle) age() int {
	return int(time.Since(p.created).Milliseconds())
}

// generateDirection generates a random direction
func generateDirection() float64 {
	// Generate a random float between 0 and 2 * pi
	randomValue := rand.Float64() * 2 * math.Pi
	return randomValue
}

// multiplyAlpha takes a color.Color and a factor (between 0 and 1), and returns a new color with the alpha channel multiplied by the factor.
func multiplyAlpha(c color.Color, factor float64) color.Color {
	if factor < 0 || factor > 1 {
		panic("factor must be between 0 and 1")
	}

	r, g, b, a := c.RGBA()
	// Since RGBA returns color components in the range [0, 65535], we need to scale them down.
	alpha := float64(a) / 65535
	alpha *= factor

	// Convert the modified alpha back to the [0, 65535] range and return a new color.
	return color.NRGBA{
		R: uint8(r >> 8),
		G: uint8(g >> 8),
		B: uint8(b >> 8),
		A: uint8(alpha * 255),
	}
}
