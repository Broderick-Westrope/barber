package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// Flags. These are set in the init() function of each command.
var (
	skipConfirm bool
	collection  string
)

const (
	metadataFilename = ".barber.yaml"
	configFilename   = ".barber.toml"
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
	// Collection
	collectionCmd.PersistentFlags().StringVarP(&collection, "collection", "c", ".", "path to collection root directory")
	rootCmd.AddCommand(collectionCmd)

	// Collection Init
	collectionCmd.AddCommand(colInitCmd)

	// Collection Reset
	colResetCmd.Flags().BoolVarP(&skipConfirm, "yes", "y", false, "confirm removal without prompting")
	collectionCmd.AddCommand(colResetCmd)

	// Collection Remove
	colRemoveCmd.Flags().BoolVarP(&skipConfirm, "yes", "y", false, "confirm removal without prompting")
	collectionCmd.AddCommand(colRemoveCmd)

	cobra.CheckErr(rootCmd.Execute())
}

func displayHelp(cmd *cobra.Command, args []string) {
	if err := cmd.Help(); err != nil {
		log.Fatalf("Failed to display help: %v", err)
	}
}