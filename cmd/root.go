package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// The base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "term-ex",
	Short: "Explore an ASCII world in your terminal!",
}

// Adds all child commands and flags to the root command
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// Defines global flags
func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.SetHelpCommand(&cobra.Command{Hidden: true})

	rootCmd.PersistentFlags().Uint64P("seed", "s", 0, "set seed for random world generation")
	rootCmd.PersistentFlags().Uint8P("radius", "r", 20, "set radius for fog-of-war that obscures unexplored terrain")
	rootCmd.PersistentFlags().Bool("no-fog", false, "disable fog-of-war that obscures unexplored terrain")
}
