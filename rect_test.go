package rect_test

import (
	r "github.com/drpaneas/rect"
	"testing"
)

func TestTop(t *testing.T) {
	// Test case 1: Check if the top value is returned correctly
	rect := r.Rect(10, 20, 30, 40)
	result := rect.Top()
	if result != 20 {
		t.Errorf("Top() returned incorrect value. Expected: 20, Got: %d", result)
	}

	// Test case 2: Check if top value is set correctly
	rect.Top(25)
	result = rect.Top()
	if result != 25 {
		t.Errorf("Top() returned incorrect value after modification. Expected: 25, Got: %d", result)
	}

	// Test case 3: getting and setting the top value of a Rectangle after moving it
	rect.Top(rect.Top() + 10)
	result = rect.Top()
	if result != 35 {
		t.Errorf("Expected top to be 35, but got %d", result)
	}
}

func TestLeft(t *testing.T) {
	// Test case 1: Check if the left value is returned correctly
	rect := r.Rect(10, 20, 30, 40)
	result := rect.Left()
	if result != 10 {
		t.Errorf("Left() returned incorrect value. Expected: 10, Got: %d", result)
	}

	// Test case 2: Check if left value is set correctly
	rect.Left(15)
	result = rect.Left()
	if result != 15 {
		t.Errorf("Left() returned incorrect value after modification. Expected: 15, Got: %d", result)
	}

	// Test case 3: getting and setting the left value of a Rectangle after moving it
	rect.Left(rect.Left() + 10)
	result = rect.Left()
	if result != 25 {
		t.Errorf("Expected left to be 25, but got %d", result)
	}
}

func TestBottom(t *testing.T) {
	// Test case 1: Check if the bottom value is returned correctly
	rect := r.Rect(10, 20, 30, 40)
	result := rect.Bottom()
	if result != 60 {
		t.Errorf("Bottom() returned incorrect value. Expected: 60, Got: %d", result)
	}

	// Test case 2: Check if bottom value is set correctly
	rect.Bottom(65)
	result = rect.Bottom()
	if result != 65 {
		t.Errorf("Bottom() returned incorrect value after modification. Expected: 65, Got: %d", result)
	}

	// Test case 3: getting and setting the bottom value of a Rectangle after moving it
	rect.Bottom(rect.Bottom() + 10)
	result = rect.Bottom()
	if result != 75 {
		t.Errorf("Expected bottom to be 75, but got %d", result)
	}
}

func TestRight(t *testing.T) {
	// Test case 1: Check if the right value is returned correctly
	rect := r.Rect(10, 20, 30, 40)
	result := rect.Right()
	if result != 40 {
		t.Errorf("Right() returned incorrect value. Expected: 40, Got: %d", result)
	}

	// Test case 2: Check if right value is set correctly
	rect.Right(45)
	result = rect.Right()
	if result != 45 {
		t.Errorf("Right() returned incorrect value after modification. Expected: 45, Got: %d", result)
	}

	// Test case 3: getting and setting the right value of a Rectangle after moving it
	rect.Right(rect.Right() + 10)
	result = rect.Right()
	if result != 55 {
		t.Errorf("Expected right to be 55, but got %d", result)
	}
}

func TestTopLeft(t *testing.T) {
	// Test case 1: Check if the top left point is returned correctly
	rect := r.Rect(10, 20, 30, 40)
	resultX, resultY := rect.TopLeft()
	if resultX != 10 || resultY != 20 {
		t.Errorf("TopLeft() returned incorrect value. Expected: (10, 20), Got: (%d, %d)", resultX, resultY)
	} else {
		t.Logf("TopLeft() returned correct value. Expected: (10, 20), Got: (%d, %d)", resultX, resultY)
	}

	// Test case 2: Check if top left point is set correctly
	rect.TopLeft(15, 25)
	resultX, resultY = rect.TopLeft()
	if resultX != 15 || resultY != 25 {
		t.Errorf("TopLeft() returned incorrect value after modification. Expected: (15, 25), Got: (%d, %d)", resultX, resultY)
	} else {
		t.Logf("TopLeft() returned correct value after modification. Expected: (15, 25), Got: (%d, %d)", resultX, resultY)
	}

	// Test case 3: getting and setting the top left point of a Rectangle after moving it
	rect.TopLeft(rect.X+10, rect.Y+10)
	resultX, resultY = rect.TopLeft()
	if resultX != 25 || resultY != 35 {
		t.Errorf("Expected top left to be (25, 35), but got (%d, %d)", resultX, resultY)
	} else {
		t.Logf("Expected top left to be (25, 35), but got (%d, %d)", resultX, resultY)
	}
}

