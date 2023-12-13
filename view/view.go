package view

import "term-ex/model"

// Defines functionality exposed by all term-ex views
type View interface {
	Update(model.Model) // Updates the view with the most recent state
}
