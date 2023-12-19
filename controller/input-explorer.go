package controller

import (
	"errors"
	"term-ex/model"
	"term-ex/view"
	"time"

	"github.com/gdamore/tcell"
)

// Explores a world in response to input events
type InputExplorer struct {
	input        chan *tcell.EventKey // the channel to read inputs from
	model        model.Model          // the model to manipulate
	view         view.View            // the view to manipulate
	introEnabled bool                 // whether to display an intro sequence
}

// Creates a new explorer that responds to an input channel
func NewInputExplorer(input chan *tcell.EventKey, model model.Model, view view.View, introEnabled bool) *InputExplorer {
	return &InputExplorer{input, model, view, introEnabled}
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
	// 2. Run intro sequence if enabled
	if e.introEnabled {
		e.runIntro(3 * time.Second)
	}
	// 4. Run the core execution loop
	// If a user triggers graceful shutdown, show the outro
	if e.runCore() {
		e.runOutro(3 * time.Second)
	}
	e.view.ShutDown()
	return nil
}

// Runs an intro sequence for the specified amount of time
func (e *InputExplorer) runIntro(duration time.Duration) {
	timer := time.NewTimer(duration)
	e.view.RenderIntro()
	for {
		select {
		case ev := <-e.input:
			switch ev.Key() {
			case tcell.KeyCtrlR: // Refresh view
				e.view.RenderIntro()
			default: // End intro
				return
			}
		case <-timer.C: // End intro
			return
		}
	}
}

// Runs the core execution loop until ended by the user
// Returns true if the user triggers a graceful shutdown, false otherwise
func (e *InputExplorer) runCore() bool {
	e.view.RenderModel(e.model)
	for {
		ev := <-e.input
		switch ev.Key() {
		case tcell.KeyESC: // Graceful shutdown
			return true
		case tcell.KeyCtrlC: // Quick shutdown
			return false
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
		case tcell.KeyRune:
			switch ev.Rune() {
			case 'w': // Move north 1 unit
				e.view.RenderModel(e.model.MoveNorth(1))
			case 's': // Move south 1 unit
				e.view.RenderModel(e.model.MoveSouth(1))
			case 'a': // Move west 1 unit
				e.view.RenderModel(e.model.MoveWest(1))
			case 'd': // Move east 1 unit
				e.view.RenderModel(e.model.MoveEast(1))
			}
		}
	}
}

// Runs an outro sequence for the specified amount of time
func (e *InputExplorer) runOutro(duration time.Duration) {
	timer := time.NewTimer(duration)
	e.view.RenderOutro()
	for {
		select {
		case ev := <-e.input:
			switch ev.Key() {
			case tcell.KeyCtrlR: // Refresh view
				e.view.RenderOutro()
			default: // End outro
				return
			}
		case <-timer.C: // End outro
			return
		}
	}
}
