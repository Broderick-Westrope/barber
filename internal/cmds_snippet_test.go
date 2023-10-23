package internal_test

import (
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/Broderick-Westrope/barber/internal"
)

func TestAddSnippet(t *testing.T) {
	// Setup a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "barber-test-add-snippet-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Define a file to be used later
	filePathBase := "test.txt"
	filePath := filepath.Join(tempDir, filePathBase)

	// Test that an error is returned if the path is not a collection
	err = internal.AddSnippet(tempDir, filePathBase)
	if err == nil {
		t.Errorf("AddSnippet error: got %v, want %v", err, "not nil")
	} else if !strings.Contains(err.Error(), "not a collection") {
		t.Errorf("AddSnippet error: got %v, want %v", err, "contains: 'not a collection'")
	}

	// Initialise test collection
	err = internal.InitCollection(tempDir)
	if err != nil {
		t.Fatalf("Failed to initialise collection: %v", err)
	}

	// Test that an error is returned if the collection path does not exist
	err = internal.AddSnippet(filepath.Join(tempDir, "not/a/dir"), filePathBase)
	if err == nil {
		t.Errorf("AddSnippet error: got %v, want %v", err, "not nil")
	}

	// Test that an error is returned if the snippet path does not exist
	err = internal.AddSnippet(tempDir, filePathBase)
	if err == nil {
		t.Errorf("AddSnippet error: got %v, want %v", err, "not nil")
	}

	// Create a file in the collection
	_, err = os.Create(filePath)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// Test that no error is returned when the path is a collection
	err = internal.AddSnippet(tempDir, filePathBase)
	if err != nil {
		t.Fatalf("should not have returned an error when the path is a collection: %v", err)
	}

	// Test that the snippet was added to the metadata
	metadata, err := internal.GetMetadata(filepath.Join(tempDir, internal.MetadataFilename))
	if err != nil {
		t.Fatalf("Failed to get metadata: %v", err)
	}
	expectedSnippet := internal.Snippet{
		Path: filePathBase,
		Tags: []string{},
		Desc: "",
	}
	if !reflect.DeepEqual(metadata.Snippets, []internal.Snippet{expectedSnippet}) {
		t.Errorf("Metadata Snippets: got %v, want %v", metadata.Snippets, []internal.Snippet{expectedSnippet})
	}
}
