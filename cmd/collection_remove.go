package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/Broderick-Westrope/barber/internal"
	"github.com/spf13/cobra"
)

var collRmCmd = &cobra.Command{
	Use:     "remove",
	Aliases: []string{"rm"},
	Short:   "Remove a collection",
	Long: `Remove a collection by removing the metadata file.
			This will not remove the git repository.
			If the metadata file does not exist, nothing will happen.`,
	Run: func(cmd *cobra.Command, args []string) {
		removeCollection()
	},
}

func init() {
	collRmCmd.PersistentFlags().BoolVarP(&skipConfirm, "yes", "y", false, "confirm removal without prompting")

	collectionCmd.AddCommand(collRmCmd)
}

func removeCollection() {
	_, err := os.Stat(defMetadataFile)
	if os.IsNotExist(err) {
		fmt.Printf("'%s' file does not exist\n", defMetadataFile)
		return
	} else if err != nil {
		log.Fatalf("Failed to check if '%s' file exists: %v", defMetadataFile, err)
	}

	if skipConfirm {
		internal.RemoveFile(defMetadataFile)
		return
	}

	fmt.Printf("Are you sure you want to remove '%s' file? [y/N] ", defMetadataFile)
	var response string
	if _, err = fmt.Scanln(&response); err != nil {
		log.Fatalf("Failed to read user input: %v", err)
	}
	if response == "y" || response == "Y" {
		internal.RemoveFile(defMetadataFile)
	} else {
		fmt.Printf("'%s' file was not removed\n", defMetadataFile)
	}
}
