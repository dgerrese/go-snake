package screens

import (
	"github.com/ebitenui/ebitenui/themes"
	"github.com/ebitenui/ebitenui/widget"

	game2 "go-snake/go_snake/types/game"
)

func CreateMenuScreen(g game2.Game) *widget.Container {
	t := themes.GetBasicDarkTheme()

	c := widget.NewContainer()

	c.AddChild(createStartButton(t, g))

	return c
}

func createStartButton(t *widget.Theme, g game2.Game) *widget.Button {
	return widget.NewButton(
		widget.ButtonOpts.TextLabel("New game"),
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			g.Start()
		}),
		widget.ButtonOpts.TextFace(t.DefaultFace),
		widget.ButtonOpts.TextColor(&widget.ButtonTextColor{
			Idle:     t.DefaultTextColor,
			Disabled: t.DefaultTextColor,
			Hover:    t.DefaultTextColor,
			Pressed:  t.DefaultTextColor,
		}),
		widget.ButtonOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(
				widget.AnchorLayoutData{
					HorizontalPosition: widget.AnchorLayoutPositionCenter,
					VerticalPosition:   widget.AnchorLayoutPositionCenter,
					StretchHorizontal:  false,
					StretchVertical:    false,
				},
			),
		),
	)
}
