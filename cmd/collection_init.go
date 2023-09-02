package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/spf13/cobra"
)

var collInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new collection",
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
	metadataFile := ".barber.yaml"
	if _, err = os.Stat(metadataFile); os.IsNotExist(err) {
		fmt.Printf("Creating a '%s' file...\n", metadataFile)
		if err := createMetadataFile(metadataFile); err != nil {
			log.Fatalf("Failed to create metadata file: %v", err)
		}
		fmt.Printf("Created '%s' file\n", metadataFile)
	} else if err != nil {
		log.Fatalf("Failed to check if '%s' file exists: %v", metadataFile, err)
	} else {
		fmt.Printf("'%s' file already exists. Resetting it...\n", metadataFile)
		if err := createMetadataFile(metadataFile); err != nil {
			log.Fatalf("Failed to reset '%s' file: %v", metadataFile, err)
		}
		fmt.Printf("Reset '%s' file\n", metadataFile)
	}
}

func createMetadataFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(`# Collection metadata`)
	if err != nil {
		return err
	}
	return nil
}
