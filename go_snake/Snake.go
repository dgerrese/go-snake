package go_snake

import (
	"github.com/hajimehoshi/ebiten/v2"
	"go-snake/go_snake/types"
	"go-snake/go_snake/util"
	"image"
	"image/color"
	"log"
)

const DirectionQueueMaxLength = 3

var validTurns = map[types.Direction][]types.Direction{
	types.North:     {types.East, types.West},
	types.East:      {types.North, types.South},
	types.South:     {types.East, types.West},
	types.West:      {types.North, types.South},
	types.Undefined: {types.North, types.East, types.South, types.West},
}

type Snake struct {
	body           []image.Point
	direction      types.Direction
	directionQueue []types.Direction
}

func NewSnake(origin image.Point, length int) *Snake {
	body := make([]image.Point, length)
	for i := 0; i < length; i++ {
		body[i] = image.Point{X: origin.X, Y: origin.Y + i}
	}

	return &Snake{
		body:      body,
		direction: types.Undefined,
	}
}

func (s *Snake) Head() image.Point {
	return s.body[0]
}

func (s *Snake) Move() {
	if len(s.directionQueue) > 0 {
		s.direction = s.directionQueue[0]
		s.directionQueue = s.directionQueue[1:]

		log.Default().Printf("Changed direction to: %v. Next in queue: %+v\n", s.direction, s.directionQueue)
	}

	nh := s.NextHead()

	if s.Head() == nh {
		// Snake is not moving
		return
	}

	for _, bp := range s.body {
		if bp == nh {
			log.Default().Println("Snake collided with itself!")
			s.directionQueue = nil
			s.direction = types.Undefined
			return
		}
	}

	s.body = append([]image.Point{nh}, s.body[:len(s.body)-1]...)
}

func (s *Snake) Turn(d types.Direction) {
	if len(s.directionQueue) >= DirectionQueueMaxLength {
		return
	}

	var ld types.Direction

	if len(s.directionQueue) > 0 {
		ld = s.directionQueue[len(s.directionQueue)-1]
	} else {
		ld = s.direction
	}

	for _, v := range validTurns[ld] {
		if d == v {
			s.directionQueue = append(s.directionQueue, d)
			return
		}
	}
}

func (s *Snake) NextHead() image.Point {
	h := s.Head()

	switch s.direction {
	case types.North:
		return image.Point{X: h.X, Y: h.Y - 1}
	case types.East:
		return image.Point{X: h.X + 1, Y: h.Y}
	case types.South:
		return image.Point{X: h.X, Y: h.Y + 1}
	case types.West:
		return image.Point{X: h.X - 1, Y: h.Y}
	default:
		return h
	}
}

func (s *Snake) Grow() {
	nh := s.NextHead()

	s.body = append([]image.Point{nh}, s.body...)
}

func (s *Snake) Render() (*ebiten.Image, *ebiten.DrawImageOptions) {
	r := image.Rectangle{
		Min: image.Point{
			X: util.IntListMin(util.MapSlice(s.body, func(point image.Point) int {
				return point.X
			})),
			Y: util.IntListMin(util.MapSlice(s.body, func(point image.Point) int {
				return point.Y
			})),
		},
		Max: image.Point{
			X: util.IntListMax(util.MapSlice(s.body, func(point image.Point) int {
				return point.X
			})),
			Y: util.IntListMax(util.MapSlice(s.body, func(point image.Point) int {
				return point.Y
			})),
		},
	}

	w := r.Dx() + 1
	h := r.Dy() + 1

	i := ebiten.NewImage(w, h)

	for _, point := range s.body {
		px := point.X - r.Min.X
		py := point.Y - r.Min.Y

		i.Set(px, py, color.White)
	}

	o := &ebiten.DrawImageOptions{}
	o.GeoM.Translate(
		float64(r.Min.X),
		float64(r.Min.Y),
	)

	return i, o
}

func (s *Snake) Occupies(p image.Point) bool {
	for _, bp := range s.body {
		if bp == p {
			return true
		}
	}

	return false
}
