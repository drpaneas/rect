package rect

import "testing"

func TestGetTop(t *testing.T) {
	testCases := []struct {
		name      string
		rect      rectangle
		expectedY int
	}{
		{
			name:      "Initial Y value",
			rect:      rectangle{X: 10, Y: 20, Width: 30, Height: 40},
			expectedY: 20,
		},
		{
			name:      "Modified Y value",
			rect:      rectangle{X: 10, Y: 25, Width: 30, Height: 40},
			expectedY: 25,
		},
		{
			name:      "Moved Y value",
			rect:      rectangle{X: 10, Y: 35, Width: 30, Height: 40},
			expectedY: 35,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			y := testCase.rect.getTop()
			if y != testCase.expectedY {
				t.Errorf("getTop() returned incorrect value. Expected: %d, Got: %d", testCase.expectedY, y)
			}
		})
	}
}

func TestGetLeft(t *testing.T) {
	testCases := []struct {
		name      string
		rect      rectangle
		expectedX int
	}{
		{
			name:      "Initial X value",
			rect:      rectangle{X: 10, Y: 20, Width: 30, Height: 40},
			expectedX: 10,
		},
		{
			name:      "Modified X value",
			rect:      rectangle{X: 15, Y: 20, Width: 30, Height: 40},
			expectedX: 15,
		},
		{
			name:      "Moved X value",
			rect:      rectangle{X: 25, Y: 20, Width: 30, Height: 40},
			expectedX: 25,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			x := testCase.rect.getLeft()
			if x != testCase.expectedX {
				t.Errorf("getLeft() returned incorrect value. Expected: %d, Got: %d", testCase.expectedX, x)
			}
		})
	}
}

func TestGetWidth(t *testing.T) {
	testCases := []struct {
		name          string
		rect          rectangle
		expectedWidth int
	}{
		{
			name:          "Initial Width value",
			rect:          rectangle{X: 10, Y: 20, Width: 30, Height: 40},
			expectedWidth: 30,
		},
		{
			name:          "Modified Width value",
			rect:          rectangle{X: 10, Y: 20, Width: 35, Height: 40},
			expectedWidth: 35,
		},
		{
			name:          "Moved Width value",
			rect:          rectangle{X: 10, Y: 20, Width: 45, Height: 40},
			expectedWidth: 45,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			width := testCase.rect.getWidth()
			if width != testCase.expectedWidth {
				t.Errorf("getWidth() returned incorrect value. Expected: %d, Got: %d", testCase.expectedWidth, width)
			}
		})
	}
}

func TestGetRight(t *testing.T) {
	testCases := []struct {
		name      string
		rect      rectangle
		expectedX int
	}{
		{
			name:      "Initial X value",
			rect:      rectangle{X: 10, Y: 20, Width: 30, Height: 40},
			expectedX: 40,
		},
		{
			name:      "Modified X value",
			rect:      rectangle{X: 15, Y: 20, Width: 30, Height: 40},
			expectedX: 45,
		},
		{
			name:      "Moved X value",
			rect:      rectangle{X: 25, Y: 20, Width: 30, Height: 40},
			expectedX: 55,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			x := testCase.rect.getRight()
			if x != testCase.expectedX {
				t.Errorf("getRight() returned incorrect value. Expected: %d, Got: %d", testCase.expectedX, x)
			}
		})
	}
}

func TestGetBottom(t *testing.T) {
	testCases := []struct {
		name      string
		rect      rectangle
		expectedY int
	}{
		{
			name:      "Initial Y value",
			rect:      rectangle{X: 10, Y: 20, Width: 30, Height: 40},
			expectedY: 60,
		},
		{
			name:      "Modified Y value",
			rect:      rectangle{X: 10, Y: 25, Width: 30, Height: 40},
			expectedY: 65,
		},
		{
			name:      "Moved Y value",
			rect:      rectangle{X: 10, Y: 35, Width: 30, Height: 40},
			expectedY: 75,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			y := testCase.rect.getBottom()
			if y != testCase.expectedY {
				t.Errorf("getBottom() returned incorrect value. Expected: %d, Got: %d", testCase.expectedY, y)
			}
		})
	}
}

