package shapes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPerimeter(t *testing.T) {
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := Perimeter(rectangle)
		expected := 40.0

		assert.Equal(t, expected, got)
	})

}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name     string
		shape    Shape
		expected float64
	}{
		{
			name:     "rectangle",
			shape:    Rectangle{12.0, 6.0},
			expected: 72.0,
		},
		{
			name:     "circle",
			shape:    Circle{10.0},
			expected: 314.1592653589793,
		},
		{
			name:     "triangle",
			shape:    Triangle{12.0, 6.0},
			expected: 36.0,
		},
	}

	for _, test := range areaTests {
		t.Run(test.name, func(t *testing.T) {
			got := test.shape.Area()
			assert.Equal(t, got, test.expected)
		})
	}
}
