package internal

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

func Scan(colPath string, shouldKeep, isDryRun bool) error {
	res, err := IsCollection(colPath)
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
		return fmt.Errorf("Failed to scan collection '%s': %v", colPath, err)
	}

	// Print the collected paths
	fmt.Printf("Found %d files:\n", len(files))
	for _, file := range files {
		fmt.Println("+",file)
	}

	if isDryRun {
		fmt.Println("\nDry run complete.")
		return nil
	}

	// Update the metadata
	err = UpdateMetadata(filepath.Join(colPath, metadataFilename), &files, shouldKeep)
	if err != nil {
		return fmt.Errorf("Failed to update metadata for collection '%s': %v", colPath, err)
	}

	fmt.Println("\nScan complete. Metadata updated.")
	return nil
}