func TestBottomLeft(t *testing.T) {
	// Test case 1: Check if the bottom left point is returned correctly
	rect := r.Rect(10, 20, 30, 40)
	resultX, resultY := rect.BottomLeft()
	if resultX != 10 || resultY != 60 {
		t.Errorf("BottomLeft() returned incorrect value. Expected: (10, 60), Got: (%d, %d)", resultX, resultY)
	} else {
		t.Logf("BottomLeft() returned correct value. Expected: (10, 60), Got: (%d, %d)", resultX, resultY)
	}

	// Test case 2: Check if bottom left point is set correctly
	rect.BottomLeft(15, 65)
	resultX, resultY = rect.BottomLeft()
	if resultX != 15 || resultY != 65 {
		t.Errorf("BottomLeft() returned incorrect value after modification. Expected: (15, 65), Got: (%d, %d)", resultX, resultY)
	} else {
		t.Logf("BottomLeft() returned correct value after modification. Expected: (15, 65), Got: (%d, %d)", resultX, resultY)
	}

	// Test case 3: getting and setting the bottom left point of a Rectangle after moving it
	rect.BottomLeft(rect.X+10, rect.Bottom()+10)
	resultX, resultY = rect.BottomLeft()
	if resultX != 25 || resultY != 75 {
		t.Errorf("Expected bottom left to be (25, 75), but got (%d, %d)", resultX, resultY)
	} else {
		t.Logf("Expected bottom left to be (25, 75), but got (%d, %d)", resultX, resultY)
	}
}

func TestTopRight(t *testing.T) {
	// Test case 1: Check if the top right point is returned correctly
	rect := r.Rect(10, 20, 30, 40)
	resultX, resultY := rect.TopRight()
	if resultX != 40 || resultY != 20 {
		t.Errorf("TopRight() returned incorrect value. Expected: (40, 20), Got: (%d, %d)", resultX, resultY)
	} else {
		t.Logf("TopRight() returned correct value. Expected: (40, 20), Got: (%d, %d)", resultX, resultY)
	}

	// Test case 2: Check if top right point is set correctly
	rect.TopRight(45, 25)
	resultX, resultY = rect.TopRight()
	if resultX != 45 || resultY != 25 {
		t.Errorf("TopRight() returned incorrect value after modification. Expected: (45, 25), Got: (%d, %d)", resultX, resultY)
	} else {
		t.Logf("TopRight() returned correct value after modification. Expected: (45, 25), Got: (%d, %d)", resultX, resultY)
	}

	// Test case 3: getting and setting the top right point of a Rectangle after moving it
	rect.TopRight(rect.Right()+10, rect.Y+10)
	resultX, resultY = rect.TopRight()
	if resultX != 55 || resultY != 35 {
		t.Errorf("Expected top right to be (55, 35), but got (%d, %d)", resultX, resultY)
	} else {
		t.Logf("Expected top right to be (55, 35), but got (%d, %d)", resultX, resultY)
	}
}

func TestBottomRight(t *testing.T) {
	// Test case 1: Check if the bottom right point is returned correctly
	rect := r.Rect(10, 20, 30, 40)
	resultX, resultY := rect.BottomRight()
	if resultX != 40 || resultY != 60 {
		t.Errorf("BottomRight() returned incorrect value. Expected: (40, 60), Got: (%d, %d)", resultX, resultY)
	} else {
		t.Logf("BottomRight() returned correct value. Expected: (40, 60), Got: (%d, %d)", resultX, resultY)
	}

	// Test case 2: Check if bottom right point is set correctly
	rect.BottomRight(45, 65)
	resultX, resultY = rect.BottomRight()
	if resultX != 45 || resultY != 65 {
		t.Errorf("BottomRight() returned incorrect value after modification. Expected: (45, 65), Got: (%d, %d)", resultX, resultY)
	} else {
		t.Logf("BottomRight() returned correct value after modification. Expected: (45, 65), Got: (%d, %d)", resultX, resultY)
	}

	// Test case 3: getting and setting the bottom right point of a Rectangle after moving it
	rect.BottomRight(rect.Right()+10, rect.Bottom()+10)
	resultX, resultY = rect.BottomRight()
	if resultX != 55 || resultY != 75 {
		t.Errorf("Expected bottom right to be (55, 75), but got (%d, %d)", resultX, resultY)
	} else {
		t.Logf("Expected bottom right to be (55, 75), but got (%d, %d)", resultX, resultY)
	}
}

