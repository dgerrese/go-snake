package game

import (
	"fmt"
	"go-snake/go_snake/config"
	"strings"

	"github.com/hajimehoshi/bitmapfont/v3"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
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
	t.WriteString(fmt.Sprintf("\tWorld scale: %.2f\n", g.worldScale))

	ww, wh := g.world.Size()
	t.WriteString(fmt.Sprintf("\tWorld size: %dx%d\n", ww, wh))

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
	ff := bitmapfont.Face
	var ls = float64(16)

	tl := []string{
		"Game Over!",
		fmt.Sprintf("Final Score: %d", g.score),
		"Press R to Restart",
	}

	var maxX int
	for _, l := range tl {
		var lxsum int
		for _, r := range l {
			rb, _, _ := ff.GlyphBounds(r)
			lxsum += (rb.Max.X - rb.Min.X).Round()
		}
		maxX = max(maxX, lxsum)
	}

	i := ebiten.NewImage(maxX, len(tl)*int(ls))

	s := (g.worldScale * float64(min(config.GameConfig.WorldWidth, config.GameConfig.WorldHeight))) / float64(max(i.Bounds().Max.X, i.Bounds().Max.Y))

	io := &ebiten.DrawImageOptions{}
	io.GeoM.Translate(
		float64(screen.Bounds().Max.X)/(2*s)-float64(i.Bounds().Max.X)/2,
		float64(screen.Bounds().Max.Y)/(2*s)-float64(i.Bounds().Max.Y)/2,
	)
	io.GeoM.Scale(s, s)

	to := ebiten.DrawImageOptions{}
	to.GeoM.Translate(float64(i.Bounds().Max.X)/2, float64(i.Bounds().Max.Y)/2)
	text.Draw(i, strings.Join(tl, "\n"), text.NewGoXFace(ff), &text.DrawOptions{
		DrawImageOptions: to,
		LayoutOptions: text.LayoutOptions{
			PrimaryAlign:   text.AlignCenter,
			SecondaryAlign: text.AlignCenter,
			LineSpacing:    ls,
		},
	})

	screen.DrawImage(i, io)
}
