package internal_test

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/Broderick-Westrope/barber/internal"
)

func TestScan(t *testing.T) {
	// Setup a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "barber-test-init-collection-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Test that an error is returned if the path is not a collection
	err = internal.Scan(filepath.Join(tempDir, "not/a/dir"), false, false)
	if err == nil {
		t.Errorf("Scan error: got %v, want %v", err, "not nil")
	}

	// Initialise test collection
	err = internal.InitCollection(tempDir)
	if err != nil {
		t.Fatalf("Failed to initialise collection: %v", err)
	}

	// Create a file in the collection
	filePath := filepath.Join(tempDir, "test.txt")
	_, err = os.Create(filePath)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// Get the metadata
	metadata, err := internal.GetMetadata(filepath.Join(tempDir, internal.MetadataFilename))
	if err != nil {
		t.Fatalf("Failed to get metadata: %v", err)
	}

	// Test that metadata is empty
	if len(metadata.Snippets) != 0 {
		t.Errorf("Metadata Snippets Length: got %v, want %v", metadata.Snippets, 0)
	}

	// Test that no error is returned when the path is a collection
	err = internal.Scan(tempDir, false, false)
	if err != nil {
		t.Fatalf("should not have returned an error when the path is a collection")
	}

	// Test that the metadata is updated
	metadata, err = internal.GetMetadata(filepath.Join(tempDir, internal.MetadataFilename))
	expectedSnippet := internal.Snippet{
		Path:        filepath.Base(filePath),
		Tags:        []string{},
		Description: "",
	}
	if !reflect.DeepEqual(metadata.Snippets, []internal.Snippet{expectedSnippet}) {
		t.Errorf("Metadata Snippets: got %v, want %v", metadata.Snippets, []internal.Snippet{expectedSnippet})
	}

	// Remove the test file
	err = os.Remove(filePath)
	if err != nil {
		t.Fatalf("Failed to remove test file: %v", err)
	}

	// Scan the collection again with the shouldKeep flag set to true
	err = internal.Scan(tempDir, true, false)
	if err != nil {
		t.Fatal("should not have returned an error when the path is a collection (shouldKeep = true))")
	}

	// Test that the metadata is not updated
	metadata, err = internal.GetMetadata(filepath.Join(tempDir, internal.MetadataFilename))
	if !reflect.DeepEqual(metadata.Snippets, []internal.Snippet{expectedSnippet}) {
		t.Errorf("Metadata Snippets: got %v, want %v", metadata.Snippets, []internal.Snippet{expectedSnippet})
	}

	// Scan the collection again with the isDryRun flag set to true
	err = internal.Scan(tempDir, false, true)
	if err != nil {
		t.Fatal("should not have returned an error when the path is a collection (isDryRun = true)")
	}

	// Test that the metadata is not updated
	metadata, err = internal.GetMetadata(filepath.Join(tempDir, internal.MetadataFilename))
	if !reflect.DeepEqual(metadata.Snippets, []internal.Snippet{expectedSnippet}) {
		t.Errorf("Metadata Snippets: got %v, want %v", metadata.Snippets, []internal.Snippet{expectedSnippet})
	}

	// Scan the collection again with the shouldKeep flag set to false
	err = internal.Scan(tempDir, false, false)
	if err != nil {
		t.Fatal("should not have returned an error when the path is a collection (no flags)")
	}

	// Test that the metadata is updated
	metadata, err = internal.GetMetadata(filepath.Join(tempDir, internal.MetadataFilename))
	if !reflect.DeepEqual(metadata.Snippets, []internal.Snippet{}) {
		t.Errorf("Metadata Snippets: got %v, want %v", metadata.Snippets, []internal.Snippet{})
	}
}