func TestMidTop(t *testing.T) {
	// Test case 1: Check if the mid top point is returned correctly
	rect := r.Rect(10, 20, 30, 40)
	resultX, resultY := rect.MidTop()
	if resultX != 25 || resultY != 20 {
		t.Errorf("MidTop() returned incorrect value. Expected: (25, 20), Got: (%d, %d)", resultX, resultY)
	} else {
		t.Logf("MidTop() returned correct value. Expected: (25, 20), Got: (%d, %d)", resultX, resultY)
	}

	// Test case 2: Check if mid top point is set correctly
	rect.MidTop(30, 25)
	resultX, resultY = rect.MidTop()
	if resultX != 30 || resultY != 25 {
		t.Errorf("MidTop() returned incorrect value after modification. Expected: (30, 25), Got: (%d, %d)", resultX, resultY)
	} else {
		t.Logf("MidTop() returned correct value after modification. Expected: (30, 25), Got: (%d, %d)", resultX, resultY)
	}

	// Test case 3: getting and setting the mid top point of a Rectangle after moving it
	rect.MidTop(rect.CenterX()+10, rect.Y+10)
	resultX, resultY = rect.MidTop()
	if resultX != 40 || resultY != 35 {
		t.Errorf("Expected mid top to be (40, 35), but got (%d, %d)", resultX, resultY)
	} else {
		t.Logf("Expected mid top to be (40, 35), but got (%d, %d)", resultX, resultY)
	}
}

func TestMidLeft(t *testing.T) {
	// Test case 1: Check if the mid left point is returned correctly
	rect := r.Rect(10, 20, 30, 40)
	resultX, resultY := rect.MidLeft()
	if resultX != 10 || resultY != 40 {
		t.Errorf("MidLeft() returned incorrect value. Expected: (10, 40), Got: (%d, %d)", resultX, resultY)
	} else {
		t.Logf("MidLeft() returned correct value. Expected: (10, 40), Got: (%d, %d)", resultX, resultY)
	}

	// Test case 2: Check if mid left point is set correctly
	rect.MidLeft(15, 45)
	resultX, resultY = rect.MidLeft()
	if resultX != 15 || resultY != 45 {
		t.Errorf("MidLeft() returned incorrect value after modification. Expected: (15, 45), Got: (%d, %d)", resultX, resultY)
	} else {
		t.Logf("MidLeft() returned correct value after modification. Expected: (15, 45), Got: (%d, %d)", resultX, resultY)
	}

	// Test case 3: getting and setting the mid left point of a Rectangle after moving it
	rect.MidLeft(rect.X+10, rect.CenterY()+10)
	resultX, resultY = rect.MidLeft()
	if resultX != 25 || resultY != 55 {
		t.Errorf("Expected mid left to be (25, 55), but got (%d, %d)", resultX, resultY)
	} else {
		t.Logf("Expected mid left to be (25, 55), but got (%d, %d)", resultX, resultY)
	}
}

func TestMidBottom(t *testing.T) {
	// Test case 1: Check if the mid bottom point is returned correctly
	rect := r.Rect(10, 20, 30, 40)
	resultX, resultY := rect.MidBottom()
	if resultX != 25 || resultY != 60 {
		t.Errorf("MidBottom() returned incorrect value. Expected: (25, 60), Got: (%d, %d)", resultX, resultY)
	} else {
		t.Logf("MidBottom() returned correct value. Expected: (25, 60), Got: (%d, %d)", resultX, resultY)
	}

	// Test case 2: Check if mid bottom point is set correctly
	rect.MidBottom(30, 65)
	resultX, resultY = rect.MidBottom()
	if resultX != 30 || resultY != 65 {
		t.Errorf("MidBottom() returned incorrect value after modification. Expected: (30, 65), Got: (%d, %d)", resultX, resultY)
	} else {
		t.Logf("MidBottom() returned correct value after modification. Expected: (30, 65), Got: (%d, %d)", resultX, resultY)
	}

	// Test case 3: getting and setting the mid bottom point of a Rectangle after moving it
	rect.MidBottom(rect.CenterX()+10, rect.Bottom()+10)
	resultX, resultY = rect.MidBottom()
	if resultX != 40 || resultY != 75 {
		t.Errorf("Expected mid bottom to be (40, 75), but got (%d, %d)", resultX, resultY)
	} else {
		t.Logf("Expected mid bottom to be (40, 75), but got (%d, %d)", resultX, resultY)
	}
}

