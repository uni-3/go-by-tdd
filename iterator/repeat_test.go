package iteration

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepeat(t *testing.T) {

	t.Run("repeat 0", func(t *testing.T) {
		repeated := Repeat("a", 0)
		expected := ""

		assert.Equal(t, repeated, expected)
	})

	t.Run("repeat 2", func(t *testing.T) {
		repeated := Repeat("a", 2)
		expected := strings.Repeat("a", 2)

		assert.Equal(t, repeated, expected)
	})

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}

func ExampleRepeat() {
	fmt.Printf("ba%s", Repeat("na", 2))
	// Output: banana
}
