package world

import (
	"go-snake/go_snake/apple"
	"go-snake/go_snake/snake"
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type World struct {
	width     int
	height    int
	apples    []*apple.Apple
	snake     *snake.Snake
	debugMode bool
}

func NewWorld(w, h int) *World {
	world := &World{
		width:     w,
		height:    h,
		apples:    make([]*apple.Apple, 0),
		debugMode: false,
	}

	world.initialize()

	return world
}

func (w *World) Process() {
	w.snake.Move()

	w.supplyApples()
}

func (w *World) Render() *ebiten.Image {
	img := ebiten.NewImage(w.width, w.height)
	img.Fill(color.Black)

	if w.debugMode {
		img.DrawImage(w.renderGrid(), &ebiten.DrawImageOptions{})
	}

	for i := range w.apples {
		img.DrawImage(w.apples[i].Render())
	}

	img.DrawImage(w.snake.Render())

	return img
}

func (w *World) initialize() {
	w.snake = snake.NewSnake(w, image.Point{X: w.width / 2, Y: w.height / 2}, 1)
	w.supplyApples()
}

func (w *World) renderGrid() *ebiten.Image {
	i := ebiten.NewImage(w.width, w.height)

	c := color.RGBA{R: 0x22, G: 0x22, B: 0x22, A: 0x22}

	for x := 0; x < w.width; x++ {
		for y := 0; y < w.height; y++ {
			if (x%2 == 0 && y%2 == 0) || (x%2 != 0 && y%2 != 0) {
				vector.DrawFilledRect(
					i,
					float32(x),
					float32(y),
					1,
					1,
					c,
					false,
				)
			}
		}
	}

	return i
}

func (w *World) Destroy() {
	w.snake.Destroy()
}
