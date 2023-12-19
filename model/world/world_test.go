package world

import (
	"term-ex/position"
	"term-ex/tile"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewWorldWithNoRules(t *testing.T) {
	world := NewWorld(Rules{}, 0, 1, 0, 0)
	assert.Len(t, world.tileMap, 1)
	ti, ok := world.Get(position.Position{})
	assert.True(t, ok)
	assert.Equal(t, tile.Tile(1), ti)

	world = NewWorld(Rules{}, 0, 1, 2, 0)
	assert.Len(t, world.tileMap, 13)
	for i := range world.tileMap {
		ti, ok := world.Get(i)
		assert.True(t, ok)
		if (i == position.Position{}) {
			assert.Equal(t, tile.Tile(1), ti)
		} else {
			assert.Equal(t, tile.Tile(0), ti)
		}
	}

	world = NewWorld(Rules{}, 2, 1, 3, 1)
	assert.Len(t, world.tileMap, 45)
	for i := range world.tileMap {
		ti, ok := world.Get(i)
		assert.True(t, ok)
		if (i == position.Position{}) {
			assert.Equal(t, tile.Tile(1), ti)
		} else {
			assert.Equal(t, tile.Tile(2), ti)
		}
	}
}
