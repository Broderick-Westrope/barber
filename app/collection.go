package app

import (
	"fmt"

	"github.com/Broderick-Westrope/barber/file"
	"github.com/spf13/cobra"
)

var collectionCmd = &cobra.Command{
	Use:     "collection",
	Aliases: []string{"col"},
	Short:   "Commands for managing snippet collections",
	Run:     displayHelp,
}

var colInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialise a new collection",
	Long: `Initialise a new collection by initialising a git repository, and creating the default metadata and config files.
			If a git repository already exists, it will not be re-initialized.
			If a file already exists, it will not be re-created.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := initCollection(collectionFlag)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
		}
	},
}

var colResetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset a collection",
	Long: `Reset a collection by setting the metadata & config files to their default.
			This will not affect the git repository.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := resetCollection(collectionFlag, yesFlag)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
		}
	},
}

var colRemoveCmd = &cobra.Command{
	Use:     "remove",
	Aliases: []string{"rm"},
	Short:   "Remove a collection",
	Long: `Remove a collection by removing the metadata & config files.
			This will not remove the git repository.
			If a file does not exist, nothing will happen.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := removeCollection(collectionFlag, yesFlag)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
		}
	},
}

// Initialises a new collection by creating a git repository, a metadata file, and a config file.
// If a git repository already exists, it will not be re-initialized.
// If a file already exists, it will not be re-created.
func initCollection(colPath string) error {
	errMsg := fmt.Errorf("Failed to remove collection at '%s'", colPath)

	// TODO: Provide a config option to disable git repository creation
	if err := file.InitGitRepo(colPath); err != nil {
		return fmt.Errorf("%v: %w", errMsg, err)
	}

	if err := file.InitMetadata(colPath); err != nil {
		return fmt.Errorf("%v: %w", errMsg, err)
	}

	if err := file.InitConfig(colPath); err != nil {
		return fmt.Errorf("%v: %w", errMsg, err)
	}

	return nil
}

// Removes a collection by removing the metadata file, and the config file.
// It does not affect the git repository.
// It uses skipConfirm to determine if the user should be prompted before removing the metadata files.
func removeCollection(colPath string, skipConfirm bool) error {
	errMsg := fmt.Errorf("Failed to remove collection at '%s'", colPath)

	if err := file.DeleteMetadata(colPath, skipConfirm); err != nil {
		return fmt.Errorf("%v: %w", errMsg, err)
	}

	if err := file.DeleteConfig(colPath, skipConfirm); err != nil {
		return fmt.Errorf("%v: %w", errMsg, err)
	}

	return nil
}

// Resets a collection by resetting the metadata file, and the config file to their defaults.
// It does not affect the git repository.
// It uses skipConfirm to determine if the user should be prompted before resetting the files.
func resetCollection(colPath string, skipConfirm bool) error {
	errMsg := fmt.Errorf("Failed to reset collection at '%s'", colPath)

	if err := file.ResetMetadata(colPath, skipConfirm); err != nil {
		return fmt.Errorf("%v: %w", errMsg, err)
	}

	if err := file.ResetConfig(colPath, skipConfirm); err != nil {
		return fmt.Errorf("%v: %w", errMsg, err)
	}

	return nil
}
