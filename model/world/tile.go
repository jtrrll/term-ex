package world

// A three-dimensional cube of terrain
type Tile struct {
	Terrain Terrain
}

// The types of terrain
type Terrain int

const (
	Air Terrain = iota
	Grass
	Water
	Mountain
)
