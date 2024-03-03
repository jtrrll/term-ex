package tile

import (
	mapset "github.com/deckarep/golang-set/v2"
)

// The types of tile
type Tile int

const (
	Fog Tile = iota

	Ocean
	Water

	Sand
	Dirt
	Grass
	Path

	Tree
	Mountain
)

// Returns a set of all tiles (excluding the zero value, Fog)
func GetAllTiles() mapset.Set[Tile] {
	return mapset.NewSet(Ocean, Water, Sand, Dirt, Grass, Path, Tree, Mountain)
}
