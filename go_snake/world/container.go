package world

import (
	euimage "github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"golang.org/x/image/colornames"
)

func (w *World) createContainer() {
	w.Container = widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(
			euimage.NewNineSliceColor(colornames.Whitesmoke),
		),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(
				widget.AnchorLayoutData{
					HorizontalPosition: widget.AnchorLayoutPositionCenter,
					VerticalPosition:   widget.AnchorLayoutPositionCenter,
					StretchHorizontal:  false,
					StretchVertical:    false,
					Padding:            nil,
				},
			),
		),
	)
}

func (w *World) scaleContainer() {
	ww, wh := w.Size()
	ww *= 2
	wh *= 2

	wid := w.Container.GetWidget()
	wid.MinWidth = ww
	wid.MinHeight = wh
}
