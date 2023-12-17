package view

import "term-ex/model"

// Defines functionality exposed by all term-ex views
type View interface {
	RenderModel(model.Model) // Renders the provided model
	RenderSplash()           // Renders a splash screen
	RenderExit()             // Renders an exit screen
}
