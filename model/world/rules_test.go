package world

import (
	"term-ex/position"
	"term-ex/tile"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApplyAll(t *testing.T) {
	rule1 := func(position.Position, World) []tile.Tile {
		return []tile.Tile{0, 1, 2}
	}
	rule2 := func(position.Position, World) []tile.Tile {
		return []tile.Tile{0}
	}
	rule3 := func(position.Position, World) []tile.Tile {
		return []tile.Tile{}
	}

	actual1 := Rules{rule1}.ApplyAll(position.Position{}, World{})
	assert.Len(t, actual1, 3)
	assert.Contains(t, actual1, 0)
	assert.Contains(t, actual1, 1)
	assert.Contains(t, actual1, 2)

	actual2 := Rules{rule1, rule1, rule1}.ApplyAll(position.Position{}, World{})
	assert.Len(t, actual2, 3)
	assert.Contains(t, actual2, 0)
	assert.Contains(t, actual2, 1)
	assert.Contains(t, actual2, 2)

	actual3 := Rules{rule1, rule2}.ApplyAll(position.Position{}, World{})
	assert.Len(t, actual3, 1)
	assert.Contains(t, actual3, 0)

	actual4 := Rules{rule1, rule3}.ApplyAll(position.Position{}, World{})
	assert.Len(t, actual4, 0)

	actual5 := Rules{rule3, rule1, rule1, rule2}.ApplyAll(position.Position{}, World{})
	assert.Len(t, actual5, 0)
}
