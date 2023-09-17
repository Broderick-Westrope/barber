package internal

import (
	"fmt"
	"path/filepath"
)

const (
	MetadataFilename = ".barber.yaml"
	ConfigFilename   = ".barber.toml"
)

// Initialises a new collection by creating a git repository, a metadata file, and a config file.
// If a git repository already exists, it will not be re-initialized.
// If a file already exists, it will not be re-created.
func InitCollection(colPath string) error {
	// TODO: Provide a config option to disable git repository creation
	if err := initGitRepo(colPath); err != nil {
		return fmt.Errorf("Failed to initialize git repository: %w", err)
	}

	metadataPath := filepath.Join(colPath, MetadataFilename)
	if err := initFile(metadataPath, MetadataFile); err != nil {
		return fmt.Errorf("Failed to initialize '%s' file: %w", metadataPath, err)
	}

	configPath := filepath.Join(colPath, ConfigFilename)
	if err := initFile(configPath, ConfigFile); err != nil {
		return fmt.Errorf("Failed to initialize '%s' file: %w", configPath, err)
	}

	return nil
}

// Removes a collection by removing the metadata file, and the config file.
// It does not affect the git repository.
// It uses skipConfirm to determine if the user should be prompted before removing the metadata files.
func RemoveCollection(colPath string, skipConfirm bool) error {
	metadataPath := filepath.Join(colPath, MetadataFilename)
	if err := DestructiveFileOp(metadataPath, MetadataFile, skipConfirm, DeleteOp); err != nil {
		return fmt.Errorf("Failed to delete '%s' file: %w", metadataPath, err)
	}

	configPath := filepath.Join(colPath, ConfigFilename)
	if err := DestructiveFileOp(configPath, ConfigFile, skipConfirm, DeleteOp); err != nil {
		return fmt.Errorf("Failed to delete '%s' file: %v", configPath, err)
	}

	return nil
}

// Resets a collection by resetting the metadata file, and the config file to their defaults.
// It does not affect the git repository.
// It uses skipConfirm to determine if the user should be prompted before resetting the files.
func ResetCollection(colPath string, skipConfirm bool) error {
	metadataPath := filepath.Join(colPath, MetadataFilename)
	if err := DestructiveFileOp(metadataPath, MetadataFile, skipConfirm, ResetOp); err != nil {
		return fmt.Errorf("Failed to reset '%s' file: %v", metadataPath, err)
	}

	configPath := filepath.Join(colPath, ConfigFilename)
	if err := DestructiveFileOp(configPath, ConfigFile, skipConfirm, ResetOp); err != nil {
		return fmt.Errorf("Failed to reset '%s' file: %v", configPath, err)
	}

	return nil
}
