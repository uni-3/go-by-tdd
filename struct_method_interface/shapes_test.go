package shapes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPerimeter(t *testing.T) {
	got := Perimeter(10.0, 10.0)
	expected := 40.0

	assert.Equal(t, expected, got)
}

func TestArea(t *testing.T) {
	got := Area(12.0, 6.0)
	expected := 72.0

	assert.Equal(t, expected, got)
}
