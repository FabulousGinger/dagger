package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGreeting(t *testing.T) {
	g := greeting()
	should := fmt.Sprintf("%c", YellowHeart)

	assert.Equal(t, should, g)
}
