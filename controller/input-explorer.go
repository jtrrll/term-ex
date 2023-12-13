package controller

import (
	"fmt"
	"term-ex/model"
	"term-ex/view"

	"github.com/gdamore/tcell"
)

// Explores a world in response to input events
type InputExplorer struct {
	model model.Model
	view  view.View
	input chan tcell.EventKey
}

// Creates a new explorer that responds to an input channel
func NewInputExplorer(model model.Model, view view.View, input chan tcell.EventKey) InputExplorer {
	return InputExplorer{model, view, input}
}

// Explores the map based on inputs.
// Panics if no input channel is defined
func (e *InputExplorer) Explore() {
	if e.input == nil {
		panic("input channel not defined")
	}
	// TODO: Implement
	for {
		ev := <-e.input
		fmt.Println("event received")
		fmt.Println(ev)
	}
}
