package integers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	assert.Equal(t, sum, expected)
}

// refer: https://blog.golang.org/examples
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
