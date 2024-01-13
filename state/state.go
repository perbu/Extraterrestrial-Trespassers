package state

import (
	"sync"
	"time"
)

type Global struct {
	mu          sync.Mutex
	width       int
	height      int
	margins     int
	scene       Scene
	actions     []Action // queue of actions to be processed
	freezeUntil time.Time
	lives       int
	dead        bool // player is dead
}

type Scene int

const (
	SceneMenu Scene = iota
	SceneGame
	SceneCredits
)

type Action int

// These are the global actions that can be queued. They're all affecting the global state of the game.
// Other events are handled locally.
const (
	Nothing Action = iota
	Quit
	NewGame
	ShowCredits
	GameOver
)

// make strings:
//go:generate stringer -type=Action

func Initial() *Global {
	return &Global{
		width:   1200,
		height:  900,
		margins: 50,
		scene:   SceneMenu,
	}
}

func (g *Global) Update() bool {
	// don't update the game if we are frozen
	return g.IsFrozen()
}

func (g *Global) QueueAction(action Action) {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.actions = append(g.actions, action)
}

func (g *Global) ShiftAction() Action {
	g.mu.Lock()
	defer g.mu.Unlock()
	if len(g.actions) == 0 {
		return Nothing
	}
	action := g.actions[0]
	g.actions = g.actions[1:]
	return action
}

func (g *Global) SetScene(scene Scene) {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.scene = scene
}

func (g *Global) GetScene() Scene {
	g.mu.Lock()
	defer g.mu.Unlock()
	return g.scene
}

func (g *Global) GetDimensions() (int, int) {
	g.mu.Lock()
	defer g.mu.Unlock()
	return g.width, g.height
}

func (g *Global) GetWidth() int {
	g.mu.Lock()
	defer g.mu.Unlock()
	return g.width
}

func (g *Global) GetHeight() int {
	g.mu.Lock()
	defer g.mu.Unlock()
	return g.height
}

func (g *Global) GetMargins() int {
	g.mu.Lock()
	defer g.mu.Unlock()
	return g.margins
}

func (g *Global) FreezeUntil(t time.Time) {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.freezeUntil = t
}

// isFrozen returns true if the freezeUntil time is in the future
// doesn't lock the mutex
func (g *Global) isFrozen() bool {
	return g.freezeUntil.After(time.Now())
}

func (g *Global) IsFrozen() bool {
	g.mu.Lock()
	defer g.mu.Unlock()
	return g.freezeUntil.After(time.Now())
}
