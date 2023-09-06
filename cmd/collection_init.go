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
	Long: `Initialize a new collection by initializing a git repository, and creating the default metadata and config files.
			If a git repository already exists, it will not be re-initialized.
			If a file already exists, it will not be re-created.`,
	Run: func(cmd *cobra.Command, args []string) {
		initCollection()
	},
}

func init() {
	collectionCmd.AddCommand(colInitCmd)
}

// Initializes a new collection by creating a git repository, a metadata file, and a config file.
// If a git repository already exists, it will not be re-initialized.
// If a file already exists, it will not be re-created.
func initCollection() {
	// TODO: Get path from collection flag
	path := "."

	// TODO: Provide a config option to disable git repository creation
	if err := initGitRepo(path); err != nil {
		log.Println(err)
	}

	metadataPath := filepath.Join(path, metadataFilename)
	if err := internal.InitFile(metadataPath, internal.Metadata); err != nil {
		log.Println(err)
	}

	configPath := filepath.Join(path, configFilename)
	if err := internal.InitFile(configPath, internal.Config); err != nil {
		log.Println(err)
	}
}

// Checks if a git repository exists at the path, if not, it will initialize a new one.
func initGitRepo(path string) error {
	_, err := git.PlainOpen(path)
	switch {
	case errors.Is(err, git.ErrRepositoryNotExists):
		fmt.Println("Initializing a new git repository...")
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
