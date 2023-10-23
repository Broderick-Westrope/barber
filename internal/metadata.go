package internal

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Metadata struct {
	Snippets []Snippet `yaml:"snippets"`
}

type Snippet struct {
	Path string   `yaml:"path"`
	Desc string   `yaml:"description"`
	Tags []string `yaml:"tags"`
}

func GetMetadata(filepath string) (*Metadata, error) {
	// Read the existing YAML file into a Metadata object
	file, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("Error reading metadata YAML file: %w", err)
	}

	var metadata Metadata
	err = yaml.Unmarshal(file, &metadata)
	if err != nil {
		return nil, fmt.Errorf("Error parsing metadata YAML: %w", err)
	}
	return &metadata, nil
}

func WriteMetadata(filepath string, metadata *Metadata) error {
	file, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("Error creating metadata YAML file: %w", err)
	}
	defer file.Close()

	encoder := yaml.NewEncoder(file)
	defer encoder.Close()
	err = encoder.Encode(metadata)
	if err != nil {
		return fmt.Errorf("Error encoding metadata YAML: %w", err)
	}
	return nil
}

// TODO: Use go-git diff to find renamed & moved files
func UpdateMetadata(metadataPath string, files *[]string, shouldKeep bool) error {
	metadata, err := GetMetadata(metadataPath)
	if err != nil {
		return fmt.Errorf("failed to get metadata from '%s': %v", metadataPath, err)
	}

	// Create a map to keep track of existing paths in metadata
	existingPaths := make(map[string]bool)
	for _, snippet := range metadata.Snippets {
		existingPaths[snippet.Path] = true
	}

	// Create a map to keep track of new files
	newFiles := make(map[string]bool)
	for _, file := range *files {
		newFiles[file] = true
	}

	// Update the metadata object with the new files but don't overwrite existing entries and don't add duplicates
	var updatedSnippets []Snippet

	if shouldKeep {
		// Keep all existing snippets (even if they no longer exist)
		updatedSnippets = append(updatedSnippets, metadata.Snippets...)
	} else {
		// only include snippets that still exist
		for _, snippet := range metadata.Snippets {
			if newFiles[snippet.Path] {
				updatedSnippets = append(updatedSnippets, snippet)
			}
		}
	}

	for _, file := range *files {
		if !existingPaths[file] {
			newSnippet := Snippet{
				Path: file,
				Desc: "",
				Tags: []string{},
			}
			updatedSnippets = append(updatedSnippets, newSnippet)
		}
	}

	// Update the metadata's Snippets with the new, filtered list
	metadata.Snippets = updatedSnippets

	// Write the updated metadata back to the file
	err = WriteMetadata(metadataPath, metadata)
	if err != nil {
		return fmt.Errorf("failed to write updated metadata to '%s': %v", metadataPath, err)
	}

	return nil
}
