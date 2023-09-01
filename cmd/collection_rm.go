package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var collRmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove a collection",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Implement
		fmt.Println("Not yet implemented")
	},
}

func init() {
	collectionCmd.AddCommand(collRmCmd)
}
