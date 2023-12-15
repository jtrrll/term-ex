package model

import "term-ex/model/world"

// Defines functionality exposed by all term-ex models
type Model interface {
	// Information retrieval functionality

	GetPosition() world.Position               // Gets the current position of the explorer in the world
	GetTile(world.Position) (world.Tile, bool) // Gets the tile at the specified position if it exists, returns false otherwise
	HasVisited(world.Position) bool            // Determines if the explorer has visited the specified position

	// Movement functionality

	MoveNorth(distance int64) Model // Moves the explorer in the negative y direction
	MoveSouth(distance int64) Model // Moves the explorer in the positive y direction
	MoveEast(distance int64) Model  // Moves the explorer in the negative x direction
	MoveWest(distance int64) Model  // Moves the explorer in the positive x direction
}
