package model

import (
	"term-ex/position"
	"term-ex/tile"
)

// Defines functionality exposed by all term-ex models
type Model interface {
	GetPosition() position.Position               // Gets the current position of the explorer in the world
	GetTile(position position.Position) tile.Tile // Gets the tile at the specified position if it exists, returns a fog tile otherwise
	HasVisited(position position.Position) bool   // Determines if the explorer has visited the specified position

	MoveNorth(distance int64) Model // Moves the explorer in the negative y direction
	MoveSouth(distance int64) Model // Moves the explorer in the positive y direction
	MoveWest(distance int64) Model  // Moves the explorer in the negative x direction
	MoveEast(distance int64) Model  // Moves the explorer in the positive x direction
}
