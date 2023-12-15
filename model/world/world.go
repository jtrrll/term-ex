package world

import (
	"math/rand"
)

// A two-dimensional, explorable world
type World struct {
	generationRules Rules               // rules that define valid tile possibilities for each position
	defaultTile     Tile                // the tile to use when generation rules fail to agree
	tileMap         map[Position][]Tile // the mapping of positions to tile possibilities
}

// Creates a new world
func NewWorld(generationRules Rules, defaultTile Tile, startingPosition Position, startingTile Tile, startingRadius uint8, genRadius uint8) World {
	// 1. Create empty world
	world := World{generationRules, defaultTile, map[Position][]Tile{}}
	// 2. Collapse the starting position and the radius around it
	world.tileMap[startingPosition] = []Tile{startingTile, startingTile}
	world.CollapseAll(startingPosition, startingRadius, genRadius)
	return world
}

// Gets the tile at the specified position if it is collapsed, returns false otherwise
func (w *World) Get(position Position) (Tile, bool) {
	possibilities, ok := w.tileMap[position]
	if ok && len(possibilities) == 1 {
		return possibilities[0], true
	}
	return Tile{}, false
}

// Collapses all uncollapsed positions within a specific radius
// into a single tile state, in order of increasing entropy
func (w *World) CollapseAll(center Position, collapseRadius uint8, genRadius uint8) {
	for w.collapse(center, collapseRadius, genRadius) {
	}
}

// Collapses the uncollapsed position with the least entropy within a collapse radius
// by randomly selecting a tile state from a slice of valid states.
// Also generates possibilities for surrounding tiles within a generation radius.
// Returns true if a position was successfully collapsed, false otherwise
func (w *World) collapse(collapseCenter Position, collapseRadius uint8, genRadius uint8) bool {
	// 1. Select position with least entropy in the collapse radius
	position, found := w.findLeastEntropicPosition(collapseCenter, collapseRadius)
	if !found {
		return false
	}
	// 2. Collapse position into a single tile state
	possibleTiles := w.tileMap[position]
	selectedTile := possibleTiles[rand.Intn(len(possibleTiles))]
	w.tileMap[position] = []Tile{selectedTile}
	// 3. Update possibilities for all positions in the generation radius
	for _, neighbor := range position.GetSurrounding(uint64(genRadius)) {
		// 3.1. Skip if already collapsed
		if len(w.tileMap[neighbor]) == 1 {
			continue
		}
		// 3.2. Generate possibilities
		neighborPossibleTiles := w.generationRules.ApplyAll(neighbor, *w)
		if len(neighborPossibleTiles) == 0 {
			neighborPossibleTiles = []Tile{w.defaultTile}
		}
		w.tileMap[neighbor] = neighborPossibleTiles
	}
	return true
}

// Finds the uncollapsed position with the lowest number of possible states within a specific radius.
// Returns false if no uncollapsed positions are found.
// Time complexity: O(n)
func (w *World) findLeastEntropicPosition(center Position, radius uint8) (Position, bool) {
	leastEntropicPosition := Position{}
	found := false
	for position, possibilities := range w.tileMap {
		if len(possibilities) > 1 {
			if !found || len(possibilities) < len(w.tileMap[leastEntropicPosition]) {
				if CalculateDistance(center, position) <= float64(radius) {
					leastEntropicPosition = position
					found = true
				}
			}
		}
	}
	return leastEntropicPosition, found
}
