package snake

import (
	"go-snake/go_snake/events"
	"go-snake/go_snake/types/direction"
	"go-snake/go_snake/types/world"
	"go-snake/go_snake/util"
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Snake struct {
	world          world.World
	body           []image.Point
	direction      direction.Direction
	directionQueue []direction.Direction
	destroyCh      chan any
}

func NewSnake(w world.World, origin image.Point, length int) *Snake {
	body := make([]image.Point, length)
	for i := 0; i < length; i++ {
		body[i] = image.Point{X: origin.X, Y: origin.Y + i}
	}

	s := &Snake{
		world:     w,
		body:      body,
		direction: direction.Undefined,
		destroyCh: make(chan any),
	}

	go s.listenForTurnEvents(s.destroyCh)

	return s
}

func (s *Snake) Head() image.Point {
	return s.body[0]
}

func (s *Snake) NextHead() image.Point {
	h := s.Head()

	switch s.direction {
	case direction.North:
		return image.Point{X: h.X, Y: h.Y - 1}
	case direction.East:
		return image.Point{X: h.X + 1, Y: h.Y}
	case direction.South:
		return image.Point{X: h.X, Y: h.Y + 1}
	case direction.West:
		return image.Point{X: h.X - 1, Y: h.Y}
	default:
		return h
	}
}

func (s *Snake) eat(p image.Point) bool {
	w, h := s.world.Size()
	if p.X < 0 || p.Y < 0 || p.X >= w || p.Y >= h || s.Occupies(p) {
		events.GameCh <- events.DeathEvent
		return true
	}

	a := s.world.GetAppleAt(p)
	if a != nil {
		s.Grow()
		events.EatingCh <- a
		return true
	}

	return false
}

func (s *Snake) Grow() {
	nh := s.NextHead()

	s.body = append([]image.Point{nh}, s.body...)
}

func (s *Snake) Render() (*ebiten.Image, *ebiten.DrawImageOptions) {
	r := image.Rectangle{
		Min: image.Point{
			X: util.IntSliceMin(util.MapSlice(s.body, func(point image.Point) int {
				return point.X
			})),
			Y: util.IntSliceMin(util.MapSlice(s.body, func(point image.Point) int {
				return point.Y
			})),
		},
		Max: image.Point{
			X: util.IntSliceMax(util.MapSlice(s.body, func(point image.Point) int {
				return point.X
			})),
			Y: util.IntSliceMax(util.MapSlice(s.body, func(point image.Point) int {
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

func (s *Snake) Destroy() {
	close(s.destroyCh)
}
