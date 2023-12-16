package model

import (
	"term-ex/model/world"
)

// A model that always returns a new copy instead of mutating
type ImmutableModel struct {
	collapseRadius uint8                       // the radius around the explorer where all tiles must be collapsed
	genRadius      uint8                       // the radius of excess tiles to cache around collapsed tiles
	world          world.World                 // the world to explore
	position       world.Position              // the current position of the explorer
	visited        map[world.Position]struct{} // a set of positions that have been visited by the explorer
}

// Creates a new model that always returns a new copy instead of mutating
func NewImmutableModel(generationRules world.Rules, defaultTile world.Tile, startingPosition world.Position, startingTile world.Tile, collapseRadius uint8, genRadius uint8) ImmutableModel {
	return ImmutableModel{
		collapseRadius: collapseRadius,
		genRadius:      genRadius,
		world:          world.NewWorld(generationRules, defaultTile, startingPosition, startingTile, collapseRadius, genRadius),
		position:       startingPosition,
		visited:        map[world.Position]struct{}{startingPosition: {}},
	}
}

// Gets the current position of the explorer in the world
func (m ImmutableModel) GetPosition() world.Position {
	return m.position
}

// Gets the tile at the specified position if it exists, returns false otherwise
func (m ImmutableModel) GetTile(position world.Position) (world.Tile, bool) {
	tile, ok := m.world.Get(position)
	return tile, ok
}

// Determines if the explorer has visited the specified position
func (m ImmutableModel) HasVisited(position world.Position) bool {
	_, hasVisited := m.visited[position]
	return hasVisited
}

// Moves the explorer in the negative y direction
func (m ImmutableModel) MoveNorth(distance int64) Model {
	m.position.Y -= distance
	m.visited[m.position] = struct{}{}
	m.world.CollapseAll(m.position, m.collapseRadius, m.genRadius)
	return m
}

// Moves the explorer in the positive y direction
func (m ImmutableModel) MoveSouth(distance int64) Model {
	m.position.Y += distance
	m.visited[m.position] = struct{}{}
	m.world.CollapseAll(m.position, m.collapseRadius, m.genRadius)
	return m
}

// Moves the explorer in the negative x direction
func (m ImmutableModel) MoveEast(distance int64) Model {
	m.position.X -= distance
	m.visited[m.position] = struct{}{}
	m.world.CollapseAll(m.position, m.collapseRadius, m.genRadius)
	return m
}

// Moves the explorer in the positive x direction
func (m ImmutableModel) MoveWest(distance int64) Model {
	m.position.X += distance
	m.visited[m.position] = struct{}{}
	m.world.CollapseAll(m.position, m.collapseRadius, m.genRadius)
	return m
}
