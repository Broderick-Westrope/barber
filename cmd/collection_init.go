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
	_, err := git.PlainOpen(".")
	switch {
	case errors.Is(err, git.ErrRepositoryNotExists):
		fmt.Println("Initializing a new git repository...")
		_, err = git.PlainInit(".", false)
		if err != nil {
			log.Fatalf("Failed to initialize git repository: %v", err)
		}
		fmt.Println("Git repository initialized")
	case err == nil:
		fmt.Println("Git repository already exists")
	default:
		log.Fatalf("Failed to open git repository: %v", err)
	}

	// Check if git repository exists, if not, initialize a new one
	_, err = git.PlainOpen(".")
	switch {
	case errors.Is(err, git.ErrRepositoryNotExists):
		fmt.Println("Initializing a new git repository...")
		_, err = git.PlainInit(".", false)
		if err != nil {
			log.Fatalf("Failed to initialize git repository: %v", err)
		}
		fmt.Println("Git repository initialized")
	case err == nil:
		fmt.Println("Git repository already exists")
	default:
		log.Fatalf("Failed to open git repository: %v", err)
	}

	// Check if metadata file exists, if not, create it
	_, err = os.Stat(defMetadataFile)
	switch {
	case errors.Is(err, fs.ErrNotExist):
		fmt.Printf("Creating '%s' metadata file...\n", defMetadataFile)
		if err = internal.CreateMetadataFile(defMetadataFile); err != nil {
			log.Fatalf("Failed to create metadata file: %v", err)
		}
		fmt.Printf("Created '%s' file\n", defMetadataFile)
	case err == nil:
		fmt.Printf("'%s' file already exists\n", defMetadataFile)
	default:
		log.Fatalf("Failed to check if '%s' file exists: %v", defMetadataFile, err)
	}
}
