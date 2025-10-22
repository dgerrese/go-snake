package game

import (
	"log"
	"time"

	"go-snake/go_snake/input"
	"go-snake/go_snake/screens"
	"go-snake/go_snake/types/gamestate"
)

func (g *Game) Start() {
	g.setGameStateStarting()
}

func (g *Game) setGameStateStarting() {
	if fn, ok := g.removeFns["menu"]; ok {
		fn()
	}

	g.state = gamestate.Starting
	log.Default().Printf("Game %s", g.state)
}

func (g *Game) handleStartingState() error {
	g.initialize()
	g.setGameStateRunning()

	return nil
}

func (g *Game) setGameStateRunning() {
	g.state = gamestate.Running
	log.Default().Printf("Game %s", g.state)
}

func (g *Game) handleRunningState() error {
	input.DebugInput()
	input.GameControlInput()

	input.SnakeInput()

	// Step every n milliseconds
	now := time.Now()
	deltaTime := now.Sub(g.lastStep)
	if deltaTime >= g.timePerStep {
		g.lastStep = now
		g.step()
	}

	return nil
}

func (g *Game) setGameStatePaused() {
	g.ui.Container.AddChild()

	g.state = gamestate.Paused
	log.Default().Printf("Game %s", g.state)
}

func (g *Game) handlePausedState() error {
	g.ui.Update()

	input.DebugInput()
	input.GameControlInput()

	return nil
}

func (g *Game) setGameStateEnded() {
	g.ui.Container.AddChild()

	g.state = gamestate.Ended
	log.Default().Printf("Game %s", g.state)
}

func (g *Game) handleEndedState() error {
	g.ui.Update()

	input.DebugInput()
	input.GameControlInput()

	return nil
}

func (g *Game) setGameStateMenu() {
	g.removeFns["menu"] = g.ui.Container.AddChild(screens.CreateMenuScreen(g))

	g.state = gamestate.Menu
	log.Default().Printf("Game %s", g.state)
}

func (g *Game) handleMenuState() error {
	g.ui.Update()

	input.DebugInput()

	return nil
}
