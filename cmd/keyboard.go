package cmd

import (
	"term-ex/model"
	"term-ex/model/world"
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
		view := view.NewTextView(screen, !noFog, 10)
		view.RenderModel(model.NewImmutableModel(world.Rules{}, world.Tile{Terrain: 0}, world.Position{}, world.Tile{Terrain: 1}, 3, 0))
		return nil
	},
}

// Defines hierarchy and flags for command
func init() {
	rootCmd.AddCommand(keyboardCmd)
}
