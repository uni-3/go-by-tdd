package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("exist key", func(t *testing.T) {

		got, err := dictionary.Search("test")
		assert.NoError(t, err)
		expected := "this is just a test"

		assert.Equal(t, expected, got)

	})

	t.Run("not exist key", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		expected := ErrNotFound

		assert.Equal(t, err, expected)
	})
}

func TestAdd(t *testing.T) {

	t.Run("add key", func(t *testing.T) {
		dictionary := Dictionary{}
		dictionary.Add("test", "this is just a test")

		got, err := dictionary.Search("test")
		assert.NoError(t, err)
		expected := "this is just a test"

		assert.Equal(t, expected, got)
	})
}
