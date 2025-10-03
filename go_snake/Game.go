package go_snake

import (
	"fmt"
	"github.com/hajimehoshi/bitmapfont/v3"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"go-snake/go_snake/config"
	"go-snake/go_snake/types"
	"image/color"
	"log"
	"time"
)

type Game struct {
	world       *World
	worldScale  float64
	score       int
	debugMode   bool
	lastStep    time.Time
	timePerStep time.Duration
}

func (g *Game) Update() error {
	g.HandleDevInput()

	// Update world debug mode
	g.world.debugMode = g.debugMode

	g.HandleSnakeInput()

	// Step every n milliseconds
	now := time.Now()
	deltaTime := now.Sub(g.lastStep)
	if deltaTime >= g.timePerStep {
		g.lastStep = now
		g.Step(deltaTime)
	}

	return nil
}

func (g *Game) Step(deltaTime time.Duration) {
	g.score++

	g.world.SupplyApples()

	if !g.EatApples() {
		g.world.snake.Move()
	}
}

func (g *Game) HandleSnakeInput() {
	if inpututil.IsKeyJustPressed(config.KeyMap[config.ActionMoveUp]) {
		g.world.snake.Turn(types.North)
	} else if inpututil.IsKeyJustPressed(config.KeyMap[config.ActionMoveDown]) {
		g.world.snake.Turn(types.South)
	} else if inpututil.IsKeyJustPressed(config.KeyMap[config.ActionMoveRight]) {
		g.world.snake.Turn(types.East)
	} else if inpututil.IsKeyJustPressed(config.KeyMap[config.ActionMoveLeft]) {
		g.world.snake.Turn(types.West)
	}
}

func (g *Game) HandleDevInput() {
	if inpututil.IsKeyJustPressed(ebiten.KeyF1) {
		g.debugMode = !g.debugMode
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyKPAdd) {
		if g.timePerStep.Milliseconds() > 100 {
			g.timePerStep -= 100 * time.Millisecond
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyKPSubtract) {
		g.timePerStep += 100 * time.Millisecond
	}
}

func (g *Game) EatApples() bool {
	for i, a := range g.world.apples {
		log.Default().Printf("Attempting to eat apple at %v", a.position)
		if g.world.snake.NextHead() == a.position {
			g.world.snake.Grow()
			g.world.apples = append(g.world.apples[:i], g.world.apples[i+1:]...)
			g.score += 10
			return true
		}
	}

	return false
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Clear()
	screen.Fill(color.RGBA{R: 0x44, G: 0x44, B: 0x44, A: 0xff})

	wo := &ebiten.DrawImageOptions{}
	wo.GeoM.Scale(g.worldScale, g.worldScale)

	wo.GeoM.Translate(
		float64(screen.Bounds().Max.X)/2-float64(g.world.width)*g.worldScale/2,
		float64(screen.Bounds().Max.Y)/2-float64(g.world.height)*g.worldScale/2,
	)
	screen.DrawImage(g.world.Render(), wo)

	// Draw score to screen
	so := &ebiten.DrawImageOptions{}
	so.GeoM.Translate(2, 2)

	text.Draw(screen, fmt.Sprintf("Score: %d", g.score), text.NewGoXFace(bitmapfont.Face), &text.DrawOptions{
		DrawImageOptions: *so,
	})

	do := &ebiten.DrawImageOptions{}
	do.GeoM.Translate(2, 20)

	dt := fmt.Sprintf("Speed: %d ms/step", g.timePerStep.Milliseconds())

	text.Draw(screen, dt, text.NewGoXFace(bitmapfont.Face), &text.DrawOptions{
		DrawImageOptions: *do,
	})
}

func (g *Game) Layout(outsideWidth int, outsideHeight int) (screenWidth int, screenHeight int) {
	scaledWidth := float64(outsideWidth) / float64(g.world.width)
	scaledHeight := float64(outsideHeight) / float64(g.world.height)

	if scaledHeight > scaledWidth {
		g.worldScale = scaledWidth
	} else {
		g.worldScale = scaledHeight
	}

	return outsideWidth, outsideHeight
}

func NewGame(width, height int, millisecondsPerStep int64) *Game {
	return &Game{
		world:       NewWorld(width, height),
		timePerStep: time.Duration(millisecondsPerStep * 1_000_000),
	}
}

func (g *Game) EnableDebugMode() {
	g.debugMode = true
}

func (g *Game) DisableDebugMode() {
	g.debugMode = false
}
