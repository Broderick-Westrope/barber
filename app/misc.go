package app

import (
	"fmt"
	"io/fs"
	"path/filepath"

	"github.com/Broderick-Westrope/barber/file"
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
		err := scan(collectionFlag, keepFlag, dryRunFlag)
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
		err := list(collectionFlag, includeMetadataFlag)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
		}
	},
}

func scan(colPath string, shouldKeep, isDryRun bool) error {
	res, err := file.IsCollection(colPath)
	if err != nil {
		return fmt.Errorf("Failed to check if '%s' is a collection: %w", colPath, err)
	} else if !res {
		return fmt.Errorf("Cannot alter metadata because '%s' is not a collection", colPath)
	}

	var files []string

	err = filepath.WalkDir(colPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Skip the colPath directory itself
		if path == colPath {
			return nil
		}

		if internal.ShouldIgnore(path) {
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
		return fmt.Errorf("Failed to scan collection '%s': %v", colPath, err)
	}

	// Print the collected paths
	fmt.Printf("Found %d files:\n", len(files))
	for _, file := range files {
		fmt.Println("+", file)
	}

	if isDryRun {
		fmt.Println("\nDry run complete.")
		return nil
	}

	// Update the metadata
	err = file.UpdateMetadata(filepath.Join(colPath, file.MetadataFilename), &files, shouldKeep)
	if err != nil {
		return fmt.Errorf("Failed to update metadata for collection '%s': %v", colPath, err)
	}

	fmt.Println("\nScan complete. Metadata updated.")
	return nil
}

func list(colPath string, includeMetadata bool) error {
	res, err := file.IsCollection(colPath)
	if err != nil {
		return fmt.Errorf("Failed to check if '%s' is a collection: %w", colPath, err)
	} else if !res {
		return fmt.Errorf("Cannot list snippets because '%s' is not a collection", colPath)
	}

	// Read the metadata
	metadataFilepath := filepath.Join(colPath, file.MetadataFilename)
	metadata, err := file.GetMetadata(metadataFilepath)
	if err != nil {
		return fmt.Errorf("Failed to read collection metadata file '%s': %v", metadataFilepath, err)
	}

	// Print the snippets
	if len(metadata.Snippets) == 0 {
		fmt.Println("No snippets found.")
		return nil
	}

	fmt.Printf("Found %d snippets:\n", len(metadata.Snippets))
	for _, snippet := range metadata.Snippets {
		msg := fmt.Sprintf("  - %s", snippet.Path)
		if !includeMetadata {
			fmt.Println(msg)
			continue
		}
		if snippet.Desc != "" {
			msg += fmt.Sprintf(" | %s", snippet.Desc)
		}
		if len(snippet.Tags) > 0 {
			msg += " | "
			for i, tag := range snippet.Tags {
				if i > 0 {
					msg += ", "
				}
				msg += tag
			}
		}
		fmt.Println(msg)
	}

	return nil
}
