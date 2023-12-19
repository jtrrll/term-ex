package model

import (
	"term-ex/tile"
)

// These tiles cannot be walked through
var (
	blockingTiles = map[tile.Tile]struct{}{
		tile.Ocean:    {},
		tile.Mountain: {},
		tile.Wall:     {},
	}
)
