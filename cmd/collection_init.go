package cmd

import (
	"errors"
	"fmt"
	"log"
	"path/filepath"

	"github.com/Broderick-Westrope/barber/internal"
	"github.com/go-git/go-git/v5"
	"github.com/spf13/cobra"
)

var colInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new collection",
	Long: `Initialize a new collection by initialising a git repository, and creating the default metadata and config files.
			If a git repository already exists, it will not be re-initialized.
			If a file already exists, it will not be re-created.`,
	Run: func(cmd *cobra.Command, args []string) {
		initCollection(collectionPath)
	},
}

// Initialises a new collection by creating a git repository, a metadata file, and a config file.
// If a git repository already exists, it will not be re-initialized.
// If a file already exists, it will not be re-created.
func initCollection(colPath string) {
	// TODO: Provide a config option to disable git repository creation
	if err := initGitRepo(colPath); err != nil {
		log.Println(err)
	}

	metadataPath := filepath.Join(colPath, metadataFilename)
	if err := internal.InitFile(metadataPath, internal.MetadataFile); err != nil {
		log.Println(err)
	}

	configPath := filepath.Join(colPath, configFilename)
	if err := internal.InitFile(configPath, internal.ConfigFile); err != nil {
		log.Println(err)
	}
}

// Checks if a git repository exists at the path, if not, it will initialize a new one.
func initGitRepo(path string) error {
	_, err := git.PlainOpen(path)
	switch {
	case errors.Is(err, git.ErrRepositoryNotExists):
		fmt.Println("Initialising a new git repository...")
		_, err = git.PlainInit(".", false)
		if err != nil {
			return fmt.Errorf("Failed to initialize git repository: %w", err)
		}
		fmt.Println("Git repository initialized")
	case err == nil:
		fmt.Println("Git repository already exists")
	default:
		return fmt.Errorf("Failed to open git repository: %w", err)
	}
	return nil
}
