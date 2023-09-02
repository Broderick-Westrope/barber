package cmd

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"

	"github.com/Broderick-Westrope/barber/internal"
	"github.com/spf13/cobra"
)

var collResetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset a collection",
	Long: `Reset a collection by removing the metadata file.
			This will not affect the git repository.`,
	Run: func(cmd *cobra.Command, args []string) {
		resetCollection()
	},
}

func init() {
	collectionCmd.AddCommand(collResetCmd)
}

func resetCollection() {
	if _, err := os.Stat(metadataFile); errors.Is(err, fs.ErrNotExist) {
		fmt.Printf("'%s' file does not exist\n", metadataFile)
	} else if err != nil {
		log.Fatalf("Failed to check if '%s' file exists: %v", metadataFile, err)
	} else {
		fmt.Printf("Found '%s' file. Resetting it...\n", metadataFile)
		if err := internal.CreateMetadataFile(metadataFile); err != nil {
			log.Fatalf("Failed to reset '%s' file: %v", metadataFile, err)
		}
		fmt.Printf("Reset '%s' file\n", metadataFile)
	}
}
