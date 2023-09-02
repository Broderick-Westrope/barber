package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var collectionCmd = &cobra.Command{
	Use:   "collection",
	Short: "Manage snippet collections",
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.Help()
		if err != nil {
			log.Fatalf("Failed to display help: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(collectionCmd)
}
