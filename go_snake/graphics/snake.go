package graphics

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type SnakePart int

const (
	SnakeGreenHeadNorth SnakePart = 1
	SnakeGreenHeadWest  SnakePart = 2
	SnakeGreenHeadSouth SnakePart = 3
	SnakeGreenHeadEast  SnakePart = 4
	SnakeGreenBody      SnakePart = 5

	SnakeGreyHeadNorth SnakePart = 16 + 1
	SnakeGreyHeadWest  SnakePart = 16 + 2
	SnakeGreyHeadSouth SnakePart = 16 + 3
	SnakeGreyHeadEast  SnakePart = 16 + 4
	SnakeGreyBody      SnakePart = 16 + 5

	SnakeBlueHeadNorth SnakePart = (4 * 16) + 1
	SnakeBlueHeadWest  SnakePart = (4 * 16) + 2
	SnakeBlueHeadSouth SnakePart = (4 * 16) + 3
	SnakeBlueHeadEast  SnakePart = (4 * 16) + 4
	SnakeBlueBody      SnakePart = (4 * 16) + 5
)

func (s SnakePart) Image() *ebiten.Image {
	return getTile(int(s))
}
