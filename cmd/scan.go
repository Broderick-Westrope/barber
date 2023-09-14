package cmd

import (
	"fmt"
	"io/fs"
	"log"
	"path/filepath"
	"strings"

	"github.com/Broderick-Westrope/barber/internal"
	"github.com/spf13/cobra"
)

var scanCmd = &cobra.Command{
	Use:  "scan",
	Aliases: []string{"a"},
	Short: "Scans a collection for snippets",
	Long: `Goes through all paths & directories in the collection, and adds them as snippets to the collection metadata.
By default, deleted paths & directories will be removed from the metadata. This can be altered with the --keep flag.
If a file or directory already exists in the collection, it will not be re-added.
If a file or directory has been renamed or moved, it will attempt to update the entry in the metadata.`,
	Run: func(cmd *cobra.Command, args []string) {
		res, err := internal.IsCollection(collectionFlag)
		if err != nil {
			log.Printf("Failed to check if '%s' is a collection: %v", collectionFlag, err)
			return
		} else if !res {
			log.Printf("Cannot alter metadata because '%s' is not a collection", collectionFlag)
			return
		}

		scan(collectionFlag, keepFlag, dryRunFlag)
	},
}

func scan(colPath string, shouldKeep, isDryRun bool){
	var files []string

	err := filepath.WalkDir(colPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Skip the colPath directory itself
		if path == colPath {
			return nil
		}

		if shouldIgnore(path) {
			if d.IsDir() {
				return filepath.SkipDir // Skip this directory and all subdirectories
			}
			return nil // Skip this file
		}

		// Skip directories but include files
		if !d.IsDir() {
			// Get the relative path
			relPath, err := filepath.Rel(colPath, path)
			if err != nil {
				return err
			}
			files = append(files, relPath)
		}
		return nil
	})

	if err != nil {
		log.Printf("Failed to scan collection '%s': %v", colPath, err)
		return
	}

	// Print the collected paths
	fmt.Printf("Found %d files:\n", len(files))
	for _, file := range files {
		fmt.Println("+",file)
	}

	if isDryRun {
		fmt.Println("\nDry run complete.")
		return
	}

	// Update the metadata
	err = internal.UpdateMetadata(filepath.Join(colPath, metadataFilename), &files, shouldKeep)
	if err != nil {
		log.Printf("Failed to update metadata for collection '%s': %v", colPath, err)
		return
	}

	fmt.Println("\nScan complete. Metadata updated.")
}

// TODO: add the ability to have a gitignore-like file that will ignore certain paths or directories
func shouldIgnore(path string) bool {
	filename := filepath.Base(path)

	if strings.HasPrefix(filename, ".") || path == metadataFilename || path == configFilename {
		return true
	}
	return false
}