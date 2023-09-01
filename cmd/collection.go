package cmd

import "github.com/spf13/cobra"

var collectionCmd = &cobra.Command{
	Use:   "collection",
	Short: "Manage snippet collections",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(collectionCmd)
}