func TestGetTopLeft(t *testing.T) {
	testCases := []struct {
		name      string
		rect      rectangle
		expectedX int
		expectedY int
	}{
		{
			name:      "Initial values",
			rect:      rectangle{X: 10, Y: 20, Width: 30, Height: 40},
			expectedX: 10,
			expectedY: 20,
		},
		{
			name:      "Modified values",
			rect:      rectangle{X: 15, Y: 25, Width: 30, Height: 40},
			expectedX: 15,
			expectedY: 25,
		},
		{
			name:      "Moved values",
			rect:      rectangle{X: 25, Y: 35, Width: 30, Height: 40},
			expectedX: 25,
			expectedY: 35,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			resultX, resultY := testCase.rect.getTopLeft()
			if resultX != testCase.expectedX {
				t.Errorf("getTopLeft() returned incorrect value. Expected: %v, Got: %v", testCase.expectedX, resultX)
			}
			if resultY != testCase.expectedY {
				t.Errorf("getTopLeft() returned incorrect value. Expected: %v, Got: %v", testCase.expectedY, resultY)
			}
		})
	}
}

func TestGetTopRight(t *testing.T) {
	testCases := []struct {
		name      string
		rect      rectangle
		expectedX int
		expectedY int
	}{
		{
			name:      "Initial values",
			rect:      rectangle{X: 10, Y: 20, Width: 30, Height: 40},
			expectedX: 40,
			expectedY: 20,
		},
		{
			name:      "Modified values",
			rect:      rectangle{X: 15, Y: 25, Width: 30, Height: 40},
			expectedX: 45,
			expectedY: 25,
		},
		{
			name:      "Moved values",
			rect:      rectangle{X: 25, Y: 35, Width: 30, Height: 40},
			expectedX: 55,
			expectedY: 35,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			resultX, resultY := testCase.rect.getTopRight()
			if resultX != testCase.expectedX {
				t.Errorf("getTopRight() returned incorrect value. Expected: %v, Got: %v", testCase.expectedX, resultX)
			}
			if resultY != testCase.expectedY {
				t.Errorf("getTopRight() returned incorrect value. Expected: %v, Got: %v", testCase.expectedY, resultY)
			}
		})
	}
}

func TestGetBottomLeft(t *testing.T) {
	testCases := []struct {
		name      string
		rect      rectangle
		expectedX int
		expectedY int
	}{
		{
			name:      "Initial values",
			rect:      rectangle{X: 10, Y: 20, Width: 30, Height: 40},
			expectedX: 10,
			expectedY: 60,
		},
		{
			name:      "Modified values",
			rect:      rectangle{X: 15, Y: 25, Width: 30, Height: 40},
			expectedX: 15,
			expectedY: 65,
		},
		{
			name:      "Moved values",
			rect:      rectangle{X: 25, Y: 35, Width: 30, Height: 40},
			expectedX: 25,
			expectedY: 75,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			resultX, resultY := testCase.rect.getBottomLeft()
			if resultX != testCase.expectedX {
				t.Errorf("getBottomLeft() returned incorrect value. Expected: %v, Got: %v", testCase.expectedX, resultX)
			}
			if resultY != testCase.expectedY {
				t.Errorf("getBottomLeft() returned incorrect value. Expected: %v, Got: %v", testCase.expectedY, resultY)
			}
		})
	}
}

func TestGetBottomRight(t *testing.T) {
	testCases := []struct {
		name      string
		rect      rectangle
		expectedX int
		expectedY int
	}{
		{
			name:      "Initial values",
			rect:      rectangle{X: 10, Y: 20, Width: 30, Height: 40},
			expectedX: 40,
			expectedY: 60,
		},
		{
			name:      "Modified values",
			rect:      rectangle{X: 15, Y: 25, Width: 30, Height: 40},
			expectedX: 45,
			expectedY: 65,
		},
		{
			name:      "Moved values",
			rect:      rectangle{X: 25, Y: 35, Width: 30, Height: 40},
			expectedX: 55,
			expectedY: 75,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			resultX, resultY := testCase.rect.getBottomRight()
			if resultX != testCase.expectedX {
				t.Errorf("getBottomRight() returned incorrect value. Expected: %v, Got: %v", testCase.expectedX, resultX)
			}
			if resultY != testCase.expectedY {
				t.Errorf("getBottomRight() returned incorrect value. Expected: %v, Got: %v", testCase.expectedY, resultY)
			}
		})
	}
}

