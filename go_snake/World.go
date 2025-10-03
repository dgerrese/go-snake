package go_snake

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image"
	"image/color"
	"log"
	"math/rand/v2"
)

type World struct {
	width     int
	height    int
	apples    []Apple
	snake     Snake
	debugMode bool
}

func NewWorld(w, h int) *World {
	log.Default().Printf("World size: %dx%d\n", w, h)
	return &World{
		width:     w,
		height:    h,
		apples:    make([]Apple, 0),
		snake:     *NewSnake(image.Point{X: w / 2, Y: h / 2}, 1),
		debugMode: false,
	}
}

func (w *World) renderGrid() *ebiten.Image {
	i := ebiten.NewImage(w.width, w.height)

	ec := color.RGBA{R: 0x7F, G: 0x7F, B: 0x7F, A: 0x7F}
	oc := color.RGBA{R: 0x3F, G: 0x3F, B: 0x3F, A: 0x7F}

	for x := 0; x < w.width; x++ {
		for y := 0; y < w.height; y++ {
			if (x%2 == 0 && y%2 == 0) || (x%2 != 0 && y%2 != 0) {
				vector.DrawFilledRect(
					i,
					float32(x),
					float32(y),
					1,
					1,
					ec,
					false,
				)
			} else {
				vector.DrawFilledRect(
					i,
					float32(x),
					float32(y),
					1,
					1,
					oc,
					false,
				)
			}
		}
	}

	return i
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

func (w *World) SupplyApples() {
	w.SupplyApplesN(1)
}

func (w *World) SupplyApplesN(n int) {
	if len(w.apples) >= n {
		return
	}

	for i := len(w.apples); i < n; i++ {

		for {
			p := image.Point{
				X: rand.IntN(w.width - 1),
				Y: rand.IntN(w.height - 1),
			}

			if !w.snake.Occupies(p) {
				log.Default().Println("New apple at: " + p.String())
				w.apples = append(w.apples, *NewApple(p))

				break
			}

		}

	}
}
