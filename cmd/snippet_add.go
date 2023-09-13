package cmd

import (
	"log"
	"path/filepath"

	"github.com/Broderick-Westrope/barber/internal"
	"github.com/spf13/cobra"
)

var snptAddCmd = &cobra.Command{
	Use:  "add",
	Aliases: []string{"a"},
	Short: "Add a snippet",
	Long: `Add a snippet to the collection`,
	Args: cobra.MinimumNArgs(1),
	ValidArgs: []string{"collection", "snippet"},
	Run: func(cmd *cobra.Command, args []string) {
		addSnippet(collectionPath, args[0])
	},
}

func addSnippet(colPath string, snptPath string){
	exists, err := internal.IsCollection(colPath)
	if err != nil {
		log.Fatalf("Failed to check if '%s' is a collection: %v", colPath, err)
	} else if !exists {
		log.Fatalf("'%s' is not a collection", colPath)
	}

	exists, err = internal.FileExists(snptPath)
	if err != nil {
		log.Fatalf("Failed to check if '%s' is a file: %v", snptPath, err)
	} else if !exists {
		log.Fatalf("'%s' is not a file", snptPath)
	}

	metadataPath := filepath.Join(colPath, internal.MetadataFilename)
	metadata, err := internal.GetMetadata(metadataPath)
	if err != nil {
		log.Fatalf("Failed to get metadata at '%s': %v", metadataPath, err)
	}

	newSnippet := internal.Snippet{
		Path: snptPath,
		Description: "",
		Tags: []string{},
	}

	metadata.Snippets = append(metadata.Snippets, newSnippet)

	err = internal.WriteMetadata(metadataPath, metadata)
	if err != nil {
		log.Fatalf("Failed to write metadata at '%s': %v", metadataPath, err)
	}

	log.Printf("Added snippet '%s' to collection '%s'", snptPath, colPath)
}