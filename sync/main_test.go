package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test .
func Test(t *testing.T) {
	t.Run("increment counter 3 times", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assert.Equal(t, 3, counter.value)
	})

}
