package cmd

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"

	"github.com/Broderick-Westrope/barber/internal"
	"github.com/go-git/go-git/v5"
	"github.com/spf13/cobra"
)

var collInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new collection",
	Long: `Initialize a new collection by creating a git repository and a metadata file.
			If a git repository already exists, it will not be re-initialized.
			If a metadata file already exists, it will not be re-created.`,
	Run: func(cmd *cobra.Command, args []string) {
		initCollection()
	},
}

func init() {
	collectionCmd.AddCommand(collInitCmd)
}

func initCollection() {
	// Check if git repository exists, if not, initialize a new one
	_, err := git.PlainOpen(".")
	if err == git.ErrRepositoryNotExists {
		fmt.Println("Initializing a new git repository...")
		_, err := git.PlainInit(".", false)
		if err != nil {
			log.Fatalf("Failed to initialize git repository: %v", err)
		}
		fmt.Println("Git repository initialized")
	} else if err != nil {
		log.Fatalf("Failed to open git repository: %v", err)
	} else {
		fmt.Println("Git repository already exists")
	}

	// Check if metadata file exists, if not, create it
	if _, err = os.Stat(metadataFile); errors.Is(err, fs.ErrNotExist) {
		fmt.Printf("Creating '%s' metadata file...\n", metadataFile)
		if err := internal.CreateMetadataFile(metadataFile); err != nil {
			log.Fatalf("Failed to create metadata file: %v", err)
		}
		fmt.Printf("Created '%s' file\n", metadataFile)
	} else if err != nil {
		log.Fatalf("Failed to check if '%s' file exists: %v", metadataFile, err)
	} else {
		fmt.Printf("'%s' file already exists\n", metadataFile)
	}
}
