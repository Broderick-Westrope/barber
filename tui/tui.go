package tui

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Broderick-Westrope/barber/file"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type Directory struct {
	Path     string
	Snippets []file.Snippet
	SubDirs  []Directory
	Parent   *Directory
}

func (d Directory) FilterValue() string {
	return d.Path
}

func GetItems[T list.Item](items []T) []list.Item {
	dirItems := make([]list.Item, len(items))
	for i, v := range items {
		dirItems[i] = list.Item(v)
	}
	return dirItems
}

// Starts the interactive terminal user interface (TUI)
func Run(colPath string) error {
	// Get snippets from collection path
	metadata, err := file.GetMetadata(filepath.Join(colPath, file.MetadataFilename))
	if err != nil {
		return err
	}

	// Put the snippets into a tree structure
	root := Directory{
		Path:   ".",
		Parent: nil,
	}

	for _, snpt := range metadata.Snippets {
		err = addSnippetToDir(&root, snpt)
		if err != nil {
			return fmt.Errorf("error adding snippet %s to directory: %w", snpt.Path, err)
		}
	}

	// Create & run the tea program
	m := NewModel(&root)
	p := tea.NewProgram(m, tea.WithAltScreen())
	var model tea.Model
	if model, err = p.Run(); err != nil {
		return fmt.Errorf("error running tea program: %w", err)
	}

	// Cast the model back to custom model type
	m, ok := model.(*Model)
	if !ok {
		return fmt.Errorf("error casting the response model to custom type: %w", err)
	}

	//TODO: Store the altered snippets

	return nil
}

func addSnippetToDir(root *Directory, snippet file.Snippet) error {
	// Check that the snippet is actually a file
	fileInfo, err := os.Stat(snippet.Path)
	switch {
	case err != nil:
		return err
	case fileInfo.IsDir():
		return fmt.Errorf("snippet %s is a directory", snippet.Path)
	}

	parts := strings.Split(snippet.Path, "/")
	currentDir := root

	for i, part := range parts {
		if i == len(parts)-1 {
			// This is a file, not a directory
			snippet.Path = part
			currentDir.Snippets = append(currentDir.Snippets, snippet)
		} else {
			// Look for existing sub-directory
			var subDir *Directory
			for i := range currentDir.SubDirs {
				if currentDir.SubDirs[i].Path == part {
					subDir = &currentDir.SubDirs[i]
					break
				}
			}

			// Create new sub-directory if it doesn't exist
			if subDir == nil {
				newDir := Directory{
					Path:    part,
					SubDirs: []Directory{},
				}
				currentDir.SubDirs = append(currentDir.SubDirs, newDir)
				subDir = &currentDir.SubDirs[len(currentDir.SubDirs)-1] // Get reference to last appended item
			}

			currentDir = subDir
		}
	}

	return nil
}
