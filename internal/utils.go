package internal

import (
	"path/filepath"
	"strings"

	"github.com/Broderick-Westrope/barber/file"
)

// Checks if a path should be ignored.
func ShouldIgnore(path string) bool {
	// TODO: add the ability to have a gitignore-like file that will ignore certain paths or directories
	filename := filepath.Base(path)

	if strings.HasPrefix(filename, ".") || path == file.MetadataFilename || path == file.ConfigFilename {
		return true
	}
	return false
}
