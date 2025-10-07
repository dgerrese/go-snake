package game

import (
	"go-snake/go_snake/config"
	"go-snake/go_snake/types/gamestate"
	"go-snake/go_snake/world"
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	world       *world.World
	worldScale  float64
	lastStep    time.Time
	timePerStep time.Duration
	score       int
	debugMode   bool
	state       gamestate.GameState
	stopCh      chan any
}

func NewGame(msps int64) *Game {
	return &Game{
		timePerStep: time.Duration(msps * int64(time.Millisecond)),
		state:       gamestate.Starting,
	}
}

func (g *Game) Update() error {
	switch g.state {
	case gamestate.Starting:
		return g.handleStartingState()
	case gamestate.Running:
		return g.handleRunningState()
	case gamestate.Paused:
		return g.handlePausedState()
	case gamestate.Ended:
		return g.handleEndedState()
	case gamestate.Menu:
		return g.handleMenuState()
	}

	panic("invalid game state")
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Clear()
	screen.Fill(color.RGBA{R: 0x44, G: 0x44, B: 0x44, A: 0xff})

	if g.state == gamestate.Ended {
		g.drawGameOver(screen)
	} else {
		g.drawWorld(screen)
		g.drawScore(screen)
	}

	if g.debugMode {
		g.drawDebugInfo(screen)
	}
}

func (g *Game) Layout(outsideWidth int, outsideHeight int) (screenWidth int, screenHeight int) {
	if g.world == nil {
		return outsideWidth, outsideHeight
	}

	ww, wh := g.world.Size()
	scaledWidth := float64(outsideWidth) / float64(ww)
	scaledHeight := float64(outsideHeight) / float64(wh)

	if scaledHeight > scaledWidth {
		g.worldScale = scaledWidth
	} else {
		g.worldScale = scaledHeight
	}

	return outsideWidth, outsideHeight
}

func (g *Game) initialize() {
	g.world = world.NewWorld(config.GameConfig.WorldWidth, config.GameConfig.WorldHeight)

	go g.listenForGameEvents(g.stopCh)
	go g.listenForActionEvents(g.stopCh)
	go g.listenForMoveEvents(g.stopCh)
	go g.listenForEatEvents(g.stopCh)
}

func (g *Game) step() {
	g.world.Process()
}

func (g *Game) restart() {
	g.score = 0
	g.world.Destroy()
	g.world = world.NewWorld(config.GameConfig.WorldWidth, config.GameConfig.WorldHeight)
	g.setGameState(gamestate.Running)
}
