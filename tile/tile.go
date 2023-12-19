package tile

// The types of tile
type Tile int

const (
	Ocean Tile = iota
	Sand
	Grass
	Dirt
	Path
	Mountain
	Wall
)
