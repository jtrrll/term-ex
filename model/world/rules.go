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

// Returns three possible tiles, ocean, sand, and grass, all of height=0
func OceanSandGrass(position position.Position, world World) []tile.Tile {
	return []tile.Tile{tile.Ocean, tile.Sand, tile.Grass}
}

// Returns two possible tiles, sand and ocean, all of height=0
func SandOcean(position position.Position, world World) []tile.Tile {
	return []tile.Tile{tile.Sand, tile.Ocean}
}

// Returns three possible tiles, ocean, sand, and grass, all of height=0
func SandBetweenOceanAndGrass(position position.Position, world World) []tile.Tile {
	for _, pos := range position.GetSurrounding(1) {
		ti, ok := world.Get(pos)
		if ok {
			if ti == tile.Ocean {
				return []tile.Tile{tile.Ocean, tile.Sand}
			} else if ti == tile.Grass {
				return []tile.Tile{tile.Sand, tile.Grass}
			}
		}
	}
	return []tile.Tile{tile.Ocean, tile.Sand, tile.Grass}
}
