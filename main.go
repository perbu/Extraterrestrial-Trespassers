package main

import (
	"errors"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/perbu/extraterrestrial_trespassers/game"
	"github.com/perbu/extraterrestrial_trespassers/intro"
	"log"
)

type Scene int

const (
	SceneMenu Scene = iota
	SceneGame
)

type App struct {
	game  *game.Game
	intro *intro.StarField
	scene Scene
}

func main() {
	ebiten.SetWindowSize(game.GameWidth, game.GameHeight)
	ebiten.SetWindowTitle("Extraterrestrial Trespassers")
	// set fullscreen:
	ebiten.SetFullscreen(true)

	app := &App{
		game:  game.NewGame(),
		intro: intro.NewStarField(),
		scene: SceneMenu,
	}
	err := ebiten.RunGame(app)
	if err != nil {
		log.Fatal(err)
	}
}

func (a *App) Update() error {
	switch a.scene {
	case SceneMenu:
		err := a.intro.Update()
		if err != nil {
			return err
		}
		switch a.intro.Menu.Selection {
		case intro.Nothing:
			break
		case intro.StartGame:
			a.scene = SceneGame

		case intro.Credits:
			break // not implemented
		case intro.Quit:
			return errors.New("Go away")
		}
		a.intro.Menu.Selection = intro.Nothing
	case SceneGame:
		err := a.game.Update()
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *App) Draw(screen *ebiten.Image) {
	switch a.scene {
	case SceneMenu:
		a.intro.Draw(screen)
	case SceneGame:
		a.game.Draw(screen)
	}
}

func (a *App) Layout(outsideWidth, outsideHeight int) (int, int) {
	return game.GameWidth, game.GameHeight
}
