package internal_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/Broderick-Westrope/barber/internal"
	"github.com/go-git/go-git/v5"
)

func TestInitCollection(t *testing.T) {
	// Setup a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "test-collection-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)
	
	// Initialise test collection
	err = internal.InitCollection(tempDir)
	if err != nil {
		t.Fatalf("Failed to initialise collection: %v", err)
	}
	
	// Check Git repository initialisation
	_, err = git.PlainOpen(tempDir)
	if err == git.ErrRepositoryNotExists || err == git.ErrRepositoryIncomplete {
		t.Fatalf("Git repository does not exist: %v", err)
	}

	// Check metadata file creation
	metadataPath := filepath.Join(tempDir, internal.MetadataFilename)
	if _, err := os.Stat(metadataPath); os.IsNotExist(err) {
		t.Fatalf("Metadata file does not exist: %v", err)
	}

	// Check config file creation
	configPath := filepath.Join(tempDir, internal.ConfigFilename)
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		t.Fatalf("Config file does not exist: %v", err)
	}

	// Check whether an error is returned when the collection already exists
	err = internal.InitCollection(tempDir)
	if err != nil {
		t.Fatalf("Failed to initialise collection when files already exist: %v", err)
	}
}