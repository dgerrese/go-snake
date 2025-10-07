package apple

import (
	"image"
	"image/color"
	"math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
)

type Apple struct {
	Position image.Point
	Color    color.RGBA
}

func NewApple(pos image.Point) *Apple {
	return &Apple{
		Position: pos,
		Color:    randomAppleColor(),
	}
}

func (a *Apple) Render() (*ebiten.Image, *ebiten.DrawImageOptions) {
	img := ebiten.NewImage(1, 1)
	img.Fill(a.Color)

	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(
		float64(a.Position.X),
		float64(a.Position.Y),
	)

	return img, opts
}

func randomAppleColor() color.RGBA {
	green := rand.IntN(1) == 1

	if green {
		r := uint8(rand.IntN(64))

		return color.RGBA{R: r, G: 0xff, B: 0, A: 255}
	}

	g := uint8(rand.IntN(64))

	return color.RGBA{R: 0xff, G: g, B: 0, A: 255}
}
