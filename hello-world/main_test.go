package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {

	t.Run("hello to people", func(t *testing.T) {
		got := Hello("Chris")
		expected := "hello, Chris"

		assert.Equal(t, expected, got)

	})

}
