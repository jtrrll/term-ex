package controller

import "term-ex/model"

// Defines functionality exposed by all term-ex controllers
type Controller interface {
	Explore() // Starts exploring a world
}

// An command that an explorer can execute
type Command func(model.Model) model.Model

/**
var (
	moveNorth Command = func(model model.Model) model.Model { return model.MoveNorth(1) }
	moveSouth Command = func(model model.Model) model.Model { return model.MoveSouth(1) }
	moveEast  Command = func(model model.Model) model.Model { return model.MoveEast(1) }
	moveWest  Command = func(model model.Model) model.Model { return model.MoveWest(1) }
	moveDown  Command = func(model model.Model) model.Model { return model.MoveDown(1) }
	moveUp    Command = func(model model.Model) model.Model { return model.MoveUp(1) }
)
**/
