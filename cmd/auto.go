package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Begins exploring with an auto-explorer
var autoCmd = &cobra.Command{
	Use:   "auto",
	Short: "Automatically explores a world based on a predefined strategy",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO: Implement
		fmt.Println("auto called")
		return nil
	},
}

// Defines hierarchy and flags for command
func init() {
	rootCmd.AddCommand(autoCmd)
	autoCmd.Flags().DurationP("delay", "d", 0, "set millisecond delay between actions")
	autoCmd.Flags().String("strategy", "greedy", "select an exploration strategy")
}
