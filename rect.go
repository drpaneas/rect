package rect

type rectangle struct {
	X, Y, Width, Height int
}

// Rect creates a new rectangle with the given X, Y, Width, and Height.
func Rect(x, y, width, height int) *rectangle {
	return &rectangle{x, y, width, height}
}

func (r *rectangle) getTop() int {
	return r.Y
}
	
func (r *rectangle) getLeft() int {
	return r.X
}

func (r *rectangle) getBottom() int {
	return r.Y + r.Height
}

func (r *rectangle) getRight() int {
	return r.X + r.Width
}

func (r *rectangle) getTopLeft() (int, int) {
	return r.X, r.Y
}

func (r *rectangle) getBottomLeft() (int, int) {
	return r.X, r.Y + r.Height
}

func (r *rectangle) getTopRight() (int, int) {
	return r.X + r.Width, r.Y
}

func (r *rectangle) getBottomRight() (int, int) {
	return r.X + r.Width, r.Y + r.Height
}

func (r *rectangle) getMidTop() (int, int) {
	return r.X + r.Width/2, r.Y
}

func (r *rectangle) getMidLeft() (int, int) {
	return r.X, r.Y + r.Height/2
}

func (r *rectangle) getMidBottom() (int, int) {
	return r.X + r.Width/2, r.Y + r.Height
}

func (r *rectangle) getMidRight() (int, int) {
	return r.X + r.Width, r.Y + r.Height/2
}

func (r *rectangle) getCenter() (int, int) {
	return r.X + r.Width/2, r.Y + r.Height/2
}

func (r *rectangle) getCenterX() int {
	return r.X + r.Width/2
}

func (r *rectangle) getCenterY() int {
	return r.Y + r.Height/2
}

func (r *rectangle) getSize() (int, int) {
	return r.Width, r.Height
}

func (r *rectangle) getWidth() int {
	return r.Width
}

func (r *rectangle) setTop(val int) {
	r.Y = val
}

func (r *rectangle) setRight(val int) {
	r.X = val - r.Width
}

func (r *rectangle) setBottom(val int) {
	r.Y = val - r.Height
}

func (r *rectangle) setX(val int) {
	r.X = val
}

func (r *rectangle) setY(val int) {
	r.Y = val
}

func (r *rectangle) setTopLeft(x, y int) {
	r.X = x
	r.Y = y
}

func (r *rectangle) setBottomLeft(x, y int) {
	r.X = x
	r.Y = y - r.Height
}

func (r *rectangle) setTopRight(x, y int) {
	r.X = x - r.Width
	r.Y = y
}

func (r *rectangle) setBottomRight(x, y int) {
	r.X = x - r.Width
	r.Y = y - r.Height
}

func (r *rectangle) setMidTop(x, y int) {
	r.X = x - r.Width/2
	r.Y = y
}

func (r *rectangle) setMidLeft(x, y int) {
	r.X = x
	r.Y = y - r.Height/2
}

func (r *rectangle) setMidBottom(x, y int) {
	r.X = x - r.Width/2
	r.Y = y - r.Height
}

func (r *rectangle) setMidRight(x, y int) {
	r.X = x - r.Width
	r.Y = y - r.Height/2
}

func (r *rectangle) setCenter(x, y int) {
	r.X = x - r.Width/2
	r.Y = y - r.Height/2
}

func (r *rectangle) setCenterX(val int) {
	r.X = val - r.Width/2
}

func (r *rectangle) setCenterY(val int) {
	r.Y = val - r.Height/2
}

func (r *rectangle) setSize(width, height int) {
	r.Width = width
	r.Height = height
}

func (r *rectangle) setHeight(val int) {
	r.Height = val
}

func (r *rectangle) Top(val ...int) int {
	if len(val) > 0 {
		r.setTop(val[0])
	}
	return r.getTop()
}

func (r *rectangle) Left(val ...int) int {
	if len(val) > 0 {
		r.setX(val[0])
	}
	return r.X
}

func (r *rectangle) Bottom(val ...int) int {
	if len(val) > 0 {
		r.setBottom(val[0])
	}
	return r.getBottom()
}

func (r *rectangle) Right(val ...int) int {
	if len(val) > 0 {
		r.setRight(val[0])
	}
	return r.getRight()
}

func (r *rectangle) TopLeft(val ...int) (int, int) {
	if len(val) > 1 {
		r.setTopLeft(val[0], val[1])
	}
	return r.getTopLeft()
}

func (r *rectangle) BottomLeft(val ...int) (int, int) {
	if len(val) > 1 {
		r.setBottomLeft(val[0], val[1])
	}
	return r.getBottomLeft()
}

func (r *rectangle) TopRight(val ...int) (int, int) {
	if len(val) > 1 {
		r.setTopRight(val[0], val[1])
	}
	return r.getTopRight()
}

func (r *rectangle) BottomRight(val ...int) (int, int) {
	if len(val) > 1 {
		r.setBottomRight(val[0], val[1])
	}
	return r.getBottomRight()
}

func (r *rectangle) MidTop(val ...int) (int, int) {
	if len(val) > 1 {
		r.setMidTop(val[0], val[1])
	}
	return r.getMidTop()
}

func (r *rectangle) MidLeft(val ...int) (int, int) {
	if len(val) > 1 {
		r.setMidLeft(val[0], val[1])
	}
	return r.getMidLeft()
}

func (r *rectangle) MidBottom(val ...int) (int, int) {
	if len(val) > 1 {
		r.setMidBottom(val[0], val[1])
	}
	return r.getMidBottom()
}

func (r *rectangle) MidRight(val ...int) (int, int) {
	if len(val) > 1 {
		r.setMidRight(val[0], val[1])
	}
	return r.getMidRight()
}

func (r *rectangle) Center(val ...int) (int, int) {
	if len(val) > 1 {
		r.setCenter(val[0], val[1])
	}
	return r.getCenter()
}

func (r *rectangle) CenterX(val ...int) int {
	if len(val) > 0 {
		r.setCenterX(val[0])
	}
	return r.getCenterX()
}

func (r *rectangle) CenterY(val ...int) int {
	if len(val) > 0 {
		r.setCenterY(val[0])
	}
	return r.getCenterY()
}

func (r *rectangle) Size(val ...int) (int, int) {
	if len(val) > 1 {
		r.setSize(val[0], val[1])
	}
	return r.getSize()
}
