package rect

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// Draw a rectangle on the screen using the ebiten vector package with white default color
func (r *Rectangle) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(r.X), float64(r.Y))
	vector.DrawFilledRect(screen, float32(r.X), float32(r.Y), float32(r.Width), float32(r.Height), r.Color, true)
}

// DrawCircle draws an ellipse on the screen using the ebiten vector package with white default color
// this is a rectangle, but for some reason we want to draw circle
func (r *Rectangle) DrawCircle(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(r.X), float64(r.Y))
	// this is a rectangle, but we want to draw circle
	// the rectangle was setup with X and Y being the upper left corner
	// but the circle is drawn with X and Y being the center
	// so we need to adjust the X and Y to be the center
	cX := r.X + r.Width/2
	cY := r.Y + r.Height/2
	radius := r.Width / 2
	vector.DrawFilledCircle(screen, float32(cX), float32(cY), float32(radius), r.Color, true)
}
