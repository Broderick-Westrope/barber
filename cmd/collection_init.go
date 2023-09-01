package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var collInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new collection",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Implement
		fmt.Println("Not yet implemented")
	},
}

func init() {
	collectionCmd.AddCommand(collInitCmd)
}
