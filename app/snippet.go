package app

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Broderick-Westrope/barber/file"
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
		err := addSnippet(collectionFlag, args[0])
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
		err := removeSnippet(collectionFlag, args[0])
		if err != nil {
			fmt.Printf("Error: %s\n", err)
		}
	},
}

// addSnippet adds a snippet to the collection.
// colPath is the absolute path to the collection.
// snptPath is the path to the snippet file you want to add to the collection.
// It is relative to the colPath.
func addSnippet(colPath string, snptPath string) error {
	exists, err := file.IsCollection(colPath)
	if err != nil {
		return fmt.Errorf("Failed to check if '%s' is a collection: %v", colPath, err)
	} else if !exists {
		return fmt.Errorf("'%s' is not a collection", colPath)
	}

	colSnptPath := filepath.Join(colPath, snptPath)
	exists, err = file.FileExists(colSnptPath)
	if err != nil {
		return fmt.Errorf("Failed to check if '%s' is a file: %v", colSnptPath, err)
	} else if !exists {
		return fmt.Errorf("'%s' is not a file", colSnptPath)
	}

	metadataPath := filepath.Join(colPath, file.MetadataFilename)
	metadata, err := file.GetMetadata(metadataPath)
	if err != nil {
		return fmt.Errorf("Failed to get metadata at '%s': %v", metadataPath, err)
	}

	newSnippet := file.Snippet{
		Path: snptPath,
		Desc: "",
		Tags: []string{},
	}

	metadata.Snippets = append(metadata.Snippets, newSnippet)

	err = file.WriteMetadata(metadataPath, metadata)
	if err != nil {
		return fmt.Errorf("Failed to write metadata at '%s': %v", metadataPath, err)
	}

	// Print the relative path to the collection from the user's home directory.
	msg := fmt.Sprintf("Added snippet '%s' to collection '%s'\n", snptPath, "%s")
	fullPath, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf(msg, colPath)
		return nil
	}
	absPath, err := filepath.Abs(colPath)
	if err != nil {
		fmt.Printf(msg, colPath)
		return nil
	}
	fullPath, err = filepath.Rel(fullPath, absPath)
	if err != nil {
		fmt.Printf(msg, colPath)
		return nil
	}
	fmt.Printf(msg, fullPath)
	return nil
}

func removeSnippet(colPath string, snptPath string) error {
	exists, err := file.IsCollection(colPath)
	if err != nil {
		return fmt.Errorf("Failed to check if '%s' is a collection: %v", colPath, err)
	} else if !exists {
		return fmt.Errorf("'%s' is not a collection", colPath)
	}

	metadataPath := filepath.Join(colPath, file.MetadataFilename)
	metadata, err := file.GetMetadata(metadataPath)
	if err != nil {
		return fmt.Errorf("Failed to get metadata at '%s': %v", metadataPath, err)
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
		return fmt.Errorf("No snippets with path '%s' found in collection '%s'", snptPath, colPath)
	}

	err = file.WriteMetadata(metadataPath, metadata)
	if err != nil {
		return fmt.Errorf("Failed to write metadata at '%s': %v", metadataPath, err)
	}

	plural := ""
	if count > 1 {
		plural = "s"
	}
	fmt.Printf("Removed %d snippet%s '%s' from collection '%s'\n", count, plural, snptPath, colPath)
	return nil
}
