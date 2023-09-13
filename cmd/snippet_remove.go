package cmd

import (
	"log"
	"path/filepath"

	"github.com/Broderick-Westrope/barber/internal"
	"github.com/spf13/cobra"
)

var snptRemoveCmd = &cobra.Command{
	Use:  "remove path",
	Aliases: []string{"rm"},
	Short: "Remove a snippet",
	Long: `Remove a snippet from the collection. If multiple snippets have the same path, they will all be removed.

	Positional Arguments:
	  path:	The path to the snippet file you want to remove from the collection. This is relative to the collection root directory.`,
	Args: cobra.MinimumNArgs(1),
	ValidArgs: []string{"snippet"},
	Run: func(cmd *cobra.Command, args []string) {
		removeSnippet(collectionPath, args[0])
	},
}

func removeSnippet(colPath string, snptPath string){
	exists, err := internal.IsCollection(colPath)
	if err != nil {
		log.Fatalf("Failed to check if '%s' is a collection: %v", colPath, err)
	} else if !exists {
		log.Fatalf("'%s' is not a collection", colPath)
	}

	metadataPath := filepath.Join(colPath, internal.MetadataFilename)
	metadata, err := internal.GetMetadata(metadataPath)
	if err != nil {
		log.Fatalf("Failed to get metadata at '%s': %v", metadataPath, err)
	}

	count := 0
	j := 0
	for i, snpt := range metadata.Snippets {
		if snpt.Path == snptPath {
			count++
			continue
		}
		metadata.Snippets[j] = metadata.Snippets[i]
		j++
	}
	metadata.Snippets = metadata.Snippets[:j]

	if count == 0 {
		log.Fatalf("No snippets with path '%s' found in collection '%s'", snptPath, colPath)
	}

	err = internal.WriteMetadata(metadataPath, metadata)
	if err != nil {
		log.Fatalf("Failed to write metadata at '%s': %v", metadataPath, err)
	}

	plural := ""
	if count > 1 {
		plural = "s"
	}
	log.Printf("Removed %d snippet%s '%s' from collection '%s'", count, plural, snptPath, colPath)
}