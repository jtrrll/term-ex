package world

// A two-dimensional square of terrain
type Tile struct {
	Terrain Terrain
	Height  int8
}

// The types of terrain
type Terrain int

const (
	Ocean Terrain = iota
	Shore
	Grass
)
