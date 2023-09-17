package internal

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// Represents a file type in relation to the application logic.
// This is not the file extension, but rather the purpose of the file.
type FileType string

const (
	MetadataFile FileType = "metadata"
	ConfigFile   FileType = "config"
)

// Represents a file operation.
// It is a verb that describes what is being done to the file.
type FileOperation string

const (
	DeleteOp FileOperation = "delete"
	ResetOp  FileOperation = "reset"
)

type fileContext string

const (
	CollectionCtx fileContext = "collection"
	AppCtx        fileContext = "app"
)

// Performs a destructive operation on a file based on the value of fileOp.
// If skipConfirm is not true, the user will be prompted before performing the destructive operation.
func DestructiveFileOp(path string, fileType FileType, skipConfirm bool, fileOp FileOperation) error {
	if err := validateFileType(fileType); err != nil {
		return fmt.Errorf("Failed to '%s' file '%s' of type '%s': %w", fileOp, path, fileType, err)
	}

	filename := filepath.Base(path)

	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		fmt.Printf("'%s' %s file does not exist\n", filename, fileType)
		return nil
	} else if err != nil {
		return fmt.Errorf("Failed to check if '%s' file exists: %w", path, err)
	}

	var operationFunc func(string, FileType) error
	switch fileOp {
	case ResetOp:
		operationFunc = resetFile
	case DeleteOp:
		operationFunc = deleteFile
	}

	performOpFunc := func() error {
		if err := operationFunc(path, fileType); err != nil {
			return fmt.Errorf("Failed to %s '%s' file: %w", fileOp, path, err)
		}
		return nil
	}

	if skipConfirm {
		return performOpFunc()
	}

	fmt.Printf("Are you sure you want to %s '%s'? [y/N] ", fileOp, filename)
	var response string
	if _, err = fmt.Scanln(&response); err != nil {
		return fmt.Errorf("Failed to read user input: %v", err)
	}
	if response == "y" || response == "Y" {
		return performOpFunc()
	} else {
		fmt.Printf("%s file was not affected\n", filename)
	}
	return nil
}

// Creates a file if it does not exist.
//
// The fileType parameter is used for logging purposes and to determine the default contents of the file. It should be one of the constants defined in this package.
func initFile(path string, fileType FileType) error {
	if err := validateFileType(fileType); err != nil {
		return fmt.Errorf("Failed to initialize file '%s' of type '%s': %w", path, fileType, err)
	}

	filename := filepath.Base(path)

	_, err := os.Stat(path)
	switch {
	case errors.Is(err, fs.ErrNotExist):
		pathDirs := filepath.Dir(path)
		if _, err := os.Stat(pathDirs); errors.Is(err, fs.ErrNotExist) {
			return fmt.Errorf("Required directory '%s' does not exist: %w", pathDirs, err)
		}

		fmt.Printf("Creating '%s' %s file...\n", filename, fileType)
		if err = createFile(path, fileType, CollectionCtx); err != nil {
			return fmt.Errorf("Failed to create '%s' %s file: %w", path, fileType, err)
		}
		fmt.Printf("Created '%s' %s file\n", filename, fileType)
	case err == nil:
		fmt.Printf("'%s' %s file already exists\n", filename, fileType)
	default:
		return fmt.Errorf("Failed to check if '%s' %s file exists: %w", path, fileType, err)
	}
	return nil
}

// Resets a file to its default if it exists.
//
// The fileType parameter is used for logging purposes and to determine the default contents of the file. It should be one of the constants defined in this package.
func resetFile(path string, fileType FileType) error {
	if err := validateFileType(fileType); err != nil {
		return fmt.Errorf("Failed to reset file '%s' of type '%s': %w", path, fileType, err)
	}

	filename := filepath.Base(path)

	_, err := os.Stat(path)
	switch {
	case errors.Is(err, fs.ErrNotExist):
		fmt.Printf("'%s' %s file does not exist\n", filename, fileType)
	case err == nil:
		fmt.Printf("Found '%s' %s file. Resetting it...\n", filename, fileType)
		if err = createFile(path, fileType, CollectionCtx); err != nil {
			return fmt.Errorf("Failed to reset '%s' %s file: %w", path, fileType, err)
		}
		fmt.Printf("Reset '%s' %s file\n", filename, fileType)
	default:
		return fmt.Errorf("Failed to check if '%s' %s file exists: %w", path, fileType, err)
	}
	return nil
}

// Deletes a file if it exists.
//
// The fileType parameter is used for logging purposes. It should be one of the constants defined in this package.
func deleteFile(path string, fileType FileType) error {
	if err := validateFileType(fileType); err != nil {
		return fmt.Errorf("Failed to remove file '%s' of type '%s': %w", path, fileType, err)
	}

	filename := filepath.Base(path)

	fmt.Printf("Removing '%s' %s file...\n", filename, fileType)
	if err := os.Remove(path); err != nil {
		return fmt.Errorf("Failed to remove '%s' %s file: %v", path, fileType, err)
	}
	fmt.Printf("Removed '%s' %s file\n", filename, fileType)
	return nil
}

// Creates a file with the default contents based on the fileType parameter.
// The
func createFile(path string, fileType FileType, context fileContext) error {
	var srcFile, dstFile *os.File
	var err error

	assetsPath := "assets"

	switch fileType {
	case MetadataFile:
		srcFile, err = os.Open(filepath.Join(assetsPath, string(context)+"-metadata.yaml"))
	case ConfigFile:
		srcFile, err = os.Open(filepath.Join(assetsPath, string(context)+"-config.toml"))
	default:
		return fmt.Errorf("Invalid file type '%s'", fileType)
	}

	if err != nil {
		return fmt.Errorf("Failed to open %s %s file '%s': %w", context, fileType, path, err)
	}
	defer srcFile.Close()

	err = os.MkdirAll(filepath.Dir(path), os.ModePerm)
	if err != nil {
		return fmt.Errorf("Failed to create directory for '%s' %s file: %w", path, fileType, err)
	}

	dstFile, err = os.Create(path)
	if err != nil {
		return fmt.Errorf("Failed to create '%s' %s file: %w", path, fileType, err)
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return fmt.Errorf("Failed to copy contents of '%s' %s file to '%s': %w", srcFile.Name(), fileType, dstFile.Name(), err)
	}
	return nil
}

// Checks if the fileType parameter is one of the constants defined in this package.
func validateFileType(fileType FileType) error {
	switch fileType {
	case MetadataFile, ConfigFile:
		return nil
	default:
		return fmt.Errorf("Invalid file type '%s'", fileType)
	}
}

// Checks if the fileOp parameter is one of the constants defined in this package.
func validateFileOperation(fileOp FileOperation) error {
	switch fileOp {
	case ResetOp, DeleteOp:
		return nil
	default:
		return fmt.Errorf("Invalid file operation '%s'", fileOp)
	}
}

// Checks if a path should be ignored.
func shouldIgnore(path string) bool {
	// TODO: add the ability to have a gitignore-like file that will ignore certain paths or directories
	filename := filepath.Base(path)

	if strings.HasPrefix(filename, ".") || path == MetadataFilename || path == ConfigFilename {
		return true
	}
	return false
}