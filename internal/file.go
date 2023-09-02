package internal

import (
	"fmt"
	"log"
	"os"
)

func CreateMetadataFile(filename string) error {
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

func RemoveFile(filename string) {
	fmt.Printf("Removing '%s' file...\n", filename)
	if err := os.Remove(filename); err != nil {
		log.Fatalf("Failed to remove '%s' file: %v", filename, err)
	}
	fmt.Printf("Removed '%s' file\n", filename)
}
