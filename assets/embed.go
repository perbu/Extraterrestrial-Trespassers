package assets

import (
	"bytes"
	_ "embed"
	"github.com/hajimehoshi/ebiten/v2"
	"image/png"
)

//go:embed player.png
var player []byte

//go:embed green.png
var green []byte

//go:embed red.png
var red []byte

//go:embed yellow.png
var yellow []byte

//go:embed blue.png
var blue []byte

//go:embed projectile.png
var projectile []byte

func GetPlayer() *ebiten.Image {
	// load a png file:
	img, err := png.Decode(bytes.NewReader(player))
	if err != nil {
		panic(err)
	}
	eimg := ebiten.NewImageFromImage(img)
	return eimg
}

func GetGreen() *ebiten.Image {
	img, err := png.Decode(bytes.NewReader(green))
	if err != nil {
		panic(err)
	}
	eimg := ebiten.NewImageFromImage(img)
	return eimg
}

func GetRed() *ebiten.Image {
	img, err := png.Decode(bytes.NewReader(red))
	if err != nil {
		panic(err)
	}
	eimg := ebiten.NewImageFromImage(img)
	return eimg
}

func GetYellow() *ebiten.Image {
	img, err := png.Decode(bytes.NewReader(yellow))
	if err != nil {
		panic(err)
	}
	eimg := ebiten.NewImageFromImage(img)
	return eimg
}

func GetBlue() *ebiten.Image {
	img, err := png.Decode(bytes.NewReader(blue))
	if err != nil {
		panic(err)
	}
	eimg := ebiten.NewImageFromImage(img)
	return eimg
}

func GetProjectile() *ebiten.Image {
	img, err := png.Decode(bytes.NewReader(projectile))
	if err != nil {
		panic(err)
	}
	eimg := ebiten.NewImageFromImage(img)
	return eimg
}
