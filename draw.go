package rect

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// Draw a rectangle on the screen using the ebiten vector package with white default color
func (r *Rectangle) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(r.X), float64(r.Y))
	vector.DrawFilledRect(screen, float32(r.X), float32(r.Y), float32(r.Width), float32(r.Height), r.Color)
}

// DrawCircle draws an ellipse on the screen using the ebiten vector package with white default color
func (r *Rectangle) DrawCircle(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(r.X), float64(r.Y))
	vector.DrawFilledCircle(screen, float32(r.X), float32(r.Y), float32(r.Width), r.Color)
}
