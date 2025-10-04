package game

import (
	"fmt"
	"github.com/hajimehoshi/bitmapfont/v3"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"strings"
)

func (g *Game) drawWorld(screen *ebiten.Image) {
	o := &ebiten.DrawImageOptions{}
	o.GeoM.Scale(g.worldScale, g.worldScale)

	w, h := g.world.Size()

	o.GeoM.Translate(
		float64(screen.Bounds().Max.X)/2-float64(w)*g.worldScale/2,
		float64(screen.Bounds().Max.Y)/2-float64(h)*g.worldScale/2,
	)
	screen.DrawImage(g.world.Render(), o)
}

func (g *Game) drawScore(screen *ebiten.Image) {
	o := &ebiten.DrawImageOptions{}
	o.GeoM.Translate(2, 2)

	text.Draw(screen, fmt.Sprintf("Score: %d", g.score), text.NewGoXFace(bitmapfont.Face), &text.DrawOptions{
		DrawImageOptions: *o,
	})
}

func (g *Game) drawDebugInfo(screen *ebiten.Image) {
	o := ebiten.DrawImageOptions{}
	o.GeoM.Translate(2, 32)

	t := strings.Builder{}
	t.WriteString("Debug Info:\n")
	t.WriteString(fmt.Sprintf("\tFPS: %.2f\n", ebiten.ActualFPS()))
	t.WriteString(fmt.Sprintf("\tGame speed: %d ms/step\n", g.timePerStep.Milliseconds()))

	text.Draw(screen, t.String(), text.NewGoXFace(bitmapfont.Face), &text.DrawOptions{
		DrawImageOptions: o,
		LayoutOptions: text.LayoutOptions{
			LineSpacing:    16,
			PrimaryAlign:   text.AlignStart,
			SecondaryAlign: text.AlignStart,
		},
	})
}

func (g *Game) drawGameOver(screen *ebiten.Image) {
	var s float64 = 2
	o := ebiten.DrawImageOptions{}
	o.GeoM.Translate(
		float64(screen.Bounds().Max.X)/(2*s),
		float64(screen.Bounds().Max.Y)/(2*s),
	)
	o.GeoM.Scale(s, s)

	t := strings.Builder{}
	t.WriteString("Game Over!\n")
	t.WriteString(fmt.Sprintf("Final Score: %d\n", g.score))
	t.WriteString("Press R to Restart")

	text.Draw(screen, t.String(), text.NewGoXFace(bitmapfont.Face), &text.DrawOptions{
		DrawImageOptions: o,
		LayoutOptions: text.LayoutOptions{
			PrimaryAlign:   text.AlignCenter,
			SecondaryAlign: text.AlignStart,
			LineSpacing:    16,
		},
	})
}
