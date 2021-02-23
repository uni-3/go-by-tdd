package main

import (
	"sync"
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

	t.Run("safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := NewCounter()

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func(w *sync.WaitGroup) {
				counter.Inc()
				w.Done()
			}(&wg)

		}
		wg.Wait()

		assert.Equal(t, wantedCount, counter.value)
	})

}

func NewCounter() *Counter {
	// muを使うときはcopyされないようにpointerで
	return &Counter{}
}
