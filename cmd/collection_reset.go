package cmd

import (
	"fmt"

	"github.com/Broderick-Westrope/barber/internal"
	"github.com/spf13/cobra"
)

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