func TestGetCenter(t *testing.T) {
	testCases := []struct {
		name      string
		rect      rectangle
		expectedX int
		expectedY int
	}{
		{
			name:      "Initial values",
			rect:      rectangle{X: 10, Y: 20, Width: 30, Height: 40},
			expectedX: 25,
			expectedY: 40,
		},
		{
			name:      "Modified values",
			rect:      rectangle{X: 15, Y: 25, Width: 30, Height: 40},
			expectedX: 30,
			expectedY: 45,
		},
		{
			name:      "Moved values",
			rect:      rectangle{X: 25, Y: 35, Width: 30, Height: 40},
			expectedX: 40,
			expectedY: 55,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			resultX, resultY := testCase.rect.getCenter()
			if resultX != testCase.expectedX {
				t.Errorf("getCenter() returned incorrect value. Expected: %v, Got: %v", testCase.expectedX, resultX)
			}
			if resultY != testCase.expectedY {
				t.Errorf("getCenter() returned incorrect value. Expected: %v, Got: %v", testCase.expectedY, resultY)
			}
		})
	}
}

func TestGetMidTop(t *testing.T) {
	testCases := []struct {
		name      string
		rect      rectangle
		expectedX int
		expectedY int
	}{
		{
			name:      "Initial values",
			rect:      rectangle{X: 10, Y: 20, Width: 30, Height: 40},
			expectedX: 25,
			expectedY: 20,
		},
		{
			name:      "Modified values",
			rect:      rectangle{X: 15, Y: 25, Width: 30, Height: 40},
			expectedX: 30,
			expectedY: 25,
		},
		{
			name:      "Moved values",
			rect:      rectangle{X: 25, Y: 35, Width: 30, Height: 40},
			expectedX: 40,
			expectedY: 35,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			resultX, resultY := testCase.rect.getMidTop()
			if resultX != testCase.expectedX {
				t.Errorf("getMidTop() returned incorrect value. Expected: %v, Got: %v", testCase.expectedX, resultX)
			}
			if resultY != testCase.expectedY {
				t.Errorf("getMidTop() returned incorrect value. Expected: %v, Got: %v", testCase.expectedY, resultY)
			}
		})
	}
}

func TestGetMidBottom(t *testing.T) {
	testCases := []struct {
		name      string
		rect      rectangle
		expectedX int
		expectedY int
	}{
		{
			name:      "Initial values",
			rect:      rectangle{X: 10, Y: 20, Width: 30, Height: 40},
			expectedX: 25,
			expectedY: 60,
		},
		{
			name:      "Modified values",
			rect:      rectangle{X: 15, Y: 25, Width: 30, Height: 40},
			expectedX: 30,
			expectedY: 65,
		},
		{
			name:      "Moved values",
			rect:      rectangle{X: 25, Y: 35, Width: 30, Height: 40},
			expectedX: 40,
			expectedY: 75,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			resultX, resultY := testCase.rect.getMidBottom()
			if resultX != testCase.expectedX {
				t.Errorf("getMidBottom() returned incorrect value. Expected: %v, Got: %v", testCase.expectedX, resultX)
			}
			if resultY != testCase.expectedY {
				t.Errorf("getMidBottom() returned incorrect value. Expected: %v, Got: %v", testCase.expectedY, resultY)
			}
		})
	}
}

func TestGetMidLeft(t *testing.T) {
	testCases := []struct {
		name      string
		rect      rectangle
		expectedX int
		expectedY int
	}{
		{
			name:      "Initial values",
			rect:      rectangle{X: 10, Y: 20, Width: 30, Height: 40},
			expectedX: 10,
			expectedY: 40,
		},
		{
			name:      "Modified values",
			rect:      rectangle{X: 15, Y: 25, Width: 30, Height: 40},
			expectedX: 15,
			expectedY: 45,
		},
		{
			name:      "Moved values",
			rect:      rectangle{X: 25, Y: 35, Width: 30, Height: 40},
			expectedX: 25,
			expectedY: 55,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			resultX, resultY := testCase.rect.getMidLeft()
			if resultX != testCase.expectedX {
				t.Errorf("getMidLeft() returned incorrect value. Expected: %v, Got: %v", testCase.expectedX, resultX)
			}
			if resultY != testCase.expectedY {
				t.Errorf("getMidLeft() returned incorrect value. Expected: %v, Got: %v", testCase.expectedY, resultY)
			}
		})
	}
}

