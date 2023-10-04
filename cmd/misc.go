package cmd

import (
	"fmt"

	"github.com/Broderick-Westrope/barber/internal"
	"github.com/spf13/cobra"
)

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scans a collection for snippets",
	Long: `Goes through all paths & directories in the collection, and adds them as snippets to the collection metadata.
By default, deleted paths & directories will be removed from the metadata. This can be altered with the --keep flag.
If a file or directory already exists in the collection, it will not be re-added.
If a file or directory has been renamed or moved, it will attempt to update the entry in the metadata.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := internal.Scan(collectionFlag, keepFlag, dryRunFlag)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
		}
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists the (relative) path of each snippet in a collection",
	Long: `Lists the relative path of each snippet in a collection by reading the collection metadata file.
This operation does not make any changes to the collection or its metadata.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := internal.List(collectionFlag)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
		}
	},
}