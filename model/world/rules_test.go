package world

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApplyAll(t *testing.T) {
	rule1 := func(Position, World) []Tile {
		return []Tile{{Terrain: Ocean}, {Terrain: Shore}, {Terrain: Grass}}
	}
	rule2 := func(Position, World) []Tile {
		return []Tile{{Terrain: Ocean}}
	}
	rule3 := func(Position, World) []Tile {
		return []Tile{}
	}

	actual1 := Rules{rule1}.ApplyAll(Position{}, World{})
	assert.Len(t, actual1, 3)
	assert.Contains(t, actual1, Tile{Terrain: Ocean})
	assert.Contains(t, actual1, Tile{Terrain: Shore})
	assert.Contains(t, actual1, Tile{Terrain: Grass})

	actual2 := Rules{rule1, rule1, rule1}.ApplyAll(Position{}, World{})
	assert.Len(t, actual2, 3)
	assert.Contains(t, actual2, Tile{Terrain: Ocean})
	assert.Contains(t, actual2, Tile{Terrain: Shore})
	assert.Contains(t, actual2, Tile{Terrain: Grass})

	actual3 := Rules{rule1, rule2}.ApplyAll(Position{}, World{})
	assert.Len(t, actual3, 1)
	assert.Contains(t, actual3, Tile{Terrain: Ocean})

	actual4 := Rules{rule1, rule3}.ApplyAll(Position{}, World{})
	assert.Len(t, actual4, 0)

	actual5 := Rules{rule3, rule1, rule1, rule2}.ApplyAll(Position{}, World{})
	assert.Len(t, actual5, 0)
}
