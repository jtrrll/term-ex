package controller

import (
	"errors"
	"term-ex/model"
	"term-ex/view"

	"github.com/gdamore/tcell"
)

// Explores a world in response to input events
type InputExplorer struct {
	input chan *tcell.EventKey // the channel to read inputs from
	model model.Model          // the model to manipulate
	view  view.View            // the view to manipulate
}

// Creates a new explorer that responds to an input channel
func NewInputExplorer(input chan *tcell.EventKey, model model.Model, view view.View) *InputExplorer {
	return &InputExplorer{input, model, view}
}

// Explores the map based on inputs
func (e *InputExplorer) Explore() error {
	// 1. Ensure required fields are initialized
	if e.input == nil {
		return errors.New("input channel not defined")
	}
	if e.model == nil {
		return errors.New("model not defined")
	}
	// 2. Render splash screen
	e.view.RenderIntro()
	// 3. Respond to user input
	for {
		ev := <-e.input
		switch ev.Key() {
		case tcell.KeyESC: // Exit process
			e.view.RenderOutro()
			fallthrough
		case tcell.KeyCtrlC: // Exit process
			e.view.ShutDown()
			return nil
		case tcell.KeyCtrlR: // Refresh view
			e.view.RenderModel(e.model)
		case tcell.KeyUp: // Move north 1 unit
			e.view.RenderModel(e.model.MoveNorth(1))
		case tcell.KeyDown: // Move south 1 unit
			e.view.RenderModel(e.model.MoveSouth(1))
		case tcell.KeyLeft: // Move west 1 unit
			e.view.RenderModel(e.model.MoveWest(1))
		case tcell.KeyRight: // Move east 1 unit
			e.view.RenderModel(e.model.MoveEast(1))
		}
	}
}
