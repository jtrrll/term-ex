package world

import (
	"term-ex/position"
	"term-ex/tile"
	"testing"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/stretchr/testify/assert"
)

func TestApplyAll(t *testing.T) {
	rule1 := func(position.Position, World) mapset.Set[tile.Tile] {
		return mapset.NewSet[tile.Tile](0, 1, 2)
	}
	rule2 := func(position.Position, World) mapset.Set[tile.Tile] {
		return mapset.NewSet[tile.Tile](0)
	}
	rule3 := func(position.Position, World) mapset.Set[tile.Tile] {
		return mapset.NewSet[tile.Tile]()
	}

	actual1 := Rules{rule1}.ApplyAll(position.Position{}, World{})
	assert.Equal(t, mapset.NewSet[tile.Tile](0, 1, 2), actual1)

	actual2 := Rules{rule1, rule1, rule1}.ApplyAll(position.Position{}, World{})
	assert.Equal(t, mapset.NewSet[tile.Tile](0, 1, 2), actual2)

	actual3 := Rules{rule1, rule2}.ApplyAll(position.Position{}, World{})
	assert.Equal(t, mapset.NewSet[tile.Tile](0), actual3)

	actual4 := Rules{rule1, rule3}.ApplyAll(position.Position{}, World{})
	assert.Equal(t, mapset.NewSet[tile.Tile](), actual4)

	actual5 := Rules{rule3, rule1, rule1, rule2}.ApplyAll(position.Position{}, World{})
	assert.Equal(t, mapset.NewSet[tile.Tile](), actual5)
}
