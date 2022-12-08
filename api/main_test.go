package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGreeting(t *testing.T) {
	g := greeting()
	should := "Hello"

	assert.Equal(t, should, g)
}
