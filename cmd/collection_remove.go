package cmd

import (
	"log"
	"path/filepath"

	"github.com/Broderick-Westrope/barber/internal"
	"github.com/spf13/cobra"
)

var colRemoveCmd = &cobra.Command{
	Use:     "remove",
	Aliases: []string{"rm"},
	Short:   "Remove a collection",
	Long: `Remove a collection by removing the metadata & config files.
			This will not remove the git repository.
			If a file does not exist, nothing will happen.`,
	Run: func(cmd *cobra.Command, args []string) {
		removeCollection()
	},
}

func init() {
	colRemoveCmd.Flags().BoolVarP(&skipConfirm, "yes", "y", false, "confirm removal without prompting")

	collectionCmd.AddCommand(colRemoveCmd)
}

// removeCollection removes a collection by removing the metadata file, and the config file.
// It does not affect the git repository.
// It uses the skipConfirm flag to determine if the user should be prompted before removing the metadata files.
func removeCollection() {
	// TODO: Get path from collection flag
	path := "."

	metadataPath := filepath.Join(path, metadataFilename)
	if err := internal.DestructiveFileOp(metadataPath, internal.Metadata, skipConfirm, internal.Delete); err != nil {
		log.Fatalf("Failed to delete '%s' file: %v", metadataPath, err)
	}

	configPath := filepath.Join(path, configFilename)
	if err := internal.DestructiveFileOp(configPath, internal.Config, skipConfirm, internal.Delete); err != nil {
		log.Fatalf("Failed to delete '%s' file: %v", configPath, err)
	}
}
