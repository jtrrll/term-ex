package model

import "term-ex/model/world"

// Defines functionality exposed by all term-ex models
type Model interface {
	// Information retrieval functionality

	GetPosition() world.Position       // Gets the current position of the explorer in the world
	GetTile(world.Position) world.Tile // Gets the tile at the specified position

	// Movement functionality

	MoveNorth(distance int) Model // Moves the explorer in the negative y direction
	MoveSouth(distance int) Model // Moves the explorer in the positive y direction
	MoveEast(distance int) Model  // Moves the explorer in the negative x direction
	MoveWest(distance int) Model  // Moves the explorer in the positive x direction
	MoveDown(distance int) Model  // Moves the explorer in the negative z direction
	MoveUp(distance int) Model    // Moves the explorer in the positive z direction
}

//TODO: random world generation
