package controller

import "term-ex/model"

// Defines functionality exposed by all term-ex controllers
type Controller interface {
	Explore() error // Starts exploring a world
}

// A command that an explorer can execute
type Command func(model.Model) model.Model
