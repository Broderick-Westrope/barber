package file

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func GetBasePath() string {
	_, b, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(b), "..")
}

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

// Checks if the fileType parameter is one of the constants defined in this package.
func validateFilePurpose(filePurpose purpose) error {
	switch filePurpose {
	case purposeMetadata, purposeConfig:
		return nil
	default:
		return fmt.Errorf("Invalid file type '%s'", filePurpose)
	}
}

// Checks if the fileOp parameter is one of the constants defined in this package.
func validateFileOperation(fileOp Operation) error {
	switch fileOp {
	case OperationReset, OperationDelete:
		return nil
	default:
		return fmt.Errorf("Invalid file operation '%s'", fileOp)
	}
}
