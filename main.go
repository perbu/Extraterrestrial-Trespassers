package main

import (
	"errors"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/perbu/extraterrestrial_trespassers/assets"
	"github.com/perbu/extraterrestrial_trespassers/game"
	"github.com/perbu/extraterrestrial_trespassers/intro"
	"github.com/perbu/extraterrestrial_trespassers/state"
	"log"
)

type App struct {
	game  *game.Game
	intro *intro.StarField
	state *state.Global
	song  *audio.Player
}

func main() {
	s := state.Initial()
	ebiten.SetWindowSize(s.GetDimensions())
	ebiten.SetWindowTitle("Extraterrestrial Trespassers")
	// set fullscreen:
	ebiten.SetFullscreen(true)

	acontext := audio.NewContext(44100)
	song, _ := acontext.NewPlayer(assets.GetSong())
	app := &App{
		game:  game.NewGame(acontext, s),
		intro: intro.NewStarField(s),
		state: s,
		song:  song,
	}
	song.Play()
	err := ebiten.RunGame(app)
	if err != nil {
		log.Fatal(err)
	}
}

func (a *App) Update() error {

	if a.song.IsPlaying() == false {
		a.song.Rewind()
		a.song.Play()
	}
	// is global state update returns true, we should abort the update,
	// since we are in a transition state.
	if a.state.Update() {
		return nil
	}

	switch a.state.GetScene() {
	case state.SceneMenu:
		a.song.SetVolume(1.0)
		action := a.state.ShiftAction()
		switch action {
		case state.Nothing:
			break
		case state.NewGame:
			a.state.SetScene(state.SceneGame)
		case state.ShowCredits:
			a.state.SetScene(state.SceneCredits)
		case state.Quit:
			return errors.New("quit")
		case state.GameOver:
			a.state.SetScene(state.SceneMenu)
		default:
			fmt.Println("unhandled action:", action.String())
		}
		return a.intro.Update()
	case state.SceneGame:
		a.song.SetVolume(0.5)
		action := a.state.ShiftAction()
		switch action {
		case state.Nothing:
			break
		case state.GameOver:
			a.state.SetScene(state.SceneMenu)
		case state.Quit:
			return errors.New("quit")
		default:
			panic("unhandled default case")
		}
		return a.game.Update()
	case state.SceneCredits:
		panic("not implemented")
	}
	return nil
}

func (a *App) Draw(screen *ebiten.Image) {
	switch a.state.GetScene() {
	case state.SceneMenu:
		a.intro.Draw(screen)
	case state.SceneGame:
		a.game.Draw(screen)
	case state.SceneCredits:
		panic("not implemented")
	}
}

func (a *App) Layout(outsideWidth, outsideHeight int) (int, int) {
	return a.state.GetDimensions()
}