func TestGetMidRight(t *testing.T) {
	testCases := []struct {
		name      string
		rect      rectangle
		expectedX int
		expectedY int
	}{
		{
			name:      "Initial values",
			rect:      rectangle{X: 10, Y: 20, Width: 30, Height: 40},
			expectedX: 40,
			expectedY: 40,
		},
		{
			name:      "Modified values",
			rect:      rectangle{X: 15, Y: 25, Width: 30, Height: 40},
			expectedX: 45,
			expectedY: 45,
		},
		{
			name:      "Moved values",
			rect:      rectangle{X: 25, Y: 35, Width: 30, Height: 40},
			expectedX: 55,
			expectedY: 55,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			resultX, resultY := testCase.rect.getMidRight()
			if resultX != testCase.expectedX {
				t.Errorf("getMidRight() returned incorrect value. Expected: %v, Got: %v", testCase.expectedX, resultX)
			}
			if resultY != testCase.expectedY {
				t.Errorf("getMidRight() returned incorrect value. Expected: %v, Got: %v", testCase.expectedY, resultY)
			}
		})
	}
}

func TestGetCenterX(t *testing.T) {
	testCases := []struct {
		name      string
		rect      rectangle
		expectedX int
	}{
		{
			name:      "Initial values",
			rect:      rectangle{X: 10, Y: 20, Width: 30, Height: 40},
			expectedX: 25,
		},
		{
			name:      "Modified values",
			rect:      rectangle{X: 15, Y: 25, Width: 30, Height: 40},
			expectedX: 30,
		},
		{
			name:      "Moved values",
			rect:      rectangle{X: 25, Y: 35, Width: 30, Height: 40},
			expectedX: 40,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			resultX := testCase.rect.getCenterX()
			if resultX != testCase.expectedX {
				t.Errorf("getCenterX() returned incorrect value. Expected: %v, Got: %v", testCase.expectedX, resultX)
			}
		})
	}
}

func TestGetCenterY(t *testing.T) {
	testCases := []struct {
		name      string
		rect      rectangle
		expectedY int
	}{
		{
			name:      "Initial values",
			rect:      rectangle{X: 10, Y: 20, Width: 30, Height: 40},
			expectedY: 40,
		},
		{
			name:      "Modified values",
			rect:      rectangle{X: 15, Y: 25, Width: 30, Height: 40},
			expectedY: 45,
		},
		{
			name:      "Moved values",
			rect:      rectangle{X: 25, Y: 35, Width: 30, Height: 40},
			expectedY: 55,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			resultY := testCase.rect.getCenterY()
			if resultY != testCase.expectedY {
				t.Errorf("getCenterY() returned incorrect value. Expected: %v, Got: %v", testCase.expectedY, resultY)
			}
		})
	}
}

func TestGetSize(t *testing.T) {
	testCases := []struct {
		name      string
		rect      rectangle
		expectedW int
		expectedH int
	}{
		{
			name:      "Initial values",
			rect:      rectangle{X: 10, Y: 20, Width: 30, Height: 40},
			expectedW: 30,
			expectedH: 40,
		},
		{
			name:      "Modified values",
			rect:      rectangle{X: 15, Y: 25, Width: 30, Height: 40},
			expectedW: 30,
			expectedH: 40,
		},
		{
			name:      "Moved values",
			rect:      rectangle{X: 25, Y: 35, Width: 30, Height: 40},
			expectedW: 30,
			expectedH: 40,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			resultW, resultH := testCase.rect.getSize()
			if resultW != testCase.expectedW {
				t.Errorf("getSize() returned incorrect value. Expected: %v, Got: %v", testCase.expectedW, resultW)
			}
			if resultH != testCase.expectedH {
				t.Errorf("getSize() returned incorrect value. Expected: %v, Got: %v", testCase.expectedH, resultH)
			}
		})
	}
}

func TestSetTopLeft(t *testing.T) {
	testCases := []struct {
		name     string
		rect     rectangle
		x        int
		y        int
		expected rectangle
	}{
		{
			name:     "Initial values",
			rect:     rectangle{X: 10, Y: 20, Width: 30, Height: 40},
			x:        15,
			y:        25,
			expected: rectangle{X: 15, Y: 25, Width: 30, Height: 40},
		},
		{
			name:     "Modified values",
			rect:     rectangle{X: 15, Y: 25, Width: 30, Height: 40},
			x:        20,
			y:        30,
			expected: rectangle{X: 20, Y: 30, Width: 30, Height: 40},
		},
		{
			name:     "Moved values",
			rect:     rectangle{X: 25, Y: 35, Width: 30, Height: 40},
			x:        30,
			y:        40,
			expected: rectangle{X: 30, Y: 40, Width: 30, Height: 40},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.rect.setTopLeft(testCase.x, testCase.y)
			if testCase.rect != testCase.expected {
				t.Errorf("setTopLeft() returned incorrect value. Expected: %v, Got: %v", testCase.expected, testCase.rect)
			}
		})
	}
}

