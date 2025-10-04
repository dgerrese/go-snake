package game

import (
	"go-snake/go_snake/input"
	"go-snake/go_snake/types/gamestate"
	"log"
	"time"
)

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

func (g *Game) handlePausedState() error {
	input.DebugInput()
	input.GameControlInput()

	return nil
}

func (g *Game) handleEndedState() error {
	input.DebugInput()
	input.GameControlInput()

	return nil
}

func (g *Game) handleStartingState() error {
	g.initialize()
	g.setGameState(gamestate.Running)

	return nil
}

func (g *Game) handleMenuState() error {
	input.DebugInput()

	return nil
}

func (g *Game) setGameState(s gamestate.GameState) {
	g.state = s
	log.Default().Printf("Game %s", g.state)
}

func (g *Game) toggleGamePause() {
	switch g.state {
	case gamestate.Running:
		g.setGameState(gamestate.Paused)
	case gamestate.Paused:
		g.setGameState(gamestate.Running)
	default:
		// Do nothing
	}
}
