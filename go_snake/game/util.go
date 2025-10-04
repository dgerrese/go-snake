package game

import (
	"log"
	"time"
)

func (g *Game) setDebugMode(mode bool) {
	g.debugMode = mode
	g.world.SetDebugMode(mode)
	log.Default().Printf("Debug mode: %t", mode)
}

func (g *Game) increaseGameSpeed(ms int64) {
	if !g.debugMode {
		log.Default().Print("Debug mode is off, cannot increase game speed")
		return
	}

	log.Default().Printf("Increasing game speed by %d msps", ms)
	if g.timePerStep.Milliseconds() > ms {
		g.timePerStep -= time.Duration(ms * int64(time.Millisecond))
	}
}

func (g *Game) decreaseGameSpeed(ms int64) {
	if !g.debugMode {
		log.Default().Print("Debug mode is off, cannot decrease game speed")
		return
	}

	log.Default().Printf("Decreasing game speed by %d msps", ms)
	g.timePerStep += time.Duration(ms * int64(time.Millisecond))
}