func TestSetTopRight(t *testing.T) {
	testCases := []struct {
		name     string
		rect     rectangle
		x        int
		y        int
		expected rectangle
	}{
		{
			name:     "Move values",
			rect:     rectangle{X: 10, Y: 20, Width: 30, Height: 40},
			x:        35,
			y:        25,
			expected: rectangle{X: 5, Y: 25, Width: 30, Height: 40},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.rect.setTopRight(testCase.x, testCase.y)
			if testCase.rect != testCase.expected {
				t.Errorf("setTopRight() returned incorrect value. Expected: %v, Got: %v", testCase.expected, testCase.rect)
			}
		})
	}
}

func TestSetBottomLeft(t *testing.T) {
	testCases := []struct {
		name     string
		rect     rectangle
		x        int
		y        int
		expected rectangle
	}{
		{
			name:     "Initial values",
			rect:     rectangle{X: 10, Y: 20, Width: 30, Height: 40},
			x:        15,
			y:        55,
			expected: rectangle{X: 15, Y: 15, Width: 30, Height: 40},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.rect.setBottomLeft(testCase.x, testCase.y)
			if testCase.rect != testCase.expected {
				t.Errorf("setBottomLeft() returned incorrect value. Expected: %v, Got: %v", testCase.expected, testCase.rect)
			}
		})
	}
}

func TestSetBottomRight(t *testing.T) {
	testCases := []struct {
		name     string
		rect     rectangle
		x        int
		y        int
		expected rectangle
	}{
		{
			name:     "Initial values",
			rect:     rectangle{X: 10, Y: 20, Width: 30, Height: 40},
			x:        35,
			y:        55,
			expected: rectangle{X: 5, Y: 15, Width: 30, Height: 40},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.rect.setBottomRight(testCase.x, testCase.y)
			if testCase.rect != testCase.expected {
				t.Errorf("setBottomRight() returned incorrect value. Expected: %v, Got: %v", testCase.expected, testCase.rect)
			}
		})
	}
}

func TestSetCenter(t *testing.T) {
	testCases := []struct {
		name     string
		rect     rectangle
		x        int
		y        int
		expected rectangle
	}{
		{
			name:     "Initial values",
			rect:     rectangle{X: 10, Y: 20, Width: 30, Height: 40},
			x:        25,
			y:        40,
			expected: rectangle{X: 10, Y: 20, Width: 30, Height: 40},
		},
		{
			name:     "Modified values",
			rect:     rectangle{X: 15, Y: 25, Width: 30, Height: 40},
			x:        30,
			y:        45,
			expected: rectangle{X: 15, Y: 25, Width: 30, Height: 40},
		},
		{
			name:     "Moved values",
			rect:     rectangle{X: 25, Y: 35, Width: 30, Height: 40},
			x:        40,
			y:        55,
			expected: rectangle{X: 25, Y: 35, Width: 30, Height: 40},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.rect.setCenter(testCase.x, testCase.y)
			if testCase.rect != testCase.expected {
				t.Errorf("setCenter() returned incorrect value. Expected: %v, Got: %v", testCase.expected, testCase.rect)
			}
		})
	}
}

func TestSetCenterX(t *testing.T) {
	testCases := []struct {
		name     string
		rect     rectangle
		x        int
		expected rectangle
	}{
		{
			name:     "Initial values",
			rect:     rectangle{X: 10, Y: 20, Width: 30, Height: 40},
			x:        25,
			expected: rectangle{X: 10, Y: 20, Width: 30, Height: 40},
		},
		{
			name:     "Modified values",
			rect:     rectangle{X: 15, Y: 25, Width: 30, Height: 40},
			x:        30,
			expected: rectangle{X: 15, Y: 25, Width: 30, Height: 40},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.rect.setCenterX(testCase.x)
			if testCase.rect != testCase.expected {
				t.Errorf("setCenterX() returned incorrect value. Expected: %v, Got: %v", testCase.expected, testCase.rect)
			}
		})
	}
}

