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

	t.Run("when an emply string is supplied", func(t *testing.T) {
		got := Hello("")
		expected := "hello, world"

		assert.Equal(t, expected, got)
	})
}
