package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var collectionCmd = &cobra.Command{
	Use:     "collection",
	Aliases: []string{"col"},
	Short:   "Manage snippet collections",
	Run: func(cmd *cobra.Command, args []string) {
		if err := cmd.Help(); err != nil {
			log.Fatalf("Failed to display help: %v", err)
		}
	},
}
