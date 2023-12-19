package view

import "term-ex/model"

// Defines functionality exposed by all term-ex views
type View interface {
	RenderIntro()                  // Renders an intro screen
	RenderModel(model model.Model) // Renders the provided model
	RenderOutro()                  // Renders an outro screen
	ShutDown()                     // Shuts down the view. This must be called before the application exits
}
