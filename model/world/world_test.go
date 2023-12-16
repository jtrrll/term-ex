package world

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewWorldWithNoRules(t *testing.T) {
	world := NewWorld(Rules{}, Tile{Terrain: 0}, Position{}, Tile{Terrain: 1}, 0, 0)
	assert.Len(t, world.tileMap, 1)
	tile, ok := world.Get(Position{})
	assert.True(t, ok)
	assert.Equal(t, Tile{Terrain: 1}, tile)

	world = NewWorld(Rules{}, Tile{Terrain: 0}, Position{}, Tile{Terrain: 1}, 2, 0)
	assert.Len(t, world.tileMap, 13)
	for i := range world.tileMap {
		tile, ok := world.Get(i)
		assert.True(t, ok)
		if (i == Position{}) {
			assert.Equal(t, Tile{Terrain: 1}, tile)
		} else {
			assert.Equal(t, Tile{Terrain: 0}, tile)
		}
	}

	world = NewWorld(Rules{}, Tile{Terrain: 2}, Position{}, Tile{Terrain: 1}, 3, 1)
	assert.Len(t, world.tileMap, 45)
	for i := range world.tileMap {
		tile, ok := world.Get(i)
		assert.True(t, ok)
		if (i == Position{}) {
			assert.Equal(t, Tile{Terrain: 1}, tile)
		} else {
			assert.Equal(t, Tile{Terrain: 2}, tile)
		}
	}
}
