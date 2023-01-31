package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGreeting(t *testing.T) {
	g := greeting()
	should := fmt.Sprintf("%c %c", wavingHand, worldMap)

	assert.Equal(t, should, g)
}
