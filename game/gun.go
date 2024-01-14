package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"time"
)

const (
	maxCoolDown = time.Second * 3
	maxGunSpeed = 20
)

type gun struct {
	lastFire time.Time
	position position
}

func (g *player) newGun() *gun {
	// find the position of the gun speed indicator, it should be in the lower right corner.
	// the speed indicator should be 10 pixels from the right and 30 pixels from the bottom
	width, height := g.game.state.GetDimensions()
	return &gun{
		position: position{
			x: width - 10,
			y: height - 30,
		},
		lastFire: time.Now(),
	}
}

func (g *gun) getSpeed() int {
	// calculate the speed based on the time since last fire
	// speed will be from 2 to maxGunSpeed, maxGunSpeed if last fire was maxCoolDown ago

	// get the time since last fire
	timeSinceLastFire := time.Since(g.lastFire)

	// calculate the speed
	// 1 / 3 = 0.333
	// speed: time / max * maxSpeed
	speed := int(timeSinceLastFire.Seconds()/maxCoolDown.Seconds()*maxGunSpeed) + 2

	// if the speed is greater than max, set it to max
	if speed > maxGunSpeed {
		speed = maxGunSpeed
	}

	return speed
}

func (g *gun) fire() int {
	speed := g.getSpeed()
	g.lastFire = time.Now()
	return speed
}

func (g *gun) Draw(screen *ebiten.Image) {
	// draw the speed indicator
	speed := g.getSpeed()
	for i := 0; i < speed; i++ {
		// if i == maxGunSpeed then j should be 255
		j := uint8(255 / maxGunSpeed * i)
		col := color.RGBA{R: 255, G: j, B: j, A: 255}
		x := g.position.x
		y := g.position.y - i*2
		vector.DrawFilledRect(screen, float32(x), float32(y), 10, 2, col, false)
	}
}
