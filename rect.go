package rect

type Rectangle struct {
	X, Y, Width, Height int
}

// Rect creates a new Rectangle with the given X, Y, Width, and Height.
func Rect(x, y, width, height int) *Rectangle {
	return &Rectangle{x, y, width, height}
}

func (r *Rectangle) getTop() int {
	return r.Y
}

func (r *Rectangle) getLeft() int {
	return r.X
}

func (r *Rectangle) getBottom() int {
	return r.Y + r.Height
}

func (r *Rectangle) getRight() int {
	return r.X + r.Width
}

func (r *Rectangle) getTopLeft() (int, int) {
	return r.X, r.Y
}

func (r *Rectangle) getBottomLeft() (int, int) {
	return r.X, r.Y + r.Height
}

func (r *Rectangle) getTopRight() (int, int) {
	return r.X + r.Width, r.Y
}

func (r *Rectangle) getBottomRight() (int, int) {
	return r.X + r.Width, r.Y + r.Height
}

func (r *Rectangle) getMidTop() (int, int) {
	return r.X + r.Width/2, r.Y
}

func (r *Rectangle) getMidLeft() (int, int) {
	return r.X, r.Y + r.Height/2
}

func (r *Rectangle) getMidBottom() (int, int) {
	return r.X + r.Width/2, r.Y + r.Height
}

func (r *Rectangle) getMidRight() (int, int) {
	return r.X + r.Width, r.Y + r.Height/2
}

func (r *Rectangle) getCenter() (int, int) {
	return r.X + r.Width/2, r.Y + r.Height/2
}

func (r *Rectangle) getCenterX() int {
	return r.X + r.Width/2
}

func (r *Rectangle) getCenterY() int {
	return r.Y + r.Height/2
}

func (r *Rectangle) getSize() (int, int) {
	return r.Width, r.Height
}

func (r *Rectangle) getWidth() int {
	return r.Width
}

func (r *Rectangle) setTop(val int) {
	r.Y = val
}

func (r *Rectangle) setRight(val int) {
	r.X = val - r.Width
}

func (r *Rectangle) setBottom(val int) {
	r.Y = val - r.Height
}

func (r *Rectangle) setX(val int) {
	r.X = val
}

func (r *Rectangle) setY(val int) {
	r.Y = val
}

func (r *Rectangle) setTopLeft(x, y int) {
	r.X = x
	r.Y = y
}

func (r *Rectangle) setBottomLeft(x, y int) {
	r.X = x
	r.Y = y - r.Height
}

func (r *Rectangle) setTopRight(x, y int) {
	r.X = x - r.Width
	r.Y = y
}

func (r *Rectangle) setBottomRight(x, y int) {
	r.X = x - r.Width
	r.Y = y - r.Height
}

func (r *Rectangle) setMidTop(x, y int) {
	r.X = x - r.Width/2
	r.Y = y
}

func (r *Rectangle) setMidLeft(x, y int) {
	r.X = x
	r.Y = y - r.Height/2
}

func (r *Rectangle) setMidBottom(x, y int) {
	r.X = x - r.Width/2
	r.Y = y - r.Height
}

func (r *Rectangle) setMidRight(x, y int) {
	r.X = x - r.Width
	r.Y = y - r.Height/2
}

func (r *Rectangle) setCenter(x, y int) {
	r.X = x - r.Width/2
	r.Y = y - r.Height/2
}

func (r *Rectangle) setCenterX(val int) {
	r.X = val - r.Width/2
}

func (r *Rectangle) setCenterY(val int) {
	r.Y = val - r.Height/2
}

func (r *Rectangle) setSize(width, height int) {
	r.Width = width
	r.Height = height
}

func (r *Rectangle) setHeight(val int) {
	r.Height = val
}

func (r *Rectangle) Top(val ...int) int {
	if len(val) > 0 {
		r.setTop(val[0])
	}
	return r.getTop()
}

func (r *Rectangle) Left(val ...int) int {
	if len(val) > 0 {
		r.setX(val[0])
	}
	return r.X
}

func (r *Rectangle) Bottom(val ...int) int {
	if len(val) > 0 {
		r.setBottom(val[0])
	}
	return r.getBottom()
}

func (r *Rectangle) Right(val ...int) int {
	if len(val) > 0 {
		r.setRight(val[0])
	}
	return r.getRight()
}

func (r *Rectangle) TopLeft(val ...int) (int, int) {
	if len(val) > 1 {
		r.setTopLeft(val[0], val[1])
	}
	return r.getTopLeft()
}

func (r *Rectangle) BottomLeft(val ...int) (int, int) {
	if len(val) > 1 {
		r.setBottomLeft(val[0], val[1])
	}
	return r.getBottomLeft()
}

func (r *Rectangle) TopRight(val ...int) (int, int) {
	if len(val) > 1 {
		r.setTopRight(val[0], val[1])
	}
	return r.getTopRight()
}

func (r *Rectangle) BottomRight(val ...int) (int, int) {
	if len(val) > 1 {
		r.setBottomRight(val[0], val[1])
	}
	return r.getBottomRight()
}

func (r *Rectangle) MidTop(val ...int) (int, int) {
	if len(val) > 1 {
		r.setMidTop(val[0], val[1])
	}
	return r.getMidTop()
}

func (r *Rectangle) MidLeft(val ...int) (int, int) {
	if len(val) > 1 {
		r.setMidLeft(val[0], val[1])
	}
	return r.getMidLeft()
}

func (r *Rectangle) MidBottom(val ...int) (int, int) {
	if len(val) > 1 {
		r.setMidBottom(val[0], val[1])
	}
	return r.getMidBottom()
}

func (r *Rectangle) MidRight(val ...int) (int, int) {
	if len(val) > 1 {
		r.setMidRight(val[0], val[1])
	}
	return r.getMidRight()
}

func (r *Rectangle) Center(val ...int) (int, int) {
	if len(val) > 1 {
		r.setCenter(val[0], val[1])
	}
	return r.getCenter()
}

func (r *Rectangle) CenterX(val ...int) int {
	if len(val) > 0 {
		r.setCenterX(val[0])
	}
	return r.getCenterX()
}

func (r *Rectangle) CenterY(val ...int) int {
	if len(val) > 0 {
		r.setCenterY(val[0])
	}
	return r.getCenterY()
}

func (r *Rectangle) Size(val ...int) (int, int) {
	if len(val) > 1 {
		r.setSize(val[0], val[1])
	}
	return r.getSize()
}

// Overlaps creates Overlaps method to check if two rectangles overlap
func (r *Rectangle) Overlaps(s *Rectangle) bool {
	return !r.Empty() && !s.Empty() &&
		r.Left() < s.Right() && s.Left() < r.Right() && r.Top() < s.Bottom() && s.Top() < r.Bottom()
}

// Empty reports whether the rectangle contains no points.
func (r *Rectangle) Empty() bool {
	return r.Left() >= r.Right() || r.Top() >= r.Bottom()
}
