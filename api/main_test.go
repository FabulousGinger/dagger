package main

import (
	"fmt"
	"testing"

	"github.com/enescakir/emoji"
	"github.com/stretchr/testify/assert"
)

func TestGreeting(t *testing.T) {
	g := greeting()
	should := fmt.Sprintf("%v %v", emoji.WavingHand, emoji.WorldMap)

	assert.Equal(t, should, g)
}
