package controller

import (
	"fmt"
	"term-ex/model"
	"term-ex/view"
	"time"
)

// Explores a world based on a predefined strategy.
// A delay can be specified to sleep between actions
type AutoExplorer struct {
	model    model.Model
	view     view.View
	strategy func(model.Model) Command
	delay    time.Duration
}

// Creates a new auto-explorer that explores based on a predefined strategy
func NewAutoExplorer(model model.Model, view view.View, strategy func(model.Model) Command, delay time.Duration) *AutoExplorer {
	return &AutoExplorer{model, view, strategy, delay}
}

// Explores the model based on the predefined strategy, waiting for a set period of time between actions.
// Panics if no strategy is defined
func (e *AutoExplorer) Explore() error {
	if e.strategy == nil {
		panic("exploration strategy not defined")
	}
	fmt.Println("exploring!")
	//TODO: Implement
	return nil
}