func TestMidRight(t *testing.T) {
	// Test case 1: Check if the mid right point is returned correctly
	rect := r.Rect(10, 20, 30, 40)
	resultX, resultY := rect.MidRight()
	if resultX != 40 || resultY != 40 {
		t.Errorf("MidRight() returned incorrect value. Expected: (40, 40), Got: (%d, %d)", resultX, resultY)
	} else {
		t.Logf("MidRight() returned correct value. Expected: (40, 40), Got: (%d, %d)", resultX, resultY)
	}

	// Test case 2: Check if mid right point is set correctly
	rect.MidRight(45, 45)
	resultX, resultY = rect.MidRight()
	if resultX != 45 || resultY != 45 {
		t.Errorf("MidRight() returned incorrect value after modification. Expected: (45, 45), Got: (%d, %d)", resultX, resultY)
	} else {
		t.Logf("MidRight() returned correct value after modification. Expected: (45, 45), Got: (%d, %d)", resultX, resultY)
	}

	// Test case 3: getting and setting the mid right point of a Rectangle after moving it
	rect.MidRight(rect.Right()+10, rect.CenterY()+10)
	resultX, resultY = rect.MidRight()
	if resultX != 55 || resultY != 55 {
		t.Errorf("Expected mid right to be (55, 55), but got (%d, %d)", resultX, resultY)
	} else {
		t.Logf("Expected mid right to be (55, 55), but got (%d, %d)", resultX, resultY)
	}
}

func TestCenter(t *testing.T) {
	// Test case 1: Check if the center point is returned correctly
	rect := r.Rect(10, 20, 30, 40)
	resultX, resultY := rect.Center()
	if resultX != 25 || resultY != 40 {
		t.Errorf("Center() returned incorrect value. Expected: (25, 40), Got: (%d, %d)", resultX, resultY)
	} else {
		t.Logf("Center() returned correct value. Expected: (25, 40), Got: (%d, %d)", resultX, resultY)
	}

	// Test case 2: Check if center point is set correctly
	rect.Center(30, 45)
	resultX, resultY = rect.Center()
	if resultX != 30 || resultY != 45 {
		t.Errorf("Center() returned incorrect value after modification. Expected: (30, 45), Got: (%d, %d)", resultX, resultY)
	} else {
		t.Logf("Center() returned correct value after modification. Expected: (30, 45), Got: (%d, %d)", resultX, resultY)
	}

	// Test case 3: getting and setting the center point of a Rectangle after moving it
	rect.Center(rect.CenterX()+10, rect.CenterY()+10)
	resultX, resultY = rect.Center()
	if resultX != 40 || resultY != 55 {
		t.Errorf("Expected center to be (40, 55), but got (%d, %d)", resultX, resultY)
	} else {
		t.Logf("Expected center to be (40, 55), but got (%d, %d)", resultX, resultY)
	}
}

func TestCenterX(t *testing.T) {
	// Test case 1: Check if the center X is returned correctly
	rect := r.Rect(10, 20, 30, 40)
	result := rect.CenterX()
	if result != 25 {
		t.Errorf("CenterX() returned incorrect value. Expected: 25, Got: %d", result)
	} else {
		t.Logf("CenterX() returned correct value. Expected: 25, Got: %d", result)
	}

	// Test case 2: Check if center X is set correctly
	rect.CenterX(30)
	result = rect.CenterX()
	if result != 30 {
		t.Errorf("CenterX() returned incorrect value after modification. Expected: 30, Got: %d", result)
	} else {
		t.Logf("CenterX() returned correct value after modification. Expected: 30, Got: %d", result)
	}

	// Test case 3: getting and setting the center X of a Rectangle after moving it
	rect.CenterX(rect.CenterX() + 10)
	result = rect.CenterX()
	if result != 40 {
		t.Errorf("Expected center X to be 40, but got %d", result)
	} else {
		t.Logf("Expected center X to be 40, but got %d", result)
	}
}

