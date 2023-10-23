package internal

import (
	"fmt"
	"path/filepath"
)

// AddSnippet adds a snippet to the collection.
// colPath is the absolute path to the collection.
// snptPath is the path to the snippet file you want to add to the collection.
// It is relative to the colPath.
func AddSnippet(colPath string, snptPath string) error {
	exists, err := IsCollection(colPath)
	if err != nil {
		return fmt.Errorf("Failed to check if '%s' is a collection: %v", colPath, err)
	} else if !exists {
		return fmt.Errorf("'%s' is not a collection", colPath)
	}

	colSnptPath := filepath.Join(colPath, snptPath)
	exists, err = FileExists(colSnptPath)
	if err != nil {
		return fmt.Errorf("Failed to check if '%s' is a file: %v", colSnptPath, err)
	} else if !exists {
		return fmt.Errorf("'%s' is not a file", colSnptPath)
	}

	metadataPath := filepath.Join(colPath, MetadataFilename)
	metadata, err := GetMetadata(metadataPath)
	if err != nil {
		return fmt.Errorf("Failed to get metadata at '%s': %v", metadataPath, err)
	}

	newSnippet := Snippet{
		Path: snptPath,
		Desc: "",
		Tags: []string{},
	}

	metadata.Snippets = append(metadata.Snippets, newSnippet)

	err = WriteMetadata(metadataPath, metadata)
	if err != nil {
		return fmt.Errorf("Failed to write metadata at '%s': %v", metadataPath, err)
	}

	fmt.Printf("Added snippet '%s' to collection '%s'\n", snptPath, colPath)
	return nil
}

func RemoveSnippet(colPath string, snptPath string) error {
	exists, err := IsCollection(colPath)
	if err != nil {
		return fmt.Errorf("Failed to check if '%s' is a collection: %v", colPath, err)
	} else if !exists {
		return fmt.Errorf("'%s' is not a collection", colPath)
	}

	metadataPath := filepath.Join(colPath, MetadataFilename)
	metadata, err := GetMetadata(metadataPath)
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

	err = WriteMetadata(metadataPath, metadata)
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
