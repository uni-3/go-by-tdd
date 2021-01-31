package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "chris")

	got := buffer.String()
	expected := "hello, chris"

	assert.Equal(t, expected, got)
}
