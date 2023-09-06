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
	Long: `Reset a collection by removing the metadata file.
			This will not affect the git repository.`,
	Run: func(cmd *cobra.Command, args []string) {
		resetCollection()
	},
}

func init() {
	colResetCmd.PersistentFlags().BoolVarP(&skipConfirm, "yes", "y", false, "confirm removal without prompting")

	collectionCmd.AddCommand(colResetCmd)
}

// resetCollection resets a collection by resetting the metadata file, and the config file to their defaults.
// It does not affect the git repository.
// It uses the skipConfirm flag to determine if the user should be prompted before resetting the files.
func resetCollection() {
	// TODO: Get path from collection flag
	path := "."

	metadataPath := filepath.Join(path, metadataFilename)
	if err := internal.DestructiveFileOp(metadataPath, internal.Metadata, skipConfirm, internal.Reset); err != nil {
		log.Fatalf("Failed to reset '%s' file: %v", metadataPath, err)
	}

	configPath := filepath.Join(path, configFilename)
	if err := internal.DestructiveFileOp(configPath, internal.Config, skipConfirm, internal.Reset); err != nil {
		log.Fatalf("Failed to reset '%s' file: %v", configPath, err)
	}
}
