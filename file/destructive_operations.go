package file

import (
	"fmt"
	"os"
	"path/filepath"
)

// Deletes the metadata file located at collection root (colPath) if it exists.
// It uses skipConfirm to determine if the user should be prompted before deleting the file.
func DeleteMetadata(colPath string, skipConfirm bool) error {
	colPath = filepath.Join(colPath, MetadataFilename)
	if err := destructiveFileOp(colPath, purposeMetadata, skipConfirm, OperationDelete); err != nil {
		return fmt.Errorf("failed to delete metadata file: %w", err)
	}
	return nil
}

// Deletes the config file located at collection root (colPath) if it exists.
// It uses skipConfirm to determine if the user should be prompted before deleting the file.
func DeleteConfig(colPath string, skipConfirm bool) error {
	colPath = filepath.Join(colPath, ConfigFilename)
	if err := destructiveFileOp(colPath, purposeConfig, skipConfirm, OperationDelete); err != nil {
		return fmt.Errorf("failed to delete config file: %w", err)
	}
	return nil
}

// Resets the metadata file located at collection root (colPath) if it exists.
// It uses skipConfirm to determine if the user should be prompted before resetting the file.
func ResetMetadata(colPath string, skipConfirm bool) error {
	colPath = filepath.Join(colPath, MetadataFilename)
	if err := destructiveFileOp(colPath, purposeMetadata, skipConfirm, OperationReset); err != nil {
		return fmt.Errorf("failed to reset metadata file: %w", err)
	}
	return nil
}

// Resets the config file located at collection root (colPath) if it exists.
// It uses skipConfirm to determine if the user should be prompted before resetting the file.
func ResetConfig(colPath string, skipConfirm bool) error {
	colPath = filepath.Join(colPath, ConfigFilename)
	if err := destructiveFileOp(colPath, purposeConfig, skipConfirm, OperationReset); err != nil {
		return fmt.Errorf("failed to reset config file: %w", err)
	}
	return nil
}

// Performs a destructive operation on a file.
// The value of fileOp defines the type of destructive operation to perform.
// The value of fileOp defines the type of destructive operation to perform.
// If skipConfirm is not true, the user will be prompted before performing the destructive operation.
func destructiveFileOp(path string, filePurpose purpose, skipConfirm bool, fileOp Operation) error {
	filename := filepath.Base(path)

	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		fmt.Printf("'%s' file does not exist\n", filename)
		return nil
	} else if err != nil {
		return fmt.Errorf("failed to check if '%s' file exists: %w", path, err)
	}

	var operationFunc func(string, purpose) error
	switch fileOp {
	case OperationReset:
		operationFunc = resetFile
	case OperationDelete:
		operationFunc = deleteFile
	}

	performOpFunc := func() error {
		if err := operationFunc(path, filePurpose); err != nil {
			return fmt.Errorf("failed to %s '%s' file: %w", fileOp, path, err)
		}
		return nil
	}

	if skipConfirm {
		return performOpFunc()
	}

	fmt.Printf("Are you sure you want to %s '%s'? [y/N] ", fileOp, filename)
	var response string
	if _, err = fmt.Scanln(&response); err != nil {
		return fmt.Errorf("failed to read user input: %w", err)
	}
	if response == "y" || response == "Y" {
		return performOpFunc()
	} else {
		fmt.Printf("'%s' file was not affected\n", filename)
	}
	return nil
}

// Resets a file to its default if it exists.
func resetFile(path string, filePurpose purpose) error {
	filename := filepath.Base(path)

	if err := createFile(path, filePurpose, contextCollection); err != nil {
		return fmt.Errorf("failed to reset '%s' file: %w", path, err)
	}
	fmt.Printf("Successfully reset '%s' %s file\n", filename, filePurpose)
	return nil
}

// Deletes a file if it exists.
func deleteFile(path string, filePurpose purpose) error {
	filename := filepath.Base(path)

	if err := os.Remove(path); err != nil {
		return fmt.Errorf("failed to delete '%s' file: %w", path, err)
	}
	fmt.Printf("Successfully deleted '%s' %s file\n", filename, filePurpose)
	return nil
}
