package cmd

import (
	"fmt"

	"github.com/Broderick-Westrope/barber/internal"
	"github.com/spf13/cobra"
)

var collectionCmd = &cobra.Command{
	Use:     "collection",
	Aliases: []string{"col"},
	Short:   "Commands for managing snippet collections",
	Run:     displayHelp,
}

var colInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialise a new collection",
	Long: `Initialise a new collection by initialising a git repository, and creating the default metadata and config files.
			If a git repository already exists, it will not be re-initialized.
			If a file already exists, it will not be re-created.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := internal.InitCollection(collectionFlag)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
		}
	},
}

var colResetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset a collection",
	Long: `Reset a collection by setting the metadata & config files to their default.
			This will not affect the git repository.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := internal.ResetCollection(collectionFlag, yesFlag)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
		}
	},
}

var colRemoveCmd = &cobra.Command{
	Use:     "remove",
	Aliases: []string{"rm"},
	Short:   "Remove a collection",
	Long: `Remove a collection by removing the metadata & config files.
			This will not remove the git repository.
			If a file does not exist, nothing will happen.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := internal.RemoveCollection(collectionFlag, yesFlag)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
		}
	},
}
