package view

import "term-ex/model"

// Defines functionality exposed by all term-ex views
type View interface {
	RenderModel(model model.Model) // Renders the provided model
	RenderIntro()                  // Renders an intro screen
	RenderOutro()                  // Renders an outro screen
	ShutDown()                     // Shuts down the view
}
