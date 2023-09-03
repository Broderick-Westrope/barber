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
	if _, err := os.Stat(defMetadataFile); errors.Is(err, fs.ErrNotExist) {
		fmt.Printf("'%s' file does not exist\n", defMetadataFile)
	} else if err != nil {
		log.Fatalf("Failed to check if '%s' file exists: %v", defMetadataFile, err)
	} else {
		fmt.Printf("Found '%s' file. Resetting it...\n", defMetadataFile)
		if err = internal.CreateMetadataFile(defMetadataFile); err != nil {
			log.Fatalf("Failed to reset '%s' file: %v", defMetadataFile, err)
		}
		fmt.Printf("Reset '%s' file\n", defMetadataFile)
	}
}
