package file

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

// Initialises a new metadata file at the collection root (colPath).
// This is done by creating a new file the correct name and the default contents for the file purpose.
func InitMetadata(colPath string) error {
	colPath = filepath.Join(colPath, MetadataFilename)
	err := initFile(colPath, purposeMetadata)
	if err != nil {
		return fmt.Errorf("failed to initialise metadata file: %w", err)
	}
	return nil
}

// Initialises a new config file at the collection root (colPath).
// This is done by creating a new file the correct name and the default contents for the file purpose.
func InitConfig(colPath string) error {
	colPath = filepath.Join(colPath, ConfigFilename)
	err := initFile(colPath, purposeConfig)
	if err != nil {
		return fmt.Errorf("failed to initialise config file: %w", err)
	}
	return nil
}

// Creates a file if it does not exist.
func initFile(path string, filePurpose purpose) error {
	filename := filepath.Base(path)

	_, err := os.Stat(path)
	switch {
	case errors.Is(err, fs.ErrNotExist):
		pathDirs := filepath.Dir(path)
		if _, err := os.Stat(pathDirs); errors.Is(err, fs.ErrNotExist) {
			return fmt.Errorf("required path '%s' does not exist: %w", pathDirs, err)
		}

		if err = createFile(path, filePurpose, contextCollection); err != nil {
			return fmt.Errorf("failed to create '%s' file: %w", path, err)
		}
		fmt.Printf("Successfully created '%s' %s file\n", filename, filePurpose)
	case err == nil:
		fmt.Printf("'%s' file already exists\n", filename)
	default:
		return fmt.Errorf("failed to check if '%s' file exists: %w", path, err)
	}
	return nil
}

// Creates a file with the default contents based on the fileType parameter.
// The file will be created at the path specified by the path parameter.
// The path parameter should not include the filename; this is determined using the filePurpose parameter.
func createFile(path string, filePurpose purpose, fileContext context) error {
	var srcFile, dstFile *os.File
	var err error

	assetPath := filepath.Join(GetBasePath(), "assets")

	switch filePurpose {
	case purposeMetadata:
		assetPath = filepath.Join(assetPath, string(fileContext)+"-metadata.yaml")
		srcFile, err = os.Open(assetPath)
	case purposeConfig:
		assetPath = filepath.Join(assetPath, string(fileContext)+"-config.toml")
		srcFile, err = os.Open(assetPath)
	default:
		return fmt.Errorf("invalid file purpose '%s'", filePurpose)
	}

	if err != nil {
		return fmt.Errorf("failed to open %s %s file '%s': %w", fileContext, filePurpose, assetPath, err)
	}
	defer srcFile.Close()

	err = os.MkdirAll(filepath.Dir(path), os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create directory for '%s' file: %w", path, err)
	}

	dstFile, err = os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create '%s' file: %w", path, err)
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return fmt.Errorf("failed to copy contents of '%s' file to '%s': %w", srcFile.Name(), dstFile.Name(), err)
	}
	return nil
}
