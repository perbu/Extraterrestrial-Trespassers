package assets

import (
	"bytes"
	_ "embed"
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
	"image"
	"image/png"
)

type Asset struct {
	Sprite *ebiten.Image
	Bounds image.Rectangle
}

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

//go:embed bomb.png
var bomb []byte

//go:embed PressStart2P-Regular.ttf
var PressStart2P []byte

func (a Asset) GetRect() image.Rectangle {
	return a.Bounds
}

func GetPlayer() Asset {
	// load a png file:
	img, err := png.Decode(bytes.NewReader(player))
	if err != nil {
		panic(err)
	}
	return Asset{
		Sprite: ebiten.NewImageFromImage(img),
		Bounds: img.Bounds(),
	}
}

func GetGreen() Asset {
	img, err := png.Decode(bytes.NewReader(green))
	if err != nil {
		panic(err)
	}
	return Asset{
		Sprite: ebiten.NewImageFromImage(img),
		Bounds: img.Bounds(),
	}
}

func GetRed() Asset {
	img, err := png.Decode(bytes.NewReader(red))
	if err != nil {
		panic(err)
	}
	return Asset{
		Sprite: ebiten.NewImageFromImage(img),
		Bounds: img.Bounds(),
	}
}

func GetYellow() Asset {
	img, err := png.Decode(bytes.NewReader(yellow))
	if err != nil {
		panic(err)
	}
	return Asset{
		Sprite: ebiten.NewImageFromImage(img),
		Bounds: img.Bounds(),
	}
}

func GetBlue() Asset {
	img, err := png.Decode(bytes.NewReader(blue))
	if err != nil {
		panic(err)
	}
	return Asset{
		Sprite: ebiten.NewImageFromImage(img),
		Bounds: img.Bounds(),
	}
}

func GetProjectile() Asset {
	img, err := png.Decode(bytes.NewReader(projectile))
	if err != nil {
		panic(err)
	}
	return Asset{
		Sprite: ebiten.NewImageFromImage(img),
		Bounds: img.Bounds(),
	}
}

func GetBomb() Asset {
	img, err := png.Decode(bytes.NewReader(bomb))
	if err != nil {
		panic(err)
	}
	return Asset{
		Sprite: ebiten.NewImageFromImage(img),
		Bounds: img.Bounds(),
	}
}

func GetFont(size float64) (font.Face, error) {
	f, err := truetype.Parse(PressStart2P)
	if err != nil {
		return nil, err
	}
	face := truetype.NewFace(f, &truetype.Options{
		Size: size,
	})
	return face, nil
}
