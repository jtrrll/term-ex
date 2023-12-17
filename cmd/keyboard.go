package cmd

import (
	"term-ex/controller"
	"term-ex/model"
	"term-ex/model/world"
	"term-ex/position"
	"term-ex/tile"
	"term-ex/view"

	"github.com/gdamore/tcell"
	"github.com/spf13/cobra"
)

// Explore in response to keyboard inputs
var keyboardCmd = &cobra.Command{
	Use:   "keyboard",
	Short: "Explore in response to keyboard inputs",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO: Implement
		noFog, flagErr := cmd.Flags().GetBool("no-fog")
		if flagErr != nil {
			panic(flagErr)
		}
		screen, newScreenErr := tcell.NewScreen()
		if newScreenErr != nil {
			panic(newScreenErr)
		}
		screenInitErr := screen.Init()
		if screenInitErr != nil {
			panic(screenInitErr)
		}
		inputChan := make(chan *tcell.EventKey)
		go func() {
			for {
				ev := screen.PollEvent()
				eventKey, ok := ev.(*tcell.EventKey)
				if ok {
					inputChan <- eventKey
					continue
				}
				_, ok = ev.(*tcell.EventResize)
				if ok {
					inputChan <- tcell.NewEventKey(tcell.KeyCtrlR, 'r', tcell.ModCtrl)
				}
			}
		}()
		model := model.NewRuleBasedModel(world.Rules{}, tile.Tile{Terrain: 0}, position.Position{}, tile.Tile{Terrain: 1}, 3, 0)
		view := view.NewTextView(screen, !noFog, 10, false, false)
		return controller.NewInputExplorer(inputChan, model, view).Explore()
	},
}

// Defines hierarchy and flags for command
func init() {
	rootCmd.AddCommand(keyboardCmd)
}
