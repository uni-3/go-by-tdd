package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	t.Run("5 nums", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		expected := 15
		assert.Equal(t, expected, got)
	})

	t.Run("use slice", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		expected := 6

		assert.Equal(t, expected, got)
	})
}

func TestSumAll(t *testing.T) {
	t.Run("2 arrays", func(t *testing.T) {
		numbers1 := []int{1, 2}
		numbers2 := []int{0, 9}

		got := SumAll(numbers1, numbers2)
		expected := []int{3, 9}

		assert.Equal(t, expected, got)
	})
}

func TestSumAllTails(t *testing.T) {
	t.Run("2 arrays", func(t *testing.T) {
		numbers1 := []int{1, 2}
		numbers2 := []int{0, 9}

		got := SumAllTails(numbers1, numbers2)
		expected := []int{2, 9}

		assert.Equal(t, expected, got)
	})

	t.Run("empty slice", func(t *testing.T) {
		numbers1 := []int{}
		numbers2 := []int{3, 4, 5}

		got := SumAllTails(numbers1, numbers2)
		expected := []int{0, 9}

		assert.Equal(t, expected, got)
	})
}
