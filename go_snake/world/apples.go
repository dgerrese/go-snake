package world

import (
	"go-snake/go_snake/apple"
	"image"
	"log"
	"math/rand/v2"
)

func (w *World) GetAppleAt(p image.Point) *apple.Apple {
	for _, a := range w.apples {
		if a.Position == p {
			return a
		}
	}

	return nil
}

func (w *World) RemoveApple(a *apple.Apple) {
	for i, wa := range w.apples {
		if wa == a {
			w.apples = append(w.apples[:i], w.apples[i+1:]...)
			return
		}
	}
}

func (w *World) supplyApples() {
	w.supplyApplesN(1)
}

func (w *World) supplyApplesN(n int) {
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
				w.apples = append(w.apples, apple.NewApple(p))

				break
			}

		}

	}
}