func TestSetCenterY(t *testing.T) {
	testCases := []struct {
		name     string
		rect     rectangle
		y        int
		expected rectangle
	}{
		{
			name:     "Initial values",
			rect:     rectangle{X: 10, Y: 20, Width: 30, Height: 40},
			y:        40,
			expected: rectangle{X: 10, Y: 20, Width: 30, Height: 40},
		},
		{
			name:     "Modified values",
			rect:     rectangle{X: 15, Y: 25, Width: 30, Height: 40},
			y:        45,
			expected: rectangle{X: 15, Y: 25, Width: 30, Height: 40},
		},
		{
			name:     "Moved values",
			rect:     rectangle{X: 25, Y: 35, Width: 30, Height: 40},
			y:        55,
			expected: rectangle{X: 25, Y: 35, Width: 30, Height: 40},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.rect.setCenterY(testCase.y)
			if testCase.rect != testCase.expected {
				t.Errorf("setCenterY() returned incorrect value. Expected: %v, Got: %v", testCase.expected, testCase.rect)
			}
		})
	}
}

func TestSetHeight(t *testing.T) {
	testCases := []struct {
		name     string
		rect     rectangle
		height   int
		expected rectangle
	}{
		{
			name:     "Initial values",
			rect:     rectangle{X: 10, Y: 20, Width: 30, Height: 40},
			height:   50,
			expected: rectangle{X: 10, Y: 20, Width: 30, Height: 50},
		},
		{
			name:     "Modified values",
			rect:     rectangle{X: 15, Y: 25, Width: 30, Height: 40},
			height:   60,
			expected: rectangle{X: 15, Y: 25, Width: 30, Height: 60},
		},
		{
			name:     "Moved values",
			rect:     rectangle{X: 25, Y: 35, Width: 30, Height: 40},
			height:   70,
			expected: rectangle{X: 25, Y: 35, Width: 30, Height: 70},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.rect.setHeight(testCase.height)
			if testCase.rect != testCase.expected {
				t.Errorf("setHeight() returned incorrect value. Expected: %v, Got: %v", testCase.expected, testCase.rect)
			}
		})
	}
}

func TestSetX(t *testing.T) {
	testCases := []struct {
		name     string
		rect     rectangle
		x        int
		expected rectangle
	}{
		{
			name:     "Initial values",
			rect:     rectangle{X: 10, Y: 20, Width: 30, Height: 40},
			x:        50,
			expected: rectangle{X: 50, Y: 20, Width: 30, Height: 40},
		},
		{
			name:     "Modified values",
			rect:     rectangle{X: 15, Y: 25, Width: 30, Height: 40},
			x:        60,
			expected: rectangle{X: 60, Y: 25, Width: 30, Height: 40},
		},
		{
			name:     "Moved values",
			rect:     rectangle{X: 25, Y: 35, Width: 30, Height: 40},
			x:        70,
			expected: rectangle{X: 70, Y: 35, Width: 30, Height: 40},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.rect.setX(testCase.x)
			if testCase.rect != testCase.expected {
				t.Errorf("setX() returned incorrect value. Expected: %v, Got: %v", testCase.expected, testCase.rect)
			}
		})
	}
}

func TestSetY(t *testing.T) {
	testCases := []struct {
		name     string
		rect     rectangle
		y        int
		expected rectangle
	}{
		{
			name:     "Initial values",
			rect:     rectangle{X: 10, Y: 20, Width: 30, Height: 40},
			y:        50,
			expected: rectangle{X: 10, Y: 50, Width: 30, Height: 40},
		},
		{
			name:     "Modified values",
			rect:     rectangle{X: 15, Y: 25, Width: 30, Height: 40},
			y:        60,
			expected: rectangle{X: 15, Y: 60, Width: 30, Height: 40},
		},
		{
			name:     "Moved values",
			rect:     rectangle{X: 25, Y: 35, Width: 30, Height: 40},
			y:        70,
			expected: rectangle{X: 25, Y: 70, Width: 30, Height: 40},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.rect.setY(testCase.y)
			if testCase.rect != testCase.expected {
				t.Errorf("setY() returned incorrect value. Expected: %v, Got: %v", testCase.expected, testCase.rect)
			}
		})
	}
}

