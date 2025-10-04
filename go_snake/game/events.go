package game

import (
	"go-snake/go_snake/config"
	"go-snake/go_snake/events"
	"go-snake/go_snake/types/gamestate"
	"log"
)

func (g *Game) listenForMoveEvents(stopCh chan any) {
	for {
		select {
		case <-stopCh:
			log.Default().Printf("Stopped listening for move events")
			return
		case <-events.MovementCh:
			g.score++
		}
	}
}

func (g *Game) listenForEatEvents(stopCh chan any) {
	for {
		select {
		case <-stopCh:
			log.Default().Printf("Stopped listening for eat events")
			return
		case a := <-events.EatingCh:
			log.Default().Printf("Apple eaten: %v", a)
			g.score += 100
			g.world.RemoveApple(a)
		}
	}
}

func (g *Game) listenForActionEvents(stopCh chan any) {
	for {
		select {
		case <-stopCh:
			log.Default().Printf("Stopped listening for action events")
			return
		case e := <-events.ActionCh:
			switch e {
			case config.ActionToggleDebugMode:
				g.setDebugMode(!g.debugMode)
			case config.ActionIncreaseGameSpeed:
				g.increaseGameSpeed(20)
			case config.ActionDecreaseGameSpeed:
				g.decreaseGameSpeed(20)
			case config.ActionPauseGame:
				g.toggleGamePause()
			case config.ActionRestart:
				if g.state == gamestate.Ended {
					g.restart()
				}
			default:
				// Put unrecognized events back to the channel
				events.ActionCh <- e
			}
		}
	}
}

func (g *Game) listenForGameEvents(stopCh chan any) {
	for {
		select {
		case <-stopCh:
			log.Default().Printf("Stopped listening for game events")
			return
		case e := <-events.GameCh:
			switch e {
			case events.DeathEvent:
				g.setGameState(gamestate.Ended)
			}
		}
	}
}
