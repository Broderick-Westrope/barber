package cmd

import (
	"log"
	"path/filepath"

	"github.com/Broderick-Westrope/barber/internal"
	"github.com/spf13/cobra"
)

var colResetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset a collection",
	Long: `Reset a collection by setting the metadata & config files to their default.
			This will not affect the git repository.`,
	Run: func(cmd *cobra.Command, args []string) {
		resetCollection(collectionFlag)
	},
}

// resetCollection resets a collection by resetting the metadata file, and the config file to their defaults.
// It does not affect the git repository.
// It uses the skipConfirm flag to determine if the user should be prompted before resetting the files.
func resetCollection(colPath string) {
	metadataPath := filepath.Join(colPath, metadataFilename)
	if err := internal.DestructiveFileOp(metadataPath, internal.MetadataFile, yesFlag, internal.ResetOp); err != nil {
		log.Fatalf("Failed to reset '%s' file: %v", metadataPath, err)
	}

	configPath := filepath.Join(colPath, configFilename)
	if err := internal.DestructiveFileOp(configPath, internal.ConfigFile, yesFlag, internal.ResetOp); err != nil {
		log.Fatalf("Failed to reset '%s' file: %v", configPath, err)
	}
}
