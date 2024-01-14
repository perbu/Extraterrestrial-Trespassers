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
	x, y      float32
	speed     float32
	direction float64
	size      float32
	color     color.Color
	created   time.Time
}

func newParticle(pos position, col color.Color) *particle {

	return &particle{
		x:         float32(pos.x),
		y:         float32(pos.y),
		speed:     projectileSpeed + rand.Float32()*5,
		direction: generateDirection(),
		size:      2,
		color:     col,
		created:   time.Now(),
	}
}

func (p *particle) Update() {
	// update the position of the particle, use the direction (radians) and speed
	// to calculate the new position
	p.x += p.speed * float32(math.Sin(p.direction))
	p.y -= p.speed * float32(math.Cos(p.direction))
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
		p.x,
		p.y,
		p.size,
		p.size, newColor, false)
}
func (p *particle) age() int {
	return int(time.Since(p.created).Milliseconds())
}

func generateDirection() float64 {
	ran := skewedRandom() * (math.Pi / 2) // a random value between 0 and pi/2
	if rand.Intn(2) == 0 {
		return ran
	} else {
		return -ran
	}
}

// skewedRandom returns a random number between 0 and 1, skewed towards 0.
func skewedRandom() float64 {
	lambda := 8.0 // Adjust lambda to control the skewness (higher values skew more towards 0)
	return rand.ExpFloat64() / lambda
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
