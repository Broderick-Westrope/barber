package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/Broderick-Westrope/barber/internal"
	"github.com/spf13/cobra"
)

var collRmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove a collection",
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
	_, err := os.Stat(metadataFile)
	if os.IsNotExist(err) {
		fmt.Printf("'%s' file does not exist\n", metadataFile)
		return
	} else if err != nil {
		log.Fatalf("Failed to check if '%s' file exists: %v", metadataFile, err)
	}

	if skipConfirm {
		internal.RemoveFile(metadataFile)
		return
	}

	fmt.Printf("Are you sure you want to remove '%s' file? [y/N] ", metadataFile)
	var response string
	if _, err = fmt.Scanln(&response); err != nil {
		log.Fatalf("Failed to read user input: %v", err)
	}
	if response == "y" || response == "Y" {
		internal.RemoveFile(metadataFile)
	} else {
		fmt.Printf("'%s' file was not removed\n", metadataFile)
	}
}
