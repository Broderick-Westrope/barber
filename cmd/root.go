package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Flags. These are set in the init() function of each command.
var skipConfirm bool

const (
	defMetadataFile = ".barber.yaml"
	defConfigFile   = ".barber.toml"
)

var rootCmd = &cobra.Command{
	Use:   "barber",
	Short: "Barber is a tool for managing snippets",
	Long: `Barber is a CLI & TUI for managing snippets.
			It is targeted towards code snippets, but can be used for any text format.
			Documentation is available at https://github.com/Broderick-Westrope/barber`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Launch interactive mode
		fmt.Println("Interactive mode not yet implemented")
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
