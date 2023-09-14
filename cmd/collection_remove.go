package cmd

import (
	"fmt"

	"github.com/Broderick-Westrope/barber/internal"
	"github.com/spf13/cobra"
)

// Removes a collection by removing the metadata & config files.
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