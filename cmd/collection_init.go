package cmd

import (
	"fmt"

	"github.com/Broderick-Westrope/barber/internal"
	"github.com/spf13/cobra"
)

// Initialises a new collection by initialising a git repository, and creating the default metadata and config files.
var colInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new collection",
	Long: `Initialize a new collection by initialising a git repository, and creating the default metadata and config files.
			If a git repository already exists, it will not be re-initialized.
			If a file already exists, it will not be re-created.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := internal.InitCollection(collectionFlag)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
		}
	},
}