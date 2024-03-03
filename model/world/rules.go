package world

import (
	"term-ex/position"
	"term-ex/tile"

	mapset "github.com/deckarep/golang-set/v2"
)

// A rule that calculates valid tile possibilities for a position
type Rule func(p position.Position, w World) mapset.Set[tile.Tile]

// Several rules that calculate valid tile possibilities for a position
type Rules []Rule

// Generates a set of tile possibilities by applying several rules and finding their intersection
func (rules Rules) ApplyAll(p position.Position, w World) mapset.Set[tile.Tile] {
	allPossibilities := tile.GetAllTiles()
	for _, rule := range rules {
		allPossibilities = allPossibilities.Intersect(rule(p, w))
	}
	return allPossibilities
}

// Ensures ocean tiles are surrounded by sand, water, or more ocean
func SmoothOceans(p position.Position, w World) mapset.Set[tile.Tile] {
	iterator := p.GetSurrounding(1).Iterator()
	for pos := range iterator.C {
		ti, ok := w.Get(pos)
		if ok && ti == tile.Ocean {
			iterator.Stop()
			return mapset.NewSet(tile.Sand, tile.Water, tile.Ocean)
		}
	}
	return tile.GetAllTiles()
}

// Ensures tiles match their direct neighbors if they are all the same
func ClumpNeighbors(p position.Position, w World) mapset.Set[tile.Tile] {
	freq := w.GetTileFrequencies(p.GetSurrounding(1))
	for ti, count := range freq {
		if count == 4 {
			return mapset.NewSet(ti)
		}
	}
	return tile.GetAllTiles()
}

// Ensures sea level (ocean and sand), ground level, and mountain levels are distinct
func LeveledTerrain(p position.Position, w World) mapset.Set[tile.Tile] {
	freq := w.GetTileFrequencies(p.GetSurrounding(3))

	invalidTiles := []tile.Tile{}

	// Mountains should not spawn near oceans or sand
	if freq[tile.Ocean] > 0 || freq[tile.Sand] > 0 {
		invalidTiles = append(invalidTiles, tile.Mountain)
	}
	// Oceans and sand should not spawn near mountains
	if freq[tile.Mountain] > 0 {
		invalidTiles = append(invalidTiles, tile.Ocean)
		invalidTiles = append(invalidTiles, tile.Sand)
	}
	validTiles := tile.GetAllTiles()
	validTiles.RemoveAll(invalidTiles...)
	return validTiles
}

// Ensures the starting area is a grassland
func StartOnGrassland(p position.Position, w World) mapset.Set[tile.Tile] {
	if p.DistanceTo(position.Position{}) <= 20 {
		return mapset.NewSet(tile.Dirt, tile.Grass, tile.Path, tile.Tree)
	}
	return tile.GetAllTiles()
}
