package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {

	t.Run("hello to people", func(t *testing.T) {
		got := Hello("Chris", "English")
		expected := "hello, Chris"

		assert.Equal(t, expected, got)
	})

	t.Run("when an emply string is supplied", func(t *testing.T) {
		got := Hello("", "English")
		expected := "hello, world"

		assert.Equal(t, expected, got)
	})

	t.Run("in spanish", func(t *testing.T) {
		got := Hello("elodie", "Spanish")
		expected := "hola, elodie"

		assert.Equal(t, expected, got)
	})

	t.Run("in french", func(t *testing.T) {
		got := Hello("lauren", "French")
		expected := "bonjour, lauren"

		assert.Equal(t, expected, got)
	})
}
