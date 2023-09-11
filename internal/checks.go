package internal

import (
	"errors"
	"os"
	"path/filepath"
)

func IsCollection(colPath string) (bool, error) {
	metadataPath := filepath.Join(colPath, MetadataFilename)
	_, err := os.Stat(metadataPath)
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	} else if err != nil {
		return false, err
	}

	configPath := filepath.Join(colPath, ConfigFilename)
	_, err = os.Stat(configPath)
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return true, nil
}

func FileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return true, nil
}