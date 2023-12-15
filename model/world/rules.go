package world

import (
	"term-ex/util"
)

// A rule that calculates valid tile possibilities for a position
type Rule func(Position, World) []Tile

// Several rules that calculate valid tile possibilities for a position
type Rules []Rule

// Generates tile possibilities by applying several rules and finding their intersection
func (rules Rules) ApplyAll(position Position, world World) []Tile {
	allPossibilities := [][]Tile{}
	for _, rule := range rules {
		allPossibilities = append(allPossibilities, rule(position, world))
	}
	return util.Intersection(allPossibilities...)
}

// Returns three possible tiles, ocean, shore, and grass, all of height=0
func OceanShoreGrass(Position, World) []Tile {
	return []Tile{{Terrain: Ocean}, {Terrain: Shore}, {Terrain: Grass}}
}
