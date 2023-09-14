package cmd

import (
	"fmt"

	"github.com/Broderick-Westrope/barber/internal"
	"github.com/spf13/cobra"
)

// Adds a snippet to the collection.
var snptAddCmd = &cobra.Command{
	Use:  "add path",
	Aliases: []string{"a"},
	Short: "Add a snippet",
	Long: `Add a snippet to the collection.

Positional Arguments:
  path:	The path to the snippet file you want to add to the collection. This is relative to the collection root directory.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := internal.AddSnippet(collectionFlag, args[0])
		if err != nil {
			fmt.Printf("Error: %s\n", err)
		}
	},
}