func TestSetTop(t *testing.T) {
	testCases := []struct {
		name     string
		rect     rectangle
		top      int
		expected rectangle
	}{
		{
			name:     "Initial values",
			rect:     rectangle{X: 10, Y: 20, Width: 30, Height: 40},
			top:      50,
			expected: rectangle{X: 10, Y: 50, Width: 30, Height: 40},
		},
		{
			name:     "Modified values",
			rect:     rectangle{X: 15, Y: 25, Width: 30, Height: 40},
			top:      60,
			expected: rectangle{X: 15, Y: 60, Width: 30, Height: 40},
		},
		{
			name:     "Moved values",
			rect:     rectangle{X: 25, Y: 35, Width: 30, Height: 40},
			top:      70,
			expected: rectangle{X: 25, Y: 70, Width: 30, Height: 40},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.rect.setTop(testCase.top)
			if testCase.rect != testCase.expected {
				t.Errorf("setTop() returned incorrect value. Expected: %v, Got: %v", testCase.expected, testCase.rect)
			}
		})
	}
}

func TestSetBottom(t *testing.T) {
	testCases := []struct {
		name     string
		rect     rectangle
		bottom   int
		expected rectangle
	}{
		{
			name:     "Initial values",
			rect:     rectangle{X: 10, Y: 20, Width: 30, Height: 40},
			bottom:   50,
			expected: rectangle{X: 10, Y: 10, Width: 30, Height: 40},
		},
		{
			name:     "Modified values",
			rect:     rectangle{X: 15, Y: 25, Width: 30, Height: 40},
			bottom:   60,
			expected: rectangle{X: 15, Y: 20, Width: 30, Height: 40},
		},
		{
			name:     "Moved values",
			rect:     rectangle{X: 25, Y: 35, Width: 30, Height: 40},
			bottom:   70,
			expected: rectangle{X: 25, Y: 30, Width: 30, Height: 40},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.rect.setBottom(testCase.bottom)
			if testCase.rect != testCase.expected {
				t.Errorf("setBottom() returned incorrect value. Expected: %v, Got: %v", testCase.expected, testCase.rect)
			}
		})
	}
}

func TestSetRight(t *testing.T) {
	testCases := []struct {
		name     string
		rect     rectangle
		right    int
		expected rectangle
	}{
		{
			name:     "Initial values",
			rect:     rectangle{X: 10, Y: 20, Width: 30, Height: 40},
			right:    50,
			expected: rectangle{X: 20, Y: 20, Width: 30, Height: 40},
		},
		{
			name:     "Moved values",
			rect:     rectangle{X: 20, Y: 35, Width: 30, Height: 40},
			right:    40,
			expected: rectangle{X: 10, Y: 35, Width: 30, Height: 40},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.rect.setRight(testCase.right)
			if testCase.rect != testCase.expected {
				t.Errorf("setRight() returned incorrect value. Expected: %v, Got: %v", testCase.expected, testCase.rect)
			}
		})
	}
}

func TestSetMidTop(t *testing.T) {
	testCases := []struct {
		name     string
		rect     rectangle
		midTopX  int
		midTopY  int
		expected rectangle
	}{
		{
			name:     "Initial values",
			rect:     rectangle{X: 10, Y: 20, Width: 30, Height: 40},
			midTopX:  50,
			midTopY:  60,
			expected: rectangle{X: 35, Y: 60, Width: 30, Height: 40},
		},
		{
			name:     "Modified values",
			rect:     rectangle{X: 15, Y: 25, Width: 30, Height: 40},
			midTopX:  60,
			midTopY:  70,
			expected: rectangle{X: 45, Y: 70, Width: 30, Height: 40},
		},
		{
			name:     "Moved values",
			rect:     rectangle{X: 25, Y: 35, Width: 30, Height: 40},
			midTopX:  70,
			midTopY:  80,
			expected: rectangle{X: 55, Y: 80, Width: 30, Height: 40},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.rect.setMidTop(testCase.midTopX, testCase.midTopY)
			if testCase.rect != testCase.expected {
				t.Errorf("setMidTop() returned incorrect value. Expected: %v, Got: %v", testCase.expected, testCase.rect)
			}
		})
	}
}

