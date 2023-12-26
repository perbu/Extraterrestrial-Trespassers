package main

import (
	"github.com/perbu/spaceinvaders/game"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 640
	screenHeight = 480
	gameWidth    = 640
	gameHeight   = 480
)

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Space Invaders")
	g := game.New(screenWidth, screenHeight, gameWidth, gameHeight)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
