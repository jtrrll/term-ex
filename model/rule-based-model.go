package model

import (
	"term-ex/model/world"
	"term-ex/position"
	"term-ex/tile"
)

// A model that generates a world according to specified rules
type RuleBasedModel struct {
	collapseRadius uint8                          // the radius around the explorer where all tiles must be collapsed
	genRadius      uint8                          // the radius of excess tiles to cache around collapsed tiles
	world          world.World                    // the world to explore
	position       position.Position              // the current position of the explorer
	visited        map[position.Position]struct{} // a set of positions that have been visited by the explorer
}

// Creates a new model that generates a world according to specified rules
func NewRuleBasedModel(generationRules world.Rules, defaultTile tile.Tile, startingTile tile.Tile, collapseRadius uint8, genRadius uint8) *RuleBasedModel {
	return &RuleBasedModel{
		collapseRadius: collapseRadius,
		genRadius:      genRadius,
		world:          world.NewWorld(generationRules, defaultTile, startingTile, collapseRadius, genRadius),
		position:       position.Position{},
		visited:        map[position.Position]struct{}{{}: {}},
	}
}

// Gets the current position of the explorer in the world
func (m *RuleBasedModel) GetPosition() position.Position {
	return m.position
}

// Gets the tile at the specified position if it exists, returns a fog tile otherwise
func (m *RuleBasedModel) GetTile(position position.Position) tile.Tile {
	ti, ok := m.world.Get(position)
	if !ok {
		return tile.Fog
	}
	return ti
}

// Determines if the explorer has visited the specified position
func (m *RuleBasedModel) HasVisited(position position.Position) bool {
	_, hasVisited := m.visited[position]
	return hasVisited
}

// Moves the explorer in the negative y direction
func (m *RuleBasedModel) MoveNorth(distance int64) Model {
	// 1. Check move validity
	nextTile := m.GetTile(position.Position{X: m.position.X, Y: m.position.Y - distance})
	_, ok := blockingTiles[nextTile]
	if ok {
		return m
	}
	// 2. Update position
	m.position.Y -= distance
	m.visited[m.position] = struct{}{}
	// 3. Generate new tiles
	m.world.CollapseAll(m.position, m.collapseRadius, m.genRadius)
	return m
}

// Moves the explorer in the positive y direction
func (m *RuleBasedModel) MoveSouth(distance int64) Model {
	// 1. Check move validity
	nextTile := m.GetTile(position.Position{X: m.position.X, Y: m.position.Y + distance})
	_, ok := blockingTiles[nextTile]
	if ok {
		return m
	}
	// 2. Update position
	m.position.Y += distance
	m.visited[m.position] = struct{}{}
	// 3. Generate new tiles
	m.world.CollapseAll(m.position, m.collapseRadius, m.genRadius)
	return m
}

// Moves the explorer in the negative x direction
func (m *RuleBasedModel) MoveWest(distance int64) Model {
	// 1. Check move validity
	nextTile := m.GetTile(position.Position{X: m.position.X - distance, Y: m.position.Y})
	_, ok := blockingTiles[nextTile]
	if ok {
		return m
	}
	// 2. Update position
	m.position.X -= distance
	m.visited[m.position] = struct{}{}
	// 3. Generate new tiles
	m.world.CollapseAll(m.position, m.collapseRadius, m.genRadius)
	return m
}

// Moves the explorer in the positive x direction
func (m *RuleBasedModel) MoveEast(distance int64) Model {
	// 1. Check move validity
	nextTile := m.GetTile(position.Position{X: m.position.X + distance, Y: m.position.Y})
	_, ok := blockingTiles[nextTile]
	if ok {
		return m
	}
	// 2. Update position
	m.position.X += distance
	m.visited[m.position] = struct{}{}
	// 3. Generate new tiles
	m.world.CollapseAll(m.position, m.collapseRadius, m.genRadius)
	return m
}