func TestSetMidLeft(t *testing.T) {
	testCases := []struct {
		name     string
		rect     rectangle
		midLeftX int
		midLeftY int
		expected rectangle
	}{
		{
			name:     "Initial values",
			rect:     rectangle{X: 10, Y: 20, Width: 30, Height: 40},
			midLeftX: 50,
			midLeftY: 60,
			expected: rectangle{X: 50, Y: 40, Width: 30, Height: 40},
		},
		{
			name:     "Modified values",
			rect:     rectangle{X: 15, Y: 25, Width: 30, Height: 40},
			midLeftX: 60,
			midLeftY: 70,
			expected: rectangle{X: 60, Y: 50, Width: 30, Height: 40},
		},
		{
			name:     "Moved values",
			rect:     rectangle{X: 25, Y: 35, Width: 30, Height: 40},
			midLeftX: 70,
			midLeftY: 80,
			expected: rectangle{X: 70, Y: 60, Width: 30, Height: 40},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.rect.setMidLeft(testCase.midLeftX, testCase.midLeftY)
			if testCase.rect != testCase.expected {
				t.Errorf("setMidLeft() returned incorrect value. Expected: %v, Got: %v", testCase.expected, testCase.rect)
			}
		})
	}
}

func TestSetMidBottom(t *testing.T) {
	testCases := []struct {
		name       string
		rect       rectangle
		midBottomX int
		midBottomY int
		expected   rectangle
	}{
		{
			name:       "Initial values",
			rect:       rectangle{X: 10, Y: 20, Width: 30, Height: 40},
			midBottomX: 50,
			midBottomY: 60,
			expected:   rectangle{X: 35, Y: 20, Width: 30, Height: 40},
		},
		{
			name:       "Modified values",
			rect:       rectangle{X: 15, Y: 25, Width: 30, Height: 40},
			midBottomX: 60,
			midBottomY: 70,
			expected:   rectangle{X: 45, Y: 30, Width: 30, Height: 40},
		},
		{
			name:       "Moved values",
			rect:       rectangle{X: 25, Y: 35, Width: 30, Height: 40},
			midBottomX: 70,
			midBottomY: 80,
			expected:   rectangle{X: 55, Y: 40, Width: 30, Height: 40},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.rect.setMidBottom(testCase.midBottomX, testCase.midBottomY)
			if testCase.rect != testCase.expected {
				t.Errorf("setMidBottom() returned incorrect value. Expected: %v, Got: %v", testCase.expected, testCase.rect)
			}
		})
	}
}

func TestSetMidRight(t *testing.T) {
	testCases := []struct {
		name      string
		rect      rectangle
		midRightX int
		midRightY int
		expected  rectangle
	}{
		{
			name:      "Initial values",
			rect:      rectangle{X: 10, Y: 20, Width: 30, Height: 40},
			midRightX: 50,
			midRightY: 60,
			expected:  rectangle{X: 20, Y: 40, Width: 30, Height: 40},
		},
		{
			name:      "Modified values",
			rect:      rectangle{X: 15, Y: 25, Width: 30, Height: 40},
			midRightX: 60,
			midRightY: 70,
			expected:  rectangle{X: 30, Y: 50, Width: 30, Height: 40},
		},
		{
			name:      "Moved values",
			rect:      rectangle{X: 25, Y: 35, Width: 30, Height: 40},
			midRightX: 70,
			midRightY: 80,
			expected:  rectangle{X: 40, Y: 60, Width: 30, Height: 40},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.rect.setMidRight(testCase.midRightX, testCase.midRightY)
			if testCase.rect != testCase.expected {
				t.Errorf("setMidRight() returned incorrect value. Expected: %v, Got: %v", testCase.expected, testCase.rect)
			}
		})
	}
}

func TestSetSize(t *testing.T) {
	testCases := []struct {
		name     string
		rect     rectangle
		width    int
		height   int
		expected rectangle
	}{
		{
			name:     "Initial values",
			rect:     rectangle{X: 10, Y: 20, Width: 30, Height: 40},
			width:    50,
			height:   60,
			expected: rectangle{X: 10, Y: 20, Width: 50, Height: 60},
		},
		{
			name:     "Modified values",
			rect:     rectangle{X: 15, Y: 25, Width: 30, Height: 40},
			width:    60,
			height:   70,
			expected: rectangle{X: 15, Y: 25, Width: 60, Height: 70},
		},
		{
			name:     "Moved values",
			rect:     rectangle{X: 25, Y: 35, Width: 30, Height: 40},
			width:    70,
			height:   80,
			expected: rectangle{X: 25, Y: 35, Width: 70, Height: 80},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.rect.setSize(testCase.width, testCase.height)
			if testCase.rect != testCase.expected {
				t.Errorf("setSize() returned incorrect value. Expected: %v, Got: %v", testCase.expected, testCase.rect)
			}
		})
	}
}
