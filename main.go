package main

import (
	_ "embed"
	"errors"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/joho/godotenv"
	"github.com/perbu/extraterrestrial_trespassers/assets"
	"github.com/perbu/extraterrestrial_trespassers/game"
	"github.com/perbu/extraterrestrial_trespassers/intro"
	"github.com/perbu/extraterrestrial_trespassers/shader"
	"github.com/perbu/extraterrestrial_trespassers/state"
	"log"
	"os"
)

type App struct {
	game         *game.Game
	intro        *intro.StarField
	state        *state.Global
	song         *audio.Player
	music        bool
	audioContext *audio.Context
	shaderBuffer *ebiten.Image
	shader       shader.Shader
}

func main() {
	_ = godotenv.Load()

	s := state.Initial()
	ebiten.SetWindowSize(s.GetDimensions())
	ebiten.SetWindowTitle("Extraterrestrial Trespassers")
	// set fullscreen:
	ebiten.SetFullscreen(true)

	lotte, err := shader.NewCrtLotte()
	if err != nil {
		panic(err)
	}

	acontext := audio.NewContext(44100)
	song, _ := acontext.NewPlayer(assets.GetSong())
	app := &App{
		audioContext: acontext,
		game:         game.NewGame(acontext, s),
		intro:        intro.NewStarField(s),
		state:        s,
		song:         song,
		music:        musicEnabled(),
		shaderBuffer: ebiten.NewImage(s.GetDimensions()),
		shader:       lotte,
	}
	err = ebiten.RunGame(app)
	if err != nil {
		log.Fatal(err)
	}
}

func (a *App) Update() error {
	if a.music && !a.song.IsPlaying() {
		_ = a.song.Rewind()
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
			a.game = game.NewGame(a.audioContext, a.state)
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
	a.shaderBuffer.Clear()
	switch a.state.GetScene() {
	case state.SceneMenu:
		a.intro.Draw(a.shaderBuffer)
	case state.SceneGame:
		a.game.Draw(a.shaderBuffer)
	case state.SceneCredits:
		panic("not implemented")
	}
	err := a.shader.Apply(screen, a.shaderBuffer)
	if err != nil {
		panic(err)
	}
	// screen.DrawRectShader(w, h, a.shader, options)

}

func (a *App) Layout(outsideWidth, outsideHeight int) (int, int) {
	return a.state.GetDimensions()
}

func musicEnabled() bool {
	_, ok := os.LookupEnv("MUSIC_DISABLED")
	fmt.Println("music enabled:", !ok)
	return !ok
}
