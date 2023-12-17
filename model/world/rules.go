package world

import (
	"term-ex/intersection"
	"term-ex/position"
	"term-ex/tile"
)

// A rule that calculates valid tile possibilities for a position
type Rule func(position position.Position, world World) []tile.Tile

// Several rules that calculate valid tile possibilities for a position
type Rules []Rule

// Generates tile possibilities by applying several rules and finding their intersection
func (rules Rules) ApplyAll(position position.Position, world World) []tile.Tile {
	allPossibilities := [][]tile.Tile{}
	for _, rule := range rules {
		allPossibilities = append(allPossibilities, rule(position, world))
	}
	return intersection.Intersection(allPossibilities...)
}

// Returns three possible tiles, ocean, shore, and grass, all of height=0
func OceanShoreGrass(position.Position, World) []tile.Tile {
	return []tile.Tile{{Terrain: tile.Ocean}, {Terrain: tile.Shore}, {Terrain: tile.Grass}}
}
