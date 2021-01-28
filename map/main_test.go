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

	t.Run("new key", func(t *testing.T) {
		dictionary := Dictionary{}
		dictionary.Add("test", "this is just a test")

		got, err := dictionary.Search("test")
		assert.NoError(t, err)
		expected := "this is just a test"

		assert.Equal(t, expected, got)
	})

	t.Run("existing key", func(t *testing.T) {
		key := "test"
		value := "this is just a test"
		dictionary := Dictionary{key: value}
		err := dictionary.Add(key, "new test")
		assert.Error(t, err, ErrWordExists)

		// testの値が更新されないこと
		got, err := dictionary.Search("test")
		assert.NoError(t, err)
		expected := "this is just a test"
		assert.Equal(t, expected, got)
	})
}

func TestUpdate(t *testing.T) {

	t.Run("existing key", func(t *testing.T) {
		key := "test"
		value := "this is just a test"
		dictionary := Dictionary{key: value}
		newValue := "new test"
		err := dictionary.Update(key, newValue)
		assert.NoError(t, err)

		// testの値が更新されること
		got, err := dictionary.Search("test")
		assert.NoError(t, err)
		expected := "new test"
		assert.Equal(t, expected, got)

	})

	t.Run("new key", func(t *testing.T) {
		key := "test"
		value := "this is just a test"
		dictionary := Dictionary{}
		err := dictionary.Update(key, value)
		assert.Error(t, err, ErrWordDoesNotExists)
	})
}

func TestDelete(t *testing.T) {

	t.Run("existing key", func(t *testing.T) {
		key := "test"
		value := "this is just a test"
		dictionary := Dictionary{key: value}
		err := dictionary.Delete(key)
		assert.NoError(t, err)

	})
}
