package cmd

import (
	"fmt"

	"github.com/Broderick-Westrope/barber/internal"
	"github.com/spf13/cobra"
)

var snippetCmd = &cobra.Command{
	Use:     "snippet",
	Aliases: []string{"snpt"},
	Short:   "Commands for managing snippets",
	Run:     displayHelp,
}

var snptAddCmd = &cobra.Command{
	Use:     "add path",
	Aliases: []string{"a"},
	Short:   "Add a snippet",
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

var snptRemoveCmd = &cobra.Command{
	Use:     "remove path",
	Aliases: []string{"rm"},
	Short:   "Remove a snippet",
	Long: `Remove a snippet from the collection. If multiple snippets have the same path, they will all be removed.

	Positional Arguments:
	  path:	The path to the snippet file you want to remove from the collection. This is relative to the collection root directory.`,
	Args:      cobra.MinimumNArgs(1),
	ValidArgs: []string{"snippet"},
	Run: func(cmd *cobra.Command, args []string) {
		err := internal.RemoveSnippet(collectionFlag, args[0])
		if err != nil {
			fmt.Printf("Error: %s\n", err)
		}
	},
}
