package internal

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

const MetadataFilename = ".barber.yaml"

type Metadata struct {
	Snippets []Snippet `yaml:"snippets"`
}

type Snippet struct {
	Path        string   `yaml:"path"`
	Description string   `yaml:"description"`
	Tags        []string `yaml:"tags"`
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