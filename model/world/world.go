package world

import (
	"math"
	"math/rand"
	"term-ex/position"
	"term-ex/tile"

	mapset "github.com/deckarep/golang-set/v2"
)

// A two-dimensional, explorable world
type World struct {
	generationRules Rules                                       // rules that define valid tile possibilities for each position
	defaultTile     tile.Tile                                   // the tile to use when generation rules fail to agree
	tileMap         map[position.Position]mapset.Set[tile.Tile] // the mapping of positions to tile possibilities
}

// Creates a new world that follows defined generation rules
func NewWorld(generationRules Rules, defaultTile tile.Tile, startingTile tile.Tile, startingRadius uint8, genRadius uint8) World {
	// 1. Create empty world
	world := World{generationRules, defaultTile, map[position.Position]mapset.Set[tile.Tile]{}}
	// 2. Collapse the starting position and the radius around it
	world.tileMap[position.Position{}] = mapset.NewSet(startingTile, defaultTile)
	world.CollapseAll(position.Position{}, startingRadius, genRadius)
	return world
}

// Gets the tile at the specified position if it is collapsed, returns false otherwise
func (w *World) Get(position position.Position) (tile.Tile, bool) {
	possibilities, ok := w.tileMap[position]
	if ok && possibilities.Cardinality() == 1 {
		return possibilities.ToSlice()[rand.Intn(possibilities.Cardinality())], true
	}
	return 0, false
}

// Gets a set of all tiles and their frequency at the specified positions if they are collapsed
func (w *World) GetTileFrequencies(positions mapset.Set[position.Position]) map[tile.Tile]uint64 {
	freq := map[tile.Tile]uint64{}
	for pos := range positions.Iter() {
		ti, ok := w.Get(pos)
		if ok {
			freq[ti] += 1
		}
	}
	return freq
}

// Collapses all uncollapsed positions within a specific radius
// into a single tile state, in order of increasing entropy
func (w *World) CollapseAll(center position.Position, collapseRadius uint8, genRadius uint8) {
	for w.collapse(center, collapseRadius, genRadius) {
	}
}

// Collapses the uncollapsed position with the least entropy within a collapse radius
// by randomly selecting a tile state from a slice of valid states.
// Also generates possibilities for surrounding tiles within a generation radius.
// Returns true if a position was successfully collapsed, false otherwise
func (w *World) collapse(collapseCenter position.Position, collapseRadius uint8, genRadius uint8) bool {
	// 1. Select position with least entropy in the collapse radius
	position, found := w.findLeastEntropicPosition(collapseCenter, collapseRadius)
	if !found {
		return false
	}
	// 2. If a non-cached position is selected, generate possibilities
	possibleTiles := w.generationRules.ApplyAll(position, *w)
	if possibleTiles.Cardinality() == 0 {
		possibleTiles = mapset.NewSet(w.defaultTile)
	}
	// 3. Collapse position into a single tile state
	selectedTile := possibleTiles.ToSlice()[rand.Intn(possibleTiles.Cardinality())]
	w.tileMap[position] = mapset.NewSet(selectedTile)
	// 4. Update and cache possibilities for all positions in the generation radius
	for neighbor := range position.GetSurrounding(uint64(genRadius)).Iter() {
		// 4.1. Skip if already collapsed
		possibilities, ok := w.tileMap[neighbor]
		if ok && possibilities.Cardinality() == 1 {
			continue
		}
		// 4.2. Generate possibilities and cache them
		neighborPossibleTiles := w.generationRules.ApplyAll(neighbor, *w)
		if neighborPossibleTiles.Cardinality() == 0 {
			neighborPossibleTiles = mapset.NewSet(w.defaultTile)
		}
		w.tileMap[neighbor] = neighborPossibleTiles
	}
	return true
}

// Finds the uncollapsed position with the lowest number of possible states within a specific radius.
// Returns false if no uncollapsed positions are found
func (w *World) findLeastEntropicPosition(center position.Position, radius uint8) (position.Position, bool) {
	leastEntropicPosition := position.Position{}
	leastEntropy := 0
	found := false
	searchZone := center.GetSurrounding(uint64(radius))
	searchZone.Add(center)
	for posn := range searchZone.Iter() {
		entropy := math.MaxInt
		possibilities, ok := w.tileMap[posn]
		if ok {
			entropy = possibilities.Cardinality()
		}
		if entropy > 1 {
			if !found || entropy < leastEntropy {
				if position.CalculateDistance(center, posn) <= float64(radius) {
					leastEntropicPosition = posn
					leastEntropy = entropy
					found = true
				}
			}
		}
	}
	return leastEntropicPosition, found
}
