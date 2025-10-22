package graphics

import (
	"bytes"
	_ "embed"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const tileSize = 8

var (
	//go:embed snake.png
	TilesetImg []byte
)

func tileIndex(x, y int) int {
	return x + y*16
}

func getTilesetImage() *ebiten.Image {
	ii, _, err := image.Decode(bytes.NewReader(TilesetImg))
	if err != nil {
		log.Fatal(err)
	}

	return ebiten.NewImageFromImage(ii)
}

func getTile(i int) *ebiten.Image {
	ti := getTilesetImage()
	w := ti.Bounds().Dx()
	tileXCount := w / tileSize

	sx := (i % tileXCount) * tileSize
	sy := (i / tileXCount) * tileSize

	return ebiten.NewImageFromImage(ti.SubImage(image.Rect(sx, sy, sx+tileSize, sy+tileSize)))
}