func TestCenterY(t *testing.T) {
	// Test case 1: Check if the center Y is returned correctly
	rect := r.Rect(10, 20, 30, 40)

	result := rect.CenterY()
	if result != 40 {
		t.Errorf("CenterY() returned incorrect value. Expected: 40, Got: %d", result)
	} else {
		t.Logf("CenterY() returned correct value. Expected: 40, Got: %d", result)
	}

	// Test case 2: Check if center Y is set correctly
	rect.CenterY(45)
	result = rect.CenterY()
	if result != 45 {
		t.Errorf("CenterY() returned incorrect value after modification. Expected: 45, Got: %d", result)
	} else {
		t.Logf("CenterY() returned correct value after modification. Expected: 45, Got: %d", result)
	}

	// Test case 3: getting and setting the center Y of a Rectangle after moving it
	rect.CenterY(rect.CenterY() + 10)
	result = rect.CenterY()
	if result != 55 {
		t.Errorf("Expected center Y to be 55, but got %d", result)
	} else {
		t.Logf("Expected center Y to be 55, but got %d", result)
	}
}

func TestSize(t *testing.T) {
	// Test case 1: Check if the size is returned correctly
	rect := r.Rect(10, 20, 30, 40)

	resultWidth, resultHeight := rect.Size()
	if resultWidth != 30 || resultHeight != 40 {
		t.Errorf("Size() returned incorrect value. Expected: (30, 40), Got: (%d, %d)", resultWidth, resultHeight)
	} else {
		t.Logf("Size() returned correct value. Expected: (30, 40), Got: (%d, %d)", resultWidth, resultHeight)
	}

	// Test case 2: Check if size is set correctly
	rect.Size(35, 45)
	resultWidth, resultHeight = rect.Size()
	if resultWidth != 35 || resultHeight != 45 {
		t.Errorf("Size() returned incorrect value after modification. Expected: (35, 45), Got: (%d, %d)", resultWidth, resultHeight)
	} else {
		t.Logf("Size() returned correct value after modification. Expected: (35, 45), Got: (%d, %d)", resultWidth, resultHeight)
	}

	// Test case 3: getting and setting the size of a Rectangle after moving it
	rect.Size(rect.Width+10, rect.Height+10)
	resultWidth, resultHeight = rect.Size()
	if resultWidth != 45 || resultHeight != 55 {
		t.Errorf("Expected size to be (45, 55), but got (%d, %d)", resultWidth, resultHeight)
	} else {
		t.Logf("Expected size to be (45, 55), but got (%d, %d)", resultWidth, resultHeight)
	}
}

func TestOverlaps(t *testing.T) {
	// Test case 1: Check if the rectangles overlap
	rect1 := r.Rect(10, 20, 30, 40)
	rect2 := r.Rect(20, 30, 40, 50)

	result := rect1.CollidesWith(rect2)
	if !result {
		t.Errorf("CollidesWith() returned incorrect value. Expected: true, Got: %t", result)
	} else {
		t.Logf("CollidesWith() returned correct value. Expected: true, Got: %t", result)
	}

	// Test case 2: Check if the rectangles do not overlap
	rect1 = r.Rect(10, 20, 30, 40)
	rect2 = r.Rect(50, 60, 70, 80)

	result = rect1.CollidesWith(rect2)
	if result {
		t.Errorf("CollidesWith() returned incorrect value. Expected: false, Got: %t", result)
	} else {
		t.Logf("CollidesWith() returned correct value. Expected: false, Got: %t", result)
	}
}

func TestEmpty(t *testing.T) {
	// Test case 1: Check if the rectangle is empty
	rect := r.Rect(0, 0, 0, 0)

	result := rect.Empty()
	if !result {
		t.Errorf("Empty() returned incorrect value. Expected: true, Got: %t", result)
	} else {
		t.Logf("Empty() returned correct value. Expected: true, Got: %t", result)
	}

	// Test case 2: Check if the rectangle is not empty
	rect = r.Rect(10, 20, 30, 40)

	result = rect.Empty()
	if result {
		t.Errorf("Empty() returned incorrect value. Expected: false, Got: %t", result)
	} else {
		t.Logf("Empty() returned correct value. Expected: false, Got: %t", result)
	}
}
