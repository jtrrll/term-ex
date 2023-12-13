package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Explore in response to keyboard inputs
var keyboardCmd = &cobra.Command{
	Use:   "keyboard",
	Short: "Explore in response to keyboard inputs",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO: Implement
		fmt.Println("keyboard called")
		return nil
	},
}

// Defines hierarchy and flags for command
func init() {
	rootCmd.AddCommand(keyboardCmd)
}
